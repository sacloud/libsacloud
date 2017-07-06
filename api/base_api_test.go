package api

import (
	"bytes"
	"github.com/sacloud/libsacloud/sacloud"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

const testTargetNoteName string = "libsacloud_base_api_test_note"
const testTargetNoteContentBefore string = `echo "sacloud_base_api_test before...`
const testTargetNoteContentAfter string = `echo "sacloud_base_api_test done!`

func TestTracer(t *testing.T) {
	requestBuf := bytes.NewBufferString("")
	responseBuf := bytes.NewBufferString("")

	client.RequestTracer = requestBuf
	client.ResponseTracer = responseBuf

	client.Archive.Reset().WithTags([]string{"os-linux", "os-windows"}).Include("ID").Include("Name").Find()

	assert.True(t, requestBuf.Len() > 0)
	assert.True(t, responseBuf.Len() > 0)

	//t.Logf("Request:%s", requestBuf.String())
	//t.Logf("Response:%s", responseBuf.String())
}

func TestCRUDByBaseAPI(t *testing.T) {

	defer initBaseAPI()()

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

func initBaseAPI() func() {
	cleanupBaseAPI()
	return cleanupBaseAPI
}

func cleanupBaseAPI() {
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
