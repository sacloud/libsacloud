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
	"testing"

	"github.com/sacloud/libsacloud/v2/helper/wait"
	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/pointer"
	"github.com/sacloud/libsacloud/v2/sacloud/testutil"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
	"github.com/stretchr/testify/require"
)

func TestLoadBalancerService_convertUpdateRequest(t *testing.T) {
	caller := testutil.SingletonAPICaller()
	ctx := context.Background()
	name := testutil.ResourceName("load-balancer-service")
	zone := testutil.TestZone()

	// setup
	swOp := sacloud.NewSwitchOp(caller)
	sw, err := swOp.Create(ctx, zone, &sacloud.SwitchCreateRequest{Name: name})
	if err != nil {
		t.Fatal(err)
	}

	current, err := New(caller).CreateWithContext(ctx, &CreateRequest{
		Zone:           zone,
		Name:           name,
		Description:    "desc",
		Tags:           types.Tags{"tag1", "tag2"},
		SwitchID:       sw.ID,
		PlanID:         types.LoadBalancerPlans.Standard,
		VRID:           10,
		IPAddresses:    []string{"192.168.0.101", "192.168.0.102"},
		NetworkMaskLen: 24,
		DefaultRoute:   "192.168.0.1",
		VirtualIPAddresses: []*sacloud.LoadBalancerVirtualIPAddress{
			{
				VirtualIPAddress: "192.168.0.201",
				Port:             80,
				DelayLoop:        10,
				SorryServer:      "192.168.0.99",
				Description:      "desc",
				Servers: []*sacloud.LoadBalancerServer{
					{
						IPAddress: "192.168.0.202",
						Port:      80,
						Enabled:   true,
						HealthCheck: &sacloud.LoadBalancerServerHealthCheck{
							Protocol:     types.LoadBalancerHealthCheckProtocols.HTTP,
							Path:         "/",
							ResponseCode: 200,
						},
					},
					{
						IPAddress: "192.168.0.203",
						Port:      80,
						Enabled:   true,
						HealthCheck: &sacloud.LoadBalancerServerHealthCheck{
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

	defer func() {
		lbOp := sacloud.NewLoadBalancerOp(caller)
		lbOp.Shutdown(ctx, zone, current.ID, &sacloud.ShutdownOption{Force: true}) // nolint
		wait.UntilLoadBalancerIsDown(ctx, lbOp, zone, current.ID)                  // nolint
		lbOp.Delete(ctx, zone, current.ID)                                         // nolint
		swOp.Delete(ctx, zone, sw.ID)                                              // nolint
	}()

	// test
	cases := []struct {
		in     *UpdateRequest
		expect *ApplyRequest
	}{
		{
			in: &UpdateRequest{
				Zone: zone,
				ID:   current.ID,
				Name: pointer.NewString(name + "-upd"),
				VirtualIPAddresses: &sacloud.LoadBalancerVirtualIPAddresses{
					{
						VirtualIPAddress: "192.168.0.202",
						Port:             80,
						DelayLoop:        10,
						SorryServer:      "192.168.0.99",
						Description:      "desc",
						Servers: []*sacloud.LoadBalancerServer{
							{
								IPAddress: "192.168.0.202",
								Port:      80,
								Enabled:   true,
								HealthCheck: &sacloud.LoadBalancerServerHealthCheck{
									Protocol:     types.LoadBalancerHealthCheckProtocols.HTTP,
									Path:         "/",
									ResponseCode: 200,
								},
							},
							{
								IPAddress: "192.168.0.203",
								Port:      80,
								Enabled:   true,
								HealthCheck: &sacloud.LoadBalancerServerHealthCheck{
									Protocol:     types.LoadBalancerHealthCheckProtocols.HTTP,
									Path:         "/",
									ResponseCode: 200,
								},
							},
						},
					},
				},
				SettingsHash: "aaaaa",
				NoWait:       true,
			},
			expect: &ApplyRequest{
				ID:             current.ID,
				Zone:           zone,
				Name:           name + "-upd",
				Description:    current.Description,
				Tags:           current.Tags,
				IconID:         current.IconID,
				SwitchID:       current.SwitchID,
				PlanID:         current.PlanID,
				VRID:           current.VRID,
				IPAddresses:    current.IPAddresses,
				NetworkMaskLen: current.NetworkMaskLen,
				DefaultRoute:   current.DefaultRoute,
				VirtualIPAddresses: sacloud.LoadBalancerVirtualIPAddresses{
					{
						VirtualIPAddress: "192.168.0.202",
						Port:             80,
						DelayLoop:        10,
						SorryServer:      "192.168.0.99",
						Description:      "desc",
						Servers: []*sacloud.LoadBalancerServer{
							{
								IPAddress: "192.168.0.202",
								Port:      80,
								Enabled:   true,
								HealthCheck: &sacloud.LoadBalancerServerHealthCheck{
									Protocol:     types.LoadBalancerHealthCheckProtocols.HTTP,
									Path:         "/",
									ResponseCode: 200,
								},
							},
							{
								IPAddress: "192.168.0.203",
								Port:      80,
								Enabled:   true,
								HealthCheck: &sacloud.LoadBalancerServerHealthCheck{
									Protocol:     types.LoadBalancerHealthCheckProtocols.HTTP,
									Path:         "/",
									ResponseCode: 200,
								},
							},
						},
					},
				},
				SettingsHash: "aaaaa",
				NoWait:       true,
			},
		},
	}

	for _, tc := range cases {
		req, err := tc.in.ApplyRequest(ctx, caller)
		require.NoError(t, err)
		require.EqualValues(t, tc.expect, req)
	}
}
