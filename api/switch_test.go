// Copyright 2016-2019 The Libsacloud Authors
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

const testSwitchName = "libsacloud_test_Switch"

func TestSwitchCRUD(t *testing.T) {
	defer initSwitch()()

	api := client.Switch

	//CREATE
	newItem := api.New()
	newItem.Name = testSwitchName
	newItem.Description = "before"

	item, err := api.Create(newItem)

	assert.NoError(t, err)
	assert.NotEmpty(t, item)

	id := item.ID

	//READ
	item, err = api.Read(id)
	assert.NoError(t, err)
	assert.NotEmpty(t, item)

	//UPDATE
	item.Description = "after"
	item, err = api.Update(id, item)

	assert.NoError(t, err)
	assert.NotEqual(t, item.Description, "before")

	//Delete
	_, err = api.Delete(id)
	assert.NoError(t, err)
}

func initSwitch() func() {
	cleanupSwitch()
	return cleanupSwitch
}

func cleanupSwitch() {
	items, _ := client.Switch.Reset().WithNameLike(testSwitchName).Find()
	for _, item := range items.Switches {
		client.Switch.Delete(item.ID)
	}
}
