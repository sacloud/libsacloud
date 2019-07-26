package test

import (
	"testing"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/testutil"
)

func TestPrivateHostOp_CRUD(t *testing.T) {
	testutil.Run(t, &testutil.CRUDTestCase{
		Parallel:           true,
		IgnoreStartupWait:  true,
		SetupAPICallerFunc: singletonAPICaller,
		Setup: func(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) error {
			planOp := sacloud.NewPrivateHostPlanOp(caller)
			searched, err := planOp.Find(ctx, privateHostTestZone, nil)
			if err != nil {
				return err
			}
			planID := searched.PrivateHostPlans[0].ID
			createPrivateHostParam.PlanID = planID
			createPrivateHostExpected.PlanID = planID
			updatePrivateHostExpected.PlanID = planID
			updatePrivateHostToMinExpected.PlanID = planID
			return nil
		},
		Create: &testutil.CRUDTestFunc{
			Func: testPrivateHostCreate,
			CheckFunc: testutil.AssertEqualWithExpected(&testutil.CRUDTestExpect{
				ExpectValue:  createPrivateHostExpected,
				IgnoreFields: ignorePrivateHostFields,
			}),
		},
		Read: &testutil.CRUDTestFunc{
			Func: testPrivateHostRead,
			CheckFunc: testutil.AssertEqualWithExpected(&testutil.CRUDTestExpect{
				ExpectValue:  createPrivateHostExpected,
				IgnoreFields: ignorePrivateHostFields,
			}),
		},
		Updates: []*testutil.CRUDTestFunc{
			{
				Func: testPrivateHostUpdate,
				CheckFunc: testutil.AssertEqualWithExpected(&testutil.CRUDTestExpect{
					ExpectValue:  updatePrivateHostExpected,
					IgnoreFields: ignorePrivateHostFields,
				}),
			},
			{
				Func: testPrivateHostUpdateToMin,
				CheckFunc: testutil.AssertEqualWithExpected(&testutil.CRUDTestExpect{
					ExpectValue:  updatePrivateHostToMinExpected,
					IgnoreFields: ignorePrivateHostFields,
				}),
			},
		},
		Delete: &testutil.CRUDTestDeleteFunc{
			Func: testPrivateHostDelete,
		},
	})
}

var (
	privateHostTestZone = "tk1a"

	ignorePrivateHostFields = []string{
		"ID",
		"CreatedAt",
		"PlanName",
		"PlanClass",
		"HostName",
		"CPU",
		"MemoryMB",
	}

	createPrivateHostParam = &sacloud.PrivateHostCreateRequest{
		Name:        "libsacloud-private-host",
		Description: "libsacloud-private-host",
		Tags:        []string{"tag1", "tag2"},
	}
	createPrivateHostExpected = &sacloud.PrivateHost{
		Name:             createPrivateHostParam.Name,
		Description:      createPrivateHostParam.Description,
		Tags:             createPrivateHostParam.Tags,
		CPU:              224,
		AssignedCPU:      0,
		AssignedMemoryMB: 0,
	}
	updatePrivateHostParam = &sacloud.PrivateHostUpdateRequest{
		Name:        "libsacloud-private-host-upd",
		Description: "libsacloud-private-host-upd",
		Tags:        []string{"tag1-upd", "tag2-upd"},
		IconID:      testIconID,
	}
	updatePrivateHostExpected = &sacloud.PrivateHost{
		Name:             updatePrivateHostParam.Name,
		Description:      updatePrivateHostParam.Description,
		Tags:             updatePrivateHostParam.Tags,
		CPU:              224,
		AssignedCPU:      0,
		AssignedMemoryMB: 0,
		IconID:           testIconID,
	}
	updatePrivateHostToMinParam = &sacloud.PrivateHostUpdateRequest{
		Name: "libsacloud-private-host-to-min",
	}
	updatePrivateHostToMinExpected = &sacloud.PrivateHost{
		Name:             updatePrivateHostToMinParam.Name,
		CPU:              224,
		AssignedCPU:      0,
		AssignedMemoryMB: 0,
	}
)

func testPrivateHostCreate(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewPrivateHostOp(caller)
	return client.Create(ctx, privateHostTestZone, createPrivateHostParam)
}

func testPrivateHostRead(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewPrivateHostOp(caller)
	return client.Read(ctx, privateHostTestZone, ctx.ID)
}

func testPrivateHostUpdate(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewPrivateHostOp(caller)
	return client.Update(ctx, privateHostTestZone, ctx.ID, updatePrivateHostParam)
}

func testPrivateHostUpdateToMin(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewPrivateHostOp(caller)
	return client.Update(ctx, privateHostTestZone, ctx.ID, updatePrivateHostToMinParam)
}

func testPrivateHostDelete(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) error {
	client := sacloud.NewPrivateHostOp(caller)
	return client.Delete(ctx, privateHostTestZone, ctx.ID)
}
