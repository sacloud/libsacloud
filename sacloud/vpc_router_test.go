// Copyright 2016-2020 The Libsacloud Authors
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
	"fmt"
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
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
	assert.NoError(t, err)

	assert.NotEmpty(t, router.Settings.Router.VRID)
	assert.NotEmpty(t, router.Settings.Router.Interfaces)

	assert.NotEmpty(t, router.Settings.Router.Interfaces[0])

	assert.NotEmpty(t, router.Settings.Router.Interfaces[0].IPAddress)
	assert.Empty(t, router.Settings.Router.Interfaces[0].NetworkMaskLen)
	assert.NotEmpty(t, router.Settings.Router.Interfaces[0].VirtualIPAddress)

}

func TestVPCRouter_RealIPAddress(t *testing.T) {
	const (
		standard = 1
		other    = 2
	)
	expects := []struct {
		plan       int
		index      int
		interfaces []*VPCRouterInterface
		expect     string
	}{
		{
			plan:  standard,
			index: 1,
			interfaces: []*VPCRouterInterface{
				{},
				{
					IPAddress:      []string{"192.168.0.1"},
					NetworkMaskLen: 24,
				},
			},
			expect: "192.168.0.1",
		},
		{
			plan:  standard,
			index: 1,
			interfaces: []*VPCRouterInterface{
				{},
				{
					IPAddress:        []string{"192.168.0.1", "192.168.0.2"},
					VirtualIPAddress: "192.168.0.3",
					NetworkMaskLen:   24,
				},
			},
			expect: "192.168.0.1",
		},
		{
			plan:  standard,
			index: 5,
			interfaces: []*VPCRouterInterface{
				{IPAddress: []string{"192.168.0.1"}, NetworkMaskLen: 24},
				{IPAddress: []string{"192.168.1.1"}, NetworkMaskLen: 24},
				{IPAddress: []string{"192.168.2.1"}, NetworkMaskLen: 24},
				{IPAddress: []string{"192.168.3.1"}, NetworkMaskLen: 24},
				{IPAddress: []string{"192.168.4.1"}, NetworkMaskLen: 24},
				{IPAddress: []string{"192.168.5.1"}, NetworkMaskLen: 24},
				{IPAddress: []string{"192.168.6.1"}, NetworkMaskLen: 24},
			},
			expect: "192.168.5.1",
		},
		{
			plan:  standard,
			index: 5,
			interfaces: []*VPCRouterInterface{
				{IPAddress: []string{"192.168.0.1"}, NetworkMaskLen: 24},
				{},
				{},
				{},
				{},
				{IPAddress: []string{"192.168.5.1"}, NetworkMaskLen: 24},
				{},
			},
			expect: "192.168.5.1",
		},
	}

	vpcRouter := CreateNewVPCRouter()

	for _, e := range expects {

		vpcRouter.InitVPCRouterSetting()
		vpcRouter.Plan.SetID(int64(e.plan))
		vpcRouter.Settings.Router.Interfaces = e.interfaces

		actual, _ := vpcRouter.RealIPAddress(e.index)
		assert.Equal(t, e.expect, actual)

	}

}

func TestVPCRouter_FindBelongsInterface(t *testing.T) {
	expects := []struct {
		ip     string
		expect int
	}{
		{
			ip:     "192.168.0.2",
			expect: 0,
		},
		{
			ip:     "192.168.1.2",
			expect: 1,
		},
		{
			ip:     "192.168.7.2",
			expect: -1,
		},
	}

	vpcRouter := CreateNewVPCRouter()
	vpcRouter.InitVPCRouterSetting()
	vpcRouter.Plan.SetID(int64(1))
	vpcRouter.Settings.Router.Interfaces = []*VPCRouterInterface{
		{IPAddress: []string{"192.168.0.1"}, NetworkMaskLen: 24},
		{IPAddress: []string{"192.168.1.1"}, NetworkMaskLen: 24},
		{IPAddress: []string{"192.168.2.1"}, NetworkMaskLen: 24},
		{IPAddress: []string{"192.168.3.1"}, NetworkMaskLen: 24},
		{IPAddress: []string{"192.168.4.1"}, NetworkMaskLen: 24},
		{IPAddress: []string{"192.168.5.1"}, NetworkMaskLen: 24},
		{IPAddress: []string{"192.168.6.1"}, NetworkMaskLen: 24},
	}
	vpcRouter.Interfaces = []Interface{
		{
			IPAddress: "192.168.0.1",
		},
	}
	vpcRouter.Interfaces[0].Switch = &Switch{
		Subnet: &Subnet{
			NetworkMaskLen: 24,
		},
	}

	for _, e := range expects {
		index, nic := vpcRouter.FindBelongsInterface(net.ParseIP(e.ip))
		assert.Equal(t, e.expect, index)
		assert.Equal(t, nic == nil, e.expect < 0)
	}

}
