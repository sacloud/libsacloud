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

	databaseBuilder "github.com/sacloud/libsacloud/v2/helper/builder/database"
	"github.com/sacloud/libsacloud/v2/helper/validate"
	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

type CreateReplicaRequest struct {
	Zone     string   `validate:"required" mapconv:"-"`
	MasterID types.ID `validate:"required"`
	Name     string   `validate:"required"`

	/* common */
	Description string `validate:"min=0,max=512"`
	Tags        types.Tags
	IconID      types.ID

	/* network */
	SwitchID       types.ID // this will be copied from the source when empty
	IPAddresses    []string `validate:"required,min=1,max=2,dive,ipv4"`
	NetworkMaskLen int      `validate:"omitempty,min=1,max=32"` // this will be copied from the source when empty
	DefaultRoute   string   `validate:"omitempty,ipv4"`         // this will be copied from the source when empty
	SourceNetworks []string `validate:"dive,cidrv4"`
}

func (r *CreateReplicaRequest) Validate() error {
	return validate.Struct(r)
}

func (r *CreateReplicaRequest) toRequestParameter(caller sacloud.APICaller, source *sacloud.Database) (*databaseBuilder.Builder, error) {
	builder := &databaseBuilder.Builder{
		PlanID:         types.SlaveDatabasePlanID(source.PlanID),
		SwitchID:       source.SwitchID,
		IPAddresses:    r.IPAddresses,
		NetworkMaskLen: source.NetworkMaskLen,
		DefaultRoute:   source.DefaultRoute,
		Conf:           source.Conf,
		CommonSetting: &sacloud.DatabaseSettingCommon{
			ServicePort:   source.CommonSetting.ServicePort,
			SourceNetwork: r.SourceNetworks,
		},
		ReplicationSetting: &sacloud.DatabaseReplicationSetting{
			Model:       types.DatabaseReplicationModels.AsyncReplica,
			IPAddress:   source.IPAddresses[0],
			Port:        source.CommonSetting.ServicePort,
			User:        source.CommonSetting.ReplicaUser,
			Password:    source.CommonSetting.ReplicaPassword,
			ApplianceID: source.ID,
		},
		Name:        r.Name,
		Description: r.Description,
		Tags:        r.Tags,
		IconID:      r.IconID,
		Client:      databaseBuilder.NewAPIClient(caller),
	}

	if r.SwitchID.IsEmpty() {
		builder.SwitchID = r.SwitchID
	}
	if r.NetworkMaskLen > 0 {
		builder.NetworkMaskLen = r.NetworkMaskLen
	}

	return builder, nil
}

func (s *Service) CreateReplica(req *CreateReplicaRequest) (*sacloud.Database, error) {
	return s.CreateReplicaWithContext(context.Background(), req)
}

func (s *Service) CreateReplicaWithContext(ctx context.Context, req *CreateReplicaRequest) (*sacloud.Database, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	source, err := sacloud.NewDatabaseOp(s.caller).Read(ctx, req.Zone, req.MasterID)
	if err != nil {
		return nil, err
	}

	builder, err := req.toRequestParameter(s.caller, source)
	if err != nil {
		return nil, err
	}

	return builder.Build(ctx, req.Zone)
}
