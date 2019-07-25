package server

import "C"
import (
	"context"
	"errors"
	"fmt"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

// Builder サーバ作成時のパラメータ
type Builder struct {
	Client *BuildersAPIClient

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
	DiskBuilders    []DiskBuilder
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

// Validate .
func (b *Builder) Validate(ctx context.Context, zone string) error {
	b.setDefaults()

	// Fields
	if b.Client == nil {
		return errors.New("field 'Client' is not set")
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
	plan, err := FindPlan(ctx, b.Client.ServerPlan, zone, &FindPlanRequest{
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
		if err := diskBuilder.Validate(ctx, b.Client, zone); err != nil {
			return err
		}
	}

	return nil
}

// Build .
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
		if err := diskReq.Validate(ctx, b.Client, zone); err != nil {
			return nil, err
		}
		builtDisk, err := diskReq.BuildDisk(ctx, b.Client, zone, server.ID)
		if err != nil {
			return nil, err
		}
		if builtDisk.GeneratedSSHKey != nil {
			result.GeneratedSSHPrivateKey = builtDisk.GeneratedSSHKey.PrivateKey
		}
	}

	// connect packet filter
	if err := b.connectPacketFilter(ctx, zone, server); err != nil {
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

type packetFilterRequest struct {
	index          int
	packetFilterID types.ID
}

func (b *Builder) collectPacketFilterIDs() []*packetFilterRequest {
	var pfs []*packetFilterRequest
	if b.NIC != nil {
		pfs = append(pfs, &packetFilterRequest{
			index:          0,
			packetFilterID: b.NIC.GetPacketFilterID(),
		})
	}
	for i, nic := range b.AdditionalNICs {
		pfs = append(pfs, &packetFilterRequest{
			index:          i + 1,
			packetFilterID: nic.GetPacketFilterID(),
		})
	}
	return pfs
}

func (b *Builder) connectPacketFilter(ctx context.Context, zone string, server *sacloud.Server) error {
	requests := b.collectPacketFilterIDs()
	for _, pfr := range requests {
		if pfr.packetFilterID.IsEmpty() {
			continue
		}
		if pfr.index < len(server.Interfaces) {
			iface := server.Interfaces[pfr.index]
			if err := b.Client.Interface.ConnectToPacketFilter(ctx, zone, iface.ID, pfr.packetFilterID); err != nil {
				return err
			}
		}
	}
	return nil
}
