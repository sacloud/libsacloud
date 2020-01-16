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
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	testFindOrCreateByName    = "libsacloud_test_note_name"
	testFindOrCreateByContent = "libsacloud_test_note_content"
)

func TestCRUDByNoteAPI(t *testing.T) {
	defer initNote()()

	noteAPI := client.Note

	//CREATE
	var note = noteAPI.New()
	note.Name = testTargetNoteName
	note.Content = testTargetNoteContentBefore

	res, err := noteAPI.Create(note)
	assert.NoError(t, err)
	assert.NotEmpty(t, res)

	//for READ
	var id = res.ID

	//READ
	res, err = noteAPI.Read(id)
	assert.NoError(t, err)
	assert.NotEmpty(t, res)
	assert.NotEmpty(t, res.Content)

	//UPDATE
	note.Content = testTargetNoteContentAfter

	res, err = noteAPI.Update(id, note)
	assert.NoError(t, err)
	assert.NotEmpty(t, res)
	assert.NotEmpty(t, res.Content)
	assert.Equal(t, res.Content, testTargetNoteContentAfter)

	//DELETE
	res, err = noteAPI.Delete(id)
	assert.NoError(t, err)
	assert.NotEmpty(t, res)
}

func initNote() func() {
	cleanupNote()
	return cleanupNote
}

func cleanupNote() {
	noteAPI := client.Note
	res, _ := noteAPI.withNameLike(testFindOrCreateByName).Find()
	if res.Count > 0 {
		noteAPI.Delete(res.Notes[0].ID)
	}
}
