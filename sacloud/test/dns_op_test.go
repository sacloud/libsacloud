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
				Func: testDNSPatch,
				CheckFunc: testutil.AssertEqualWithExpected(&testutil.CRUDTestExpect{
					ExpectValue:  patchDNSExpected,
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
	patchDNSParam = &sacloud.DNSPatchRequest{
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
	patchDNSExpected = &sacloud.DNS{
		Name:         createDNSParam.Name,
		Description:  createDNSParam.Description,
		Tags:         createDNSParam.Tags,
		Availability: types.Availabilities.Available,
		DNSZone:      createDNSParam.Name,
		Records:      patchDNSParam.Records,
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

func testDNSPatch(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewDNSOp(caller)
	return client.Patch(ctx, ctx.ID, patchDNSParam)
}

func testDNSUpdate(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewDNSOp(caller)
	return client.Update(ctx, ctx.ID, updateDNSParam)
}

func testDNSUpdateToMin(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewDNSOp(caller)
	return client.Update(ctx, ctx.ID, updateDNSToMinParam)
}

func testDNSDelete(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) error {
	client := sacloud.NewDNSOp(caller)
	return client.Delete(ctx, ctx.ID)
}
