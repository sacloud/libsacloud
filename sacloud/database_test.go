package sacloud

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	testDatabaseJSON = `
{
        "ID": 999999999999,
        "Class": "database",
        "Name": "\u540d\u79f0\u672a\u8a2d\u5b9a \u30c7\u30fc\u30bf\u30d9\u30fc\u30b9 15690ceb5d8",
        "Description": "",
        "Plan": {
            "ID": 1
        },
        "Settings": {
            "DBConf": {
                "Common": {
                    "AdminPassword": "HogeHogeooo",
                    "DefaultUser": "HogeHogeooo",
                    "UserPassword": "HogeHogeooo",
                    "ServicePort": "",
                    "SourceNetwork": %s,
                    "WebUI": "8.8.8.8"
                },
                "Backup": {
                    "Rotate": "8",
                    "Time": "00:45"
                }
            }
        },
        "SettingsHash": "589300a0e9c9c8a6c6055ffb27f26155",
        "Remark": {
            "Zone": {
                "ID": 21001
            },
            "Network": %s
            ,
            "Servers": [
                [
                ]
            ],
            "DBConf": {
                "Common": {
                    "DatabaseTitle": "PostgreSQL 9.4.7",
                    "DatabaseName": "postgres",
                    "DatabaseVersion": "9.4",
                    "DatabaseRevision": "9.4.7",
                    "ReplicaUser": "replica",
                    "ReplicaPassword": "HogeHogeooo"
                }
            },
            "Switch": {
                "Scope": "shared"
            },
            "Plan": {
                "ID": 1
            }
        },
        "Availability": "available",
        "Instance": {
            "Status": "down",
            "StatusChangedAt": "2016-08-16T09:59:49+09:00"
        },
        "ServiceClass": "cloud\/appliance\/database\/mini",
        "CreatedAt": "2016-08-16T09:44:42+09:00",
        "Icon": null,
        "Switch": null,
        "Interfaces": [
            {
                "IPAddress": "8.8.8.8",
                "UserIPAddress": null,
                "HostName": null,
                "Switch": {
                    "ID": 999999999999,
                    "Name": "\u30b9\u30a4\u30c3\u30c1",
                    "Scope": "shared",
                    "Subnet": {
                        "NetworkAddress": "8.8.8.0",
                        "NetworkMaskLen": 24,
                        "DefaultRoute": "8.8.8.1",
                        "Internet": {
                            "BandWidthMbps": 100
                        }
                    },
                    "UserSubnet": {
                        "DefaultRoute": "8.8.8.1",
                        "NetworkMaskLen": 24
                    }
                }
            }
        ],
        "Tags": [
        ]
    }
`
	testEmptySourceNetworkJSON = `""`
	testArraySourceNetworkJSON = `["192.168.11.1","192.168.11.2"]`
	testEmptyNetworkJSON       = `[]`
	testObjectNetworkJSON      = `{
                    "NetworkMaskLen": 24,
                    "DefaultRoute": "192.168.11.1"
                }`
)

func TestMarshalDatabaseJSON(t *testing.T) {
	//standard plan
	var db Database
	err := json.Unmarshal([]byte(fmt.Sprintf(testDatabaseJSON, testEmptySourceNetworkJSON, testEmptyNetworkJSON)), &db)

	assert.NoError(t, err)
	assert.NotEmpty(t, db)

	assert.NotEmpty(t, db.ID)
	assert.NotEmpty(t, db.Remark)

	assert.NotEmpty(t, db.Remark.Zone)
	assert.NotEmpty(t, db.Remark.DBConf)

	assert.NotEmpty(t, db.Instance)
	assert.NotEmpty(t, db.Interfaces)

	assert.NotEmpty(t, db.Settings.DBConf)
	assert.NotEmpty(t, db.Settings.DBConf.Backup)
	assert.NotEmpty(t, db.Settings.DBConf.Common)

}

func TestMarshalDatabaseJSONWithSourceNetwork(t *testing.T) {
	//standard plan
	var db Database
	err := json.Unmarshal([]byte(fmt.Sprintf(testDatabaseJSON, testArraySourceNetworkJSON, testEmptyNetworkJSON)), &db)

	assert.NoError(t, err)
	assert.NotEmpty(t, db)

	assert.Equal(t, db.Settings.DBConf.Common.SourceNetwork[0], "192.168.11.1")
	assert.Equal(t, db.Settings.DBConf.Common.SourceNetwork[1], "192.168.11.2")

	//add
	db.AddSourceNetwork("192.168.11.3")
	assert.Len(t, db.Settings.DBConf.Common.SourceNetwork, 3)
	assert.Equal(t, db.Settings.DBConf.Common.SourceNetwork[2], "192.168.11.3")

	//del
	db.DeleteSourceNetwork("192.168.11.2")
	assert.Len(t, db.Settings.DBConf.Common.SourceNetwork, 2)
	assert.Equal(t, db.Settings.DBConf.Common.SourceNetwork[1], "192.168.11.3")

}

func TestMarshalDatabaseJSONWithObjectNetwork(t *testing.T) {
	//standard plan
	var db Database
	err := json.Unmarshal([]byte(fmt.Sprintf(testDatabaseJSON, testArraySourceNetworkJSON, testObjectNetworkJSON)), &db)

	assert.NoError(t, err)
	assert.NotEmpty(t, db)

	assert.Equal(t, db.Remark.Network.DefaultRoute, "192.168.11.1")
	assert.Equal(t, db.Remark.Network.NetworkMaskLen, 24)
}
