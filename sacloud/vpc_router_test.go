package sacloud

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	testVPCRouterStandardJSON = `
	{
		"Interfaces": [
			null,
			{
				"IPAddress": [
					"192.168.200.1"
				],
				"NetworkMaskLen": 24
			}
		]
	}
	`

	testVPCRouterPremiumJSON = `
	{
		"Interfaces": [
			{
				"IPAddress": [
					"133.242.253.101",
					"133.242.253.102"
				],
				"VirtualIPAddress": "133.242.253.100"
			}
		],
		"VRID": 1
	}
	`

	testVPCRouterJSONTemplate = `
	{
            "ID": 123456789012,
            "Class": "vpcrouter",
            "Name": "\u308b\u30fc\u305f\u30fc",
            "Description": "\u308b\u30fc\u305f\u30fc\u306e\u307b\u3052\u307b\u3052",
            "Plan": {
                "ID": 1
            },
            "Settings": {
                "Router": %s
            },
            "SettingsHash": "569c7da372d30123a22763d33a7dee15",
            "Remark": {
                "Servers": [
                    [
                    ]
                ],
                "Switch": {
                    "Scope": "shared"
                },
                "Zone": {
                    "ID": 31002
                }
            },
            "Availability": "available",
            "Instance": {
                "Status": "up",
                "StatusChangedAt": "2016-04-29T18:28:27+09:00"
            },
            "ServiceClass": "cloud\/appliance\/vpc\/1",
            "CreatedAt": "2016-04-29T18:25:23+09:00",
            "Icon": {
                "ID": 112300511988,
                "URL": "https:\/\/secure.sakura.ad.jp\/cloud\/zone\/is1b\/api\/cloud\/1.1\/icon\/112300511988.png",
                "Name": "Wall",
                "Scope": "shared"
            },
            "Switch": null,
            "Interfaces": [
                {
                    "IPAddress": "153.127.196.78",
                    "UserIPAddress": null,
                    "HostName": null,
                    "Switch": {
                        "ID": 112800387183,
                        "Name": "\u30b9\u30a4\u30c3\u30c1",
                        "Scope": "shared",
                        "Subnet": {
                            "NetworkAddress": "153.127.196.0",
                            "NetworkMaskLen": 24,
                            "DefaultRoute": "153.127.196.1",
                            "Internet": {
                                "BandWidthMbps": 100
                            }
                        },
                        "UserSubnet": null
                    }
                },
                {
                    "IPAddress": null,
                    "UserIPAddress": null,
                    "HostName": null,
                    "Switch": {
                        "ID": 112800442260,
                        "Name": "\u3059\u3046\u3043\u3063\u3061",
                        "Scope": "user",
                        "Subnet": null,
                        "UserSubnet": {
                            "DefaultRoute": "192.168.200.1",
                            "NetworkMaskLen": 24
                        }
                    }
                },
                null,
                null,
                null,
                null,
                null,
                null
            ],
            "Tags": [
                "\u307b\u3052",
                "\u307b\u3052\u304a",
                "\u307b\u3052\u307e"
            ]
        }
	`
)

func TestMarshalVPCRouterJSON(t *testing.T) {
	//standard plan
	var router VPCRouter
	err := json.Unmarshal([]byte(fmt.Sprintf(testVPCRouterJSONTemplate, testVPCRouterStandardJSON)), &router)

	assert.NoError(t, err)
	assert.NotEmpty(t, router)

	assert.NotEmpty(t, router.ID)
	assert.NotEmpty(t, router.Remark)

	assert.NotEmpty(t, router.Remark.Switch)
	//TODO Zone
	//assert.NotEmpty(t, router.Remark.Zone)

	assert.Nil(t, router.Remark.Network)
	assert.NotEmpty(t, router.Remark.Servers)
	assert.True(t, len(router.Remark.Servers) > 0)
	assert.Nil(t, router.Remark.VRRP)

	assert.NotEmpty(t, router.Instance)
	assert.NotEmpty(t, router.Interfaces)

	assert.NotEmpty(t, router.Settings.Router)

	//for standard
	assert.Nil(t, router.Settings.Router.VRID)
	assert.NotEmpty(t, router.Settings.Router.Interfaces)

	assert.Nil(t, router.Settings.Router.Interfaces[0])
	assert.NotEmpty(t, router.Settings.Router.Interfaces[1])

	assert.NotEmpty(t, router.Settings.Router.Interfaces[1].IPAddress)
	assert.NotEmpty(t, router.Settings.Router.Interfaces[1].NetworkMaskLen)
	assert.Empty(t, router.Settings.Router.Interfaces[1].VirtualIPAddress)

	//for premium
	err = json.Unmarshal([]byte(fmt.Sprintf(testVPCRouterJSONTemplate, testVPCRouterPremiumJSON)), &router)

	assert.NotEmpty(t, router.Settings.Router.VRID)
	assert.NotEmpty(t, router.Settings.Router.Interfaces)

	assert.NotEmpty(t, router.Settings.Router.Interfaces[0])

	assert.NotEmpty(t, router.Settings.Router.Interfaces[0].IPAddress)
	assert.Empty(t, router.Settings.Router.Interfaces[0].NetworkMaskLen)
	assert.NotEmpty(t, router.Settings.Router.Interfaces[0].VirtualIPAddress)

}
