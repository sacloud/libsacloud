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

type UpdateServerRequest struct {
	Zone             string   `validate:"required" mapconv:"-"`
	ID               types.ID `validate:"required" mapconv:"-"`
	VirtualIPAddress string   `validate:"required,ipv4" mapconv:"-"`
	CurrentIPAddress string   `validate:"required,ipv4" mapconv:"-"`

	IPAddress    *string `validate:"omitempty,ipv4"`
	Enabled      *bool
	Protocol     *types.ELoadBalancerHealthCheckProtocol `validate:"omitempty,oneof=http https ping tcp" mapconv:"HealthCheck.Protocol"`
	Path         *string                                 `mapconv:"HealthCheck.Path"`
	ResponseCode *types.StringNumber                     `mapconv:"HealthCheck.ResponseCode"`
}

func (r *UpdateServerRequest) Validate() error {
	return validate.Struct(r)
}

func (s *Service) UpdateServer(req *UpdateServerRequest) (*sacloud.LoadBalancerServer, error) {
	return s.UpdateServerWithContext(context.Background(), req)
}

func (s *Service) UpdateServerWithContext(ctx context.Context, req *UpdateServerRequest) (*sacloud.LoadBalancerServer, error) {
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

	server := vip.Servers.FindAt(req.CurrentIPAddress)
	if server == nil {
		return nil, fmt.Errorf("not found: %s", req.CurrentIPAddress)
	}

	if err := mapconv.ConvertTo(req, server); err != nil {
		return nil, err
	}

	updated, err := client.UpdateSettings(ctx, req.Zone, req.ID, &sacloud.LoadBalancerUpdateSettingsRequest{
		VirtualIPAddresses: current.VirtualIPAddresses,
		SettingsHash:       current.SettingsHash,
	})
	if err != nil {
		return nil, err
	}
	updVIP := updated.VirtualIPAddresses.FindAt(req.VirtualIPAddress)
	if updVIP == nil {
		return nil, fmt.Errorf("not found: %s", req.VirtualIPAddress)
	}
	return updVIP.Servers.FindAt(server.IPAddress), nil
}
