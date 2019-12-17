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

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

type dummyPlanFinder struct {
	plans []*sacloud.ServerPlan
	err   error
}

func (f *dummyPlanFinder) Find(ctx context.Context, zone string, conditions *sacloud.FindCondition) (*sacloud.ServerPlanFindResult, error) {
	if f.err != nil {
		return nil, f.err
	}

	return &sacloud.ServerPlanFindResult{
		Total:       len(f.plans),
		Count:       len(f.plans),
		ServerPlans: f.plans,
	}, nil
}

type dummyInterfaceHandler struct {
	iface *sacloud.Interface
	err   error
}

func (d *dummyInterfaceHandler) Create(ctx context.Context, zone string, param *sacloud.InterfaceCreateRequest) (*sacloud.Interface, error) {
	if d.err != nil {
		return nil, d.err
	}
	return d.iface, nil
}
func (d *dummyInterfaceHandler) Update(ctx context.Context, zone string, id types.ID, param *sacloud.InterfaceUpdateRequest) (*sacloud.Interface, error) {
	if d.err != nil {
		return nil, d.err
	}
	return d.iface, nil
}
func (d *dummyInterfaceHandler) Delete(ctx context.Context, zone string, id types.ID) error {
	return d.err
}
func (d *dummyInterfaceHandler) ConnectToSharedSegment(ctx context.Context, zone string, id types.ID) error {
	return d.err
}
func (d *dummyInterfaceHandler) ConnectToSwitch(ctx context.Context, zone string, id types.ID, switchID types.ID) error {
	return d.err
}
func (d *dummyInterfaceHandler) DisconnectFromSwitch(ctx context.Context, zone string, id types.ID) error {
	return d.err
}
func (d *dummyInterfaceHandler) ConnectToPacketFilter(ctx context.Context, zone string, id types.ID, packetFilterID types.ID) error {
	return d.err
}
func (d *dummyInterfaceHandler) DisconnectFromPacketFilter(ctx context.Context, zone string, id types.ID) error {
	return d.err
}

type dummyCreateServerHandler struct {
	server      *sacloud.Server
	err         error
	cdromErr    error
	bootErr     error
	shutdownErr error
}

func (d *dummyCreateServerHandler) Create(ctx context.Context, zone string, param *sacloud.ServerCreateRequest) (*sacloud.Server, error) {
	if d.err != nil {
		return nil, d.err
	}
	return d.server, nil
}

func (d *dummyCreateServerHandler) Read(ctx context.Context, zone string, id types.ID) (*sacloud.Server, error) {
	if d.err != nil {
		return nil, d.err
	}
	return d.server, nil
}

func (d *dummyCreateServerHandler) Update(ctx context.Context, zone string, id types.ID, param *sacloud.ServerUpdateRequest) (*sacloud.Server, error) {
	if d.err != nil {
		return nil, d.err
	}
	return d.server, nil
}

func (d *dummyCreateServerHandler) InsertCDROM(ctx context.Context, zone string, id types.ID, insertParam *sacloud.InsertCDROMRequest) error {
	return d.cdromErr
}

func (d *dummyCreateServerHandler) EjectCDROM(ctx context.Context, zone string, id types.ID, ejectParam *sacloud.EjectCDROMRequest) error {
	return d.cdromErr
}

func (d *dummyCreateServerHandler) Boot(ctx context.Context, zone string, id types.ID) error {
	return d.bootErr
}

func (d *dummyCreateServerHandler) Shutdown(ctx context.Context, zone string, id types.ID, shutdownOption *sacloud.ShutdownOption) error {
	return d.shutdownErr
}

func (d *dummyCreateServerHandler) ChangePlan(ctx context.Context, zone string, id types.ID, plan *sacloud.ServerChangePlanRequest) (*sacloud.Server, error) {
	if d.err != nil {
		return nil, d.err
	}
	return d.server, nil
}
