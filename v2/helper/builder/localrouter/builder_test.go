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

package localrouter

import (
	"testing"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/testutil"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

func TestLocalRouterBuilder_Build(t *testing.T) {
	var testZone = testutil.TestZone()
	var peerLocalRouter *sacloud.LocalRouter
	var sw *sacloud.Switch

	testutil.RunCRUD(t, &testutil.CRUDTestCase{
		SetupAPICallerFunc: func() sacloud.APICaller {
			return testutil.SingletonAPICaller()
		},
		Parallel:          true,
		IgnoreStartupWait: true,
		Setup: func(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) error {
			lrOp := sacloud.NewLocalRouterOp(caller)
			lr, err := lrOp.Create(ctx, &sacloud.LocalRouterCreateRequest{
				Name: testutil.ResourceName("local-router-builder"),
			})
			if err != nil {
				return err
			}
			peerLocalRouter = lr

			swOp := sacloud.NewSwitchOp(caller)
			sw, err = swOp.Create(ctx, testZone, &sacloud.SwitchCreateRequest{
				Name: testutil.ResourceName("local-router-builder"),
			})
			return err
		},
		Create: &testutil.CRUDTestFunc{
			Func: func(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
				builder := &Builder{
					Name:        testutil.ResourceName("local-router-builder"),
					Description: "description",
					Tags:        types.Tags{"tag1", "tag2"},
					Switch: &sacloud.LocalRouterSwitch{
						Code:     sw.ID.String(),
						Category: "cloud",
						ZoneID:   testZone,
					},
					Interface: &sacloud.LocalRouterInterface{
						VirtualIPAddress: "192.168.0.1",
						IPAddress:        []string{"192.168.0.11", "192.168.0.12"},
						NetworkMaskLen:   24,
						VRID:             101,
					},
					Peers: []*sacloud.LocalRouterPeer{
						{
							ID:        peerLocalRouter.ID,
							SecretKey: peerLocalRouter.SecretKeys[0],
							Enabled:   true,
						},
					},
					StaticRoutes: []*sacloud.LocalRouterStaticRoute{
						{
							Prefix:  "192.168.1.0/24",
							NextHop: "192.168.0.101",
						},
					},
					Client: NewAPIClient(caller),
				}
				return builder.Build(ctx)
			},
		},
		Read: &testutil.CRUDTestFunc{
			Func: func(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
				return sacloud.NewLocalRouterOp(caller).Read(ctx, ctx.ID)
			},
			CheckFunc: func(t testutil.TestT, ctx *testutil.CRUDTestContext, value interface{}) error {
				lr := value.(*sacloud.LocalRouter)
				return testutil.DoAsserts(
					testutil.AssertNotNilFunc(t, lr, "LocalRouter"),
					testutil.AssertLenFunc(t, lr.Peers, 1, "LocalRouter.Peers"),
				)
			},
		},
		Delete: &testutil.CRUDTestDeleteFunc{
			Func: func(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) error {
				lrOp := sacloud.NewLocalRouterOp(caller)
				if err := lrOp.Delete(ctx, ctx.ID); err != nil {
					return err
				}
				if err := lrOp.Delete(ctx, peerLocalRouter.ID); err != nil {
					return err
				}
				sacloud.NewSwitchOp(caller).Delete(ctx, testZone, sw.ID) // nolint
				return nil
			},
		},
	})
}

func TestLocalRouterBuilder_minimum(t *testing.T) {
	testutil.RunCRUD(t, &testutil.CRUDTestCase{
		SetupAPICallerFunc: func() sacloud.APICaller {
			return testutil.SingletonAPICaller()
		},
		Parallel:          true,
		IgnoreStartupWait: true,
		Create: &testutil.CRUDTestFunc{
			Func: func(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
				builder := &Builder{
					Name:   testutil.ResourceName("local-router-builder"),
					Client: NewAPIClient(caller),
				}
				return builder.Build(ctx)
			},
		},
		Read: &testutil.CRUDTestFunc{
			Func: func(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
				return sacloud.NewLocalRouterOp(caller).Read(ctx, ctx.ID)
			},
			CheckFunc: func(t testutil.TestT, ctx *testutil.CRUDTestContext, value interface{}) error {
				lr := value.(*sacloud.LocalRouter)
				return testutil.DoAsserts(
					testutil.AssertNotNilFunc(t, lr, "LocalRouter"),
				)
			},
		},
		Delete: &testutil.CRUDTestDeleteFunc{
			Func: func(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) error {
				lrOp := sacloud.NewLocalRouterOp(caller)
				return lrOp.Delete(ctx, ctx.ID)
			},
		},
	})
}
