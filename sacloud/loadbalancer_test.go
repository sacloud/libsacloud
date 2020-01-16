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
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	testLoadBalancerJSON = `
{
        "ID": 123456789012,
        "Class": "loadbalancer",
        "Name": "\u308d\u304a\u3069\u3070\u3089\u3093\u3055",
        "Description": "\u30ed\u30aa\u30c9\u30d0\u30e9\u30f3\u3055\u306e\u8aac\u660e",
        "Plan": {
            "ID": 1
        },
        "Settings": {
            "LoadBalancer": [
                {
                    "VirtualIPAddress": "192.168.200.50",
                    "Port": "80",
                    "DelayLoop": "1000",
                    "SorryServer": "",
                    "Description": "description",
                    "Servers": [
                       {
                            "IPAddress": "192.168.200.51",
                            "Port": "80",
                            "HealthCheck": {
                                "Protocol": "http",
                                "Path": "\/index.html",
                                "Status": "200"
                            },
                            "Enabled": "True",
                            "Status": "DOWN",
                            "ActiveConn": "0"
                        },
                        {
                            "IPAddress": "192.168.200.52",
                            "Port": "80",
                            "HealthCheck": {
                                "Protocol": "ping"
                            },
                            "Enabled": "True"
                        }
                    ]
                }
            ]
        },
        "SettingsHash": "924521e812f96157a83d138c79d423fb",
        "Remark": {
            "Zone": {
                "ID": 31002
            },
            "Switch": {
                "ID": "123456789012"
            },
            "VRRP": {
                "VRID": 1
            },
            "Network": {
                "NetworkMaskLen": 24,
                "DefaultRoute": "192.168.200.1"
            },
            "Servers": [
                {
                    "IPAddress": "192.168.200.11"
                }
            ],
            "Plan": {
                "ID": 1
            }
        },
        "Availability": "available",
        "Instance": {
            "Status": "up",
            "StatusChangedAt": "2016-04-29T18:29:17+09:00"
        },
        "ServiceClass": "cloud\/appliance\/loadbalancer\/1",
        "CreatedAt": "2016-04-29T18:27:18+09:00",
        "Icon": {
            "ID": 9999999999,
            "URL": "https:\/\/secure.sakura.ad.jp\/cloud\/zone\/is1b\/api\/cloud\/1.1\/icon\/112300511981.png",
            "Name": "CentOS",
            "Scope": "shared"
        },
        "Switch": {
            "ID": 112800442260,
            "Name": "\u3059\u3046\u3043\u3063\u3061",
            "Internet": null,
            "Scope": "user",
            "Availability": "available",
            "Zone": {
                "ID": 31002,
                "Name": "is1b",
                "Region": {
                    "ID": 310,
                    "Name": "\u77f3\u72e9"
                }
            }
        },
        "Interfaces": [
            {
                "IPAddress": null,
                "UserIPAddress": "192.168.200.11",
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
            }
        ],
        "Tags": [
            "\u3042\u3042",
            "\u3044\u3044",
            "\u3046\u3046"
        ]
    }
    	`

	createLoadBalancerValues = &CreateLoadBalancerValue{
		SwitchID:     "9999999999",
		VRID:         1,
		Plan:         LoadBalancerPlanStandard,
		IPAddress1:   "192.168.11.11",
		MaskLen:      28,
		DefaultRoute: "192.168.11.1",
		Name:         "TestLoadBalancer",
		Description:  "TestDescription",
		Tags:         []string{"tag1", "tag2", "tag3"},
		Icon:         &Resource{ID: 9999999999},
	}
	loadBalancerSettings = []*LoadBalancerSetting{
		{
			VirtualIPAddress: "192.168.11.101",
			Port:             "8080",
			DelayLoop:        "30",
			SorryServer:      "192.168.11.201",
			Servers: []*LoadBalancerServer{
				{
					IPAddress: "192.168.11.51",
					Port:      "8080",
					HealthCheck: &LoadBalancerHealthCheck{
						Protocol: "http",
						Path:     "/",
						Status:   "200",
					},
				},
				{
					IPAddress: "192.168.11.52",
					Port:      "8080",
					HealthCheck: &LoadBalancerHealthCheck{
						Protocol: "http",
						Path:     "/",
						Status:   "200",
					},
				},
			},
		},
	}
)

