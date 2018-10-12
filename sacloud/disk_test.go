package sacloud

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testDiskJSON = `
{
	"ID": 123456789012,
	"Name": "libsacloud-disk-test-name",
	"Connection": "virtio",
	"ConnectionOrder": 1,
	"ReinstallCount": 0,
	"Availability": "available",
	"SizeMB": 20480,
	"Plan": {
		"ID": 4
	},
	"Storage": {
        "Class": "iscsi1204",
        "DiskPlan": {
            "ID": 4,
            "Name": "SSD",
            "StorageClass": "iscsi1204"
        },
        "ID": 1234567890,
        "Name": "sac-isxx-iscsi4-st999",
        "Zone": {
            "ID": 31002,
            "Name": "is1b",
            "Region": {
                "ID": 310,
                "Name": "\u77f3\u72e9"
            }
         }
    },
	"BundleInfo": null
}
`

var testDiskMigratingJSON = `
{
            "ID": 123456789012,
            "Name": "hogedisk",
            "Description": "aaaaa",
            "Connection": "virtio",
            "ConnectionOrder": null,
            "ReinstallCount": 0,
            "Availability": "migrating",
            "SizeMB": 20480,
            "MigratedMB": 7744,
            "WaitingJobCount": null,
            "JobStatus": {
                "Status": "running",
                "Delays": {
                    "Finish": {
                        "Max": 99,
                        "Min": 99
                    }
                }
            },
            "ServiceClass": "cloud\/disk\/ssd\/20g",
            "BundleID": null,
            "CreatedAt": "2016-04-29T21:26:05+09:00",
            "Icon": ` + testIconJSON + `,
            "Plan": {
                "ID": 4,
                "StorageClass": "iscsi1204",
                "Name": "SSD\u30d7\u30e9\u30f3"
            },
            "SourceDisk": ` + testDiskJSON + `,
            "SourceArchive": null,
            "BundleInfo": null,
            "Storage": ` + testStorageJSON + `,
            "Appliance": null,
            "Server": null,
            "Tags": [
                "aa",
                "bb",
                "cc"
            ]
        }

`

func TestMarshalDiskJSON(t *testing.T) {
	var disk Disk
	err := json.Unmarshal([]byte(testDiskJSON), &disk)

	assert.NoError(t, err)
	assert.NotEmpty(t, disk)

	assert.NotEmpty(t, disk.ID)
	assert.NotEmpty(t, disk.Plan.ID)
	assert.NotEmpty(t, disk.Storage.ID)
}

func TestMarshalMigratingDiskJSON(t *testing.T) {
	var disk Disk
	err := json.Unmarshal([]byte(testDiskMigratingJSON), &disk)

	assert.NoError(t, err)
	assert.NotEmpty(t, disk)

	assert.Nil(t, disk.SourceArchive)
	assert.NotEmpty(t, disk.SourceDisk)
	assert.Equal(t, disk.Availability, EAMigrating)

}
