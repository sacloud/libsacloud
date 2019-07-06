package test

import (
	"context"
	"testing"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

func TestNFSOpCRUD(t *testing.T) {
	Run(t, &CRUDTestCase{
		Parallel: true,

		SetupAPICallerFunc: singletonAPICaller,
		Setup: func(testContext *CRUDTestContext, caller sacloud.APICaller) error {
			swClient := sacloud.NewSwitchOp(caller)
			swCreateResult, err := swClient.Create(context.Background(), testZone, &sacloud.SwitchCreateRequest{
				Name: "libsacloud-switch-for-nfs",
			})
			if err != nil {
				return err
			}
			sw := swCreateResult.Switch

			testContext.Values["nfs/switch"] = sw.ID
			createNFSParam.SwitchID = sw.ID
			createNFSExpected.SwitchID = sw.ID
			updateNFSExpected.SwitchID = sw.ID
			return nil
		},

		Create: &CRUDTestFunc{
			Func: testNFSCreate,
			Expect: &CRUDTestExpect{
				ExpectValue:  createNFSExpected,
				IgnoreFields: ignoreNFSFields,
			},
		},

		Read: &CRUDTestFunc{
			Func: testNFSRead,
			Expect: &CRUDTestExpect{
				ExpectValue:  createNFSExpected,
				IgnoreFields: ignoreNFSFields,
			},
		},

		Update: &CRUDTestFunc{
			Func: testNFSUpdate,
			Expect: &CRUDTestExpect{
				ExpectValue:  updateNFSExpected,
				IgnoreFields: ignoreNFSFields,
			},
		},

		Shutdown: func(testContext *CRUDTestContext, caller sacloud.APICaller) error {
			client := sacloud.NewNFSOp(caller)
			return client.Shutdown(context.Background(), testZone, testContext.ID, &sacloud.ShutdownOption{Force: true})
		},

		Delete: &CRUDTestDeleteFunc{
			Func: testNFSDelete,
		},

		Cleanup: func(testContext *CRUDTestContext, caller sacloud.APICaller) error {

			switchID, ok := testContext.Values["nfs/switch"]
			if !ok {
				return nil
			}

			swClient := sacloud.NewSwitchOp(caller)
			return swClient.Delete(context.Background(), testZone, switchID.(types.ID))
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

func testNFSCreate(testContext *CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewNFSOp(caller)
	res, err := client.Create(context.Background(), testZone, createNFSParam)
	if err != nil {
		return nil, err
	}
	return res.NFS, nil
}

func testNFSRead(testContext *CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewNFSOp(caller)
	res, err := client.Read(context.Background(), testZone, testContext.ID)
	if err != nil {
		return nil, err
	}
	return res.NFS, nil
}

func testNFSUpdate(testContext *CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewNFSOp(caller)
	res, err := client.Update(context.Background(), testZone, testContext.ID, updateNFSParam)
	if err != nil {
		return nil, err
	}
	return res.NFS, nil
}

func testNFSDelete(testContext *CRUDTestContext, caller sacloud.APICaller) error {
	client := sacloud.NewNFSOp(caller)
	return client.Delete(context.Background(), testZone, testContext.ID)
}
