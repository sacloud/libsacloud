// Copyright 2016-2022 The Libsacloud Authors
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
	"testing"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/testutil"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

func TestGSLBOp_CRUD(t *testing.T) {
	testutil.RunCRUD(t, &testutil.CRUDTestCase{
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
				Func: testGSLBUpdateSettings,
				CheckFunc: testutil.AssertEqualWithExpected(&testutil.CRUDTestExpect{
					ExpectValue:  updateGSLBSettingsExpected,
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
			Protocol:     types.GSLBHealthCheckProtocols.HTTP,
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
			Protocol:     types.GSLBHealthCheckProtocols.HTTPS,
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
	updateGSLBSettingsParam = &sacloud.GSLBUpdateSettingsRequest{
		HealthCheck: &sacloud.GSLBHealthCheck{
			Protocol:     types.GSLBHealthCheckProtocols.HTTP,
			HostHeader:   "upd2.usacloud.jp",
			Path:         "/index-upd2.html",
			ResponseCode: types.StringNumber(202),
		},
		DelayLoop:   22,
		Weighted:    types.StringFalse,
		SorryServer: "1.1.1.1",
		DestinationServers: []*sacloud.GSLBServer{
			{
				IPAddress: "192.2.0.12",
				Enabled:   types.StringFalse,
				Weight:    types.StringNumber(100),
			},
			{
				IPAddress: "192.2.0.22",
				Enabled:   types.StringFalse,
				Weight:    types.StringNumber(200),
			},
		},
	}
	updateGSLBSettingsExpected = &sacloud.GSLB{
		Name:               updateGSLBParam.Name,
		Description:        updateGSLBParam.Description,
		Tags:               updateGSLBParam.Tags,
		Availability:       types.Availabilities.Available,
		DelayLoop:          updateGSLBSettingsParam.DelayLoop,
		Weighted:           updateGSLBSettingsParam.Weighted,
		HealthCheck:        updateGSLBSettingsParam.HealthCheck,
		SorryServer:        updateGSLBSettingsParam.SorryServer,
		DestinationServers: updateGSLBSettingsParam.DestinationServers,
		IconID:             testIconID,
	}
	updateGSLBToMinParam = &sacloud.GSLBUpdateRequest{
		Name: testutil.ResourceName("gslb-to-min"),
		HealthCheck: &sacloud.GSLBHealthCheck{
			Protocol: types.GSLBHealthCheckProtocols.Ping,
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

func testGSLBUpdateSettings(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewGSLBOp(caller)
	return client.UpdateSettings(ctx, ctx.ID, updateGSLBSettingsParam)
}

func testGSLBUpdateToMin(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewGSLBOp(caller)
	return client.Update(ctx, ctx.ID, updateGSLBToMinParam)
}

func testGSLBDelete(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) error {
	client := sacloud.NewGSLBOp(caller)
	return client.Delete(ctx, ctx.ID)
}
