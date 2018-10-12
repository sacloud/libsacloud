package sacloud

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testPublicInterfaceJSON = `
{
	"ID": 112800430558,
	"MACAddress": "9C:A3:BA:30:E4:D0",
	"IPAddress": "133.242.224.41",
	"UserIPAddress": null,
	"HostName": null,
	"Switch": ` + testPublicNetSwitchJSON + `,
	"PacketFilter": ` + testPacketFilterJSON + `
}
`

var testPrivateInterfaceJSON = `
{
	"ID": 112800442276,
	"MACAddress": "9C:A3:BA:30:86:9C",
	"IPAddress": null,
	"UserIPAddress": "192.168.200.50",
	"HostName": null,
	"Switch": ` + testPrivateNetSwitchJSON + `,
	"PacketFilter": null
}
`

func TestMarshalInterfaceJSON(t *testing.T) {
	var publicInterface, privateInterface Interface

	err := json.Unmarshal([]byte(testPublicInterfaceJSON), &publicInterface)
	assert.NoError(t, err)
	assert.NotEmpty(t, publicInterface)
	assert.NotEmpty(t, publicInterface.ID)
	assert.NotEmpty(t, publicInterface.Switch.ID)
	assert.NotEmpty(t, publicInterface.PacketFilter.ID)

	err = json.Unmarshal([]byte(testPrivateInterfaceJSON), &privateInterface)
	assert.NoError(t, err)
	assert.NotEmpty(t, privateInterface)
	assert.NotEmpty(t, privateInterface.ID)
	assert.NotEmpty(t, privateInterface.Switch.ID)
	assert.Nil(t, privateInterface.PacketFilter)

}
