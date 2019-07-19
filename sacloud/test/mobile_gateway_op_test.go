package test

import (
	"context"
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
				Func: func(testContext *CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
					mgwOp := sacloud.NewMobileGatewayOp(caller)
					// shutdown
					if err := mgwOp.Shutdown(context.Background(), testZone, testContext.ID, &sacloud.ShutdownOption{Force: true}); err != nil {
						return nil, err
					}

					waiter := sacloud.WaiterForDown(func() (interface{}, error) {
						return mgwOp.Read(context.Background(), testZone, testContext.ID)
					})

					return waiter.WaitForState(context.Background())
				},
				SkipExtractID: true,
			},
			// connect to switch
			{
				Func: func(testContext *CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
					// prepare switch
					swOp := sacloud.NewSwitchOp(caller)
					sw, err := swOp.Create(context.Background(), testZone, &sacloud.SwitchCreateRequest{
						Name: "libsacloud-test-mobile-gateway",
					})
					if err != nil {
						return nil, err
					}

					// connect
					mgwOp := sacloud.NewMobileGatewayOp(caller)
					if err := mgwOp.ConnectToSwitch(context.Background(), testZone, testContext.ID, sw.ID); err != nil {
						return nil, err
					}

					return mgwOp.Read(context.Background(), testZone, testContext.ID)
				},
				CheckFunc: func(t TestT, testContext *CRUDTestContext, i interface{}) error {
					mgw := i.(*sacloud.MobileGateway)
					return DoAsserts(
						AssertLenFunc(t, mgw.Interfaces, 2, "len(MobileGateway.Interfaces)"),
					)
				},
				SkipExtractID: true,
			},
			// set IPAddress to eth1
			{
				Func: func(testContext *CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
					mgwOp := sacloud.NewMobileGatewayOp(caller)
					return mgwOp.Update(context.Background(), testZone, testContext.ID, &sacloud.MobileGatewayUpdateRequest{
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
				CheckFunc: func(t TestT, testContext *CRUDTestContext, i interface{}) error {
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
				Func: func(testContext *CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
					mgwOp := sacloud.NewMobileGatewayOp(caller)
					if err := mgwOp.SetDNS(context.Background(), testZone, testContext.ID, &sacloud.MobileGatewayDNSSetting{
						DNS1: "8.8.8.8",
						DNS2: "8.8.4.4",
					}); err != nil {
						return nil, err
					}
					return mgwOp.GetDNS(context.Background(), testZone, testContext.ID)
				},
				CheckFunc: func(t TestT, testContext *CRUDTestContext, i interface{}) error {
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
				Func: func(testContext *CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
					simOp := sacloud.NewSIMOp(caller)
					sim, err := simOp.Create(context.Background(), sacloud.APIDefaultZone, &sacloud.SIMCreateRequest{
						Name:     "libsacloud-switch-for-mobile-gateway",
						ICCID:    iccid,
						PassCode: passcode,
					})
					if err != nil {
						return nil, err
					}

					mgwOp := sacloud.NewMobileGatewayOp(caller)
					if err := mgwOp.AddSIM(context.Background(), testZone, testContext.ID, &sacloud.MobileGatewayAddSIMRequest{
						SIMID: sim.ID.String(),
					}); err != nil {
						return nil, err
					}

					testContext.Values["mobile-gateway/sim"] = sim.ID
					return mgwOp.ListSIM(context.Background(), testZone, testContext.ID)
				},
				CheckFunc: func(t TestT, testContext *CRUDTestContext, i interface{}) error {
					sims := i.([]*sacloud.MobileGatewaySIMInfo)
					return DoAsserts(
						AssertLenFunc(t, sims, 1, "len(SIM)"),
					)
				},
				SkipExtractID: true,
			},
			// SIMOp: Assign IP
			{
				Func: func(testContext *CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
					client := sacloud.NewSIMOp(caller)
					simID := testContext.Values["mobile-gateway/sim"].(types.ID)
					if err := client.AssignIP(context.Background(), sacloud.APIDefaultZone, simID, &sacloud.SIMAssignIPRequest{
						IP: "192.168.2.1",
					}); err != nil {
						return nil, err
					}
					return client.Status(context.Background(), sacloud.APIDefaultZone, simID)
				},
				CheckFunc: func(t TestT, testContext *CRUDTestContext, v interface{}) error {
					simInfo := v.(*sacloud.SIMInfo)
					return DoAsserts(
						AssertEqualFunc(t, "192.168.2.1", simInfo.IP, "SIMInfo.IP"),
					)
				},
				SkipExtractID: true,
			},
			// SIMOp: clear IP
			{
				Func: func(testContext *CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
					client := sacloud.NewSIMOp(caller)
					simID := testContext.Values["mobile-gateway/sim"].(types.ID)
					if err := client.ClearIP(context.Background(), sacloud.APIDefaultZone, simID); err != nil {
						return nil, err
					}
					return client.Status(context.Background(), sacloud.APIDefaultZone, simID)
				},
				CheckFunc: func(t TestT, testContext *CRUDTestContext, v interface{}) error {
					simInfo := v.(*sacloud.SIMInfo)
					return DoAsserts(
						AssertEmptyFunc(t, simInfo.IP, "SIMInfo.IP"),
					)
				},
				SkipExtractID: true,
			},

			// Get/Set SIMRoutes
			{
				Func: func(testContext *CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
					mgwOp := sacloud.NewMobileGatewayOp(caller)
					if err := mgwOp.SetSIMRoutes(context.Background(), testZone, testContext.ID, []*sacloud.MobileGatewaySIMRouteParam{
						{
							ResourceID: testContext.Values["mobile-gateway/sim"].(types.ID).String(),
							Prefix:     "192.168.3.0/24",
						},
					}); err != nil {
						return nil, err
					}
					return mgwOp.GetSIMRoutes(context.Background(), testZone, testContext.ID)
				},
				CheckFunc: func(t TestT, testContext *CRUDTestContext, i interface{}) error {
					routes := i.([]*sacloud.MobileGatewaySIMRoute)
					simID := testContext.Values["mobile-gateway/sim"].(types.ID)
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
				Func: func(testContext *CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
					mgwOp := sacloud.NewMobileGatewayOp(caller)
					if err := mgwOp.SetSIMRoutes(context.Background(), testZone, testContext.ID, []*sacloud.MobileGatewaySIMRouteParam{}); err != nil {
						return nil, err
					}
					return mgwOp.GetSIMRoutes(context.Background(), testZone, testContext.ID)
				},
				CheckFunc: func(t TestT, testContext *CRUDTestContext, i interface{}) error {
					routes := i.([]*sacloud.MobileGatewaySIMRoute)
					return DoAsserts(
						AssertLenFunc(t, routes, 0, "len(SIMRoutes)"),
					)
				},
				SkipExtractID: true,
			},

			// Get/Set TrafficConfig
			{
				Func: func(testContext *CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
					mgwOp := sacloud.NewMobileGatewayOp(caller)
					if err := mgwOp.SetTrafficConfig(context.Background(), testZone, testContext.ID, &sacloud.MobileGatewayTrafficControl{
						TrafficQuotaInMB:       10,
						BandWidthLimitInKbps:   20,
						EmailNotifyEnabled:     true,
						SlackNotifyEnabled:     true,
						SlackNotifyWebhooksURL: "https://hooks.slack.com/services/XXXXXXXXX/XXXXXXXXX/XXXXXXXXXXXXXXXXXXXXXXXX",
						AutoTrafficShaping:     true,
					}); err != nil {
						return nil, err
					}
					return mgwOp.GetTrafficConfig(context.Background(), testZone, testContext.ID)
				},
				CheckFunc: func(t TestT, testContext *CRUDTestContext, i interface{}) error {
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
				Func: func(testContext *CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
					mgwOp := sacloud.NewMobileGatewayOp(caller)
					return nil, mgwOp.DeleteTrafficConfig(context.Background(), testZone, testContext.ID)
				},
				SkipExtractID: true,
			},

			// Get TrafficStatus
			{
				Func: func(testContext *CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
					mgwOp := sacloud.NewMobileGatewayOp(caller)
					return mgwOp.TrafficStatus(context.Background(), testZone, testContext.ID)
				},
				CheckFunc: func(t TestT, testContext *CRUDTestContext, i interface{}) error {
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
				Func: func(testContext *CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
					simID := testContext.Values["mobile-gateway/sim"].(types.ID)
					mgwOp := sacloud.NewMobileGatewayOp(caller)
					if err := mgwOp.DeleteSIM(context.Background(), testZone, testContext.ID, simID); err != nil {
						return nil, err
					}

					simOp := sacloud.NewSIMOp(caller)
					if err := simOp.Delete(context.Background(), testZone, simID); err != nil {
						return nil, err
					}

					return mgwOp.ListSIM(context.Background(), testZone, testContext.ID)
				},
				CheckFunc: func(t TestT, testContext *CRUDTestContext, i interface{}) error {
					sims := i.([]*sacloud.MobileGatewaySIMInfo)
					return DoAsserts(
						AssertLenFunc(t, sims, 0, "len(SIM)"),
					)
				},
				SkipExtractID: true,
			},
			// disconnect from switch
			{
				Func: func(testContext *CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
					mgwOp := sacloud.NewMobileGatewayOp(caller)
					if err := mgwOp.DisconnectFromSwitch(context.Background(), testZone, testContext.ID); err != nil {
						return nil, err
					}
					return mgwOp.Read(context.Background(), testZone, testContext.ID)
				},
				CheckFunc: func(t TestT, testContext *CRUDTestContext, i interface{}) error {
					mgw := i.(*sacloud.MobileGateway)
					return DoAsserts(
						AssertLenFunc(t, mgw.Interfaces, 1, "len(MobileGateway.Interfaces)"),
					)
				},
				SkipExtractID: true,
			},
		},
		Shutdown: func(testContext *CRUDTestContext, caller sacloud.APICaller) error {
			client := sacloud.NewMobileGatewayOp(caller)
			return client.Shutdown(context.Background(), testZone, testContext.ID, &sacloud.ShutdownOption{Force: true})
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
		PlanID:      types.ID(1),
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
		PlanID:       createMobileGatewayParam.PlanID,
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
		PlanID:       createMobileGatewayParam.PlanID,
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

func testMobileGatewayCreate(testContext *CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewMobileGatewayOp(caller)
	v, err := client.Create(context.Background(), testZone, createMobileGatewayParam)
	if err != nil {
		return nil, err
	}
	value, err := sacloud.WaiterForReady(func() (interface{}, error) {
		return client.Read(context.Background(), testZone, v.ID)
	}).WaitForState(context.Background())
	if err != nil {
		return nil, err
	}
	if err := client.Boot(context.Background(), testZone, v.ID); err != nil {
		return nil, err
	}
	return value, nil
}

func testMobileGatewayRead(testContext *CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewMobileGatewayOp(caller)
	return client.Read(context.Background(), testZone, testContext.ID)
}

func testMobileGatewayUpdate(testContext *CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewMobileGatewayOp(caller)
	return client.Update(context.Background(), testZone, testContext.ID, updateMobileGatewayParam)
}

func testMobileGatewayDelete(testContext *CRUDTestContext, caller sacloud.APICaller) error {
	client := sacloud.NewMobileGatewayOp(caller)
	return client.Delete(context.Background(), testZone, testContext.ID)
}
