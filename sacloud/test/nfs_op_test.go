package test

import (
	"testing"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

func TestNFSOpCRUD(t *testing.T) {
	Run(t, &CRUDTestCase{
		Parallel: true,

		SetupAPICallerFunc: singletonAPICaller,
		Setup: func(ctx *CRUDTestContext, caller sacloud.APICaller) error {
			swClient := sacloud.NewSwitchOp(caller)
			sw, err := swClient.Create(ctx, testZone, &sacloud.SwitchCreateRequest{
				Name: "libsacloud-switch-for-nfs",
			})
			if err != nil {
				return err
			}

			ctx.Values["nfs/switch"] = sw.ID
			createNFSParam.SwitchID = sw.ID
			createNFSExpected.SwitchID = sw.ID
			updateNFSExpected.SwitchID = sw.ID
			return nil
		},

		Create: &CRUDTestFunc{
			Func: testNFSCreate,
			CheckFunc: AssertEqualWithExpected(&CRUDTestExpect{
				ExpectValue:  createNFSExpected,
				IgnoreFields: ignoreNFSFields,
			}),
		},

		Read: &CRUDTestFunc{
			Func: testNFSRead,
			CheckFunc: AssertEqualWithExpected(&CRUDTestExpect{
				ExpectValue:  createNFSExpected,
				IgnoreFields: ignoreNFSFields,
			}),
		},

		Updates: []*CRUDTestFunc{
			{
				Func: testNFSUpdate,
				CheckFunc: AssertEqualWithExpected(&CRUDTestExpect{
					ExpectValue:  updateNFSExpected,
					IgnoreFields: ignoreNFSFields,
				}),
			},
		},

		Shutdown: func(ctx *CRUDTestContext, caller sacloud.APICaller) error {
			client := sacloud.NewNFSOp(caller)
			return client.Shutdown(ctx, testZone, ctx.ID, &sacloud.ShutdownOption{Force: true})
		},

		Delete: &CRUDTestDeleteFunc{
			Func: testNFSDelete,
		},

		Cleanup: func(ctx *CRUDTestContext, caller sacloud.APICaller) error {

			switchID, ok := ctx.Values["nfs/switch"]
			if !ok {
				return nil
			}

			swClient := sacloud.NewSwitchOp(caller)
			return swClient.Delete(ctx, testZone, switchID.(types.ID))
		},
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
		"IconID",
		"CreatedAt",
		"ModifiedAt",
	}
	createNFSParam = &sacloud.NFSCreateRequest{
		PlanID:         types.ID(1001012001),
		IPAddresses:    []string{"192.168.0.11"},
		NetworkMaskLen: 24,
		DefaultRoute:   "192.168.0.1",
		Name:           "libsacloud-nfs",
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
		Name:        "libsacloud-nfs-upd",
		Tags:        []string{"tag1-upd", "tag2-upd"},
		Description: "desc-upd",
	}
	updateNFSExpected = &sacloud.NFS{
		Name:           updateNFSParam.Name,
		Description:    updateNFSParam.Description,
		Tags:           updateNFSParam.Tags,
		PlanID:         createNFSParam.PlanID,
		DefaultRoute:   createNFSParam.DefaultRoute,
		NetworkMaskLen: createNFSParam.NetworkMaskLen,
		IPAddresses:    createNFSParam.IPAddresses,
	}
)

func testNFSCreate(ctx *CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewNFSOp(caller)
	return client.Create(ctx, testZone, createNFSParam)
}

func testNFSRead(ctx *CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewNFSOp(caller)
	return client.Read(ctx, testZone, ctx.ID)
}

func testNFSUpdate(ctx *CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewNFSOp(caller)
	return client.Update(ctx, testZone, ctx.ID, updateNFSParam)
}

func testNFSDelete(ctx *CRUDTestContext, caller sacloud.APICaller) error {
	client := sacloud.NewNFSOp(caller)
	return client.Delete(ctx, testZone, ctx.ID)
}
