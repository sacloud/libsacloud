package test

import (
	"os"
	"testing"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

func TestMobileGatewayOpCRUD(t *testing.T) {

	PreCheckEnvsFunc("SAKURACLOUD_SIM_ICCID", "SAKURACLOUD_SIM_PASSCODE")(t)

	initMobileGatewayVariables()

	Run(t, &CRUDTestCase{
		Parallel: true,

		SetupAPICallerFunc: singletonAPICaller,

		Create: &CRUDTestFunc{
			Func: testMobileGatewayCreate,
			CheckFunc: AssertEqualWithExpected(&CRUDTestExpect{
				ExpectValue:  createMobileGatewayExpected,
				IgnoreFields: ignoreMobileGatewayFields,
			}),
		},

		Read: &CRUDTestFunc{
			Func: testMobileGatewayRead,
			CheckFunc: AssertEqualWithExpected(&CRUDTestExpect{
				ExpectValue:  createMobileGatewayExpected,
				IgnoreFields: ignoreMobileGatewayFields,
			}),
		},

		Updates: []*CRUDTestFunc{
			{
				Func: testMobileGatewayUpdate,
				CheckFunc: AssertEqualWithExpected(&CRUDTestExpect{
					ExpectValue:  updateMobileGatewayExpected,
					IgnoreFields: ignoreMobileGatewayFields,
				}),
			},
			// shutdown(no check)
			{
				Func: func(ctx *CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
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
				Func: func(ctx *CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
					// prepare switch
					swOp := sacloud.NewSwitchOp(caller)
					sw, err := swOp.Create(ctx, testZone, &sacloud.SwitchCreateRequest{
						Name: "libsacloud-test-mobile-gateway",
					})
					if err != nil {
						return nil, err
					}

					// connect
					mgwOp := sacloud.NewMobileGatewayOp(caller)
					if err := mgwOp.ConnectToSwitch(ctx, testZone, ctx.ID, sw.ID); err != nil {
						return nil, err
					}

					return mgwOp.Read(ctx, testZone, ctx.ID)
				},
				CheckFunc: func(t TestT, ctx *CRUDTestContext, i interface{}) error {
					mgw := i.(*sacloud.MobileGateway)
					return DoAsserts(
						AssertLenFunc(t, mgw.Interfaces, 2, "len(MobileGateway.Interfaces)"),
					)
				},
				SkipExtractID: true,
			},
			// set IPAddress to eth1
			{
				Func: func(ctx *CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
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
				CheckFunc: func(t TestT, ctx *CRUDTestContext, i interface{}) error {
					mgw := i.(*sacloud.MobileGateway)
					return DoAsserts(
						AssertNotNilFunc(t, mgw.Settings.Interfaces, "MobileGateway.Settings.Interfaces"),
						AssertEqualFunc(t, 1, mgw.Settings.Interfaces[0].Index, "MobileGateway.Settings.Interfaces.Index"),
						AssertEqualFunc(t, "192.168.2.11", mgw.Settings.Interfaces[0].IPAddress[0], "MobileGateway.Settings.Interfaces.IPAddress"),
						AssertEqualFunc(t, 24, mgw.Settings.Interfaces[0].NetworkMaskLen, "MobileGateway.Settings.Interfaces.NetworkMaskLen"),
					)
				},
				SkipExtractID: true,
			},

			// Get/Set DNS
			{
				Func: func(ctx *CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
					mgwOp := sacloud.NewMobileGatewayOp(caller)
					if err := mgwOp.SetDNS(ctx, testZone, ctx.ID, &sacloud.MobileGatewayDNSSetting{
						DNS1: "8.8.8.8",
						DNS2: "8.8.4.4",
					}); err != nil {
						return nil, err
					}
					return mgwOp.GetDNS(ctx, testZone, ctx.ID)
				},
				CheckFunc: func(t TestT, ctx *CRUDTestContext, i interface{}) error {
					dns := i.(*sacloud.MobileGatewayDNSSetting)
					return DoAsserts(
						AssertEqualFunc(t, "8.8.8.8", dns.DNS1, "DNS1"),
						AssertEqualFunc(t, "8.8.4.4", dns.DNS2, "DNS2"),
					)
				},
				SkipExtractID: true,
			},
			// Add/List SIM
			{
				Func: func(ctx *CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
					simOp := sacloud.NewSIMOp(caller)
					sim, err := simOp.Create(ctx, &sacloud.SIMCreateRequest{
						Name:     "libsacloud-switch-for-mobile-gateway",
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
				CheckFunc: func(t TestT, ctx *CRUDTestContext, i interface{}) error {
					sims := i.([]*sacloud.MobileGatewaySIMInfo)
					return DoAsserts(
						AssertLenFunc(t, sims, 1, "len(SIM)"),
					)
				},
				SkipExtractID: true,
			},
			// SIMOp: Assign IP
			{
				Func: func(ctx *CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
					client := sacloud.NewSIMOp(caller)
					simID := ctx.Values["mobile-gateway/sim"].(types.ID)
					if err := client.AssignIP(ctx, simID, &sacloud.SIMAssignIPRequest{
						IP: "192.168.2.1",
					}); err != nil {
						return nil, err
					}
					return client.Status(ctx, simID)
				},
				CheckFunc: func(t TestT, ctx *CRUDTestContext, v interface{}) error {
					simInfo := v.(*sacloud.SIMInfo)
					return DoAsserts(
						AssertEqualFunc(t, "192.168.2.1", simInfo.IP, "SIMInfo.IP"),
					)
				},
				SkipExtractID: true,
			},
			// SIMOp: clear IP
			{
				Func: func(ctx *CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
					client := sacloud.NewSIMOp(caller)
					simID := ctx.Values["mobile-gateway/sim"].(types.ID)
					if err := client.ClearIP(ctx, simID); err != nil {
						return nil, err
					}
					return client.Status(ctx, simID)
				},
				CheckFunc: func(t TestT, ctx *CRUDTestContext, v interface{}) error {
					simInfo := v.(*sacloud.SIMInfo)
					return DoAsserts(
						AssertEmptyFunc(t, simInfo.IP, "SIMInfo.IP"),
					)
				},
				SkipExtractID: true,
			},

			// Get/Set SIMRoutes
			{
				Func: func(ctx *CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
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
				CheckFunc: func(t TestT, ctx *CRUDTestContext, i interface{}) error {
					routes := i.([]*sacloud.MobileGatewaySIMRoute)
					simID := ctx.Values["mobile-gateway/sim"].(types.ID)
					return DoAsserts(
						AssertLenFunc(t, routes, 1, "len(SIMRoutes)"),
						AssertEqualFunc(t, "192.168.3.0/24", routes[0].Prefix, "SIMRoute.Prefix"),
						AssertEqualFunc(t, simID.String(), routes[0].ResourceID, "SIMRoute.ResourceID"),
					)
				},
				SkipExtractID: true,
			},
			// Delete SIMRoutes
			{
				Func: func(ctx *CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
					mgwOp := sacloud.NewMobileGatewayOp(caller)
					if err := mgwOp.SetSIMRoutes(ctx, testZone, ctx.ID, []*sacloud.MobileGatewaySIMRouteParam{}); err != nil {
						return nil, err
					}
					return mgwOp.GetSIMRoutes(ctx, testZone, ctx.ID)
				},
				CheckFunc: func(t TestT, ctx *CRUDTestContext, i interface{}) error {
					routes := i.([]*sacloud.MobileGatewaySIMRoute)
					return DoAsserts(
						AssertLenFunc(t, routes, 0, "len(SIMRoutes)"),
					)
				},
				SkipExtractID: true,
			},

			// Get/Set TrafficConfig
			{
				Func: func(ctx *CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
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
				CheckFunc: func(t TestT, ctx *CRUDTestContext, i interface{}) error {
					slackURL := "https://hooks.slack.com/services/XXXXXXXXX/XXXXXXXXX/XXXXXXXXXXXXXXXXXXXXXXXX"
					config := i.(*sacloud.MobileGatewayTrafficControl)
					return DoAsserts(
						AssertEqualFunc(t, 10, config.TrafficQuotaInMB, "TrafficConfig.TrafficQuotaInMB"),
						AssertEqualFunc(t, 20, config.BandWidthLimitInKbps, "TrafficConfig.BandWidthLimitInKbps"),
						AssertEqualFunc(t, true, config.EmailNotifyEnabled, "TrafficConfig.EmailNotifyEnabled"),
						AssertEqualFunc(t, true, config.SlackNotifyEnabled, "TrafficConfig.SlackNotifyEnabled"),
						AssertEqualFunc(t, slackURL, config.SlackNotifyWebhooksURL, "TrafficConfig.SlackNotifyWebhooksURL"),
						AssertEqualFunc(t, true, config.AutoTrafficShaping, "TrafficConfig.AutoTrafficShaping"),
					)
				},
				SkipExtractID: true,
			},
			// Delete TrafficConfig(no check)
			{
				Func: func(ctx *CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
					mgwOp := sacloud.NewMobileGatewayOp(caller)
					return nil, mgwOp.DeleteTrafficConfig(ctx, testZone, ctx.ID)
				},
				SkipExtractID: true,
			},

			// Get TrafficStatus
			{
				Func: func(ctx *CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
					mgwOp := sacloud.NewMobileGatewayOp(caller)
					return mgwOp.TrafficStatus(ctx, testZone, ctx.ID)
				},
				CheckFunc: func(t TestT, ctx *CRUDTestContext, i interface{}) error {
					status := i.(*sacloud.MobileGatewayTrafficStatus)
					return DoAsserts(
						AssertNotNilFunc(t, status, "TrafficStatus"),
						AssertEqualFunc(t, types.StringNumber(0), status.UplinkBytes, "TrafficStatus.UplinkBytes"),
						AssertEqualFunc(t, types.StringNumber(0), status.DownlinkBytes, "TrafficStatus.DownlinkBytes"),
					)
				},
				SkipExtractID: true,
			},

			// Delete SIM
			{
				Func: func(ctx *CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
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
				CheckFunc: func(t TestT, ctx *CRUDTestContext, i interface{}) error {
					sims := i.([]*sacloud.MobileGatewaySIMInfo)
					return DoAsserts(
						AssertLenFunc(t, sims, 0, "len(SIM)"),
					)
				},
				SkipExtractID: true,
			},
			// disconnect from switch
			{
				Func: func(ctx *CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
					mgwOp := sacloud.NewMobileGatewayOp(caller)
					if err := mgwOp.DisconnectFromSwitch(ctx, testZone, ctx.ID); err != nil {
						return nil, err
					}
					return mgwOp.Read(ctx, testZone, ctx.ID)
				},
				CheckFunc: func(t TestT, ctx *CRUDTestContext, i interface{}) error {
					mgw := i.(*sacloud.MobileGateway)
					return DoAsserts(
						AssertLenFunc(t, mgw.Interfaces, 1, "len(MobileGateway.Interfaces)"),
					)
				},
				SkipExtractID: true,
			},
		},
		Shutdown: func(ctx *CRUDTestContext, caller sacloud.APICaller) error {
			client := sacloud.NewMobileGatewayOp(caller)
			return client.Shutdown(ctx, testZone, ctx.ID, &sacloud.ShutdownOption{Force: true})
		},
		Delete: &CRUDTestDeleteFunc{
			Func: testMobileGatewayDelete,
		},
	})
}

func initMobileGatewayVariables() {

	iccid = os.Getenv("SAKURACLOUD_SIM_ICCID")
	passcode = os.Getenv("SAKURACLOUD_SIM_PASSCODE")

	createMobileGatewayParam = &sacloud.MobileGatewayCreateRequest{
		Name:        "libsacloud-mobile-gateway",
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
		Name:        "libsacloud-mobile-gateway-upd",
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
	iccid                       string
	passcode                    string
	createMobileGatewayParam    *sacloud.MobileGatewayCreateRequest
	createMobileGatewayExpected *sacloud.MobileGateway
	updateMobileGatewayParam    *sacloud.MobileGatewayUpdateRequest
	updateMobileGatewayExpected *sacloud.MobileGateway
)

func testMobileGatewayCreate(ctx *CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
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

func testMobileGatewayRead(ctx *CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewMobileGatewayOp(caller)
	return client.Read(ctx, testZone, ctx.ID)
}

func testMobileGatewayUpdate(ctx *CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewMobileGatewayOp(caller)
	return client.Update(ctx, testZone, ctx.ID, updateMobileGatewayParam)
}

func testMobileGatewayDelete(ctx *CRUDTestContext, caller sacloud.APICaller) error {
	client := sacloud.NewMobileGatewayOp(caller)
	return client.Delete(ctx, testZone, ctx.ID)
}