func TestMarshalLoadBalancerJSON(t *testing.T) {
	//standard plan
	var lb LoadBalancer
	err := json.Unmarshal([]byte(testLoadBalancerJSON), &lb)

	assert.NoError(t, err)
	assert.NotEmpty(t, lb)

	assert.NotEmpty(t, lb.ID)
	assert.NotEmpty(t, lb.Remark)

	assert.NotEmpty(t, lb.Remark.Servers)
	assert.NotEmpty(t, lb.Remark.Network)
	assert.NotEmpty(t, lb.Remark.Switch)
	assert.NotEmpty(t, lb.Remark.VRRP)
	//TODO Zone
	//assert.NotEmpty(t, lb.Remark.Zone)
	//assert.NotEmpty(t, lb.Remark.Plan)

	assert.NotEmpty(t, lb.Instance)
	assert.NotEmpty(t, lb.Interfaces)

	assert.NotEmpty(t, lb.Settings.LoadBalancer)
	assert.NotEmpty(t, lb.Settings.LoadBalancer[0].VirtualIPAddress)
	assert.NotEmpty(t, lb.Settings.LoadBalancer[0].Description)

	assert.NotEmpty(t, lb.Settings.LoadBalancer[0].Servers)
	assert.NotEmpty(t, lb.Settings.LoadBalancer[0].Servers[0])

	assert.NotEmpty(t, lb.Settings.LoadBalancer[0].Servers[0].IPAddress)
	assert.NotEmpty(t, lb.Settings.LoadBalancer[0].Servers[0].HealthCheck.Protocol)
	assert.NotEmpty(t, lb.Settings.LoadBalancer[0].Servers[0].HealthCheck.Path)
	assert.NotEmpty(t, lb.Settings.LoadBalancer[0].Servers[0].HealthCheck.Status)

	assert.NotEmpty(t, lb.Settings.LoadBalancer[0].Servers[1])
	assert.NotEmpty(t, lb.Settings.LoadBalancer[0].Servers[1].HealthCheck.Protocol)
	assert.Empty(t, lb.Settings.LoadBalancer[0].Servers[1].HealthCheck.Path)
	assert.Empty(t, lb.Settings.LoadBalancer[0].Servers[1].HealthCheck.Status)

}

func TestCreateNewLoadBalancerSingle(t *testing.T) {

	lb, err := CreateNewLoadBalancerSingle(createLoadBalancerValues, loadBalancerSettings)
	assert.NoError(t, err)
	assert.NotEmpty(t, lb)

	assert.Equal(t, lb.Class, "loadbalancer")

	assert.Equal(t, lb.Remark.Switch.ID, "9999999999")
	assert.Equal(t, lb.Remark.VRRP.VRID, 1)
	plan := lb.Plan.ID
	assert.NoError(t, err)
	assert.Equal(t, LoadBalancerPlan(plan), LoadBalancerPlanStandard)

	assert.Equal(t, lb.Remark.Servers[0].(map[string]string)["IPAddress"], "192.168.11.11")
	assert.Equal(t, lb.Remark.Network.NetworkMaskLen, 28)
	assert.Equal(t, lb.Remark.Network.DefaultRoute, "192.168.11.1")
	assert.Equal(t, lb.Name, "TestLoadBalancer")
	assert.Equal(t, lb.Description, "TestDescription")
	assert.Equal(t, lb.Tags, []string{"tag1", "tag2", "tag3"})
	assert.Equal(t, lb.propTags, propTags{Tags: []string{"tag1", "tag2", "tag3"}})
	//assert.Equal(t, lb.Icon.ID, 9999999999)

	assert.Equal(t, len(lb.Settings.LoadBalancer), 1)
	setting := lb.Settings.LoadBalancer[0]

	assert.Equal(t, setting.VirtualIPAddress, "192.168.11.101")
	assert.Equal(t, setting.Port, "8080")
	assert.Equal(t, setting.DelayLoop, "30")
	assert.Equal(t, setting.SorryServer, "192.168.11.201")
	assert.Equal(t, len(setting.Servers), 2)

	assert.Equal(t, setting.Servers[0].IPAddress, "192.168.11.51")
	assert.Equal(t, setting.Servers[0].Port, "8080")
	assert.Equal(t, setting.Servers[0].HealthCheck.Protocol, "http")
	assert.Equal(t, setting.Servers[0].HealthCheck.Path, "/")
	assert.Equal(t, setting.Servers[0].HealthCheck.Status, "200")

	assert.Equal(t, setting.Servers[1].IPAddress, "192.168.11.52")
	assert.Equal(t, setting.Servers[1].Port, "8080")
	assert.Equal(t, setting.Servers[1].HealthCheck.Protocol, "http")
	assert.Equal(t, setting.Servers[1].HealthCheck.Path, "/")
	assert.Equal(t, setting.Servers[1].HealthCheck.Status, "200")
}

