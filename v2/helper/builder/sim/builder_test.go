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

package sim

import (
	"os"
	"testing"

	"github.com/sacloud/libsacloud/v2/helper/query"

	"github.com/sacloud/libsacloud/v2/helper/cleanup"
	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/testutil"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

func TestBuilder_Build(t *testing.T) {
	testutil.PreCheckEnvsFunc("SAKURACLOUD_SIM_ICCID", "SAKURACLOUD_SIM_PASSCODE")(t)

	iccid := os.Getenv("SAKURACLOUD_SIM_ICCID")
	passcode := os.Getenv("SAKURACLOUD_SIM_PASSCODE")
	imei := "123456789012345"

	testutil.Run(t, &testutil.CRUDTestCase{
		SetupAPICallerFunc: func() sacloud.APICaller {
			return testutil.SingletonAPICaller()
		},
		Parallel:          true,
		IgnoreStartupWait: true,
		Create: &testutil.CRUDTestFunc{
			Func: func(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
				builder := &Builder{
					Name:        testutil.ResourceName("sim-builder"),
					Description: "description",
					Tags:        types.Tags{"tag1", "tag2"},
					ICCID:       iccid,
					PassCode:    passcode,
					Activate:    true,
					IMEI:        imei,
					Carrier: []*sacloud.SIMNetworkOperatorConfig{
						{
							Allow: true,
							Name:  types.SIMOperators.SoftBank.String(),
						},
					},
					Client: NewAPIClient(caller),
				}
				return builder.Build(ctx)
			},
		},
		Read: &testutil.CRUDTestFunc{
			Func: func(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
				simOp := sacloud.NewSIMOp(caller)
				return query.FindSIMByID(ctx, simOp, ctx.ID)
			},
			CheckFunc: func(t testutil.TestT, ctx *testutil.CRUDTestContext, value interface{}) error {
				sim := value.(*sacloud.SIM)
				return testutil.DoAsserts(
					testutil.AssertNotNilFunc(t, sim, "SIM"),
					testutil.AssertNotNilFunc(t, sim.Info, "SIM.Info"),
					testutil.AssertTrueFunc(t, sim.Info.Activated, "SIM.Info.Activated"),
					testutil.AssertTrueFunc(t, sim.Info.IMEILock, "SIM.Info.IMEILock"),
				)
			},
		},
		Delete: &testutil.CRUDTestDeleteFunc{
			Func: func(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) error {
				simOp := sacloud.NewSIMOp(caller)
				return cleanup.DeleteSIM(ctx, simOp, ctx.ID)
			},
		},
	})
}
