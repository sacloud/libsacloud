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

package sacloud

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCDROMJSON = `
{
	"ID": 123456789012,
	"DisplayOrder": 100904201011,
	"StorageClass": "iscsi1204",
	"Name": "openSUSE 42.1 64bit",
	"Description": "openSUSE-Leap-42.1-DVD-x86_64.iso",
	"SizeMB": 5120,
	"Scope": "shared",
	"Availability": "available",
	"ServiceClass": null,
	"CreatedAt": "2015-12-24T18:14:57+09:00",
	"Icon": null,
	"Storage": ` + testStorageJSON + `
}
`

func TestMarshalCDROMSON(t *testing.T) {
	var cd CDROM
	err := json.Unmarshal([]byte(testCDROMJSON), &cd)

	assert.NoError(t, err)
	assert.NotEmpty(t, cd)

	assert.NotEmpty(t, cd.ID)
	assert.NotEmpty(t, cd.Storage)
}
