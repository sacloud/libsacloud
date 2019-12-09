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

package disk

import (
	"context"
	"errors"
	"fmt"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/ostype"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
	archiveUtil "github.com/sacloud/libsacloud/v2/utils/archive"
)

// Builder ディスクの構築インターフェース
type Builder interface {
	Validate(ctx context.Context, zone string) error
	BuildDisk(ctx context.Context, zone string, serverID types.ID) (*BuildResult, error)
}

// BuildResult ディスク構築結果
type BuildResult struct {
	DiskID          types.ID
	GeneratedSSHKey *sacloud.SSHKeyGenerated
}

// FromUnixBuilder Unix系パブリックアーカイブからディスクを作成するリクエスト
type FromUnixBuilder struct {
	OSType ostype.ArchiveOSType

	Name        string
	SizeGB      int
	DistantFrom []types.ID
	PlanID      types.ID
	Connection  types.EDiskConnection
	Description string
	Tags        types.Tags
	IconID      types.ID

	EditParameter *UnixEditRequest

	Client *APIClient

	generatedSSHKey *sacloud.SSHKeyGenerated
	generatedNotes  []*sacloud.Note
}

// Validate 設定値の検証
func (d *FromUnixBuilder) Validate(ctx context.Context, zone string) error {
	if !d.OSType.IsSupportDiskEdit() {
		return fmt.Errorf("invalid OSType: %s", d.OSType.String())
	}
	if err := validateDiskPlan(ctx, d.Client, zone, d.PlanID, d.SizeGB); err != nil {
		return err
	}

	if d.EditParameter != nil {
		return d.EditParameter.Validate(ctx, d.Client)
	}
	return nil
}

// BuildDisk ディスクの構築
func (d *FromUnixBuilder) BuildDisk(ctx context.Context, zone string, serverID types.ID) (*BuildResult, error) {
	res, err := build(ctx, d.Client, zone, serverID, d.DistantFrom, d)
	if err != nil {
		return nil, err
	}
	if d.generatedSSHKey != nil {
		res.GeneratedSSHKey = d.generatedSSHKey
	}

	if d.EditParameter != nil {
		if d.EditParameter.IsSSHKeysEphemeral {
			if err := d.Client.SSHKey.Delete(ctx, d.generatedSSHKey.ID); err != nil {
				return nil, err
			}
		}
		if d.EditParameter.IsNotesEphemeral {
			for _, note := range d.generatedNotes {
				if err := d.Client.Note.Delete(ctx, note.ID); err != nil {
					return nil, err
				}
			}
		}
	}
	return res, nil
}

func (d *FromUnixBuilder) createDiskParameter(ctx context.Context, client *APIClient, zone string, serverID types.ID) (*sacloud.DiskCreateRequest, *sacloud.DiskEditRequest, error) {
	archive, err := archiveUtil.FindByOSType(ctx, client.Archive, zone, d.OSType)
	if err != nil {
		return nil, nil, err
	}

	createReq := &sacloud.DiskCreateRequest{
		DiskPlanID:      d.PlanID,
		SizeMB:          d.SizeGB * 1024,
		Connection:      d.Connection,
		SourceArchiveID: archive.ID,
		ServerID:        serverID,
		Name:            d.Name,
		Description:     d.Description,
		Tags:            d.Tags,
		IconID:          d.IconID,
	}

	var editReq *sacloud.DiskEditRequest
	if d.EditParameter != nil {
		req, sshKey, notes, err := d.EditParameter.prepareDiskEditParameter(ctx, client)
		if err != nil {
			return nil, nil, err
		}
		editReq = req
		if sshKey != nil {
			d.generatedSSHKey = sshKey
		}
		if len(notes) > 0 {
			d.generatedNotes = notes
		}
	}

	return createReq, editReq, nil
}

