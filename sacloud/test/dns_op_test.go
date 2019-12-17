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

package test

import (
	"testing"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/testutil"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

func TestDNSOp_CRUD(t *testing.T) {
	testutil.Run(t, &testutil.CRUDTestCase{
		Parallel: true,

		SetupAPICallerFunc: singletonAPICaller,

		Create: &testutil.CRUDTestFunc{
			Func: testDNSCreate,
			CheckFunc: testutil.AssertEqualWithExpected(&testutil.CRUDTestExpect{
				ExpectValue:  createDNSExpected,
				IgnoreFields: ignoreDNSFields,
			}),
		},

		Read: &testutil.CRUDTestFunc{
			Func: testDNSRead,
			CheckFunc: testutil.AssertEqualWithExpected(&testutil.CRUDTestExpect{
				ExpectValue:  createDNSExpected,
				IgnoreFields: ignoreDNSFields,
			}),
		},

		Updates: []*testutil.CRUDTestFunc{
			{
				Func: testDNSUpdateSettings,
				CheckFunc: testutil.AssertEqualWithExpected(&testutil.CRUDTestExpect{
					ExpectValue:  updateDNSSettingsExpected,
					IgnoreFields: ignoreDNSFields,
				}),
			},
			{
				Func: testDNSUpdate,
				CheckFunc: testutil.AssertEqualWithExpected(&testutil.CRUDTestExpect{
					ExpectValue:  updateDNSExpected,
					IgnoreFields: ignoreDNSFields,
				}),
			},
			{
				Func: testDNSUpdateToMin,
				CheckFunc: testutil.AssertEqualWithExpected(&testutil.CRUDTestExpect{
					ExpectValue:  updateDNSToMinExpected,
					IgnoreFields: ignoreDNSFields,
				}),
			},
		},

		Delete: &testutil.CRUDTestDeleteFunc{
			Func: testDNSDelete,
		},
	})
}

var (
	ignoreDNSFields = []string{
		"ID",
		"Class",
		"SettingsHash",
		"FQDN",
		"CreatedAt",
		"ModifiedAt",
		"DNSNameServers",
	}
	createDNSParam = &sacloud.DNSCreateRequest{
		Name:        testutil.ResourceName("dns.com"),
		Description: "desc",
		Tags:        []string{"tag1", "tag2"},
		Records: []*sacloud.DNSRecord{
			{
				Name:  "host1",
				Type:  types.DNSRecordTypes.A,
				RData: "192.0.2.1",
			},
			{
				Name:  "host2",
				Type:  types.DNSRecordTypes.A,
				RData: "192.0.2.2",
			},
		},
	}
	createDNSExpected = &sacloud.DNS{
		Name:         createDNSParam.Name,
		Description:  createDNSParam.Description,
		Tags:         createDNSParam.Tags,
		Availability: types.Availabilities.Available,
		DNSZone:      createDNSParam.Name,
		Records:      createDNSParam.Records,
	}
	updateDNSSettingsParam = &sacloud.DNSUpdateSettingsRequest{
		Records: []*sacloud.DNSRecord{
			{
				Name:  "host1",
				Type:  types.DNSRecordTypes.A,
				RData: "192.0.2.11",
			},
			{
				Name:  "host2",
				Type:  types.DNSRecordTypes.A,
				RData: "192.0.2.12",
			},
			{
				Name:  "host3",
				Type:  types.DNSRecordTypes.A,
				RData: "192.0.2.13",
			},
			{
				Name:  "host4",
				Type:  types.DNSRecordTypes.A,
				RData: "192.0.2.14",
			},
		},
	}
	updateDNSSettingsExpected = &sacloud.DNS{
		Name:         createDNSParam.Name,
		Description:  createDNSParam.Description,
		Tags:         createDNSParam.Tags,
		Availability: types.Availabilities.Available,
		DNSZone:      createDNSParam.Name,
		Records:      updateDNSSettingsParam.Records,
	}
	updateDNSParam = &sacloud.DNSUpdateRequest{
		Description: "desc-upd",
		Tags:        []string{"tag1-upd", "tag2-upd"},
		IconID:      testIconID,
		Records: []*sacloud.DNSRecord{
			{
				Name:  "host1",
				Type:  types.DNSRecordTypes.A,
				RData: "192.0.2.11",
			},
			{
				Name:  "host2",
				Type:  types.DNSRecordTypes.A,
				RData: "192.0.2.12",
			},
			{
				Name:  "host3",
				Type:  types.DNSRecordTypes.A,
				RData: "192.0.2.13",
			},
		},
	}
	updateDNSExpected = &sacloud.DNS{
		Name:         createDNSParam.Name,
		Description:  updateDNSParam.Description,
		IconID:       testIconID,
		Tags:         updateDNSParam.Tags,
		Availability: types.Availabilities.Available,
		DNSZone:      createDNSParam.Name,
		Records:      updateDNSParam.Records,
	}
	updateDNSToMinParam    = &sacloud.DNSUpdateRequest{}
	updateDNSToMinExpected = &sacloud.DNS{
		Name:         createDNSParam.Name,
		Availability: types.Availabilities.Available,
		DNSZone:      createDNSParam.Name,
	}
)

func testDNSCreate(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewDNSOp(caller)
	return client.Create(ctx, createDNSParam)
}

func testDNSRead(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewDNSOp(caller)
	return client.Read(ctx, ctx.ID)
}

func testDNSUpdate(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewDNSOp(caller)
	return client.Update(ctx, ctx.ID, updateDNSParam)
}

func testDNSUpdateSettings(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewDNSOp(caller)
	return client.UpdateSettings(ctx, ctx.ID, updateDNSSettingsParam)
}

func testDNSUpdateToMin(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewDNSOp(caller)
	return client.Update(ctx, ctx.ID, updateDNSToMinParam)
}

func testDNSDelete(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) error {
	client := sacloud.NewDNSOp(caller)
	return client.Delete(ctx, ctx.ID)
}
