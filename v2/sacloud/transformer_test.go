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

package sacloud

import (
	"testing"

	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

func TestTransformer_transformLoadBalancerCreateArgs(t *testing.T) {
	op := &LoadBalancerOp{}

	ret, err := op.transformCreateArgs(&LoadBalancerCreateRequest{
		VirtualIPAddresses: LoadBalancerVirtualIPAddresses{
			{
				VirtualIPAddress: "192.168.0.1",
				Servers: LoadBalancerServers{
					{
						IPAddress: "192.168.0.11",
						Port:      80,
						Enabled:   true,
						HealthCheck: &LoadBalancerServerHealthCheck{
							Protocol:     types.LoadBalancerHealthCheckProtocols.HTTP,
							Path:         "/",
							ResponseCode: 200,
						},
					},
				},
			},
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	v := ret.Appliance.Settings.LoadBalancer[0].Servers[0].HealthCheck.Status
	if v != types.StringNumber(200) {
		t.Fatal("unexpected value:", v)
	}
}
