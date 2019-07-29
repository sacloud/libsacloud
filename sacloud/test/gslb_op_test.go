package test

import (
	"testing"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/testutil"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

func TestGSLBOp_CRUD(t *testing.T) {
	testutil.Run(t, &testutil.CRUDTestCase{
		Parallel: true,

		SetupAPICallerFunc: singletonAPICaller,

		Create: &testutil.CRUDTestFunc{
			Func: testGSLBCreate,
			CheckFunc: testutil.AssertEqualWithExpected(&testutil.CRUDTestExpect{
				ExpectValue:  createGSLBExpected,
				IgnoreFields: ignoreGSLBFields,
			}),
		},

		Read: &testutil.CRUDTestFunc{
			Func: testGSLBRead,
			CheckFunc: testutil.AssertEqualWithExpected(&testutil.CRUDTestExpect{
				ExpectValue:  createGSLBExpected,
				IgnoreFields: ignoreGSLBFields,
			}),
		},

		Updates: []*testutil.CRUDTestFunc{
			{
				Func: testGSLBUpdate,
				CheckFunc: testutil.AssertEqualWithExpected(&testutil.CRUDTestExpect{
					ExpectValue:  updateGSLBExpected,
					IgnoreFields: ignoreGSLBFields,
				}),
			},
			{
				Func: testGSLBUpdateToMin,
				CheckFunc: testutil.AssertEqualWithExpected(&testutil.CRUDTestExpect{
					ExpectValue:  updateGSLBToMinExpected,
					IgnoreFields: ignoreGSLBFields,
				}),
			},
		},

		Delete: &testutil.CRUDTestDeleteFunc{
			Func: testGSLBDelete,
		},
	})
}

var (
	ignoreGSLBFields = []string{
		"ID",
		"Class",
		"SettingsHash",
		"FQDN",
		"CreatedAt",
		"ModifiedAt",
	}
	createGSLBParam = &sacloud.GSLBCreateRequest{
		Name:        testutil.ResourceName("gslb"),
		Description: "desc",
		Tags:        []string{"tag1", "tag2"},
		HealthCheck: &sacloud.GSLBHealthCheck{
			Protocol:     "http",
			HostHeader:   "usacloud.jp",
			Path:         "/index.html",
			ResponseCode: types.StringNumber(200),
		},
		DelayLoop:   20,
		Weighted:    types.StringTrue,
		SorryServer: "8.8.8.8",
		DestinationServers: []*sacloud.GSLBServer{
			{
				IPAddress: "192.2.0.1",
				Enabled:   types.StringTrue,
			},
			{
				IPAddress: "192.2.0.2",
				Enabled:   types.StringTrue,
			},
		},
	}
	createGSLBExpected = &sacloud.GSLB{
		Name:         createGSLBParam.Name,
		Description:  createGSLBParam.Description,
		Tags:         createGSLBParam.Tags,
		Availability: types.Availabilities.Available,
		DelayLoop:    createGSLBParam.DelayLoop,
		Weighted:     createGSLBParam.Weighted,
		HealthCheck:  createGSLBParam.HealthCheck,
		SorryServer:  createGSLBParam.SorryServer,
		DestinationServers: []*sacloud.GSLBServer{
			{
				IPAddress: "192.2.0.1",
				Enabled:   types.StringTrue,
			},
			{
				IPAddress: "192.2.0.2",
				Enabled:   types.StringTrue,
			},
		},
	}
	updateGSLBParam = &sacloud.GSLBUpdateRequest{
		Name:        testutil.ResourceName("gslb-upd"),
		Description: "desc-upd",
		Tags:        []string{"tag1-upd", "tag2-upd"},
		HealthCheck: &sacloud.GSLBHealthCheck{
			Protocol:     "https",
			HostHeader:   "upd.usacloud.jp",
			Path:         "/index-upd.html",
			ResponseCode: types.StringNumber(201),
		},
		DelayLoop:   21,
		Weighted:    types.StringTrue,
		SorryServer: "8.8.4.4",
		DestinationServers: []*sacloud.GSLBServer{
			{
				IPAddress: "192.2.0.11",
				Enabled:   types.StringFalse,
				Weight:    types.StringNumber(100),
			},
			{
				IPAddress: "192.2.0.21",
				Enabled:   types.StringFalse,
				Weight:    types.StringNumber(200),
			},
		},
		IconID: testIconID,
	}
	updateGSLBExpected = &sacloud.GSLB{
		Name:               updateGSLBParam.Name,
		Description:        updateGSLBParam.Description,
		Tags:               updateGSLBParam.Tags,
		Availability:       types.Availabilities.Available,
		DelayLoop:          updateGSLBParam.DelayLoop,
		Weighted:           updateGSLBParam.Weighted,
		HealthCheck:        updateGSLBParam.HealthCheck,
		SorryServer:        updateGSLBParam.SorryServer,
		DestinationServers: updateGSLBParam.DestinationServers,
		IconID:             testIconID,
	}
	updateGSLBToMinParam = &sacloud.GSLBUpdateRequest{
		Name: testutil.ResourceName("gslb-to-min"),
		HealthCheck: &sacloud.GSLBHealthCheck{
			Protocol: "ping",
		},
	}
	updateGSLBToMinExpected = &sacloud.GSLB{
		Name:         updateGSLBToMinParam.Name,
		DelayLoop:    10, // default value
		Availability: types.Availabilities.Available,
		HealthCheck:  updateGSLBToMinParam.HealthCheck,
	}
)

func testGSLBCreate(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewGSLBOp(caller)
	return client.Create(ctx, createGSLBParam)
}

func testGSLBRead(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewGSLBOp(caller)
	return client.Read(ctx, ctx.ID)
}

func testGSLBUpdate(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewGSLBOp(caller)
	return client.Update(ctx, ctx.ID, updateGSLBParam)
}

func testGSLBUpdateToMin(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewGSLBOp(caller)
	return client.Update(ctx, ctx.ID, updateGSLBToMinParam)
}

func testGSLBDelete(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) error {
	client := sacloud.NewGSLBOp(caller)
	return client.Delete(ctx, ctx.ID)
}
