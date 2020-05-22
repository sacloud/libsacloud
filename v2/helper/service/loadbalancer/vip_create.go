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

	"github.com/sacloud/libsacloud/v2/sacloud/types"

	"github.com/sacloud/libsacloud/v2/helper/validate"
	"github.com/sacloud/libsacloud/v2/pkg/mapconv"
	"github.com/sacloud/libsacloud/v2/sacloud"
)

type CreateVirtualIPAddressRequest struct {
	Zone string   `validate:"required" mapconv:"-"`
	ID   types.ID `validate:"required"`

	VirtualIPAddress string `validate:"required,ipv4"`
	Port             int    `validate:"required,min=1,max=65535"`
	DelayLoop        int    `validate:"min=0,max=10000"`
	SorryServer      string `validate:"omitempty,ipv4"`
	Description      string `validate:"min=0,max=512"`
}

func (r *CreateVirtualIPAddressRequest) Validate() error {
	return validate.Struct(r)
}

func (s *Service) CreateVirtualIPAddress(req *CreateVirtualIPAddressRequest) (*sacloud.LoadBalancerVirtualIPAddress, error) {
	return s.CreateVirtualIPAddressWithContext(context.Background(), req)
}

func (s *Service) CreateVirtualIPAddressWithContext(ctx context.Context, req *CreateVirtualIPAddressRequest) (*sacloud.LoadBalancerVirtualIPAddress, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	client := sacloud.NewLoadBalancerOp(s.caller)
	current, err := client.Read(ctx, req.Zone, req.ID)
	if err != nil {
		return nil, fmt.Errorf("reading load balancer[%s] failed: %s", req.ID, err)
	}

	if current.VirtualIPAddresses.ExistAt(req.VirtualIPAddress) {
		return nil, fmt.Errorf("VirtualIPAddress[%s] already exists", req.VirtualIPAddress)
	}
	vip := &sacloud.LoadBalancerVirtualIPAddress{}
	if err := mapconv.ConvertTo(req, vip); err != nil {
		return nil, err
	}
	current.VirtualIPAddresses.Add(vip)

	updated, err := client.UpdateSettings(ctx, req.Zone, req.ID, &sacloud.LoadBalancerUpdateSettingsRequest{
		VirtualIPAddresses: current.VirtualIPAddresses,
		SettingsHash:       current.SettingsHash,
	})
	if err != nil {
		return nil, err
	}
	return updated.VirtualIPAddresses.FindAt(req.VirtualIPAddress), nil
}
