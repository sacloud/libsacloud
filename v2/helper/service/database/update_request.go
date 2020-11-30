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

package database

import (
	"context"

	"github.com/sacloud/libsacloud/v2/helper/service"

	"github.com/sacloud/libsacloud/v2/helper/validate"
	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

type UpdateRequest struct {
	Zone string   `validate:"required"`
	ID   types.ID `validate:"required"`

	Name        *string     `request:",omitempty" validate:"omitempty,min=1"`
	Description *string     `request:",omitempty" validate:"omitempty,min=1,max=512"`
	Tags        *types.Tags `request:",omitempty"`
	IconID      *types.ID   `request:",omitempty"`

	SourceNetwork         *[]string                   `request:",omitempty" validate:"omitempty,dive,cidrv4"`
	EnableReplication     *bool                       `request:",omitempty"`
	ReplicaUserPassword   *string                     `request:",omitempty" validate:"required_with=EnableReplication"`
	EnableWebUI           *bool                       `request:",omitempty"`
	EnableBackup          *bool                       `request:",omitempty"`
	BackupWeekdays        *[]types.EBackupSpanWeekday `request:",omitempty" validate:"required_with=EnableBackup,max=7"`
	BackupStartTimeHour   *int                        `request:",omitempty" validate:"omitempty,min=0,max=23"`
	BackupStartTimeMinute *int                        `request:",omitempty" validate:"omitempty,oneof=0 15 30 45"`

	SettingsHash string
	NoWait       bool
}

func (req *UpdateRequest) Validate() error {
	return validate.Struct(req)
}

func (req *UpdateRequest) Builder(ctx context.Context, caller sacloud.APICaller) (*Builder, error) {
	builder, err := BuilderFromResource(ctx, caller, req.Zone, req.ID)
	if err != nil {
		return nil, err
	}
	if err := service.RequestConvertTo(req, builder); err != nil {
		return nil, err
	}
	return builder, nil
}
