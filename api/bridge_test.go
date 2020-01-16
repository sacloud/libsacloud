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

const testBridgeName = "libsacloud_test_archive"

func TestBridgeCRUD(t *testing.T) {
	defer initBridge()()

	currentRegion := client.Zone
	defer func() { client.Zone = currentRegion }()
	client.Zone = "is1a"

	api := client.Bridge

	//CREATE
	newBr := api.New()
	newBr.Name = testBridgeName
	newBr.Description = "before"

	br, err := api.Create(newBr)

	assert.NoError(t, err)
	assert.NotEmpty(t, br)

	id := br.ID

	//READ
	br, err = api.Read(id)
	assert.NoError(t, err)
	assert.NotEmpty(t, br)

	//UPDATE
	br.Description = "after"
	br, err = api.Update(id, br)

	assert.NoError(t, err)
	assert.NotEqual(t, br.Description, "before")

	//Delete
	_, err = api.Delete(id)
	assert.NoError(t, err)
}

func initBridge() func() {
	cleanupBridge()
	return cleanupBridge
}

func cleanupBridge() {
	items, _ := client.Bridge.Reset().WithNameLike(testBridgeName).Find()
	for _, item := range items.Bridges {
		client.Bridge.Delete(item.ID)
	}
}
