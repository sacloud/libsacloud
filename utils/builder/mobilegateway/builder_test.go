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

package mobilegateway

import (
	"testing"
	"time"

	"github.com/sacloud/libsacloud/v2/utils/builder"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/testutil"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

func getSetupOption() *builder.RetryableSetupParameter {
	if testutil.IsAccTest() {
		return nil
	}
	return &builder.RetryableSetupParameter{
		DeleteRetryInterval:       10 * time.Millisecond,
		ProvisioningRetryInterval: 10 * time.Millisecond,
		PollingInterval:           10 * time.Millisecond,
		NICUpdateWaitDuration:     10 * time.Millisecond,
	}
}

func TestBuilder_Build(t *testing.T) {
	var switchID types.ID
	var testZone = testutil.TestZone()

	testutil.Run(t, &testutil.CRUDTestCase{
		SetupAPICallerFunc: func() sacloud.APICaller {
			return testutil.SingletonAPICaller()
		},
		Parallel:          true,
		IgnoreStartupWait: true,
		Setup: func(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) error {
			swOp := sacloud.NewSwitchOp(caller)

			sw, err := swOp.Create(ctx, testZone, &sacloud.SwitchCreateRequest{
				Name: testutil.ResourceName("mobile-gateway-builder"),
			})
			if err != nil {
				return err
			}
			switchID = sw.ID
			return nil
		},
		Create: &testutil.CRUDTestFunc{
			Func: func(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
				builder := &Builder{
					Name:        testutil.ResourceName("mobile-gateway-builder"),
					Description: "description",
					Tags:        types.Tags{"tag1", "tag2"},
					PrivateInterface: &PrivateInterfaceSetting{
						SwitchID:       switchID,
						IPAddress:      "192.168.0.1",
						NetworkMaskLen: 24,
					},
					StaticRoutes: []*sacloud.MobileGatewayStaticRoute{
						{
							Prefix:  "192.168.1.0/24",
							NextHop: "192.168.0.1",
						},
						{
							Prefix:  "192.168.2.0/24",
							NextHop: "192.168.0.1",
						},
					},
					SIMs:                            nil,
					SIMRoutes:                       nil,
					InternetConnectionEnabled:       true,
					InterDeviceCommunicationEnabled: true,
					DNS: &sacloud.MobileGatewayDNSSetting{
						DNS1: "1.1.1.1",
						DNS2: "2.2.2.2",
					},
					TrafficConfig: &sacloud.MobileGatewayTrafficControl{
						TrafficQuotaInMB:     1024,
						BandWidthLimitInKbps: 128,
						EmailNotifyEnabled:   true,
						AutoTrafficShaping:   true,
					},
					SetupOptions: getSetupOption(),
					Client:       NewAPIClient(caller),
				}
				return builder.Build(ctx, testZone)
			},
		},
		Read: &testutil.CRUDTestFunc{
			Func: func(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
				mgwOp := sacloud.NewMobileGatewayOp(caller)
				return mgwOp.Read(ctx, testZone, ctx.ID)
			},
			CheckFunc: func(t testutil.TestT, ctx *testutil.CRUDTestContext, value interface{}) error {
				mgw := value.(*sacloud.MobileGateway)
				return testutil.DoAsserts(
					testutil.AssertNotNilFunc(t, mgw, "MobileGateway"),
					testutil.AssertNotNilFunc(t, mgw.Settings, "MobileGateway.Settings"),
					testutil.AssertLenFunc(t, mgw.Settings.Interfaces, 1, "MobileGateway.Settings.Interfaces"),
				)
			},
		},
		Delete: &testutil.CRUDTestDeleteFunc{
			Func: func(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) error {
				mgwOp := sacloud.NewMobileGatewayOp(caller)
				return mgwOp.Delete(ctx, testZone, ctx.ID)
			},
		},
		Cleanup: func(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) error {
			swOp := sacloud.NewSwitchOp(caller)
			return swOp.Delete(ctx, testZone, switchID)
		},
	})
}
