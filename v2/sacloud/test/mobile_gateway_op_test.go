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

package test

import (
	"os"
	"testing"

	"github.com/sacloud/libsacloud/v2/helper/power"
	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/testutil"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

func TestMobileGatewayOpCRUD(t *testing.T) {
	testutil.PreCheckEnvsFunc("SAKURACLOUD_SIM_ICCID", "SAKURACLOUD_SIM_PASSCODE")(t)

	initMobileGatewayVariables()

	testutil.Run(t, &testutil.CRUDTestCase{
		Parallel: true,

		SetupAPICallerFunc: singletonAPICaller,

		Create: &testutil.CRUDTestFunc{
			Func: testMobileGatewayCreate,
			CheckFunc: testutil.AssertEqualWithExpected(&testutil.CRUDTestExpect{
				ExpectValue:  createMobileGatewayExpected,
				IgnoreFields: ignoreMobileGatewayFields,
			}),
		},

		Read: &testutil.CRUDTestFunc{
			Func: testMobileGatewayRead,
			CheckFunc: testutil.AssertEqualWithExpected(&testutil.CRUDTestExpect{
				ExpectValue:  createMobileGatewayExpected,
				IgnoreFields: ignoreMobileGatewayFields,
			}),
		},

		Updates: []*testutil.CRUDTestFunc{
			{
				Func: testMobileGatewayUpdate,
				CheckFunc: testutil.AssertEqualWithExpected(&testutil.CRUDTestExpect{
					ExpectValue:  updateMobileGatewayExpected,
					IgnoreFields: ignoreMobileGatewayFields,
				}),
			},
			{
				Func: testMobileGatewayUpdateSettings,
				CheckFunc: testutil.AssertEqualWithExpected(&testutil.CRUDTestExpect{
					ExpectValue:  updateMobileGatewaySettingsExpected,
					IgnoreFields: ignoreMobileGatewayFields,
				}),
			},
			// shutdown(no check)
			{
				Func: func(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
					mgwOp := sacloud.NewMobileGatewayOp(caller)
					// shutdown
					if err := mgwOp.Shutdown(ctx, testZone, ctx.ID, &sacloud.ShutdownOption{Force: true}); err != nil {
						return nil, err
					}

					waiter := sacloud.WaiterForDown(func() (interface{}, error) {
						return mgwOp.Read(ctx, testZone, ctx.ID)
					})

					return waiter.WaitForState(ctx)
				},
				SkipExtractID: true,
			},
			// connect to switch
			{
				Func: func(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
					// prepare switch
					swOp := sacloud.NewSwitchOp(caller)
					sw, err := swOp.Create(ctx, testZone, &sacloud.SwitchCreateRequest{
						Name: testutil.ResourceName("switch-for-mobile-gateway"),
					})
					if err != nil {
						return nil, err
					}

					ctx.Values["mobile-gateway/switch"] = sw.ID

					// connect
					mgwOp := sacloud.NewMobileGatewayOp(caller)
					if err := mgwOp.ConnectToSwitch(ctx, testZone, ctx.ID, sw.ID); err != nil {
						return nil, err
					}

					return mgwOp.Read(ctx, testZone, ctx.ID)
				},
				CheckFunc: func(t testutil.TestT, ctx *testutil.CRUDTestContext, i interface{}) error {
					mgw := i.(*sacloud.MobileGateway)
					return testutil.DoAsserts(
						testutil.AssertLenFunc(t, mgw.Interfaces, 2, "len(MobileGateway.Interfaces)"),
					)
				},
				SkipExtractID: true,
			},
			// set IPAddress to eth1
			{
				Func: func(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
					mgwOp := sacloud.NewMobileGatewayOp(caller)
					return mgwOp.Update(ctx, testZone, ctx.ID, &sacloud.MobileGatewayUpdateRequest{
						Settings: &sacloud.MobileGatewaySetting{
							Interfaces: []*sacloud.MobileGatewayInterfaceSetting{
								{
									IPAddress:      []string{"192.168.2.11"},
									NetworkMaskLen: 24,
									Index:          1,
								},
							},
						},
					})
				},
				CheckFunc: func(t testutil.TestT, ctx *testutil.CRUDTestContext, i interface{}) error {
					mgw := i.(*sacloud.MobileGateway)
					return testutil.DoAsserts(
						testutil.AssertNotNilFunc(t, mgw.Settings.Interfaces, "MobileGateway.Settings.Interfaces"),
						testutil.AssertEqualFunc(t, 1, mgw.Settings.Interfaces[0].Index, "MobileGateway.Settings.Interfaces.Index"),
						testutil.AssertEqualFunc(t, "192.168.2.11", mgw.Settings.Interfaces[0].IPAddress[0], "MobileGateway.Settings.Interfaces.IPAddress"),
						testutil.AssertEqualFunc(t, 24, mgw.Settings.Interfaces[0].NetworkMaskLen, "MobileGateway.Settings.Interfaces.NetworkMaskLen"),
					)
				},
				SkipExtractID: true,
			},

			// Get/Set DNS
			{
				Func: func(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
					mgwOp := sacloud.NewMobileGatewayOp(caller)
					if err := mgwOp.SetDNS(ctx, testZone, ctx.ID, &sacloud.MobileGatewayDNSSetting{
						DNS1: "8.8.8.8",
						DNS2: "8.8.4.4",
					}); err != nil {
						return nil, err
					}
					return mgwOp.GetDNS(ctx, testZone, ctx.ID)
				},
				CheckFunc: func(t testutil.TestT, ctx *testutil.CRUDTestContext, i interface{}) error {
					dns := i.(*sacloud.MobileGatewayDNSSetting)
					return testutil.DoAsserts(
						testutil.AssertEqualFunc(t, "8.8.8.8", dns.DNS1, "DNS1"),
						testutil.AssertEqualFunc(t, "8.8.4.4", dns.DNS2, "DNS2"),
					)
				},
				SkipExtractID: true,
			},
			// Add/List SIM
			{
				Func: func(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
					simOp := sacloud.NewSIMOp(caller)
					sim, err := simOp.Create(ctx, &sacloud.SIMCreateRequest{
						Name:     testutil.ResourceName("switch-for-mobile-gateway"),
						ICCID:    iccid,
						PassCode: passcode,
					})
					if err != nil {
						return nil, err
					}

					mgwOp := sacloud.NewMobileGatewayOp(caller)
					if err := mgwOp.AddSIM(ctx, testZone, ctx.ID, &sacloud.MobileGatewayAddSIMRequest{
						SIMID: sim.ID.String(),
					}); err != nil {
						return nil, err
					}

					ctx.Values["mobile-gateway/sim"] = sim.ID
					return mgwOp.ListSIM(ctx, testZone, ctx.ID)
				},
				CheckFunc: func(t testutil.TestT, ctx *testutil.CRUDTestContext, i interface{}) error {
					sims := i.([]*sacloud.MobileGatewaySIMInfo)
					return testutil.DoAsserts(
						testutil.AssertLenFunc(t, sims, 1, "len(SIM)"),
					)
				},
				SkipExtractID: true,
			},
			// SIMOp: Assign IP
			{
				Func: func(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
					client := sacloud.NewSIMOp(caller)
					simID := ctx.Values["mobile-gateway/sim"].(types.ID)
					if err := client.AssignIP(ctx, simID, &sacloud.SIMAssignIPRequest{
						IP: "192.168.2.1",
					}); err != nil {
						return nil, err
					}
					return client.Status(ctx, simID)
				},
				CheckFunc: func(t testutil.TestT, ctx *testutil.CRUDTestContext, v interface{}) error {
					simInfo := v.(*sacloud.SIMInfo)
					return testutil.DoAsserts(
						testutil.AssertEqualFunc(t, "192.168.2.1", simInfo.IP, "SIMInfo.IP"),
					)
				},
				SkipExtractID: true,
			},
			// SIMOp: clear IP
			{
				Func: func(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
					client := sacloud.NewSIMOp(caller)
					simID := ctx.Values["mobile-gateway/sim"].(types.ID)
					if err := client.ClearIP(ctx, simID); err != nil {
						return nil, err
					}
					return client.Status(ctx, simID)
				},
				CheckFunc: func(t testutil.TestT, ctx *testutil.CRUDTestContext, v interface{}) error {
					simInfo := v.(*sacloud.SIMInfo)
					return testutil.DoAsserts(
						testutil.AssertEmptyFunc(t, simInfo.IP, "SIMInfo.IP"),
					)
				},
				SkipExtractID: true,
			},

			// Get/Set SIMRoutes
			{
				Func: func(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
					mgwOp := sacloud.NewMobileGatewayOp(caller)
					if err := mgwOp.SetSIMRoutes(ctx, testZone, ctx.ID, []*sacloud.MobileGatewaySIMRouteParam{
						{
							ResourceID: ctx.Values["mobile-gateway/sim"].(types.ID).String(),
							Prefix:     "192.168.3.0/24",
						},
					}); err != nil {
						return nil, err
					}
					return mgwOp.GetSIMRoutes(ctx, testZone, ctx.ID)
				},
				CheckFunc: func(t testutil.TestT, ctx *testutil.CRUDTestContext, i interface{}) error {
					routes := i.([]*sacloud.MobileGatewaySIMRoute)
					simID := ctx.Values["mobile-gateway/sim"].(types.ID)
					return testutil.DoAsserts(
						testutil.AssertLenFunc(t, routes, 1, "len(SIMRoutes)"),
						testutil.AssertEqualFunc(t, "192.168.3.0/24", routes[0].Prefix, "SIMRoute.Prefix"),
						testutil.AssertEqualFunc(t, simID.String(), routes[0].ResourceID, "SIMRoute.ResourceID"),
					)
				},
				SkipExtractID: true,
			},
			// Delete SIMRoutes
			{
				Func: func(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
					mgwOp := sacloud.NewMobileGatewayOp(caller)
					if err := mgwOp.SetSIMRoutes(ctx, testZone, ctx.ID, []*sacloud.MobileGatewaySIMRouteParam{}); err != nil {
						return nil, err
					}
					return mgwOp.GetSIMRoutes(ctx, testZone, ctx.ID)
				},
				CheckFunc: func(t testutil.TestT, ctx *testutil.CRUDTestContext, i interface{}) error {
					routes := i.([]*sacloud.MobileGatewaySIMRoute)
					return testutil.DoAsserts(
						testutil.AssertLenFunc(t, routes, 0, "len(SIMRoutes)"),
					)
				},
				SkipExtractID: true,
			},

			// Get/Set TrafficConfig
			{
				Func: func(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
					mgwOp := sacloud.NewMobileGatewayOp(caller)
					if err := mgwOp.SetTrafficConfig(ctx, testZone, ctx.ID, &sacloud.MobileGatewayTrafficControl{
						TrafficQuotaInMB:       10,
						BandWidthLimitInKbps:   20,
						EmailNotifyEnabled:     true,
						SlackNotifyEnabled:     true,
						SlackNotifyWebhooksURL: "https://hooks.slack.com/services/XXXXXXXXX/XXXXXXXXX/XXXXXXXXXXXXXXXXXXXXXXXX",
						AutoTrafficShaping:     true,
					}); err != nil {
						return nil, err
					}
					return mgwOp.GetTrafficConfig(ctx, testZone, ctx.ID)
				},
				CheckFunc: func(t testutil.TestT, ctx *testutil.CRUDTestContext, i interface{}) error {
					slackURL := "https://hooks.slack.com/services/XXXXXXXXX/XXXXXXXXX/XXXXXXXXXXXXXXXXXXXXXXXX"
					config := i.(*sacloud.MobileGatewayTrafficControl)
					return testutil.DoAsserts(
						testutil.AssertEqualFunc(t, 10, config.TrafficQuotaInMB, "TrafficConfig.TrafficQuotaInMB"),
						testutil.AssertEqualFunc(t, 20, config.BandWidthLimitInKbps, "TrafficConfig.BandWidthLimitInKbps"),
						testutil.AssertEqualFunc(t, true, config.EmailNotifyEnabled, "TrafficConfig.EmailNotifyEnabled"),
						testutil.AssertEqualFunc(t, true, config.SlackNotifyEnabled, "TrafficConfig.SlackNotifyEnabled"),
						testutil.AssertEqualFunc(t, slackURL, config.SlackNotifyWebhooksURL, "TrafficConfig.SlackNotifyWebhooksURL"),
						testutil.AssertEqualFunc(t, true, config.AutoTrafficShaping, "TrafficConfig.AutoTrafficShaping"),
					)
				},
				SkipExtractID: true,
			},
			// Delete TrafficConfig(no check)
			{
				Func: func(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
					mgwOp := sacloud.NewMobileGatewayOp(caller)
					return nil, mgwOp.DeleteTrafficConfig(ctx, testZone, ctx.ID)
				},
				SkipExtractID: true,
			},

			// Get TrafficStatus
			{
				Func: func(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
					mgwOp := sacloud.NewMobileGatewayOp(caller)
					return mgwOp.TrafficStatus(ctx, testZone, ctx.ID)
				},
				CheckFunc: func(t testutil.TestT, ctx *testutil.CRUDTestContext, i interface{}) error {
					status := i.(*sacloud.MobileGatewayTrafficStatus)
					return testutil.DoAsserts(
						testutil.AssertNotNilFunc(t, status, "TrafficStatus"),
						testutil.AssertEqualFunc(t, types.StringNumber(0), status.UplinkBytes, "TrafficStatus.UplinkBytes"),
						testutil.AssertEqualFunc(t, types.StringNumber(0), status.DownlinkBytes, "TrafficStatus.DownlinkBytes"),
					)
				},
				SkipExtractID: true,
			},

			// Delete SIM
			{
				Func: func(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
					simID := ctx.Values["mobile-gateway/sim"].(types.ID)
					mgwOp := sacloud.NewMobileGatewayOp(caller)
					if err := mgwOp.DeleteSIM(ctx, testZone, ctx.ID, simID); err != nil {
						return nil, err
					}

					simOp := sacloud.NewSIMOp(caller)
					if err := simOp.Delete(ctx, simID); err != nil {
						return nil, err
					}

					return mgwOp.ListSIM(ctx, testZone, ctx.ID)
				},
				CheckFunc: func(t testutil.TestT, ctx *testutil.CRUDTestContext, i interface{}) error {
					sims := i.([]*sacloud.MobileGatewaySIMInfo)
					return testutil.DoAsserts(
						testutil.AssertLenFunc(t, sims, 0, "len(SIM)"),
					)
				},
				SkipExtractID: true,
			},
			// disconnect from switch
			{
				Func: func(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
					mgwOp := sacloud.NewMobileGatewayOp(caller)
					if err := mgwOp.DisconnectFromSwitch(ctx, testZone, ctx.ID); err != nil {
						return nil, err
					}

					swID := ctx.Values["mobile-gateway/switch"].(types.ID)
					swOp := sacloud.NewSwitchOp(caller)
					if err := swOp.Delete(ctx, testZone, swID); err != nil {
						return nil, err
					}

					return mgwOp.Read(ctx, testZone, ctx.ID)
				},
				CheckFunc: func(t testutil.TestT, ctx *testutil.CRUDTestContext, i interface{}) error {
					mgw := i.(*sacloud.MobileGateway)
					return testutil.DoAsserts(
						testutil.AssertLenFunc(t, mgw.Interfaces, 1, "len(MobileGateway.Interfaces)"),
					)
				},
				SkipExtractID: true,
			},
		},
		Shutdown: func(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) error {
			client := sacloud.NewMobileGatewayOp(caller)
			return power.ShutdownMobileGateway(ctx, client, testZone, ctx.ID, true)
		},
		Delete: &testutil.CRUDTestDeleteFunc{
			Func: testMobileGatewayDelete,
		},
	})
}

func initMobileGatewayVariables() {
	iccid = os.Getenv("SAKURACLOUD_SIM_ICCID")
	passcode = os.Getenv("SAKURACLOUD_SIM_PASSCODE")

	createMobileGatewayParam = &sacloud.MobileGatewayCreateRequest{
		Name:        testutil.ResourceName("mobile-gateway"),
		Description: "desc",
		Tags:        []string{"tag1", "tag2"},
		Settings: &sacloud.MobileGatewaySettingCreate{
			InternetConnectionEnabled:       true,
			InterDeviceCommunicationEnabled: true,
		},
	}
	createMobileGatewayExpected = &sacloud.MobileGateway{
		Name:         createMobileGatewayParam.Name,
		Description:  createMobileGatewayParam.Description,
		Tags:         createMobileGatewayParam.Tags,
		Availability: types.Availabilities.Available,
		Settings: &sacloud.MobileGatewaySetting{
			InternetConnectionEnabled:       true,
			InterDeviceCommunicationEnabled: true,
		},
	}
	updateMobileGatewayParam = &sacloud.MobileGatewayUpdateRequest{
		Name:        testutil.ResourceName("mobile-gateway-upd"),
		Description: "desc-upd",
		Tags:        []string{"tag1-upd", "tag2-upd"},
		Settings: &sacloud.MobileGatewaySetting{
			InternetConnectionEnabled:       false,
			InterDeviceCommunicationEnabled: false,
		},
	}
	updateMobileGatewayExpected = &sacloud.MobileGateway{
		Name:         updateMobileGatewayParam.Name,
		Description:  updateMobileGatewayParam.Description,
		Tags:         updateMobileGatewayParam.Tags,
		Availability: types.Availabilities.Available,
		Settings: &sacloud.MobileGatewaySetting{
			InternetConnectionEnabled:       false,
			InterDeviceCommunicationEnabled: false,
		},
	}
	updateMobileGatewaySettingsParam = &sacloud.MobileGatewayUpdateSettingsRequest{
		Settings: &sacloud.MobileGatewaySetting{
			InternetConnectionEnabled:       true,
			InterDeviceCommunicationEnabled: true,
		},
	}
	updateMobileGatewaySettingsExpected = &sacloud.MobileGateway{
		Name:         updateMobileGatewayParam.Name,
		Description:  updateMobileGatewayParam.Description,
		Tags:         updateMobileGatewayParam.Tags,
		Availability: types.Availabilities.Available,
		Settings: &sacloud.MobileGatewaySetting{
			InternetConnectionEnabled:       true,
			InterDeviceCommunicationEnabled: true,
		},
	}
}

var (
	ignoreMobileGatewayFields = []string{
		"ID",
		"Class",
		"IconID",
		"CreatedAt",
		"Availability",
		"InstanceHostName",
		"InstanceHostInfoURL",
		"InstanceStatus",
		"InstanceStatusChangedAt",
		"Interfaces",
		"ZoneID",
		"SettingsHash",
	}
	iccid                               string
	passcode                            string
	createMobileGatewayParam            *sacloud.MobileGatewayCreateRequest
	createMobileGatewayExpected         *sacloud.MobileGateway
	updateMobileGatewayParam            *sacloud.MobileGatewayUpdateRequest
	updateMobileGatewayExpected         *sacloud.MobileGateway
	updateMobileGatewaySettingsParam    *sacloud.MobileGatewayUpdateSettingsRequest
	updateMobileGatewaySettingsExpected *sacloud.MobileGateway
)

func testMobileGatewayCreate(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewMobileGatewayOp(caller)
	v, err := client.Create(ctx, testZone, createMobileGatewayParam)
	if err != nil {
		return nil, err
	}
	value, err := sacloud.WaiterForReady(func() (interface{}, error) {
		return client.Read(ctx, testZone, v.ID)
	}).WaitForState(ctx)
	if err != nil {
		return nil, err
	}
	if err := client.Boot(ctx, testZone, v.ID); err != nil {
		return nil, err
	}
	return value, nil
}

func testMobileGatewayRead(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewMobileGatewayOp(caller)
	return client.Read(ctx, testZone, ctx.ID)
}

func testMobileGatewayUpdate(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewMobileGatewayOp(caller)
	return client.Update(ctx, testZone, ctx.ID, updateMobileGatewayParam)
}

func testMobileGatewayUpdateSettings(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewMobileGatewayOp(caller)
	return client.UpdateSettings(ctx, testZone, ctx.ID, updateMobileGatewaySettingsParam)
}

func testMobileGatewayDelete(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) error {
	client := sacloud.NewMobileGatewayOp(caller)
	return client.Delete(ctx, testZone, ctx.ID)
}
