// Copyright 2016-2021 The Libsacloud Authors
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

func TestVPCRouterService_convertCreateRequest(t *testing.T) {
	cases := []struct {
		in     *CreateRequest
		expect *ApplyRequest
	}{
		{
			in: &CreateRequest{
				Zone:        "tk1a",
				Name:        "name",
				Description: "desc",
				Tags:        types.Tags{"tag1", "tag2"},
				IconID:      101,
				PlanID:      types.VPCRouterPlans.Premium,
				NICSetting: &PremiumNICSetting{
					SwitchID:         102,
					IPAddresses:      []string{"192.168.0.101", "192.168.0.102"},
					VirtualIPAddress: "192.168.0.11",
					IPAliases:        []string{"192.168.0.12", "192.168.0.13"},
				},
				AdditionalNICSettings: []*AdditionalPremiumNICSetting{
					{
						SwitchID:         103,
						IPAddresses:      []string{"192.168.1.101", "192.168.1.102"},
						VirtualIPAddress: "192.168.1.11",
						NetworkMaskLen:   24,
						Index:            1,
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
				PlanID:      types.VPCRouterPlans.Premium,
				NICSetting: &PremiumNICSetting{
					SwitchID:         102,
					IPAddresses:      []string{"192.168.0.101", "192.168.0.102"},
					VirtualIPAddress: "192.168.0.11",
					IPAliases:        []string{"192.168.0.12", "192.168.0.13"},
				},
				AdditionalNICSettings: []AdditionalNICSettingHolder{
					&AdditionalPremiumNICSetting{
						SwitchID:         103,
						IPAddresses:      []string{"192.168.1.101", "192.168.1.102"},
						VirtualIPAddress: "192.168.1.11",
						NetworkMaskLen:   24,
						Index:            1,
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
