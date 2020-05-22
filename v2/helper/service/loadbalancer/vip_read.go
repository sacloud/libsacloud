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
	"github.com/sacloud/libsacloud/v2/sacloud"
)

type ReadVirtualIPAddressRequest struct {
	Zone string   `validate:"required" mapconv:"-"`
	ID   types.ID `validate:"required"`

	VirtualIPAddress string `validate:"required,ipv4"`
}

func (r *ReadVirtualIPAddressRequest) Validate() error {
	return validate.Struct(r)
}

func (s *Service) ReadVirtualIPAddress(req *ReadVirtualIPAddressRequest) (*sacloud.LoadBalancerVirtualIPAddress, error) {
	return s.ReadVirtualIPAddressWithContext(context.Background(), req)
}

func (s *Service) ReadVirtualIPAddressWithContext(ctx context.Context, req *ReadVirtualIPAddressRequest) (*sacloud.LoadBalancerVirtualIPAddress, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	client := sacloud.NewLoadBalancerOp(s.caller)
	current, err := client.Read(ctx, req.Zone, req.ID)
	if err != nil {
		return nil, fmt.Errorf("reading load balancer[%s] failed: %s", req.ID, err)
	}
	vip := current.VirtualIPAddresses.FindAt(req.VirtualIPAddress)

	if vip == nil {
		return nil, fmt.Errorf("not found: %s", req.VirtualIPAddress)
	}
	return vip, nil
}
