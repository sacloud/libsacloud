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
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testPublicNetSwitchJSON = `
{
	"ID": 123456789012,
	"Name": "\u30b9\u30a4\u30c3\u30c1",
	"Scope": "shared",
	"Subnet": {
		"ID": null,
		"NetworkAddress": "133.242.224.0",
		"NetworkMaskLen": 24,
		"DefaultRoute": "133.242.224.1",
		"Internet": {
		"BandWidthMbps": 100
		}
	},
	"UserSubnet": {
		"DefaultRoute": "133.242.224.1",
		"NetworkMaskLen": 24
	}
}
`

var testPrivateNetSwitchJSON = `
{
	"ID": 123456789012,
	"Name": "\u3059\u3046\u3043\u3063\u3061",
	"Scope": "user",
	"Subnet": null,
	"UserSubnet": {
		"DefaultRoute": "192.168.200.1",
		"NetworkMaskLen": 24
	}
}
`

var testSwitchWithBridgeJSON = `
{
            "Index": 1,
            "ID": 123456789012,
            "Name": "\u3059\u3046\u3043\u3063\u3061",
            "Description": "\u3059\u3046\u3043\u3063\u3061\u306e\u8aac\u660e\u306a\u306e",
            "ServerCount": 1,
            "ApplianceCount": 2,
            "Scope": "user",
            "UserSubnet": {
                "DefaultRoute": "192.168.200.1",
                "NetworkMaskLen": 24
            },
            "HybridConnection": null,
            "ServiceClass": "cloud\/switch\/default",
            "CreatedAt": "2016-04-29T18:26:16+09:00",
            "Icon": {
                "ID": 112300512546,
                "URL": "https:\/\/secure.sakura.ad.jp\/cloud\/zone\/is1b\/api\/cloud\/1.1\/icon\/112300512546.png",
                "Name": "Scientific Linux",
                "Scope": "shared"
            },
            "Zone": ` + testZoneJSON + `,
            "Subnets": [
            ],
            "IPv6Nets": [
            ],
            "Internet": null,
            "Bridge": {
                "ID": 123456789012,
                "Name": "sakura_hyb",
                "Info": {
                    "Switches": [
                        {
                            "ID": "123456789012",
                            "Name": "\u3059\u3046\u3043\u3063\u3061",
                            "Zone": {
                                "ID": 31002,
                                "Name": "is1b"
                            }
                        }
                    ]
                }
            },
            "Tags": [
                "\u307b",
                "\u307b\u3052",
                "\u307b\u3052\u306a"
            ]
        }
`

var testRouterJSON = `
{
            "Index": 0,
            "ID": 123456789012,
            "Name": "\u308b\u30fc\u305f",
            "Description": "\u30eb\u30fc\u30bf\u306e\u8aac\u660e",
            "ServerCount": 0,
            "ApplianceCount": 0,
            "Scope": "user",
            "UserSubnet": null,
            "HybridConnection": null,
            "ServiceClass": "cloud\/switch\/default",
            "CreatedAt": "2016-05-02T10:51:33+09:00",
            "Icon": {
                "ID": 112300511380,
                "URL": "https:\/\/secure.sakura.ad.jp\/cloud\/zone\/is1b\/api\/cloud\/1.1\/icon\/112300511380.png",
                "Name": "CGI",
                "Scope": "shared"
            },
            "Zone": ` + testZoneJSON + `,
            "Subnets": [
                {
                    "ID": 9999,
                    "NetworkAddress": "133.242.253.96",
                    "NetworkMaskLen": 28,
                    "DefaultRoute": "133.242.253.97",
                    "NextHop": null,
                    "StaticRoute": null,
                    "ServiceClass": "cloud\/global-ipaddress-v4\/28",
                    "IPAddresses": {
                        "Min": "133.242.253.100",
                        "Max": "133.242.253.110"
                    },
                    "Internet": {
                        "ID": 123456789012,
                        "Name": "\u308b\u30fc\u305f",
                        "BandWidthMbps": 100,
                        "ServiceClass": "cloud\/internet\/router\/100m"
                    }
                }
            ],
            "IPv6Nets": [
                {
                    "ID": 999,
                    "IPv6Prefix": "2401:2500:10a:101e::",
                    "IPv6PrefixLen": 64
                }
            ],
            "Internet": {
                "ID": 123456789012,
                "Name": "\u308b\u30fc\u305f",
                "BandWidthMbps": 100,
                "Scope": "user",
                "ServiceClass": "cloud\/internet\/router\/100m"
            },
            "Bridge": null,
            "Tags": [
                "hoge",
                "hoge2"
            ]
        }
`

