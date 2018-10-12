package sacloud

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testZoneJSON = `
{
	"ID": 31002,
	"DisplayOrder": 20031002,
	"Name": "is1b",
	"Description": "\u77f3\u72e9\u7b2c2\u30be\u30fc\u30f3",
	"IsDummy": false,
	"VNCProxy": {
	    "HostName": "sac-is1b-ssl.sakura.ad.jp",
	    "IPAddress": "133.242.239.244"
	},
	"FTPServer": {
	    "HostName": "sac-is1b-ssl.sakura.ad.jp",
	    "IPAddress": "133.242.239.244"
	},
	"Settings": {
	    "Subnet": {
		"Plan": {
		    "Member": [28,27,26],
		    "Staff": [28,27,26,25,24]
		}
	    }
	},
	"Region": ` + testRegionJSON + `
}
`

func TestMarshalZoneJSON(t *testing.T) {
	var zone Zone
	err := json.Unmarshal([]byte(testZoneJSON), &zone)

	assert.NoError(t, err)
	assert.NotEmpty(t, zone)

	assert.NotEmpty(t, zone.ID)
	assert.NotEmpty(t, zone.VNCProxy.HostName)
	assert.NotEmpty(t, zone.FTPServer.HostName)
	assert.NotEmpty(t, zone.Region.ID)
	assert.NotEmpty(t, zone.Region.NameServers)
}
