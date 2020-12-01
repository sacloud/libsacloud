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
	"testing"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
	"github.com/stretchr/testify/require"
)

func TestVPCRouterService_convertCreateStandardRequest(t *testing.T) {
	cases := []struct {
		in     *CreateStandardRequest
		expect *ApplyRequest
	}{
		{
			in: &CreateStandardRequest{
				Zone:        "tk1a",
				Name:        "name",
				Description: "desc",
				Tags:        types.Tags{"tag1", "tag2"},
				IconID:      101,
				AdditionalNICSettings: []*AdditionalStandardNICSetting{
					{
						SwitchID:       103,
						IPAddress:      "192.168.1.101",
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
				NoWait:          true,
				BootAfterCreate: true,
			},
			expect: &ApplyRequest{
				Zone:        "tk1a",
				Name:        "name",
				Description: "desc",
				Tags:        types.Tags{"tag1", "tag2"},
				IconID:      101,
				PlanID:      types.VPCRouterPlans.Standard,
				NICSetting:  &StandardNICSetting{},
				AdditionalNICSettings: []AdditionalNICSettingHolder{
					&AdditionalStandardNICSetting{
						SwitchID:       103,
						IPAddress:      "192.168.1.101",
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
				NoWait:          true,
				BootAfterCreate: true,
			},
		},
	}

	for _, tc := range cases {
		require.EqualValues(t, tc.expect, tc.in.ApplyRequest())
	}
}
