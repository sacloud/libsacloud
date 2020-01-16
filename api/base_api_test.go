// Copyright 2016-2020 The Libsacloud Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package api

import (
	"bytes"
	"log"
	"testing"

	"github.com/sacloud/libsacloud/sacloud"
	"github.com/stretchr/testify/assert"
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
