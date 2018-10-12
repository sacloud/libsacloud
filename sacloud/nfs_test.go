package sacloud

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testNFSJSON = `
{
       "Availability": "available",
        "Class": "nfs",
        "CreatedAt": "2017-09-07T12:20:15+09:00",
        "Description": "nfs-description",
        "ID": 112901132074,
        "Icon": {
            "ID": 112901044094,
            "Name": "myicon",
            "Scope": "user",
            "URL": "https://secure.sakura.ad.jp/cloud/zone/is1b/api/cloud/1.1/icon/112901044094.png"
        },
        "Instance": {
            "Status": "up",
            "StatusChangedAt": "2017-09-07T12:22:33+09:00"
        },
        "Interfaces": [
            {
                "HostName": null,
                "IPAddress": null,
                "Switch": {
                    "ID": 112900591751,
                    "Name": "vpc_router_for_vpn",
                    "Scope": "user",
                    "Subnet": null,
                    "UserSubnet": {
                        "DefaultRoute": "192.168.150.1",
                        "NetworkMaskLen": 24
                    }
                },
                "UserIPAddress": "192.168.150.101"
            }
        ],
        "Name": "nfs-test",
        "Plan": {
            "ID": 100
        },
        "Remark": {
            "Network": {
                "DefaultRoute": "192.168.150.1",
                "NetworkMaskLen": 24
            },
            "Plan": {
                "ID": 100
            },
            "Servers": [
                {
                    "IPAddress": "192.168.150.101"
                }
            ],
            "Switch": {
                "ID": "112900591751"
            },
            "Zone": {
                "ID": "31002"
            }
        },
        "ServiceClass": "cloud/appliance/nfs/100gb",
        "Settings": null,
        "SettingsHash": null,
        "Switch": {
            "Availability": "available",
            "ID": 112900591751,
            "Internet": null,
            "Name": "vpc_router_for_vpn",
            "Scope": "user",
            "Zone": {
                "ID": 31002,
                "Name": "is1b",
                "Region": {
                    "ID": 310,
                    "Name": "\u77f3\u72e9"
                }
            }
        },
        "Tags": [
            "tag1",
            "tag2"
        ]
    }
    	`

func TestMarshalNFSJSON(t *testing.T) {
	var nfs NFS
	err := json.Unmarshal([]byte(testLoadBalancerJSON), &nfs)

	assert.NoError(t, err)
	assert.NotEmpty(t, nfs)

	assert.NotEmpty(t, nfs.ID)
	assert.NotEmpty(t, nfs.Remark)

	assert.NotEmpty(t, nfs.Remark.Servers)
	assert.NotEmpty(t, nfs.Remark.Network)
	assert.NotEmpty(t, nfs.Remark.Switch)
	//TODO Zone
	//assert.NotEmpty(t, nfs.Remark.Zone)
	//assert.NotEmpty(t, nfs.Remark.Plan)

	assert.NotEmpty(t, nfs.Instance)
	assert.NotEmpty(t, nfs.Interfaces)

}
