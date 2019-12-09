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

type dummyArchiveFinder struct {
	archive *sacloud.Archive
	err     error
}

func (d *dummyArchiveFinder) Find(ctx context.Context, zone string, conditions *sacloud.FindCondition) (*sacloud.ArchiveFindResult, error) {
	if d.err != nil {
		return nil, d.err
	}

	count := 0
	if d.archive != nil {
		count = 1
	}
	return &sacloud.ArchiveFindResult{
		Total:    count,
		Count:    count,
		Archives: []*sacloud.Archive{d.archive},
	}, nil
}
func (d *dummyArchiveFinder) Read(ctx context.Context, zone string, id types.ID) (*sacloud.Archive, error) {
	if d.err != nil {
		return nil, d.err
	}
	return d.archive, nil
}

type dummyDiskHandler struct {
	disk *sacloud.Disk
	err  error
}

func (d *dummyDiskHandler) Create(ctx context.Context, zone string, createParam *sacloud.DiskCreateRequest, distantFrom []types.ID) (*sacloud.Disk, error) {
	if d.err != nil {
		return nil, d.err
	}
	return d.disk, nil
}

func (d *dummyDiskHandler) CreateWithConfig(ctx context.Context, zone string, createParam *sacloud.DiskCreateRequest, editParam *sacloud.DiskEditRequest, bootAtAvailable bool, distantFrom []types.ID) (*sacloud.Disk, error) {
	return d.Create(ctx, zone, createParam, distantFrom)
}

func (d *dummyDiskHandler) Read(ctx context.Context, zone string, id types.ID) (*sacloud.Disk, error) {
	if d.err != nil {
		return nil, d.err
	}
	return d.disk, nil
}

type dummyDiskPlanReader struct {
	diskPlan *sacloud.DiskPlan
	err      error
}

func (d *dummyDiskPlanReader) Read(ctx context.Context, zone string, id types.ID) (*sacloud.DiskPlan, error) {
	if d.err != nil {
		return nil, d.err
	}
	return d.diskPlan, nil
}

type dummySwitchReader struct {
	sw  *sacloud.Switch
	err error
}

func (d *dummySwitchReader) Read(ctx context.Context, zone string, id types.ID) (*sacloud.Switch, error) {
	if d.err != nil {
		return nil, d.err
	}
	return d.sw, nil
}

type dummyInterfaceHandler struct {
	iface *sacloud.Interface
	err   error
}

func (d *dummyInterfaceHandler) ConnectToPacketFilter(ctx context.Context, zone string, id types.ID, packetFilterID types.ID) error {
	return d.err
}

func (d *dummyInterfaceHandler) Update(ctx context.Context, zone string, id types.ID, param *sacloud.InterfaceUpdateRequest) (*sacloud.Interface, error) {
	if d.err != nil {
		return nil, d.err
	}
	return d.iface, nil
}

type dummyPacketFilterReader struct {
	packetFilter *sacloud.PacketFilter
	err          error
}

func (d *dummyPacketFilterReader) Read(ctx context.Context, zone string, id types.ID) (*sacloud.PacketFilter, error) {
	if d.err != nil {
		return nil, d.err
	}
	return d.packetFilter, nil
}

type dummyCreateServerHandler struct {
	server   *sacloud.Server
	err      error
	cdromErr error
	bootErr  error
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

func (d *dummyCreateServerHandler) InsertCDROM(ctx context.Context, zone string, id types.ID, insertParam *sacloud.InsertCDROMRequest) error {
	return d.cdromErr
}

func (d *dummyCreateServerHandler) Boot(ctx context.Context, zone string, id types.ID) error {
	return d.bootErr
}

type dummyNoteHandler struct {
	note *sacloud.Note
	err  error
}

func (d *dummyNoteHandler) Read(ctx context.Context, id types.ID) (*sacloud.Note, error) {
	if d.err != nil {
		return nil, d.err
	}
	return d.note, nil
}

func (d *dummyNoteHandler) Create(ctx context.Context, param *sacloud.NoteCreateRequest) (*sacloud.Note, error) {
	if d.err != nil {
		return nil, d.err
	}
	return d.note, nil
}

func (d *dummyNoteHandler) Delete(ctx context.Context, id types.ID) error {
	return d.err
}

type dummySSHKeyHandler struct {
	sshKey          *sacloud.SSHKey
	sshKeyGenerated *sacloud.SSHKeyGenerated
	err             error
}

func (d *dummySSHKeyHandler) Read(ctx context.Context, id types.ID) (*sacloud.SSHKey, error) {
	if d.err != nil {
		return nil, d.err
	}
	return d.sshKey, nil
}

func (d *dummySSHKeyHandler) Generate(ctx context.Context, param *sacloud.SSHKeyGenerateRequest) (*sacloud.SSHKeyGenerated, error) {
	if d.err != nil {
		return nil, d.err
	}
	return d.sshKeyGenerated, nil
}

func (d *dummySSHKeyHandler) Delete(ctx context.Context, id types.ID) error {
	return d.err
}
