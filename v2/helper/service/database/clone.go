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

type CloneRequest struct {
	Zone     string   `validate:"required" mapconv:"-"`
	SourceID types.ID `validate:"required"`
	Name     string   `validate:"required"`

	/* common */
	Description string `validate:"min=0,max=512"`
	Tags        types.Tags
	IconID      types.ID
	PlanID      types.ID // this will be copied from the source when empty

	/* network */
	SwitchID       types.ID // this will be copied from the source when empty
	IPAddresses    []string `validate:"required,min=1,max=2,dive,ipv4"`
	NetworkMaskLen int      `validate:"omitempty,min=1,max=32"`    // this will be copied from the source when empty
	DefaultRoute   string   `validate:"omitempty,ipv4"`            // this will be copied from the source when empty
	Port           int      `validate:"omitempty,min=1,max=65535"` // this will be copied from the source when empty
	SourceNetworks []string `validate:"dive,cidrv4"`

	/* WebUI */
	EnableWebUI bool

	/* Backup */
	EnableBackup          bool
	BackupWeekdays        []types.EBackupSpanWeekday `validate:"required_with=EnableBackup,max=7"`
	BackupStartTimeHour   int                        `validate:"omitempty,min=0,max=23"`
	BackupStartTimeMinute int                        `validate:"omitempty,oneof=0 15 30 45"`
}

func (r *CloneRequest) Validate() error {
	return validate.Struct(r)
}

func (r *CloneRequest) toRequestParameter(caller sacloud.APICaller, source *sacloud.Database) (*databaseBuilder.Builder, error) {
	builder := &databaseBuilder.Builder{
		PlanID:         source.PlanID,
		SwitchID:       source.SwitchID,
		IPAddresses:    r.IPAddresses,
		NetworkMaskLen: source.NetworkMaskLen,
		DefaultRoute:   source.DefaultRoute,
		SourceID:       source.ID,
		Conf:           source.Conf,
		CommonSetting: &sacloud.DatabaseSettingCommon{
			WebUI:           types.ToWebUI(r.EnableWebUI),
			ServicePort:     source.CommonSetting.ServicePort,
			SourceNetwork:   r.SourceNetworks,
			DefaultUser:     source.CommonSetting.DefaultUser,
			UserPassword:    source.CommonSetting.UserPassword,
			ReplicaUser:     source.CommonSetting.ReplicaUser,
			ReplicaPassword: source.CommonSetting.ReplicaPassword,
		},
		ReplicationSetting: source.ReplicationSetting,
		Name:               r.Name,
		Description:        r.Description,
		Tags:               r.Tags,
		IconID:             r.IconID,
		Client:             databaseBuilder.NewAPIClient(caller),
	}

	if !r.PlanID.IsEmpty() {
		builder.PlanID = r.PlanID
	}
	if !r.SwitchID.IsEmpty() {
		builder.SwitchID = r.SwitchID
	}
	if r.NetworkMaskLen > 0 {
		builder.NetworkMaskLen = r.NetworkMaskLen
	}
	if r.Port > 0 {
		builder.CommonSetting.ServicePort = r.Port
	}

	if r.EnableBackup {
		builder.BackupSetting = &sacloud.DatabaseSettingBackup{
			Time:      fmt.Sprintf("%02d:%02d", r.BackupStartTimeHour, r.BackupStartTimeMinute),
			DayOfWeek: r.BackupWeekdays,
		}
	}

	return builder, nil
}

func (s *Service) Clone(req *CloneRequest) (*sacloud.Database, error) {
	return s.CloneWithContext(context.Background(), req)
}

func (s *Service) CloneWithContext(ctx context.Context, req *CloneRequest) (*sacloud.Database, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	source, err := sacloud.NewDatabaseOp(s.caller).Read(ctx, req.Zone, req.SourceID)
	if err != nil {
		return nil, err
	}

	builder, err := req.toRequestParameter(s.caller, source)
	if err != nil {
		return nil, err
	}

	return builder.Build(ctx, req.Zone)
}
