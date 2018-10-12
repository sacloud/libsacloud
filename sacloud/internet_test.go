package sacloud

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testInternetJSON = `
{
            "Index": 0,
            "ID": 112800452760,
            "Name": "\u308b\u30fc\u305f",
            "Description": "\u30eb\u30fc\u30bf\u306e\u8aac\u660e",
            "BandWidthMbps": 100,
            "NetworkMaskLen": 28,
            "Scope": "user",
            "ServiceClass": "cloud\/internet\/router\/100m",
            "CreatedAt": "2016-05-02T10:51:49+09:00",
            "Icon": ` + testIconJSON + `,
            "Zone": ` + testZoneJSON + `,
            "Switch": {
                "ID": 112800452761,
                "Name": "\u308b\u30fc\u305f",
                "Scope": "user",
                "UserSubnet": null,
                "HybridConnection": null,
                "Subnets": [
                    {
                        "ID": 3981,
                        "NetworkAddress": "133.242.253.96",
                        "NetworkMaskLen": 28,
                        "DefaultRoute": "133.242.253.97",
                        "NextHop": null,
                        "StaticRoute": null
                    }
                ],
                "IPv6Nets": [
                    {
                        "ID": 216,
                        "IPv6Prefix": "2401:2500:10a:101e::",
                        "IPv6PrefixLen": 64
                    }
                ],
                "Bridge": null
            },
            "Tags": [
                "hoge",
                "hoge2"
            ]
        }
`

func TestMarshalInternetJSON(t *testing.T) {
	var router Internet
	err := json.Unmarshal([]byte(testInternetJSON), &router)

	assert.NoError(t, err)
	assert.NotEmpty(t, router)

	assert.NotEmpty(t, router.ID)
	assert.NotEmpty(t, router.Scope)
	assert.NotEmpty(t, router.Icon)
	//TODO Zone
	//assert.NotEmpty(t, router.Zone)
	assert.NotEmpty(t, router.Switch)
}
