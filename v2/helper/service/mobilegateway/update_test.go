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

package mobilegateway

import (
	"context"
	"testing"

	mobileGatewayBuilder "github.com/sacloud/libsacloud/v2/helper/builder/mobilegateway"
	"github.com/sacloud/libsacloud/v2/helper/cleanup"
	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/pointer"
	"github.com/sacloud/libsacloud/v2/sacloud/testutil"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
	"github.com/stretchr/testify/require"
)

func TestMobileGatewayService_validate(t *testing.T) {
	cases := []struct {
		in          *UpdateRequest
		errorExists bool // 有無だけチェック
	}{
		{
			in:          &UpdateRequest{},
			errorExists: true,
		},
		{
			in: &UpdateRequest{
				Zone: "tk1a",
				ID:   1,
			},
			errorExists: false,
		},
		{
			in: &UpdateRequest{
				Zone: "tk1a",
				ID:   1,
				DNS: &DNSSettingUpdate{
					DNS1: pointer.NewString("8.8.8.8"),
					DNS2: nil,
				},
			},
			errorExists: true,
		},
		{
			in: &UpdateRequest{
				Zone: "tk1a",
				ID:   1,
				DNS: &DNSSettingUpdate{
					DNS1: pointer.NewString("8.8.8.8"),
					DNS2: pointer.NewString("8.8.4.4"),
				},
			},
			errorExists: false,
		},
	}

	for _, tc := range cases {
		err := tc.in.Validate()
		require.EqualValues(t, tc.errorExists, err != nil, "in: %#+v error: %s", tc.in, err)
	}
}

func TestMobileGatewayService_convertUpdateRequest(t *testing.T) {
	ctx := context.Background()
	name := testutil.ResourceName("mobile-gateway-service-create")
	zone := testutil.TestZone()
	caller := testutil.SingletonAPICaller()

	// setup
	swOp := sacloud.NewSwitchOp(caller)
	sw, err := swOp.Create(ctx, zone, &sacloud.SwitchCreateRequest{Name: name})
	if err != nil {
		t.Fatal(err)
	}

	builder := &mobileGatewayBuilder.Builder{
		Name:                            name,
		Description:                     "description",
		Tags:                            types.Tags{"tag1", "tag2"},
		InternetConnectionEnabled:       true,
		InterDeviceCommunicationEnabled: true,
		Client:                          mobileGatewayBuilder.NewAPIClient(caller),
		TrafficConfig: &sacloud.MobileGatewayTrafficControl{
			TrafficQuotaInMB: 1,
		},
	}
	mgw, err := builder.Build(ctx, zone)
	if err != nil {
		t.Fatal(err)
	}

	defer func() {
		cleanup.DeleteMobileGateway(ctx, sacloud.NewMobileGatewayOp(caller), sacloud.NewSIMOp(caller), zone, mgw.ID) // nolint
		swOp.Delete(ctx, zone, sw.ID)                                                                                // nolint
	}()

	// test
	cases := []struct {
		in     *UpdateRequest
		expect *ApplyRequest
	}{
		{
			in: &UpdateRequest{
				ID:   mgw.ID,
				Zone: zone,
				Name: pointer.NewString(name + "-upd"),
				PrivateInterface: &PrivateInterfaceSettingUpdate{
					SwitchID:       &sw.ID,
					IPAddress:      pointer.NewString("192.168.0.1"),
					NetworkMaskLen: pointer.NewInt(24),
				},
				StaticRoutes: &[]*sacloud.MobileGatewayStaticRoute{
					{
						Prefix:  "192.168.1.0/24",
						NextHop: "192.168.0.2",
					},
				},
				InternetConnectionEnabled:       pointer.NewBool(false),
				InterDeviceCommunicationEnabled: pointer.NewBool(false),
				DNS: &DNSSettingUpdate{
					DNS1: pointer.NewString("8.8.8.8"),
					DNS2: pointer.NewString("8.8.4.4"),
				},
				SIMs: nil,
				TrafficConfig: &TrafficConfigUpdate{
					BandWidthLimitInKbps: pointer.NewInt(128),
					EmailNotifyEnabled:   pointer.NewBool(true),
					AutoTrafficShaping:   pointer.NewBool(true),
				},
				NoWait: true,
			},
			expect: &ApplyRequest{
				ID:          mgw.ID,
				Zone:        zone,
				Name:        name + "-upd",
				Description: "description",
				Tags:        types.Tags{"tag1", "tag2"},
				PrivateInterface: &PrivateInterfaceSetting{
					SwitchID:       sw.ID,
					IPAddress:      "192.168.0.1",
					NetworkMaskLen: 24,
				},
				StaticRoutes: []*sacloud.MobileGatewayStaticRoute{
					{
						Prefix:  "192.168.1.0/24",
						NextHop: "192.168.0.2",
					},
				},
				InternetConnectionEnabled:       false,
				InterDeviceCommunicationEnabled: false,
				DNS: &DNSSetting{
					DNS1: "8.8.8.8",
					DNS2: "8.8.4.4",
				},
				SIMs: nil,
				TrafficConfig: &TrafficConfig{
					TrafficQuotaInMB:     1,
					BandWidthLimitInKbps: 128,
					EmailNotifyEnabled:   true,
					AutoTrafficShaping:   true,
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
