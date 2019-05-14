package sacloud

import (
	"context"
	"testing"

	"github.com/sacloud/libsacloud-v2/sacloud/types"
)

func TestNFSOpCRUD(t *testing.T) {
	Test(t, &CRUDTestCase{
		Parallel: true,

		SetupAPICaller: singletonAPICaller,
		Setup: func(testContext *CRUDTestContext, caller APICaller) error {
			swClient := NewSwitchOp(caller)
			sw, err := swClient.Create(context.Background(), testZone, &SwitchCreateRequest{
				Name: "libsacloud-v2-switch-for-nfs",
			})
			if err != nil {
				return err
			}

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

		Shutdown: func(testContext *CRUDTestContext, caller APICaller) error {
			client := NewNFSOp(caller)
			return client.Shutdown(context.Background(), testZone, testContext.ID, &ShutdownOption{Force: true})
		},

		Delete: &CRUDTestDeleteFunc{
			Func: testNFSDelete,
		},

		Cleanup: func(testContext *CRUDTestContext, caller APICaller) error {

			switchID, ok := testContext.Values["nfs/switch"]
			if !ok {
				return nil
			}

			swClient := NewSwitchOp(caller)
			return swClient.Delete(context.Background(), testZone, switchID.(types.ID))
		},
	})
}

var (
	ignoreNFSFields = []string{
		"ID",
		"Class",
		"InstanceHostName",
		"InstanceHostInfoURL",
		"InstanceStatusChangedAt",
		"Interfaces",
		"Switch",
		"ZoneID",
		"Icon",
		"CreatedAt",
		"ModifiedAt",
	}
	createNFSParam = &NFSCreateRequest{
		PlanID:         types.ID(1001012001),
		IPAddress:      "192.168.0.11",
		NetworkMaskLen: 24,
		DefaultRoute:   "192.168.0.1",
		Name:           "libsacloud-v2-nfs",
		Description:    "desc",
		Tags:           []string{"tag1", "tag2"},
	}
	createNFSExpected = &NFS{
		Name:           createNFSParam.Name,
		Description:    createNFSParam.Description,
		Tags:           createNFSParam.Tags,
		Availability:   types.Availabilities.Available,
		InstanceStatus: types.ServerInstanceStatuses.Up,
		PlanID:         createNFSParam.PlanID,
		DefaultRoute:   createNFSParam.DefaultRoute,
		NetworkMaskLen: createNFSParam.NetworkMaskLen,
		IPAddresses:    []string{createNFSParam.IPAddress},
	}
	updateNFSParam = &NFSUpdateRequest{
		Name:        "libsacloud-v2-nfs-upd",
		Tags:        []string{"tag1-upd", "tag2-upd"},
		Description: "desc-upd",
	}
	updateNFSExpected = &NFS{
		Name:           updateNFSParam.Name,
		Description:    updateNFSParam.Description,
		Tags:           updateNFSParam.Tags,
		Availability:   types.Availabilities.Available,
		InstanceStatus: types.ServerInstanceStatuses.Up,
		PlanID:         createNFSParam.PlanID,
		DefaultRoute:   createNFSParam.DefaultRoute,
		NetworkMaskLen: createNFSParam.NetworkMaskLen,
		IPAddresses:    []string{createNFSParam.IPAddress},
	}
)

func testNFSCreate(testContext *CRUDTestContext, caller APICaller) (interface{}, error) {
	client := NewNFSOp(caller)
	return client.Create(context.Background(), testZone, createNFSParam)
}

func testNFSRead(testContext *CRUDTestContext, caller APICaller) (interface{}, error) {
	client := NewNFSOp(caller)
	return client.Read(context.Background(), testZone, testContext.ID)
}

func testNFSUpdate(testContext *CRUDTestContext, caller APICaller) (interface{}, error) {
	client := NewNFSOp(caller)
	return client.Update(context.Background(), testZone, testContext.ID, updateNFSParam)
}

func testNFSDelete(testContext *CRUDTestContext, caller APICaller) error {
	client := NewNFSOp(caller)
	return client.Delete(context.Background(), testZone, testContext.ID)
}
