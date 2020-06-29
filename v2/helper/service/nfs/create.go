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

package nfs

import (
	"context"

	"github.com/sacloud/libsacloud/v2/sacloud/types"

	"github.com/sacloud/libsacloud/v2/helper/validate"
	"github.com/sacloud/libsacloud/v2/pkg/mapconv"
	"github.com/sacloud/libsacloud/v2/sacloud"
)

type CreateRequest struct {
	Zone        string `validate:"required" mapconv:"-"`
	Name        string `validate:"required"`
	Description string `validate:"min=0,max=512"`
	Tags        types.Tags
	IconID      types.ID `mapconv:"Icon.ID"`

	SwitchID       types.ID
	PlanID         types.ID
	IPAddresses    []string `validate:"required,min=1,max=2,dive,ipv4"`
	NetworkMaskLen int      `validate:"required,min=1,max=32"`
	DefaultRoute   string   `validate:"omitempty,ipv4"`
}

func (r *CreateRequest) Validate() error {
	return validate.Struct(r)
}

func (s *Service) Create(req *CreateRequest) (*sacloud.NFS, error) {
	return s.CreateWithContext(context.Background(), req)
}

func (s *Service) CreateWithContext(ctx context.Context, req *CreateRequest) (*sacloud.NFS, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	params := &sacloud.NFSCreateRequest{}
	if err := mapconv.ConvertTo(req, params); err != nil {
		return nil, err
	}

	client := sacloud.NewNFSOp(s.caller)
	return client.Create(ctx, req.Zone, params)
}
