package sacloud

import (
	"context"
	"testing"
)

func TestSwitchOpCRUD(t *testing.T) {
	Test(t, &CRUDTestCase{
		Parallel: true,

		SetupAPICaller: singletonAPICaller,
		Create: &CRUDTestFunc{
			Func: testSwitchCreate,
			Expect: &CRUDTestExpect{
				ExpectValue:  createSwitchExpected,
				IgnoreFields: ignoreSwitchFields,
			},
		},
		Read: &CRUDTestFunc{
			Func: testSwitchRead,
			Expect: &CRUDTestExpect{
				ExpectValue:  createSwitchExpected,
				IgnoreFields: ignoreSwitchFields,
			},
		},
		Update: &CRUDTestFunc{
			Func: testSwitchUpdate,
			Expect: &CRUDTestExpect{
				ExpectValue:  updateSwitchExpected,
				IgnoreFields: ignoreSwitchFields,
			},
		},
		Delete: &CRUDTestDeleteFunc{
			Func: testSwitchDelete,
		},
	})
}

var (
	ignoreSwitchFields = []string{
		"ID",
		"Icon",
		"CreatedAt",
		"ModifiedAt",
		"Zone",
	}
	createSwitchParam = &SwitchCreateRequest{
		Name:           "libsacloud-v2-switch",
		Description:    "desc",
		Tags:           []string{"tag1", "tag2"},
		DefaultRoute:   "192.168.0.1",
		NetworkMaskLen: 24,
	}
	createSwitchExpected = &Switch{
		Name:           createSwitchParam.Name,
		Description:    createSwitchParam.Description,
		Tags:           createSwitchParam.Tags,
		DefaultRoute:   createSwitchParam.DefaultRoute,
		NetworkMaskLen: createSwitchParam.NetworkMaskLen,
	}
	updateSwitchParam = &SwitchUpdateRequest{
		Name:           "libsacloud-v2-switch-upd",
		Tags:           []string{"tag1-upd", "tag2-upd"},
		Description:    "desc-upd",
		DefaultRoute:   "192.168.0.2",
		NetworkMaskLen: 28,
	}
	updateSwitchExpected = &Switch{
		Name:           updateSwitchParam.Name,
		Description:    updateSwitchParam.Description,
		Tags:           updateSwitchParam.Tags,
		DefaultRoute:   updateSwitchParam.DefaultRoute,
		NetworkMaskLen: updateSwitchParam.NetworkMaskLen,
	}
)

func testSwitchCreate(testContext *CRUDTestContext, caller APICaller) (interface{}, error) {
	client := NewSwitchOp(caller)
	return client.Create(context.Background(), testZone, createSwitchParam)
}

func testSwitchRead(testContext *CRUDTestContext, caller APICaller) (interface{}, error) {
	client := NewSwitchOp(caller)
	return client.Read(context.Background(), testZone, testContext.ID)
}

func testSwitchUpdate(testContext *CRUDTestContext, caller APICaller) (interface{}, error) {
	client := NewSwitchOp(caller)
	return client.Update(context.Background(), testZone, testContext.ID, updateSwitchParam)
}

func testSwitchDelete(testContext *CRUDTestContext, caller APICaller) error {
	client := NewSwitchOp(caller)
	return client.Delete(context.Background(), testZone, testContext.ID)
}
