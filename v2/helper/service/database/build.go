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
	"fmt"

	databaseBuilder "github.com/sacloud/libsacloud/v2/helper/builder/database"
	"github.com/sacloud/libsacloud/v2/helper/validate"
	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

type BuildRequest struct {
	Zone string `validate:"required" mapconv:"-"`
	ID   types.ID
	Name string `validate:"required"`

	/* common */
	Description string `validate:"min=0,max=512"`
	Tags        types.Tags
	IconID      types.ID
	PlanID      types.ID `validate:"required"`

	/* network */
	SwitchID       types.ID `validate:"required"`
	IPAddresses    []string `validate:"required,min=1,max=2,dive,ipv4"`
	NetworkMaskLen int      `validate:"required,min=1,max=32"`
	DefaultRoute   string   `validate:"omitempty,ipv4"`
	Port           int      `validate:"omitempty,min=1,max=65535"`
	SourceNetworks []string `validate:"dive,cidrv4"`

	/* RDBMS */
	DatabaseType string `validate:"required,oneof=mariadb postgresql"`
	Username     string `validate:"required"`
	Password     string `validate:"required"`

	/* Replication */
	EnableReplication   bool
	ReplicaUserPassword string `validate:"required_with=EnableReplication"`

	/* WebUI */
	EnableWebUI bool

	/* Backup */
	EnableBackup          bool
	BackupWeekdays        []types.EBackupSpanWeekday `validate:"required_with=EnableBackup,max=7"`
	BackupStartTimeHour   int                        `validate:"omitempty,min=0,max=23"`
	BackupStartTimeMinute int                        `validate:"omitempty,oneof=0 15 30 45"`
}

func (r *BuildRequest) Validate() error {
	return validate.Struct(r)
}

func (r *BuildRequest) toRequestParameter(caller sacloud.APICaller) (*databaseBuilder.Builder, error) {
	builder := &databaseBuilder.Builder{
		PlanID:         r.PlanID,
		SwitchID:       r.SwitchID,
		IPAddresses:    r.IPAddresses,
		NetworkMaskLen: r.NetworkMaskLen,
		DefaultRoute:   r.DefaultRoute,
		Conf: &sacloud.DatabaseRemarkDBConfCommon{
			DatabaseName:     types.RDBMSVersions[types.RDBMSTypesPostgreSQL].Name,
			DatabaseVersion:  types.RDBMSVersions[types.RDBMSTypesPostgreSQL].Version,
			DatabaseRevision: types.RDBMSVersions[types.RDBMSTypesPostgreSQL].Revision,
		},
		CommonSetting: &sacloud.DatabaseSettingCommon{
			WebUI:           types.ToWebUI(r.EnableWebUI),
			ServicePort:     r.Port,
			SourceNetwork:   r.SourceNetworks,
			DefaultUser:     r.Username,
			UserPassword:    r.Password,
			ReplicaUser:     "",
			ReplicaPassword: r.ReplicaUserPassword,
		},
		Name:        r.Name,
		Description: r.Description,
		Tags:        r.Tags,
		IconID:      r.IconID,
		Client:      databaseBuilder.NewAPIClient(caller),
	}

	if r.EnableBackup {
		builder.BackupSetting = &sacloud.DatabaseSettingBackup{
			Time:      fmt.Sprintf("%02d:%02d", r.BackupStartTimeHour, r.BackupStartTimeMinute),
			DayOfWeek: r.BackupWeekdays,
		}
	}
	if r.EnableReplication {
		builder.ReplicationSetting = &sacloud.DatabaseReplicationSetting{
			Model: types.DatabaseReplicationModels.MasterSlave,
		}
	}
	return builder, nil
}

func (s *Service) Build(req *BuildRequest) (*sacloud.Database, error) {
	return s.BuildWithContext(context.Background(), req)
}

func (s *Service) BuildWithContext(ctx context.Context, req *BuildRequest) (*sacloud.Database, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	builder, err := req.toRequestParameter(s.caller)
	if err != nil {
		return nil, err
	}

	if req.ID.IsEmpty() {
		return builder.Build(ctx, req.Zone)
	}
	return builder.Update(ctx, req.Zone, req.ID)
}
