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

var testAutoBackupJSON = `
{
        "CreatedAt": "2016-07-07T10:58:36+09:00",
        "Description": "\u30d0\u30c3\u30af\u30a2\u30c3\u30d7\u306a\u306e\u306a\u306e",
        "ID": 112800693628,
        "Icon": null,
        "ModifiedAt": "2016-07-07T10:58:36+09:00",
        "Name": "sakura-windows",
        "Provider": {
            "Class": "autobackup",
            "ID": 4000001,
            "Name": "autobackup1",
            "ServiceClass": "cloud/autobackup"
        },
        "ServiceClass": "cloud/autobackup",
        "Settings": {
            "AccountId": "999999999999",
            "Autobackup": {
                "BackupSpanType": "weekdays",
                "BackupSpanWeekdays": [
                    "sun",
                    "mon",
                    "tue",
                    "wed",
                    "thu",
                    "fri",
                    "sat"
                ],
                "MaximumNumberOfArchives": 2
            },
            "DiskId": "112800666498",
            "ZoneId": 31002,
            "ZoneName": "is1b"
        },
        "SettingsHash": "9f3442b29229f67981bc8971861edc37",
        "Status": {
            "AccountId": "999999999999",
            "DiskId": "112800666498",
            "ZoneId": 31002,
            "ZoneName": "is1b"
        },
        "Tags": [
            "bak2",
            "bk1",
            "javascript:type:hogehoge:alert"
        ]
}
`

func TestMarshalAutoBackupJSON(t *testing.T) {
	var ab AutoBackup
	err := json.Unmarshal([]byte(testAutoBackupJSON), &ab)

	assert.NoError(t, err)
	assert.NotEmpty(t, ab)

	assert.NotEmpty(t, ab.ID)
	assert.Equal(t, ab.Provider.Class, "autobackup")
	assert.Equal(t, ab.Name, "sakura-windows")
}
