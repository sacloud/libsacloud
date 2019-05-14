package sacloud

import (
	"context"
	"testing"

	"github.com/sacloud/libsacloud-v2/sacloud/types"
	"github.com/stretchr/testify/require"
)

func TestNoteOpCRUD(t *testing.T) {
	Test(t, &CRUDTestCase{
		Parallel: true,

		SetupAPICaller: singletonAPICaller,
		Create: &CRUDTestFunc{
			Func: testNoteCreate,
			Expect: &CRUDTestExpect{
				ExpectValue:  createNoteExpected,
				IgnoreFields: []string{"ID", "CreatedAt", "ModifiedAt"},
			},
		},
		Read: &CRUDTestFunc{
			Func: testNoteRead,
			Expect: &CRUDTestExpect{
				ExpectValue:  createNoteExpected,
				IgnoreFields: []string{"ID", "CreatedAt", "ModifiedAt"},
			},
		},
		Update: &CRUDTestFunc{
			Func: testNoteUpdate,
			Expect: &CRUDTestExpect{
				ExpectValue:  updateNoteExpected,
				IgnoreFields: []string{"ID", "CreatedAt", "ModifiedAt"},
			},
		},
		Delete: &CRUDTestDeleteFunc{
			Func: testNoteDelete,
		},
	})
}

var (
	createNoteParam = &NoteCreateRequest{
		Name:    "libsacloud-v2-note",
		Tags:    []string{"tag1", "tag2"},
		Class:   "shell",
		Content: "test-content",
	}
	createNoteExpected = &Note{
		Name:         createNoteParam.Name,
		Tags:         createNoteParam.Tags,
		Class:        createNoteParam.Class,
		Content:      createNoteParam.Content,
		Scope:        types.Scopes.User,
		Availability: types.Availabilities.Available,
	}
	updateNoteParam = &NoteUpdateRequest{
		Name:    "libsacloud-v2-note-upd",
		Tags:    []string{"tag1-upd", "tag2-upd"},
		Class:   "shell",
		Content: "test-content-upd",
	}
	updateNoteExpected = &Note{
		Name:         updateNoteParam.Name,
		Tags:         updateNoteParam.Tags,
		Class:        updateNoteParam.Class,
		Content:      updateNoteParam.Content,
		Scope:        types.Scopes.User,
		Availability: types.Availabilities.Available,
	}
)

func testNoteCreate(testContext *CRUDTestContext, caller APICaller) (interface{}, error) {
	client := NewNoteOp(caller)
	return client.Create(context.Background(), DefaultZone, createNoteParam)
}

func testNoteRead(testContext *CRUDTestContext, caller APICaller) (interface{}, error) {
	client := NewNoteOp(caller)
	return client.Read(context.Background(), DefaultZone, testContext.ID)
}

func testNoteUpdate(testContext *CRUDTestContext, caller APICaller) (interface{}, error) {
	client := NewNoteOp(caller)
	return client.Update(context.Background(), DefaultZone, testContext.ID, updateNoteParam)
}

func testNoteDelete(testContext *CRUDTestContext, caller APICaller) error {
	client := NewNoteOp(caller)
	return client.Delete(context.Background(), DefaultZone, testContext.ID)
}

func TestFindNote(t *testing.T) {
	if !isAccTest() {
		t.Skip("TESTACC is not set. skip")
	}

	client := NewNoteOp(singletonAPICaller())

	notes, err := client.Find(context.Background(), DefaultZone, &FindCondition{Count: 1})
	require.NoError(t, err)
	require.Len(t, notes, 1)
}
