package test

import (
	"context"
	"testing"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

func TestNoteOpCRUD(t *testing.T) {
	Run(t, &CRUDTestCase{
		Parallel:           true,
		IgnoreStartupWait:  true,
		SetupAPICallerFunc: singletonAPICaller,
		Create: &CRUDTestFunc{
			Func: testNoteCreate,
			CheckFunc: AssertEqualWithExpected(&CRUDTestExpect{
				ExpectValue:  createNoteExpected,
				IgnoreFields: []string{"ID", "CreatedAt", "ModifiedAt"},
			}),
		},
		Read: &CRUDTestFunc{
			Func: testNoteRead,
			CheckFunc: AssertEqualWithExpected(&CRUDTestExpect{
				ExpectValue:  createNoteExpected,
				IgnoreFields: []string{"ID", "CreatedAt", "ModifiedAt"},
			}),
		},
		Updates: []*CRUDTestFunc{
			{
				Func: testNoteUpdate,
				CheckFunc: AssertEqualWithExpected(&CRUDTestExpect{
					ExpectValue:  updateNoteExpected,
					IgnoreFields: []string{"ID", "CreatedAt", "ModifiedAt"},
				}),
			},
		},
		Delete: &CRUDTestDeleteFunc{
			Func: testNoteDelete,
		},
	})
}

var (
	createNoteParam = &sacloud.NoteCreateRequest{
		Name:    "libsacloud-note",
		Tags:    []string{"tag1", "tag2"},
		Class:   "shell",
		Content: "test-content",
	}
	createNoteExpected = &sacloud.Note{
		Name:         createNoteParam.Name,
		Tags:         createNoteParam.Tags,
		Class:        createNoteParam.Class,
		Content:      createNoteParam.Content,
		Scope:        types.Scopes.User,
		Availability: types.Availabilities.Available,
	}
	updateNoteParam = &sacloud.NoteUpdateRequest{
		Name:    "libsacloud-note-upd",
		Tags:    []string{"tag1-upd", "tag2-upd"},
		Class:   "shell",
		Content: "test-content-upd",
	}
	updateNoteExpected = &sacloud.Note{
		Name:         updateNoteParam.Name,
		Tags:         updateNoteParam.Tags,
		Class:        updateNoteParam.Class,
		Content:      updateNoteParam.Content,
		Scope:        types.Scopes.User,
		Availability: types.Availabilities.Available,
	}
)

func testNoteCreate(testContext *CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewNoteOp(caller)
	return client.Create(context.Background(), createNoteParam)
}

func testNoteRead(testContext *CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewNoteOp(caller)
	return client.Read(context.Background(), testContext.ID)
}

func testNoteUpdate(testContext *CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewNoteOp(caller)
	return client.Update(context.Background(), testContext.ID, updateNoteParam)
}

func testNoteDelete(testContext *CRUDTestContext, caller sacloud.APICaller) error {
	client := sacloud.NewNoteOp(caller)
	return client.Delete(context.Background(), testContext.ID)
}
