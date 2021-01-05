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

package internet

import (
	"testing"

	"github.com/sacloud/libsacloud/v2/helper/cleanup"
	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/testutil"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

func TestBuilder_Build(t *testing.T) {
	var testZone = testutil.TestZone()

	testutil.RunCRUD(t, &testutil.CRUDTestCase{
		SetupAPICallerFunc: func() sacloud.APICaller {
			return testutil.SingletonAPICaller()
		},
		Parallel:          true,
		IgnoreStartupWait: true,
		Create: &testutil.CRUDTestFunc{
			Func: func(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
				builder := &Builder{
					Name:           testutil.ResourceName("internet-builder"),
					Description:    "description",
					Tags:           types.Tags{"tag1", "tag2"},
					NetworkMaskLen: 28,
					BandWidthMbps:  100,
					EnableIPv6:     true,
					Client:         NewAPIClient(caller),
				}
				return builder.Build(ctx, testZone)
			},
		},
		Read: &testutil.CRUDTestFunc{
			Func: func(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
				return sacloud.NewInternetOp(caller).Read(ctx, testZone, ctx.ID)
			},
			CheckFunc: func(t testutil.TestT, ctx *testutil.CRUDTestContext, value interface{}) error {
				internet := value.(*sacloud.Internet)
				return testutil.DoAsserts(
					testutil.AssertNotNilFunc(t, internet, "Internet"),
					testutil.AssertEqualFunc(t, 28, internet.NetworkMaskLen, "Internet.NetworkMaskLen"),
					testutil.AssertEqualFunc(t, 100, internet.BandWidthMbps, "Internet.BandWidthMbps"),
					testutil.AssertTrueFunc(t, len(internet.Switch.IPv6Nets) > 0, "Internet.Switch.IPv6Nets"),
				)
			},
		},
		Updates: []*testutil.CRUDTestFunc{
			{
				Func: func(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
					builder := &Builder{
						Name:           testutil.ResourceName("internet-builder"),
						Description:    "description",
						Tags:           types.Tags{"tag1", "tag2"},
						NetworkMaskLen: 28,
						BandWidthMbps:  500,
						EnableIPv6:     false,
						Client:         NewAPIClient(caller),
					}
					return builder.Update(ctx, testZone, ctx.ID)
				},
				CheckFunc: func(t testutil.TestT, ctx *testutil.CRUDTestContext, value interface{}) error {
					internet := value.(*sacloud.Internet)
					return testutil.DoAsserts(
						testutil.AssertNotNilFunc(t, internet, "Internet"),
						testutil.AssertEqualFunc(t, 28, internet.NetworkMaskLen, "Internet.NetworkMaskLen"),
						testutil.AssertEqualFunc(t, 500, internet.BandWidthMbps, "Internet.BandWidthMbps"),
						testutil.AssertTrueFunc(t, len(internet.Switch.IPv6Nets) == 0, "Internet.Switch.IPv6Nets"),
					)
				},
			},
			{
				Func: func(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
					internetOp := sacloud.NewInternetOp(caller)
					swOp := sacloud.NewSwitchOp(caller)

					internet, err := internetOp.Read(ctx, testZone, ctx.ID)
					if err != nil {
						return nil, err
					}
					sw, err := swOp.Read(ctx, testZone, internet.Switch.ID)
					if err != nil {
						return nil, err
					}

					_, err = internetOp.AddSubnet(ctx, testZone, ctx.ID, &sacloud.InternetAddSubnetRequest{
						NetworkMaskLen: 28,
						NextHop:        sw.Subnets[0].AssignedIPAddressMin,
					})
					return nil, err
				},
				CheckFunc: func(t testutil.TestT, ctx *testutil.CRUDTestContext, value interface{}) error {
					internet := value.(*sacloud.Internet)
					return testutil.DoAsserts(
						testutil.AssertNotNilFunc(t, internet, "Internet"),
						testutil.AssertTrueFunc(t, len(internet.Switch.Subnets) == 2, "Internet.Switch.Subnets"),
					)
				},
				SkipExtractID: true,
			},
		},
		Delete: &testutil.CRUDTestDeleteFunc{
			Func: func(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) error {
				return cleanup.DeleteInternet(ctx, sacloud.NewInternetOp(caller), testZone, ctx.ID)
			},
		},
	})
}
