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

type DeleteServerRequest struct {
	Zone             string   `validate:"required" mapconv:"-"`
	ID               types.ID `validate:"required" mapconv:"-"`
	VirtualIPAddress string   `validate:"required,ipv4" mapconv:"-"`
	IPAddress        string   `validate:"required,ipv4"`
}

func (r *DeleteServerRequest) Validate() error {
	return validate.Struct(r)
}

func (s *Service) DeleteServer(req *DeleteServerRequest) error {
	return s.DeleteServerWithContext(context.Background(), req)
}

func (s *Service) DeleteServerWithContext(ctx context.Context, req *DeleteServerRequest) error {
	if err := req.Validate(); err != nil {
		return err
	}

	client := sacloud.NewLoadBalancerOp(s.caller)
	current, err := client.Read(ctx, req.Zone, req.ID)
	if err != nil {
		return fmt.Errorf("reading load balancer[%s] failed: %s", req.ID, err)
	}

	vip := current.VirtualIPAddresses.FindAt(req.VirtualIPAddress)
	if vip == nil {
		return fmt.Errorf("not found: %s", req.VirtualIPAddress)
	}
	if !vip.Servers.ExistAt(req.IPAddress) {
		return fmt.Errorf("not found: %s", req.IPAddress)
	}

	vip.Servers.DeleteAt(req.IPAddress)

	_, err = client.UpdateSettings(ctx, req.Zone, req.ID, &sacloud.LoadBalancerUpdateSettingsRequest{
		VirtualIPAddresses: current.VirtualIPAddresses,
		SettingsHash:       current.SettingsHash,
	})
	return err
}
