package sacloud

import (
	"encoding/json"
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
      "ID": 1
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

func TestUnmarshalMGWJSON(t *testing.T) {

	var mgw MobileGateway
	err := json.Unmarshal([]byte(testMobileGatewayJSON), &mgw)
	assert.NoError(t, err)
	assert.NotEmpty(t, mgw)

}
