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
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

const testMobileGatewayJSON = `
{
    "ID": 123456789012,
    "Class": "mobilegateway",
    "Name": "example-mgw",
    "Description": "description",
    "Plan": {
      "ID": 2
    },
    "Settings": {
      "MobileGateway": {
        "Interfaces": [
          null,
          {
            "IPAddress": [
              "192.168.0.1"
            ],
            "NetworkMaskLen": 24
          }
        ],
        "InternetConnection": {
          "Enabled": "True"
        },
        "StaticRoutes": [
          {
            "Prefix": "172.16.1.0/24",
            "NextHop": "192.168.0.254"
          },
          {
            "Prefix": "172.16.2.0/24",
            "NextHop": "192.168.0.254"
          }
        ]
      }
    },
    "SettingsHash": "6dab2e7a0a47bdec844d36db79420b76",
    "Remark": {
      "Servers": [
        []
      ],
      "Switch": {
        "Scope": "shared"
      },
      "Zone": {
        "ID": "30293"
      }
    },
    "Availability": "available",
    "Instance": {
      "Status": "down",
      "StatusChangedAt": null
    },
    "ServiceClass": "cloud/appliance/mobilegateway/1",
    "CreatedAt": "2018-04-04T16:04:14+09:00",
    "Icon": null,
    "Switch": null,
    "Interfaces": [
      {
        "IPAddress": "133.242.32.y",
        "UserIPAddress": null,
        "HostName": null,
        "Switch": {
          "ID": 123456789012,
          "Name": "スイッチ",
          "Scope": "shared",
          "Subnet": {
            "NetworkAddress": "133.242.32.0",
            "NetworkMaskLen": 24,
            "DefaultRoute": "133.242.32.1",
            "Internet": {
              "BandWidthMbps": 100
            }
          },
          "UserSubnet": {
            "DefaultRoute": "133.242.32.1",
            "NetworkMaskLen": 24
          }
        }
      },
      null
    ],
    "Tags": []
}
`

const testTrafficMonitoringJSON = `
{
  "traffic_quota_in_mb": 512,
  "bandwidth_limit_in_kbps": 64,
  "email_config": {
    "enabled": true
  },
  "slack_config": {
    "enabled": true,
    "slack_url": "https://hooks.slack.com/services/xxxxxxx/xxxxx/xxxx"
  },
  "auto_traffic_shaping": true
}
`

const testTrafficStatusJSON = `
{
  "uplink_bytes": "9223372036854775808",
  "downlink_bytes": "21989271821",
  "traffic_shaping": true 
}
`

func TestUnmarshalMGWJSON(t *testing.T) {
	var mgw MobileGateway
	err := json.Unmarshal([]byte(testMobileGatewayJSON), &mgw)
	assert.NoError(t, err)
	assert.NotEmpty(t, mgw)
}

func TestUnmarshalTrafficMonitoring(t *testing.T) {
	var tm TrafficMonitoringConfig
	err := json.Unmarshal([]byte(testTrafficMonitoringJSON), &tm)
	assert.NoError(t, err)
	assert.NotEmpty(t, tm)
	assert.Equal(t, 512, tm.TrafficQuotaInMB)
	assert.Equal(t, 64, tm.BandWidthLimitInKbps)
	assert.NotNil(t, tm.EMailConfig)
	assert.Equal(t, tm.EMailConfig.Enabled, true)
	assert.NotNil(t, tm.SlackConfig)
	assert.Equal(t, tm.SlackConfig.Enabled, true)
	assert.Equal(t, tm.SlackConfig.IncomingWebhooksURL, "https://hooks.slack.com/services/xxxxxxx/xxxxx/xxxx")
	assert.Equal(t, true, tm.AutoTrafficShaping)
}

func TestUnmarshalTrafficStatus(t *testing.T) {
	var ts TrafficStatus
	err := json.Unmarshal([]byte(testTrafficStatusJSON), &ts)

	assert.NoError(t, err)
	assert.NotEmpty(t, ts)

	uplink := uint64(math.MaxInt64) + uint64(1)
	assert.Equal(t, uplink, ts.UplinkBytes)
	assert.Equal(t, uint64(21989271821), ts.DownlinkBytes)
	assert.Equal(t, true, ts.TrafficShaping)

}
