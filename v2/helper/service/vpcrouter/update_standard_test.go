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

package vpcrouter

import (
	"context"
	"testing"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/pointer"
	"github.com/sacloud/libsacloud/v2/sacloud/testutil"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
	"github.com/stretchr/testify/require"
)

func TestVPCRouterService_convertUpdateStandardRequest(t *testing.T) {
	ctx := context.Background()
	caller := testutil.SingletonAPICaller()
	zone := testutil.TestZone()
	name := testutil.ResourceName("vpc-router-service-update")

	// setup
	swOp := sacloud.NewSwitchOp(caller)
	additionalSwitch, err := swOp.Create(ctx, zone, &sacloud.SwitchCreateRequest{Name: name})
	if err != nil {
		t.Fatal(err)
	}

	vpcRouter, err := New(caller).CreateStandardWithContext(ctx, &CreateStandardRequest{
		Zone:        zone,
		Name:        name,
		Description: "desc",
		Tags:        types.Tags{"tag1", "tag2"},
		AdditionalNICSettings: []*AdditionalStandardNICSetting{
			{
				SwitchID:       additionalSwitch.ID,
				IPAddress:      "192.168.0.101",
				NetworkMaskLen: 24,
				Index:          1,
			},
		},
		RouterSetting: &RouterSetting{
			VRID:                      1,
			InternetConnectionEnabled: true,
			L2TPIPsecServer: &sacloud.VPCRouterL2TPIPsecServer{
				RangeStart:      "192.168.0.250",
				RangeStop:       "192.168.0.254",
				PreSharedSecret: "presharedsecret",
			},
			RemoteAccessUsers: []*sacloud.VPCRouterRemoteAccessUser{
				{
					UserName: "username",
					Password: "password",
				},
			},
		},
		NoWait: false,
	})

	if err != nil {
		t.Fatal(err)
	}

	defer func() {
		sacloud.NewVPCRouterOp(caller).Delete(ctx, zone, vpcRouter.ID) // nolint
		swOp.Delete(ctx, zone, additionalSwitch.ID)                    // nolint
	}()

	// test
	cases := []struct {
		in     *UpdateStandardRequest
		expect *ApplyRequest
	}{
		{
			in: &UpdateStandardRequest{
				ID:     vpcRouter.ID,
				Zone:   zone,
				Name:   pointer.NewString(name + "-upd"),
				NoWait: true,
			},
			expect: &ApplyRequest{
				ID:          vpcRouter.ID,
				Zone:        zone,
				Name:        name + "-upd",
				Description: "desc",
				Tags:        types.Tags{"tag1", "tag2"},
				PlanID:      types.VPCRouterPlans.Standard,
				NICSetting:  &StandardNICSetting{},
				AdditionalNICSettings: []AdditionalNICSettingHolder{
					&AdditionalStandardNICSetting{
						SwitchID:       additionalSwitch.ID,
						IPAddress:      "192.168.0.101",
						NetworkMaskLen: 24,
						Index:          1,
					},
				},
				RouterSetting: &RouterSetting{
					VRID:                      1,
					InternetConnectionEnabled: true,
					L2TPIPsecServer: &sacloud.VPCRouterL2TPIPsecServer{
						RangeStart:      "192.168.0.250",
						RangeStop:       "192.168.0.254",
						PreSharedSecret: "presharedsecret",
					},
					RemoteAccessUsers: []*sacloud.VPCRouterRemoteAccessUser{
						{
							UserName: "username",
							Password: "password",
						},
					},
				},
				NoWait: true,
			},
		},
	}

	for _, tc := range cases {
		req, err := tc.in.ApplyRequest(ctx, caller)
		require.NoError(t, err)
		require.EqualValues(t, tc.expect, req)
	}
}
