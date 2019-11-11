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

const testCDROMName = "libsacloud_test_iso_image"

func TestCRUDCDROM(t *testing.T) {

	defer initCDROM()()

	api := client.CDROM

	//CREATE
	newCD := api.New()
	newCD.Name = testCDROMName
	newCD.Description = "hoge"
	newCD.SizeMB = 5120

	cd, _, err := api.Create(newCD)

	assert.NoError(t, err)
	assert.NotEmpty(t, cd)
	id := cd.ID

	//READ
	cd, err = api.Read(id)
	assert.NoError(t, err)
	assert.NotEmpty(t, cd)

	//Delete
	_, err = api.Delete(id)
	assert.NoError(t, err)
}

func initCDROM() func() {
	cleanupCDROM()
	return cleanupCDROM
}

func cleanupCDROM() {
	items, _ := client.CDROM.Reset().WithNameLike(testCDROMName).Find()
	for _, item := range items.CDROMs {
		client.CDROM.Delete(item.ID)
	}
}