func TestCreateNewLoadBalancerDouble(t *testing.T) {
	values := &CreateDoubleLoadBalancerValue{
		CreateLoadBalancerValue: createLoadBalancerValues,
		IPAddress2:              "192.168.11.12",
	}

	lb, err := CreateNewLoadBalancerDouble(values, loadBalancerSettings)
	assert.NoError(t, err)
	assert.NotEmpty(t, lb)

	assert.Equal(t, lb.Class, "loadbalancer")

	assert.Equal(t, lb.Remark.Switch.ID, "9999999999")
	assert.Equal(t, lb.Remark.VRRP.VRID, 1)
	plan := lb.Plan.ID
	assert.NoError(t, err)
	assert.Equal(t, LoadBalancerPlan(plan), LoadBalancerPlanStandard)

	assert.Equal(t, lb.Remark.Servers[0].(map[string]string)["IPAddress"], "192.168.11.11")
	assert.Equal(t, lb.Remark.Servers[1].(map[string]string)["IPAddress"], "192.168.11.12")

	assert.Equal(t, lb.Remark.Network.NetworkMaskLen, 28)
	assert.Equal(t, lb.Remark.Network.DefaultRoute, "192.168.11.1")
	assert.Equal(t, lb.Name, "TestLoadBalancer")
	assert.Equal(t, lb.Description, "TestDescription")
	assert.Equal(t, lb.propTags, propTags{Tags: []string{"tag1", "tag2", "tag3"}})
	//assert.Equal(t, lb.Icon.ID, 9999999999)

	assert.Equal(t, len(lb.Settings.LoadBalancer), 1)
	setting := lb.Settings.LoadBalancer[0]

	assert.Equal(t, setting.VirtualIPAddress, "192.168.11.101")
	assert.Equal(t, setting.Port, "8080")
	assert.Equal(t, setting.DelayLoop, "30")
	assert.Equal(t, setting.SorryServer, "192.168.11.201")
	assert.Equal(t, len(setting.Servers), 2)

	assert.Equal(t, setting.Servers[0].IPAddress, "192.168.11.51")
	assert.Equal(t, setting.Servers[0].Port, "8080")
	assert.Equal(t, setting.Servers[0].HealthCheck.Protocol, "http")
	assert.Equal(t, setting.Servers[0].HealthCheck.Path, "/")
	assert.Equal(t, setting.Servers[0].HealthCheck.Status, "200")

	assert.Equal(t, setting.Servers[1].IPAddress, "192.168.11.52")
	assert.Equal(t, setting.Servers[1].Port, "8080")
	assert.Equal(t, setting.Servers[1].HealthCheck.Protocol, "http")
	assert.Equal(t, setting.Servers[1].HealthCheck.Path, "/")
	assert.Equal(t, setting.Servers[1].HealthCheck.Status, "200")

}
