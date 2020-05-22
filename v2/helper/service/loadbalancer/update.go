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

package loadbalancer

import (
	"context"
	"fmt"

	"github.com/sacloud/libsacloud/v2/helper/validate"
	"github.com/sacloud/libsacloud/v2/pkg/mapconv"
	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

type UpdateRequest struct {
	Zone string   `validate:"required" mapconv:"-"`
	ID   types.ID `validate:"required" mapconv:"-"`

	Name        *string `validate:"omitempty,min=1"`
	Description *string `validate:"omitempty,min=1,max=512"`

	Tags         *types.Tags
	IconID       *types.ID
	SettingsHash string
}

func (r *UpdateRequest) Validate() error {
	return validate.Struct(r)
}

func (r *UpdateRequest) toRequestParameter(current *sacloud.LoadBalancer) (*sacloud.LoadBalancerUpdateRequest, error) {
	req := &sacloud.LoadBalancerUpdateRequest{}
	if err := mapconv.ConvertFrom(current, req); err != nil {
		return nil, err
	}
	if err := mapconv.ConvertFrom(r, req); err != nil {
		return nil, err
	}
	return req, nil
}

func (s *Service) Update(req *UpdateRequest) (*sacloud.LoadBalancer, error) {
	return s.UpdateWithContext(context.Background(), req)
}

func (s *Service) UpdateWithContext(ctx context.Context, req *UpdateRequest) (*sacloud.LoadBalancer, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	client := sacloud.NewLoadBalancerOp(s.caller)
	current, err := client.Read(ctx, req.Zone, req.ID)
	if err != nil {
		return nil, fmt.Errorf("reading load balancer[%s] failed: %s", req.ID, err)
	}

	params, err := req.toRequestParameter(current)
	if err != nil {
		return nil, fmt.Errorf("processing request parameter failed: %s", err)
	}
	return client.Update(ctx, req.Zone, req.ID, params)
}
