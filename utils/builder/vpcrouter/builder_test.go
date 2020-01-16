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
	"time"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/testutil"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
	"github.com/sacloud/libsacloud/v2/utils/builder"
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
					Name:        testutil.ResourceName("vpc-router-builder"),
					Description: "description",
					Tags:        types.Tags{"tag1", "tag2"},
					PlanID:      types.VPCRouterPlans.Standard,
					NICSetting:  &StandardNICSetting{},
					AdditionalNICSettings: []AdditionalNICSettingHolder{
						&AdditionalStandardNICSetting{
							SwitchID:       switchID,
							IPAddress:      "192.168.0.1",
							NetworkMaskLen: 24,
							Index:          2,
						},
					},
					RouterSetting: &RouterSetting{
						InternetConnectionEnabled: types.StringTrue,
					},
					SetupOptions: getSetupOption(),
					Client:       sacloud.NewVPCRouterOp(caller),
				}
				return builder.Build(ctx, testZone)
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
	var routerID, routerSwitchID, switchID, updSwitchID types.ID
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

			updSwitch, err := swOp.Create(ctx, testZone, &sacloud.SwitchCreateRequest{
				Name: testutil.ResourceName("vpc-router-builder-upd"),
			})
			if err != nil {
				return err
			}
			updSwitchID = updSwitch.ID

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
					Name:        testutil.ResourceName("vpc-router-builder"),
					Description: "description",
					Tags:        types.Tags{"tag1", "tag2"},
					PlanID:      types.VPCRouterPlans.Premium,
					NICSetting: &PremiumNICSetting{
						SwitchID:         routerSwitchID,
						VirtualIPAddress: addresses[0],
						IPAddress1:       addresses[1],
						IPAddress2:       addresses[2],
						IPAliases:        []string{addresses[3], addresses[4]},
					},
					AdditionalNICSettings: []AdditionalNICSettingHolder{
						&AdditionalPremiumNICSetting{
							SwitchID:         switchID,
							IPAddress1:       "192.168.0.11",
							IPAddress2:       "192.168.0.12",
							VirtualIPAddress: "192.168.0.1",
							NetworkMaskLen:   24,
							Index:            2,
						},
					},
					RouterSetting: &RouterSetting{
						VRID:                      1,
						InternetConnectionEnabled: types.StringTrue,
					},
					SetupOptions: getSetupOption(),
					Client:       sacloud.NewVPCRouterOp(caller),
				}
				return builder.Build(ctx, testZone)
			},
			CheckFunc: func(t testutil.TestT, ctx *testutil.CRUDTestContext, value interface{}) error {
				vpcRouter := value.(*sacloud.VPCRouter)
				found := false
				for _, iface := range vpcRouter.Interfaces {
					if iface.Index == 2 {
						found = true
						if err := testutil.AssertEqual(t, switchID, iface.SwitchID, "VPCRouter.Interfaces[index=2].SwitchID"); err != nil {
							return err
						}
					}
				}
				return testutil.AssertTrue(t, found, "VPCRouter.Interfaces[index=2]")
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
		Updates: []*testutil.CRUDTestFunc{
			{
				Func: func(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
					builder := &Builder{
						Name:        testutil.ResourceName("vpc-router-builder"),
						Description: "description",
						Tags:        types.Tags{"tag1", "tag2"},
						PlanID:      types.VPCRouterPlans.Premium,
						NICSetting: &PremiumNICSetting{
							SwitchID:         routerSwitchID,
							VirtualIPAddress: addresses[0],
							IPAddress1:       addresses[1],
							IPAddress2:       addresses[2],
							IPAliases:        []string{addresses[3], addresses[4]},
						},
						AdditionalNICSettings: []AdditionalNICSettingHolder{
							&AdditionalPremiumNICSetting{
								SwitchID:         updSwitchID,
								VirtualIPAddress: "192.168.0.5",
								IPAddress1:       "192.168.0.6",
								IPAddress2:       "192.168.0.7",
								NetworkMaskLen:   28,
								Index:            3,
							},
						},
						RouterSetting: &RouterSetting{
							VRID:                      1,
							InternetConnectionEnabled: types.StringTrue,
						},
						SetupOptions: getSetupOption(),
						Client:       sacloud.NewVPCRouterOp(caller),
					}
					return builder.Update(ctx, testZone, ctx.ID)
				},
				CheckFunc: func(t testutil.TestT, ctx *testutil.CRUDTestContext, value interface{}) error {
					vpcRouter := value.(*sacloud.VPCRouter)
					found := false
					for _, iface := range vpcRouter.Interfaces {
						if iface.Index == 3 {
							found = true
							if err := testutil.AssertEqual(t, updSwitchID, iface.SwitchID, "VPCRouter.Interfaces[index=2].SwitchID"); err != nil {
								return err
							}
						}
					}
					if err := testutil.AssertTrue(t, found, "VPCRouter.Interfaces[index=2]"); err != nil {
						return err
					}

					found = false
					for _, nicSetting := range vpcRouter.Settings.Interfaces {
						if nicSetting.Index == 3 {
							found = true
							err := testutil.DoAsserts(
								testutil.AssertEqualFunc(t, "192.168.0.5", nicSetting.VirtualIPAddress, "VPCRouter.Settings.Interfaces.VirtualIPAddress"),
								testutil.AssertEqualFunc(t, []string{"192.168.0.6", "192.168.0.7"}, nicSetting.IPAddress, "VPCRouter.Settings.Interfaces.IPAddress"),
								testutil.AssertEqualFunc(t, 28, nicSetting.NetworkMaskLen, "VPCRouter.Settings.Interfaces.NetworkMaskLen"),
							)
							if err != nil {
								return err
							}
						}
					}
					return testutil.AssertTrue(t, found, "VPCRouter.Setting.Interfaces[index=2]")
				},
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
				if err := swOp.Delete(ctx, testZone, switchID); err != nil {
					return err
				}
				return swOp.Delete(ctx, testZone, updSwitchID)
			},
		},
	})
}
