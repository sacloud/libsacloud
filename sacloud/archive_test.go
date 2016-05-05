package sacloud

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

var testArchiveJSON = `
        {
            "Index": 19,
            "ID": "112800288256",
            "DisplayOrder": "400201220002",
            "Name": "Windows Server 2012 R2 Datacenter Edition",
            "Description": "\u003Cdiv class=\u0022lang ja\u0022\u003E\n\u30fb\u30b5\u30fc\u30d0\u3001\u30c7\u30a3\u30b9\u30af\u6599\u91d1\u306e\u4ed6\u3001OS\u306e\u5229\u7528\u6599\u91d1\u304c\u5fc5\u8981\u3068\u306a\u308a\u307e\u3059\u3002\n\u30fb\u30c7\u30a3\u30b9\u30af\u4fee\u6b63\u6a5f\u80fd\u306f\u3054\u5229\u7528\u3044\u305f\u3060\u3051\u307e\u305b\u3093\u3002\n\u3000\u30b5\u30fc\u30d0\u4f5c\u6210\u5f8c\u3001\u304a\u5ba2\u69d8\u306b\u3066\u521d\u671f\u8a2d\u5b9a\u3092\u884c\u3063\u3066\u3044\u305f\u3060\u304f\u5fc5\u8981\u304c\u3042\u308a\u307e\u3059\u3002\n\u3000\u624b\u9806\u306b\u3064\u3044\u3066\u306f\u4ee5\u4e0bURL\u3092\u53c2\u7167\u3057\u3066\u304f\u3060\u3055\u3044\u3002\n\u3000http:\/\/cloud-news.sakura.ad.jp\/windows-server-plan\/windows-server-setup\/\n\u003C\/div\u003E\n\u003Cdiv class=\u0022lang en\u0022\u003E\n\u30fbLicense fee for operating system is not included to your server and disk charge.\n\u3000An additional charge is required to use Windows OS.\n\u30fbDisk modification function is not available for Windows servsers.\n\u3000You will need to perform initial configuration after creating a server.\n\u3000Please refer to the link below for instructions.\n\u3000http:\/\/cloud-news.sakura.ad.jp\/windows-server-plan\/windows-server-setup\/\n\u003C\/div\u003E",
            "Scope": "shared",
            "Availability": "available",
            "SizeMB": 102400,
            "MigratedMB": 102400,
            "WaitingJobCount": null,
            "JobStatus": {
		"Status":"waiting",
		"Delays":
		{
			"Start":{"Max":0,"Min":0},
			"Finish":{"Max":900,"Min":205}
		}
	    },
            "OriginalArchive": {
                "ID": "112800288256"
            },
            "SourceInfo": null,
            "ServiceClass": "cloud\/archive\/100g",
            "CreatedAt": "2016-03-16T10:59:35+09:00",
            "Icon": null,
            "Plan": {
                "ID": 2,
                "StorageClass": "iscsi1204",
                "Name": "\u6a19\u6e96\u30d7\u30e9\u30f3"
            },
            "SourceDisk": ` + testDiskJSON + `,
            "SourceArchive": null,
            "BundleInfo": {
                "ID": 10000,
                "Group": null,
                "HostClass": "ms_windows",
                "Name": "Windows_DC",
                "ServiceClass": "cloud\/os\/windows\/datacenter",
                "Attr": null
            },
            "Storage": ` + testStorageJSON + `,
            "Tags": [
                "arch-64bit",
                "current-stable",
                "distro-ver-2012.2",
                "os-windows"
            ]
        }

`

func TestMarshalArchiveJSON(t *testing.T) {
	var archive Archive
	err := json.Unmarshal([]byte(testArchiveJSON), &archive)

	assert.NoError(t, err)
	assert.NotEmpty(t, archive)

	assert.NotEmpty(t, archive.ID)
	assert.NotEmpty(t, archive.OriginalArchive.ID)
}
