package test

import (
	"testing"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/testutil"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

func TestNoteOp_CRUD(t *testing.T) {
	testutil.Run(t, &testutil.CRUDTestCase{
		Parallel:           true,
		IgnoreStartupWait:  true,
		SetupAPICallerFunc: singletonAPICaller,
		Create: &testutil.CRUDTestFunc{
			Func: testNoteCreate,
			CheckFunc: testutil.AssertEqualWithExpected(&testutil.CRUDTestExpect{
				ExpectValue:  createNoteExpected,
				IgnoreFields: ignoreNoteFields,
			}),
		},
		Read: &testutil.CRUDTestFunc{
			Func: testNoteRead,
			CheckFunc: testutil.AssertEqualWithExpected(&testutil.CRUDTestExpect{
				ExpectValue:  createNoteExpected,
				IgnoreFields: ignoreNoteFields,
			}),
		},
		Updates: []*testutil.CRUDTestFunc{
			{
				Func: testNoteUpdate,
				CheckFunc: testutil.AssertEqualWithExpected(&testutil.CRUDTestExpect{
					ExpectValue:  updateNoteExpected,
					IgnoreFields: ignoreNoteFields,
				}),
			},
			{
				Func: testNoteUpdateToMin,
				CheckFunc: testutil.AssertEqualWithExpected(&testutil.CRUDTestExpect{
					ExpectValue:  updateNoteToMinExpected,
					IgnoreFields: ignoreNoteFields,
				}),
			},
		},
		Delete: &testutil.CRUDTestDeleteFunc{
			Func: testNoteDelete,
		},
	})
}

var (
	ignoreNoteFields = []string{"ID", "CreatedAt", "ModifiedAt"}
	createNoteParam  = &sacloud.NoteCreateRequest{
		Name:    testutil.ResourceName("note"),
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
		Name:    testutil.ResourceName("note-upd"),
		Tags:    []string{"tag1-upd", "tag2-upd"},
		Class:   "shell",
		Content: "test-content-upd",
		IconID:  testIconID,
	}
	updateNoteExpected = &sacloud.Note{
		Name:         updateNoteParam.Name,
		Tags:         updateNoteParam.Tags,
		Class:        updateNoteParam.Class,
		Content:      updateNoteParam.Content,
		Scope:        types.Scopes.User,
		Availability: types.Availabilities.Available,
		IconID:       updateNoteParam.IconID,
	}
	updateNoteToMinParam = &sacloud.NoteUpdateRequest{
		Name:    testutil.ResourceName("note-to-min"),
		Class:   "shell",
		Content: "test-content-upd",
	}
	updateNoteToMinExpected = &sacloud.Note{
		Name:         updateNoteToMinParam.Name,
		Class:        updateNoteToMinParam.Class,
		Content:      updateNoteToMinParam.Content,
		Scope:        types.Scopes.User,
		Availability: types.Availabilities.Available,
	}
)

func testNoteCreate(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewNoteOp(caller)
	return client.Create(ctx, createNoteParam)
}

func testNoteRead(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewNoteOp(caller)
	return client.Read(ctx, ctx.ID)
}

func testNoteUpdate(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewNoteOp(caller)
	return client.Update(ctx, ctx.ID, updateNoteParam)
}

func testNoteUpdateToMin(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewNoteOp(caller)
	return client.Update(ctx, ctx.ID, updateNoteToMinParam)
}

func testNoteDelete(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) error {
	client := sacloud.NewNoteOp(caller)
	return client.Delete(ctx, ctx.ID)
}
