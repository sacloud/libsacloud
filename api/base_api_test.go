package api

import (
	"github.com/stretchr/testify/assert"
	sakura "github.com/yamamoto-febc/libsacloud/resources"
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
	var res = &sakura.Response{}
	var note = &sakura.Request{
		Note: &sakura.Note{
			Name:    testTargetNoteName,
			Content: testTargetNoteContentBefore,
		},
	}

	err := baseAPI.create(note, res)
	assert.NoError(t, err)
	assert.NotEmpty(t, res)
	assert.NotEmpty(t, res.Note)

	//for READ
	var id = res.Note.ID

	//READ
	res = &sakura.Response{}
	note = &sakura.Request{}

	err = baseAPI.read(id, nil, res)
	assert.NoError(t, err)
	assert.NotEmpty(t, res)
	assert.NotEmpty(t, res.Note.Content)

	//for UPDATE
	note.Note = res.Note

	//UPDATE
	res = &sakura.Response{}
	note.Note.Content = testTargetNoteContentAfter

	err = baseAPI.update(id, note, res)
	assert.NoError(t, err)
	assert.NotEmpty(t, res)
	assert.NotEmpty(t, res.Note.Content)
	assert.Equal(t, res.Note.Content, testTargetNoteContentAfter)

	//DELETE
	res = &sakura.Response{}

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
	res, _ := baseAPI.Find(&sakura.Request{
		Filter: map[string]interface{}{
			"Name": testTargetNoteName,
		},
	})

	if res != nil && res.Count > 0 && res.Notes[0].Name == testTargetNoteName {
		err := baseAPI.delete(res.Notes[0].ID, nil, nil)
		if err != nil {
			log.Fatalf("Cleanup Notes error! %v", err)
		} else {
			log.Println("Cleanup notes done.")
		}

	}
}