// FromFixedArchiveBuilder ディスクの修正をサポートしないパブリックアーカイブからディスクを作成するリクエスト
type FromFixedArchiveBuilder struct {
	OSType ostype.ArchiveOSType

	Name        string
	SizeGB      int
	DistantFrom []types.ID
	PlanID      types.ID
	Connection  types.EDiskConnection
	Description string
	Tags        types.Tags
	IconID      types.ID

	Client *APIClient

	generatedSSHKey *sacloud.SSHKeyGenerated
	generatedNotes  []*sacloud.Note
}

// Validate 設定値の検証
func (d *FromFixedArchiveBuilder) Validate(ctx context.Context, zone string) error {
	if d.OSType.IsSupportDiskEdit() || d.OSType.IsWindows() {
		return fmt.Errorf("invalid OSType: %s", d.OSType.String())
	}
	if err := validateDiskPlan(ctx, d.Client, zone, d.PlanID, d.SizeGB); err != nil {
		return err
	}

	return nil
}

// BuildDisk ディスクの構築
func (d *FromFixedArchiveBuilder) BuildDisk(ctx context.Context, zone string, serverID types.ID) (*BuildResult, error) {
	res, err := build(ctx, d.Client, zone, serverID, d.DistantFrom, d)
	if err != nil {
		return nil, err
	}
	if d.generatedSSHKey != nil {
		res.GeneratedSSHKey = d.generatedSSHKey
	}
	return res, nil
}

func (d *FromFixedArchiveBuilder) createDiskParameter(ctx context.Context, client *APIClient, zone string, serverID types.ID) (*sacloud.DiskCreateRequest, *sacloud.DiskEditRequest, error) {
	archive, err := archiveUtil.FindByOSType(ctx, client.Archive, zone, d.OSType)
	if err != nil {
		return nil, nil, err
	}

	createReq := &sacloud.DiskCreateRequest{
		DiskPlanID:      d.PlanID,
		SizeMB:          d.SizeGB * 1024,
		Connection:      d.Connection,
		SourceArchiveID: archive.ID,
		ServerID:        serverID,
		Name:            d.Name,
		Description:     d.Description,
		Tags:            d.Tags,
		IconID:          d.IconID,
	}
	return createReq, nil, nil
}

// FromWindowsBuilder Windows系パブリックアーカイブからディスクを作成するリクエスト
type FromWindowsBuilder struct {
	OSType ostype.ArchiveOSType

	Name        string
	SizeGB      int
	DistantFrom []types.ID
	PlanID      types.ID
	Connection  types.EDiskConnection
	Description string
	Tags        types.Tags
	IconID      types.ID

	EditParameter *WindowsEditRequest

	Client *APIClient
}

// Validate 設定値の検証
func (d *FromWindowsBuilder) Validate(ctx context.Context, zone string) error {
	if !d.OSType.IsWindows() {
		return fmt.Errorf("invalid OSType: %s", d.OSType.String())
	}
	if err := validateDiskPlan(ctx, d.Client, zone, d.PlanID, d.SizeGB); err != nil {
		return err
	}
	return nil
}

