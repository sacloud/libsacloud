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
