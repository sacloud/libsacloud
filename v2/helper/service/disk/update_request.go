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

	diskBuilder "github.com/sacloud/libsacloud/v2/helper/builder/disk"
	"github.com/sacloud/libsacloud/v2/helper/service"
	"github.com/sacloud/libsacloud/v2/helper/validate"
	"github.com/sacloud/libsacloud/v2/pkg/size"
	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

type UpdateRequest struct {
	Zone string   `request:"-" validate:"required"`
	ID   types.ID `request:"-" validate:"required"`

	Name          *string                `request:",omitempty" validate:"omitempty,min=1"`
	Description   *string                `request:",omitempty" validate:"omitempty,min=1,max=512"`
	Tags          *types.Tags            `request:",omitempty"`
	IconID        *types.ID              `request:",omitempty"`
	Connection    *types.EDiskConnection `request:",omitempty"`
	EditParameter *EditParameter         `request:"-"`

	NoWait bool
}

func (req *UpdateRequest) Validate() error {
	return validate.Struct(req)
}

func (req *UpdateRequest) Builder(ctx context.Context, caller sacloud.APICaller) (diskBuilder.Builder, error) {
	disk, err := sacloud.NewDiskOp(caller).Read(ctx, req.Zone, req.ID)
	if err != nil {
		return nil, err
	}

	editParameter := &diskBuilder.EditRequest{}
	if req.EditParameter != nil {
		if err := service.RequestConvertTo(req.EditParameter, editParameter); err != nil {
			return nil, err
		}
	}

	director := &diskBuilder.Director{
		DiskID:        req.ID,
		Name:          disk.Name,
		SizeGB:        size.MiBToGiB(disk.SizeMB),
		PlanID:        disk.DiskPlanID,
		Connection:    disk.Connection,
		Description:   disk.Description,
		Tags:          disk.Tags,
		IconID:        disk.IconID,
		EditParameter: editParameter,
		NoWait:        req.NoWait,
		Client:        diskBuilder.NewBuildersAPIClient(caller),
	}
	builder := director.Builder()

	if err := service.RequestConvertTo(req, builder); err != nil {
		return nil, err
	}
	return builder, err
}
