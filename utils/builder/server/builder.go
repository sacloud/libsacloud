// Copyright 2016-2019 The Libsacloud Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package server

import (
	"context"
	"errors"
	"fmt"

	"github.com/sacloud/libsacloud/v2/utils/builder/disk"

	"github.com/sacloud/libsacloud/v2/utils/server"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

// Builder サーバ作成時のパラメータ
type Builder struct {
	Name            string
	CPU             int
	MemoryGB        int
	Commitment      types.ECommitment
	Generation      types.EPlanGeneration
	InterfaceDriver types.EInterfaceDriver
	Description     string
	IconID          types.ID
	Tags            types.Tags
	BootAfterCreate bool
	CDROMID         types.ID
	PrivateHostID   types.ID
	NIC             NICSettingHolder
	AdditionalNICs  []AdditionalNICSettingHolder
	DiskBuilders    []disk.Builder

	Client *APIClient
}

// BuildResult サーバ構築結果
type BuildResult struct {
	ServerID               types.ID
	GeneratedSSHPrivateKey string
}

var (
	defaultCPU             = 1
	defaultMemoryGB        = 1
	defaultCommitment      = types.Commitments.Standard
	defaultGeneration      = types.PlanGenerations.Default
	defaultInterfaceDriver = types.InterfaceDrivers.VirtIO
)

func (b *Builder) setDefaults() {
	if b.CPU == 0 {
		b.CPU = defaultCPU
	}
	if b.MemoryGB == 0 {
		b.MemoryGB = defaultMemoryGB
	}
	if b.Commitment == types.ECommitment("") {
		b.Commitment = defaultCommitment
	}
	if b.Generation == types.EPlanGeneration(0) {
		b.Generation = defaultGeneration
	}
	if b.InterfaceDriver == types.EInterfaceDriver("") {
		b.InterfaceDriver = defaultInterfaceDriver
	}
}

// Validate 入力値の検証
//
// 各種IDの存在確認のためにAPIリクエストが行われます。
func (b *Builder) Validate(ctx context.Context, zone string) error {
	b.setDefaults()

	// Fields
	if b.Client == nil {
		return errors.New("client is empty")
	}

	if b.NIC == nil && len(b.AdditionalNICs) > 0 {
		return errors.New("NIC is required when AdditionalNICs is specified")
	}

	if len(b.AdditionalNICs) > 3 {
		return errors.New("AdditionalNICs must be less than 4")
	}

	if b.InterfaceDriver != types.InterfaceDrivers.E1000 && b.InterfaceDriver != types.InterfaceDrivers.VirtIO {
		return fmt.Errorf("invalid InterfaceDriver: %s", b.InterfaceDriver)
	}

	// Field values
	plan, err := server.FindPlan(ctx, b.Client.ServerPlan, zone, &server.FindPlanRequest{
		CPU:        b.CPU,
		MemoryGB:   b.MemoryGB,
		Commitment: b.Commitment,
		Generation: b.Generation,
	})
	if err != nil {
		return err
	}
	b.CPU = plan.CPU
	b.MemoryGB = plan.GetMemoryGB()
	b.Commitment = plan.Commitment
	b.Generation = plan.Generation

	for _, diskBuilder := range b.DiskBuilders {
		if err := diskBuilder.Validate(ctx, zone); err != nil {
			return err
		}
	}

	return nil
}

// Build サーバ構築を行う
func (b *Builder) Build(ctx context.Context, zone string) (*BuildResult, error) {
	// validate
	if err := b.Validate(ctx, zone); err != nil {
		return nil, err
	}

	// create server
	server, err := b.createServer(ctx, zone)
	if err != nil {
		return nil, err
	}
	result := &BuildResult{
		ServerID: server.ID,
	}

	// create&connect disk(s)
	for _, diskReq := range b.DiskBuilders {
		if err := diskReq.Validate(ctx, zone); err != nil {
			return nil, err
		}
		builtDisk, err := diskReq.BuildDisk(ctx, zone, server.ID)
		if err != nil {
			return nil, err
		}
		if builtDisk.GeneratedSSHKey != nil {
			result.GeneratedSSHPrivateKey = builtDisk.GeneratedSSHKey.PrivateKey
		}
	}

	// connect packet filter
	if err := b.updateInterfaces(ctx, zone, server); err != nil {
		return nil, err
	}

	// insert CD-ROM
	if !b.CDROMID.IsEmpty() {
		req := &sacloud.InsertCDROMRequest{ID: b.CDROMID}
		if err := b.Client.Server.InsertCDROM(ctx, zone, server.ID, req); err != nil {
			return nil, err
		}
	}

	// bool
	if b.BootAfterCreate {
		if err := b.Client.Server.Boot(ctx, zone, server.ID); err != nil {
			return nil, err
		}
		// wait
		waiter := sacloud.WaiterForUp(func() (interface{}, error) {
			return b.Client.Server.Read(ctx, zone, server.ID)
		})

		lastState, err := waiter.WaitForState(ctx)
		if err != nil {
			return nil, err
		}
		server = lastState.(*sacloud.Server)
	}

	return result, nil
}

// createServer サーバ作成
func (b *Builder) createServer(ctx context.Context, zone string) (*sacloud.Server, error) {
	param := &sacloud.ServerCreateRequest{
		CPU:                  b.CPU,
		MemoryMB:             b.MemoryGB * 1024,
		ServerPlanCommitment: b.Commitment,
		ServerPlanGeneration: b.Generation,
		InterfaceDriver:      b.InterfaceDriver,
		Name:                 b.Name,
		Description:          b.Description,
		Tags:                 b.Tags,
		IconID:               b.IconID,
		WaitDiskMigration:    false,
		ConnectedSwitches:    []*sacloud.ConnectedSwitch{},
	}
	if b.NIC != nil {
		cs := b.NIC.GetConnectedSwitchParam()
		if cs == nil {
			param.ConnectedSwitches = append(param.ConnectedSwitches, nil)
		} else {
			param.ConnectedSwitches = append(param.ConnectedSwitches, cs)
		}
	}
	if len(b.AdditionalNICs) > 0 {
		for _, nic := range b.AdditionalNICs {
			switchID := nic.GetSwitchID()
			if switchID.IsEmpty() {
				param.ConnectedSwitches = append(param.ConnectedSwitches, nil)
			} else {
				param.ConnectedSwitches = append(param.ConnectedSwitches, &sacloud.ConnectedSwitch{ID: switchID})
			}
		}
	}
	return b.Client.Server.Create(ctx, zone, param)
}

type updateInterfaceRequest struct {
	index          int
	packetFilterID types.ID
	displayIP      string
}

func (b *Builder) collectInterfaceParameters() []*updateInterfaceRequest {
	var reqs []*updateInterfaceRequest
	if b.NIC != nil {
		reqs = append(reqs, &updateInterfaceRequest{
			index:          0,
			packetFilterID: b.NIC.GetPacketFilterID(),
		})
	}
	for i, nic := range b.AdditionalNICs {
		reqs = append(reqs, &updateInterfaceRequest{
			index:          i + 1,
			packetFilterID: nic.GetPacketFilterID(),
		})
	}
	return reqs
}

func (b *Builder) updateInterfaces(ctx context.Context, zone string, server *sacloud.Server) error {
	requests := b.collectInterfaceParameters()
	for _, req := range requests {
		if req.index < len(server.Interfaces) {
			iface := server.Interfaces[req.index]

			if !req.packetFilterID.IsEmpty() {
				if err := b.Client.Interface.ConnectToPacketFilter(ctx, zone, iface.ID, req.packetFilterID); err != nil {
					return err
				}
			}

			if req.displayIP != "" {
				if _, err := b.Client.Interface.Update(ctx, zone, iface.ID, &sacloud.InterfaceUpdateRequest{
					UserIPAddress: req.displayIP,
				}); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
