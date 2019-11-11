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