// BuildDisk ディスクの構築
func (d *FromWindowsBuilder) BuildDisk(ctx context.Context, zone string, serverID types.ID) (*BuildResult, error) {
	res, err := build(ctx, d.Client, zone, serverID, d.DistantFrom, d)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (d *FromWindowsBuilder) createDiskParameter(ctx context.Context, client *APIClient, zone string, serverID types.ID) (*sacloud.DiskCreateRequest, *sacloud.DiskEditRequest, error) {
	archive, err := archiveUtil.FindByOSType(ctx, client.Archive, zone, d.OSType)
	if err != nil {
		return nil, nil, err
	}

	createReq := &sacloud.DiskCreateRequest{
		DiskPlanID:      d.PlanID,
		SizeMB:          d.SizeGB * 1024,
		Connection:      d.Connection,
		SourceArchiveID: archive.ID,
		ServerID:        serverID,
		Name:            d.Name,
		Description:     d.Description,
		Tags:            d.Tags,
		IconID:          d.IconID,
	}

	var editReq *sacloud.DiskEditRequest
	if d.EditParameter != nil {
		editReq = d.EditParameter.prepareDiskEditParameter()
	}

	return createReq, editReq, nil
}

// FromDiskOrArchiveBuilder ディスクorアーカイブからディスクを作成するリクエスト
//
// ディスクの修正が可能かは実行時にさくらのクラウドAPI側にて判定される
type FromDiskOrArchiveBuilder struct {
	SourceDiskID    types.ID
	SourceArchiveID types.ID

	Name        string
	SizeGB      int
	DistantFrom []types.ID
	PlanID      types.ID
	Connection  types.EDiskConnection
	Description string
	Tags        types.Tags
	IconID      types.ID

	EditParameter *UnixEditRequest

	Client *APIClient

	generatedSSHKey *sacloud.SSHKeyGenerated
	generatedNotes  []*sacloud.Note
}

// Validate 設定値の検証
func (d *FromDiskOrArchiveBuilder) Validate(ctx context.Context, zone string) error {
	if d.SourceArchiveID.IsEmpty() && d.SourceDiskID.IsEmpty() {
		return errors.New("SourceArchiveID or SourceDiskID is required")
	}
	if err := validateDiskPlan(ctx, d.Client, zone, d.PlanID, d.SizeGB); err != nil {
		return err
	}

	if !d.SourceArchiveID.IsEmpty() {
		if _, err := d.Client.Archive.Read(ctx, zone, d.SourceArchiveID); err != nil {
			return err
		}
	}
	if !d.SourceDiskID.IsEmpty() {
		if _, err := d.Client.Disk.Read(ctx, zone, d.SourceDiskID); err != nil {
			return err
		}
	}
	return nil
}

// BuildDisk ディスクの構築
func (d *FromDiskOrArchiveBuilder) BuildDisk(ctx context.Context, zone string, serverID types.ID) (*BuildResult, error) {
	res, err := build(ctx, d.Client, zone, serverID, d.DistantFrom, d)
	if err != nil {
		return nil, err
	}
	if d.generatedSSHKey != nil {
		res.GeneratedSSHKey = d.generatedSSHKey
	}

	if d.EditParameter != nil {
		if d.EditParameter.IsSSHKeysEphemeral {
			if err := d.Client.SSHKey.Delete(ctx, d.generatedSSHKey.ID); err != nil {
				return nil, err
			}
		}
		if d.EditParameter.IsNotesEphemeral {
			for _, note := range d.generatedNotes {
				if err := d.Client.Note.Delete(ctx, note.ID); err != nil {
					return nil, err
				}
			}
		}
	}
	return res, nil
}

func (d *FromDiskOrArchiveBuilder) createDiskParameter(ctx context.Context, client *APIClient, zone string, serverID types.ID) (*sacloud.DiskCreateRequest, *sacloud.DiskEditRequest, error) {
	createReq := &sacloud.DiskCreateRequest{
		DiskPlanID:      d.PlanID,
		SizeMB:          d.SizeGB * 1024,
		Connection:      d.Connection,
		SourceArchiveID: d.SourceArchiveID,
		SourceDiskID:    d.SourceDiskID,
		ServerID:        serverID,
		Name:            d.Name,
		Description:     d.Description,
		Tags:            d.Tags,
		IconID:          d.IconID,
	}

	var editReq *sacloud.DiskEditRequest
	if d.EditParameter != nil {
		req, sshKey, notes, err := d.EditParameter.prepareDiskEditParameter(ctx, client)
		if err != nil {
			return nil, nil, err
		}
		editReq = req
		if sshKey != nil {
			d.generatedSSHKey = sshKey
		}
		if len(notes) > 0 {
			d.generatedNotes = notes
		}
	}

	return createReq, editReq, nil
}

// BlankBuilder ブランクディスクを作成する場合のリクエスト
type BlankBuilder struct {
	Name        string
	SizeGB      int
	DistantFrom []types.ID
	PlanID      types.ID
	Connection  types.EDiskConnection
	Description string
	Tags        types.Tags
	IconID      types.ID

	Client *APIClient
}

// Validate 設定値の検証
func (d *BlankBuilder) Validate(ctx context.Context, zone string) error {
	if err := validateDiskPlan(ctx, d.Client, zone, d.PlanID, d.SizeGB); err != nil {
		return err
	}
	return nil
}

// BuildDisk ディスクの構築
func (d *BlankBuilder) BuildDisk(ctx context.Context, zone string, serverID types.ID) (*BuildResult, error) {
	return build(ctx, d.Client, zone, serverID, d.DistantFrom, d)
}

func (d *BlankBuilder) createDiskParameter(ctx context.Context, client *APIClient, zone string, serverID types.ID) (*sacloud.DiskCreateRequest, *sacloud.DiskEditRequest, error) {
	createReq := &sacloud.DiskCreateRequest{
		DiskPlanID:  d.PlanID,
		SizeMB:      d.SizeGB * 1024,
		Connection:  d.Connection,
		ServerID:    serverID,
		Name:        d.Name,
		Description: d.Description,
		Tags:        d.Tags,
		IconID:      d.IconID,
	}
	return createReq, nil, nil
}

// ConnectedDiskBuilder 既存ディスクを接続する場合のリクエスト
type ConnectedDiskBuilder struct {
	DiskID        types.ID
	EditParameter *UnixEditRequest
	Client        *APIClient
}

// Validate 設定値の検証
func (d *ConnectedDiskBuilder) Validate(ctx context.Context, zone string) error {
	if d.DiskID.IsEmpty() {
		return errors.New("DiskID is required")
	}

	if _, err := d.Client.Disk.Read(ctx, zone, d.DiskID); err != nil {
		return err
	}
	return nil
}

// BuildDisk ディスクの構築
func (d *ConnectedDiskBuilder) BuildDisk(ctx context.Context, zone string, serverID types.ID) (*BuildResult, error) {
	return &BuildResult{
		DiskID: d.DiskID,
	}, nil
}

type diskBuilder interface {
	createDiskParameter(
		ctx context.Context,
		client *APIClient,
		zone string,
		serverID types.ID,
	) (*sacloud.DiskCreateRequest, *sacloud.DiskEditRequest, error)
}

func build(ctx context.Context, client *APIClient, zone string, serverID types.ID, distantFrom []types.ID, builder diskBuilder) (*BuildResult, error) {
	var err error

	diskReq, editReq, err := builder.createDiskParameter(ctx, client, zone, serverID)
	if err != nil {
		return nil, err
	}
	if diskReq == nil {
		return nil, fmt.Errorf("disk create request is nil")
	}
	diskReq.ServerID = serverID

	var disk *sacloud.Disk

	if editReq == nil {
		disk, err = client.Disk.Create(ctx, zone, diskReq, distantFrom)
	} else {
		disk, err = client.Disk.CreateWithConfig(ctx, zone, diskReq, editReq, false, distantFrom)
	}
	if err != nil {
		return nil, err
	}

	waiter := sacloud.WaiterForReady(func() (interface{}, error) {
		return client.Disk.Read(ctx, zone, disk.ID)
	})
	lastState, err := waiter.WaitForState(ctx)
	if err != nil {
		return nil, err
	}
	disk = lastState.(*sacloud.Disk)

	return &BuildResult{DiskID: disk.ID}, nil
}

func validateDiskPlan(ctx context.Context, client *APIClient, zone string, diskPlanID types.ID, sizeGB int) error {
	plan, err := client.DiskPlan.Read(ctx, zone, diskPlanID)
	if err != nil {
		return err
	}
	found := false
	for _, size := range plan.Size {
		if size.Availability.IsAvailable() && size.GetSizeGB() == sizeGB {
			found = true
			break
		}
	}
	if !found {
		return fmt.Errorf("disk plan[%s:%dGB] is not found", plan.Name, sizeGB)
	}
	return nil
}