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

var testInstanceJSON = `
{
	"Server": {
		"ID": 123456789012
	},
	"Status": "up",
	"BeforeStatus": "down",
	"StatusChangedAt": "2016-04-29T18:33:40+09:00",
	"MigrationProgress": null,
	"MigrationSchedule": null,
	"IsMigrating": null,
	"MigrationAllowed": null,
	"ModifiedAt": "2016-04-29T18:33:40+09:00",
	"Host": {
		"Name": "sac-is1b-sv053",
		"InfoURL": "http://support.sakura.ad.jp/mainte/mainteentry.php?id=22178",
		"Class": "dynamic",
		"Version": 200,
		"SystemVersion": "SAKURA Internet [CLOUD SERVICE 2.0]"
	},
	"CDROM": ` + testCDROMJSON + `,
	"CDROMStorage": ` + testStorageJSON + `
}
`

var testStorageJSON = `
{
	"ID": 1234567890,
	"Class": "iscsi1204",
	"Name": "sac-is1b-arc-st01",
	"Description": "",
	"Zone": ` + testZoneJSON + `,
	"DiskPlan": {
	    "ID": 2,
	    "StorageClass": "iscsi1204",
	    "Name": "\u6a19\u6e96\u30d7\u30e9\u30f3"
	},
	"Capacity": []
}
`

func TestInstance(t *testing.T) {
	var instance Instance
	err := json.Unmarshal([]byte(testInstanceJSON), &instance)

	t.Run("MarshalJSON", func(t *testing.T) {
		assert.NoError(t, err)
		assert.NotEmpty(t, instance)

		assert.NotEmpty(t, instance.Server.ID)
		assert.NotEmpty(t, instance.Host.Name)
		assert.NotEmpty(t, instance.CDROM.ID)
		assert.NotEmpty(t, instance.CDROMStorage.ID)
	})
	t.Run("MaintenanceInfo", func(t *testing.T) {
		assert.True(t, instance.HasInfoURL())
		assert.Equal(t, instance.HasInfoURL(), instance.MaintenanceScheduled())
	})
}

func TestMarshalStorageJSON(t *testing.T) {
	var storage Storage
	err := json.Unmarshal([]byte(testStorageJSON), &storage)

	assert.NoError(t, err)
	assert.NotEmpty(t, storage)
	assert.NotEmpty(t, storage.Zone.ID)
	assert.NotEmpty(t, storage.DiskPlan.ID)
}
