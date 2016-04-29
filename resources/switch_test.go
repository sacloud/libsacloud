package resources

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

var testPublicNetSwitchJSON = `
{
	"ID": "123456789012",
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
	"ID": "123456789012",
	"Name": "\u3059\u3046\u3043\u3063\u3061",
	"Scope": "user",
	"Subnet": null,
	"UserSubnet": {
		"DefaultRoute": "192.168.200.1",
		"NetworkMaskLen": 24
	}
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
