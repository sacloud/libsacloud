// Copyright 2016-2019 The Libsacloud Authors
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
	"time"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/testutil"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

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
				Name: testutil.ResourceName("vpc-router-builder"),
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
					VPCRouterCreateRequest: &sacloud.VPCRouterCreateRequest{
						Name:        testutil.ResourceName("vpc-router-builder"),
						Description: "description",
						Tags:        types.Tags{"tag1", "tag2"},
						PlanID:      types.VPCRouterPlans.Standard,
						Switch:      &sacloud.ApplianceConnectedSwitch{Scope: types.Scopes.Shared},
						Settings: &sacloud.VPCRouterSetting{
							VRID:                      1,
							InternetConnectionEnabled: types.StringTrue,
							Interfaces: []*sacloud.VPCRouterInterfaceSetting{
								{
									IPAddress:      []string{"192.168.0.1"},
									NetworkMaskLen: 24,
									Index:          2,
								},
							},
						},
					},
					AdditionalSwitches: []*AdditionalSwitch{
						{
							Index: 2,
							ID:    switchID,
						},
					},
				}
				return builder.Build(ctx, sacloud.NewVPCRouterOp(caller), testZone)
			},
		},
		Read: &testutil.CRUDTestFunc{
			Func: func(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
				vpcRouterOp := sacloud.NewVPCRouterOp(caller)
				return vpcRouterOp.Read(ctx, testZone, ctx.ID)
			},
			CheckFunc: func(t testutil.TestT, ctx *testutil.CRUDTestContext, value interface{}) error {
				vpcRouter := value.(*sacloud.VPCRouter)
				return testutil.DoAsserts(
					testutil.AssertNotNilFunc(t, vpcRouter, "VPCRouter"),
					testutil.AssertNotNilFunc(t, vpcRouter.Settings, "VPCRouter.Settings"),
					testutil.AssertLenFunc(t, vpcRouter.Settings.Interfaces, 1, "VPCRouter.Settings.Interfaces"),
				)
			},
		},
		Delete: &testutil.CRUDTestDeleteFunc{
			Func: func(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) error {
				vpcRouterOp := sacloud.NewVPCRouterOp(caller)
				return vpcRouterOp.Delete(ctx, testZone, ctx.ID)
			},
		},
		Cleanup: func(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) error {
			swOp := sacloud.NewSwitchOp(caller)
			return swOp.Delete(ctx, testZone, switchID)
		},
	})

}

func TestBuilder_BuildWithRouter(t *testing.T) {
	var routerID, routerSwitchID, switchID types.ID
	var addresses []string
	var testZone = testutil.TestZone()

	testutil.Run(t, &testutil.CRUDTestCase{
		SetupAPICallerFunc: func() sacloud.APICaller {
			return testutil.SingletonAPICaller()
		},
		Parallel:          true,
		IgnoreStartupWait: true,
		Setup: func(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) error {
			internetOp := sacloud.NewInternetOp(caller)

			created, err := internetOp.Create(ctx, testZone, &sacloud.InternetCreateRequest{
				Name:           testutil.ResourceName("vpc-router-builder"),
				NetworkMaskLen: 28,
				BandWidthMbps:  100,
			})
			if err != nil {
				return err
			}

			routerID = created.ID
			routerSwitchID = created.Switch.ID
			max := 30
			for {
				if max == 0 {
					break
				}
				_, err := internetOp.Read(ctx, testZone, routerID)
				if err != nil || sacloud.IsNotFoundError(err) {
					max--
					time.Sleep(3 * time.Second)
					continue
				}
				break
			}

			swOp := sacloud.NewSwitchOp(caller)
			sw, err := swOp.Create(ctx, testZone, &sacloud.SwitchCreateRequest{
				Name: testutil.ResourceName("vpc-router-builder"),
			})
			if err != nil {
				return err
			}
			switchID = sw.ID

			routerSwitch, err := swOp.Read(ctx, testZone, created.Switch.ID)
			if err != nil {
				return err
			}
			addresses = routerSwitch.Subnets[0].GetAssignedIPAddresses()
			return nil
		},
		Create: &testutil.CRUDTestFunc{
			Func: func(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
				builder := &Builder{
					VPCRouterCreateRequest: &sacloud.VPCRouterCreateRequest{
						Name:        testutil.ResourceName("vpc-router-builder"),
						Description: "description",
						Tags:        types.Tags{"tag1", "tag2"},
						PlanID:      types.VPCRouterPlans.Premium,
						Switch: &sacloud.ApplianceConnectedSwitch{
							ID: routerSwitchID,
						},
						IPAddresses: []string{addresses[1], addresses[2]},
						Settings: &sacloud.VPCRouterSetting{
							VRID:                      1,
							InternetConnectionEnabled: types.StringTrue,
							Interfaces: []*sacloud.VPCRouterInterfaceSetting{
								{
									IPAddress:        []string{addresses[1], addresses[2]},
									VirtualIPAddress: addresses[0],
									IPAliases:        []string{addresses[3], addresses[4]},
									NetworkMaskLen:   28,
									Index:            0,
								},
								{
									IPAddress:        []string{"192.168.0.11", "192.168.0.12"},
									VirtualIPAddress: "192.168.0.1",
									NetworkMaskLen:   24,
									Index:            2,
								},
							},
						},
					},
					AdditionalSwitches: []*AdditionalSwitch{
						{
							Index: 2,
							ID:    switchID,
						},
					},
				}
				return builder.Build(ctx, sacloud.NewVPCRouterOp(caller), testZone)
			},
		},
		Read: &testutil.CRUDTestFunc{
			Func: func(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
				vpcRouterOp := sacloud.NewVPCRouterOp(caller)
				return vpcRouterOp.Read(ctx, testZone, ctx.ID)
			},
			CheckFunc: func(t testutil.TestT, ctx *testutil.CRUDTestContext, value interface{}) error {
				vpcRouter := value.(*sacloud.VPCRouter)
				return testutil.DoAsserts(
					testutil.AssertNotNilFunc(t, vpcRouter, "VPCRouter"),
					testutil.AssertNotNilFunc(t, vpcRouter.Settings, "VPCRouter.Settings"),
					testutil.AssertLenFunc(t, vpcRouter.Settings.Interfaces, 2, "VPCRouter.Settings.Interfaces"),
				)
			},
		},
		Delete: &testutil.CRUDTestDeleteFunc{
			Func: func(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) error {
				vpcRouterOp := sacloud.NewVPCRouterOp(caller)
				if err := vpcRouterOp.Delete(ctx, testZone, ctx.ID); err != nil {
					return err
				}

				internetOp := sacloud.NewInternetOp(caller)
				if err := internetOp.Delete(ctx, testZone, routerID); err != nil {
					return err
				}
				swOp := sacloud.NewSwitchOp(caller)
				return swOp.Delete(ctx, testZone, switchID)
			},
		},
	})

}
