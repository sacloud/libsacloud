package test

import (
	"testing"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/testutil"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
	"github.com/sacloud/libsacloud/v2/utils/nfs"
)

func TestNFSOp_CRUD(t *testing.T) {
	testutil.Run(t, &testutil.CRUDTestCase{
		Parallel: true,

		SetupAPICallerFunc: singletonAPICaller,
		Setup: func(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) error {
			setupSwitchFunc("nfs",
				createNFSParam,
				createNFSExpected,
				updateNFSExpected,
				updateNFSToMinExpected,
			)(ctx, caller)

			// find plan id
			planID, err := nfs.FindNFSPlanID(ctx, sacloud.NewNoteOp(caller), types.NFSPlans.HDD, types.NFSHDDSizes.Size100GB)
			if err != nil {
				return err
			}
			createNFSParam.PlanID = planID
			createNFSExpected.PlanID = planID
			updateNFSExpected.PlanID = planID
			updateNFSToMinExpected.PlanID = planID
			return nil
		},

		Create: &testutil.CRUDTestFunc{
			Func: testNFSCreate,
			CheckFunc: testutil.AssertEqualWithExpected(&testutil.CRUDTestExpect{
				ExpectValue:  createNFSExpected,
				IgnoreFields: ignoreNFSFields,
			}),
		},

		Read: &testutil.CRUDTestFunc{
			Func: testNFSRead,
			CheckFunc: testutil.AssertEqualWithExpected(&testutil.CRUDTestExpect{
				ExpectValue:  createNFSExpected,
				IgnoreFields: ignoreNFSFields,
			}),
		},

		Updates: []*testutil.CRUDTestFunc{
			{
				Func: testNFSUpdate,
				CheckFunc: testutil.AssertEqualWithExpected(&testutil.CRUDTestExpect{
					ExpectValue:  updateNFSExpected,
					IgnoreFields: ignoreNFSFields,
				}),
			},
			{
				Func: testNFSUpdateToMin,
				CheckFunc: testutil.AssertEqualWithExpected(&testutil.CRUDTestExpect{
					ExpectValue:  updateNFSToMinExpected,
					IgnoreFields: ignoreNFSFields,
				}),
			},
		},

		Shutdown: func(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) error {
			client := sacloud.NewNFSOp(caller)
			return client.Shutdown(ctx, testZone, ctx.ID, &sacloud.ShutdownOption{Force: true})
		},

		Delete: &testutil.CRUDTestDeleteFunc{
			Func: testNFSDelete,
		},

		Cleanup: cleanupSwitchFunc("nfs"),
	})
}

var (
	ignoreNFSFields = []string{
		"ID",
		"Class",
		"Availability",
		"InstanceStatus",
		"InstanceHostName",
		"InstanceHostInfoURL",
		"InstanceStatusChangedAt",
		"Interfaces",
		"Switch",
		"ZoneID",
		"CreatedAt",
		"ModifiedAt",
	}
	createNFSParam = &sacloud.NFSCreateRequest{
		// PlanID:      type.ID(0), // プランIDはSetUpで設定する
		IPAddresses:    []string{"192.168.0.11"},
		NetworkMaskLen: 24,
		DefaultRoute:   "192.168.0.1",
		Name:           testutil.ResourceName("nfs"),
		Description:    "desc",
		Tags:           []string{"tag1", "tag2"},
	}
	createNFSExpected = &sacloud.NFS{
		Name:           createNFSParam.Name,
		Description:    createNFSParam.Description,
		Tags:           createNFSParam.Tags,
		PlanID:         createNFSParam.PlanID,
		DefaultRoute:   createNFSParam.DefaultRoute,
		NetworkMaskLen: createNFSParam.NetworkMaskLen,
		IPAddresses:    createNFSParam.IPAddresses,
	}
	updateNFSParam = &sacloud.NFSUpdateRequest{
		Name:        testutil.ResourceName("nfs-upd"),
		Tags:        []string{"tag1-upd", "tag2-upd"},
		Description: "desc-upd",
		IconID:      testIconID,
	}
	updateNFSExpected = &sacloud.NFS{
		Name:           updateNFSParam.Name,
		Description:    updateNFSParam.Description,
		Tags:           updateNFSParam.Tags,
		DefaultRoute:   createNFSParam.DefaultRoute,
		NetworkMaskLen: createNFSParam.NetworkMaskLen,
		IPAddresses:    createNFSParam.IPAddresses,
		IconID:         testIconID,
	}
	updateNFSToMinParam = &sacloud.NFSUpdateRequest{
		Name: testutil.ResourceName("nfs-to-min"),
	}
	updateNFSToMinExpected = &sacloud.NFS{
		Name:           updateNFSToMinParam.Name,
		DefaultRoute:   createNFSParam.DefaultRoute,
		NetworkMaskLen: createNFSParam.NetworkMaskLen,
		IPAddresses:    createNFSParam.IPAddresses,
	}
)

func testNFSCreate(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewNFSOp(caller)
	return client.Create(ctx, testZone, createNFSParam)
}

func testNFSRead(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewNFSOp(caller)
	return client.Read(ctx, testZone, ctx.ID)
}

func testNFSUpdate(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewNFSOp(caller)
	return client.Update(ctx, testZone, ctx.ID, updateNFSParam)
}

func testNFSUpdateToMin(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewNFSOp(caller)
	return client.Update(ctx, testZone, ctx.ID, updateNFSToMinParam)
}

func testNFSDelete(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) error {
	client := sacloud.NewNFSOp(caller)
	return client.Delete(ctx, testZone, ctx.ID)
}