func TestMarshalSwitchJSON(t *testing.T) {
	var publicSwitch, privateSwitch Switch

	err := json.Unmarshal([]byte(testPublicNetSwitchJSON), &publicSwitch)
	assert.NoError(t, err)
	assert.NotEmpty(t, publicSwitch)
	assert.NotEmpty(t, publicSwitch.ID)
	assert.NotEmpty(t, publicSwitch.Subnet.NetworkAddress)
	assert.NotEmpty(t, publicSwitch.Subnet.Internet)
	assert.NotEmpty(t, publicSwitch.UserSubnet.DefaultRoute)

	err = json.Unmarshal([]byte(testPrivateNetSwitchJSON), &privateSwitch)
	assert.NoError(t, err)
	assert.NotEmpty(t, privateSwitch)
	assert.NotEmpty(t, privateSwitch.ID)
	assert.Nil(t, privateSwitch.Subnet)
	assert.NotEmpty(t, privateSwitch.UserSubnet.DefaultRoute)

}

func TestMarshalSwitchWithBridgeJSON(t *testing.T) {
	var sw Switch
	err := json.Unmarshal([]byte(testSwitchWithBridgeJSON), &sw)

	assert.NoError(t, err)
	assert.NotEmpty(t, sw)
	assert.NotEmpty(t, sw.ID)
	assert.NotEmpty(t, sw.Bridge)

}

func TestMarshalRouterJSON(t *testing.T) {
	var router Switch
	err := json.Unmarshal([]byte(testRouterJSON), &router)

	assert.NoError(t, err)
	assert.NotEmpty(t, router)
	assert.NotEmpty(t, router.ID)
	assert.NotEmpty(t, router.Subnets)
	assert.NotEmpty(t, router.Subnets[0].IPAddresses.Min)

	ipaddresses, err := router.GetIPAddressList()
	assert.NoError(t, err)
	assert.NotEmpty(t, ipaddresses)
	assert.Equal(t, router.Subnets[0].IPAddresses.Min, ipaddresses[0])
	assert.Equal(t, router.Subnets[0].IPAddresses.Max, ipaddresses[len(ipaddresses)-1])

	assert.NotEmpty(t, router.Internet)
	assert.NotEmpty(t, router.IPv6Nets)

}

func TestIPHandling(t *testing.T) {
	ip := net.ParseIP("192.168.0.1").To4()
	assert.Equal(t, ip.String(), "192.168.0.1")
	assert.Equal(t, byte(192), ip[0])
	assert.Equal(t, byte(168), ip[1])
	assert.Equal(t, byte(0), ip[2])
	assert.Equal(t, byte(1), ip[3])

	assert.Equal(t, byte(2), ip[3]+1)

}

func TestSwitchProp(t *testing.T) {

	type PropID interface {
		GetID() int64
	}
	type PropName interface {
		GetName() string
	}

	sw := &Switch{}
	var i interface{} = sw

	prop, ok := i.(PropID)
	assert.True(t, ok)
	assert.NotNil(t, prop)

	prop2, ok := i.(PropName)
	assert.True(t, ok)
	assert.NotNil(t, prop2)
}
