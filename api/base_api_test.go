package api

import (
	"github.com/sacloud/libsacloud/sacloud"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

const testTargetNoteName string = "libsacloud_base_api_test_note"
const testTargetNoteContentBefore string = `echo "sacloud_base_api_test before...`
const testTargetNoteContentAfter string = `echo "sacloud_base_api_test done!`

func TestCRUDByBaseAPI(t *testing.T) {
	baseAPI := &baseAPI{
		client: client,
		FuncGetResourceURL: func() string {
			return "note"
		},
	}

	//CREATE
	res := &sacloud.Response{}
	req := &sacloud.Request{}

	note := &sacloud.Note{}
	note.Name = testTargetNoteName
	note.Content = testTargetNoteContentBefore

	req.Note = note

	err := baseAPI.create(req, res)
	assert.NoError(t, err)
	assert.NotEmpty(t, res)
	assert.NotEmpty(t, res.Note)

	//for READ
	var id = res.Note.ID

	//READ
	res = &sacloud.Response{}
	req = &sacloud.Request{}

	err = baseAPI.read(id, nil, res)
	assert.NoError(t, err)
	assert.NotEmpty(t, res)
	assert.NotEmpty(t, res.Note.Content)

	//for UPDATE
	req.Note = res.Note

	//UPDATE
	res = &sacloud.Response{}
	req.Note.Content = testTargetNoteContentAfter

	err = baseAPI.update(id, req, res)
	assert.NoError(t, err)
	assert.NotEmpty(t, res)
	assert.NotEmpty(t, res.Note.Content)
	assert.Equal(t, res.Note.Content, testTargetNoteContentAfter)

	//DELETE
	res = &sacloud.Response{}

	err = baseAPI.delete(id, nil, res)
	assert.NoError(t, err)
	assert.NotEmpty(t, res)
}

func init() {
	testSetupHandlers = append(testSetupHandlers, cleanupTestNote)
	testTearDownHandlers = append(testTearDownHandlers, cleanupTestNote)
}

func cleanupTestNote() {
	baseAPI := &baseAPI{
		client: client,
		FuncGetResourceURL: func() string {
			return "note"
		},
	}

	//Find
	//res, _ := baseAPI.Find(&sacloud.Request{
	//	Filter: map[string]interface{}{
	//		"Name": testTargetNoteName,
	//	},
	//})

	res, _ := baseAPI.withNameLike(testTargetNoteName).Find()
	if res != nil && res.Count > 0 && res.Notes[0].Name == testTargetNoteName {
		err := baseAPI.delete(res.Notes[0].ID, nil, nil)
		if err != nil {
			log.Fatalf("Cleanup Notes error! %v", err)
		} else {
			log.Println("Cleanup notes done.")
		}

	}
}
