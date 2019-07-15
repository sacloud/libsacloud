package test

import (
	"context"
	"testing"

	"github.com/sacloud/libsacloud/v2/sacloud"
)

func TestPrivateHostOpCRUD(t *testing.T) {
	Run(t, &CRUDTestCase{
		Parallel:           true,
		IgnoreStartupWait:  true,
		SetupAPICallerFunc: singletonAPICaller,
		Setup: func(testContext *CRUDTestContext, caller sacloud.APICaller) error {
			planOp := sacloud.NewPrivateHostPlanOp(caller)
			searched, err := planOp.Find(context.Background(), privateHostTestZone, nil)
			if err != nil {
				return err
			}
			planID := searched.PrivateHostPlans[0].ID
			createPrivateHostParam.PlanID = planID
			createPrivateHostExpected.PlanID = planID
			updatePrivateHostExpected.PlanID = planID
			return nil
		},
		Create: &CRUDTestFunc{
			Func: testPrivateHostCreate,
			CheckFunc: AssertEqualWithExpected(&CRUDTestExpect{
				ExpectValue:  createPrivateHostExpected,
				IgnoreFields: ignorePrivateHostFields,
			}),
		},
		Read: &CRUDTestFunc{
			Func: testPrivateHostRead,
			CheckFunc: AssertEqualWithExpected(&CRUDTestExpect{
				ExpectValue:  createPrivateHostExpected,
				IgnoreFields: ignorePrivateHostFields,
			}),
		},
		Updates: []*CRUDTestFunc{
			{
				Func: testPrivateHostUpdate,
				CheckFunc: AssertEqualWithExpected(&CRUDTestExpect{
					ExpectValue:  updatePrivateHostExpected,
					IgnoreFields: ignorePrivateHostFields,
				}),
			},
		},
		Delete: &CRUDTestDeleteFunc{
			Func: testPrivateHostDelete,
		},
	})
}

var (
	privateHostTestZone = "tk1a"

	ignorePrivateHostFields = []string{
		"ID",
		"IconID",
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
	}
	updatePrivateHostExpected = &sacloud.PrivateHost{
		Name:             updatePrivateHostParam.Name,
		Description:      updatePrivateHostParam.Description,
		Tags:             updatePrivateHostParam.Tags,
		CPU:              224,
		AssignedCPU:      0,
		AssignedMemoryMB: 0,
	}
)

func testPrivateHostCreate(testContext *CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewPrivateHostOp(caller)
	return client.Create(context.Background(), privateHostTestZone, createPrivateHostParam)
}

func testPrivateHostRead(testContext *CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewPrivateHostOp(caller)
	return client.Read(context.Background(), privateHostTestZone, testContext.ID)
}

func testPrivateHostUpdate(testContext *CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewPrivateHostOp(caller)
	return client.Update(context.Background(), privateHostTestZone, testContext.ID, updatePrivateHostParam)
}

func testPrivateHostDelete(testContext *CRUDTestContext, caller sacloud.APICaller) error {
	client := sacloud.NewPrivateHostOp(caller)
	return client.Delete(context.Background(), privateHostTestZone, testContext.ID)
}
