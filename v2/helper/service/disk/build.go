// Copyright 2016-2020 The Libsacloud Authors
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

	"github.com/sacloud/libsacloud/v2/sacloud/ostype"

	diskBuilder "github.com/sacloud/libsacloud/v2/helper/builder/disk"
	"github.com/sacloud/libsacloud/v2/helper/validate"
	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

/* TODO あとでbuilderと統合する  */

type BuildRequest struct {
	Zone string `validate:"required" mapconv:"-"`
	Name string `validate:"required"`

	ID types.ID

	Description string `validate:"min=0,max=512"`
	Tags        types.Tags
	IconID      types.ID

	OSType ostype.ArchiveOSType

	DiskPlanID      types.ID
	Connection      types.EDiskConnection
	SourceDiskID    types.ID
	SourceArchiveID types.ID
	ServerID        types.ID
	SizeGB          int

	EditParameter *diskBuilder.EditRequest

	DistantFrom []types.ID
}

func (r *BuildRequest) Validate() error {
	return validate.Struct(r)
}

func (r *BuildRequest) toRequestParameter(caller sacloud.APICaller) (diskBuilder.Builder, error) {
	director := &diskBuilder.Director{
		OSType:          r.OSType,
		Name:            r.Name,
		SizeGB:          r.SizeGB,
		DistantFrom:     r.DistantFrom,
		PlanID:          r.DiskPlanID,
		Connection:      r.Connection,
		Description:     r.Description,
		Tags:            r.Tags,
		IconID:          r.IconID,
		DiskID:          r.ID,
		SourceDiskID:    r.SourceDiskID,
		SourceArchiveID: r.SourceArchiveID,
		EditParameter:   r.EditParameter,
		Client:          diskBuilder.NewBuildersAPIClient(caller),
	}
	return director.Builder(), nil
}

type BuildResult struct {
	Disk            *sacloud.Disk
	GeneratedSSHKey *sacloud.SSHKeyGenerated
}

// Build ディスクの作成を行います。作成後は利用可能になるまでWaitForReady()などで待機する必要があります。
//
// SourceDiskID,SourceArchiveIDの両方を指定しないことでブランクディスクの作成が可能です。
func (s *Service) Build(req *BuildRequest) (*BuildResult, error) {
	return s.BuildWithContext(context.Background(), req)
}

// BuildWithContext ディスクの作成を行います。作成後は利用可能になるまでWaitForReady()などで待機する必要があります。
//
// SourceDiskID,SourceArchiveIDの両方を指定しないことでブランクディスクの作成が可能です。
func (s *Service) BuildWithContext(ctx context.Context, req *BuildRequest) (*BuildResult, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	builder, err := req.toRequestParameter(s.caller)
	if err != nil {
		return nil, err
	}

	// create
	if req.ID.IsEmpty() {
		res, err := builder.Build(ctx, req.Zone, req.ServerID)
		if err != nil {
			return nil, err
		}
		diskOp := sacloud.NewDiskOp(s.caller)
		disk, err := diskOp.Read(ctx, req.Zone, res.DiskID)
		if err != nil {
			return nil, err
		}
		return &BuildResult{
			Disk:            disk,
			GeneratedSSHKey: res.GeneratedSSHKey,
		}, nil
	}

	// update
	res, err := builder.Update(ctx, req.Zone)
	if err != nil {
		return nil, err
	}
	return &BuildResult{Disk: res.Disk}, nil
}
