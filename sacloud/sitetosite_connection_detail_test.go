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
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testSiteToSiteConnectionDetailJSON = `
   {
      "ESP": {
          "AuthenticationProtocol": "sha1",
          "DHGroup": "modp1024",
          "EncryptionProtocol": "aes128",
          "Lifetime": "28800",
          "Mode": "tunnel",
          "PerfectForwardSecrecy": "enabled"
      },
      "IKE": {
          "AuthenticationProtocol": "sha1",
          "EncryptionProtocol": "aes128",
          "Lifetime": "1800",
          "Mode": "main",
          "PerfectForwardSecrecy": "enabled",
          "PreSharedSecret": "preShared"
      },
      "Index": 1,
      "Peer": {
          "ID": "8.8.8.8",
          "InsideNetworks": [
              "10.0.0.0/24"
          ],
          "OutsideIPAddress": "8.8.8.8"
      },
      "VPCRouter": {
          "ID": "133.242.231.180",
          "InsideNetworks": [
              "192.168.150.0/26"
          ],
          "OutsideIPAddress": "133.242.231.180"
      }
  }

`

var testSiteToSiteInfoJSON = fmt.Sprintf(`{ "Details": { "Config" : [ %s ] }}`, testSiteToSiteConnectionDetailJSON)
var testSiteToSiteInfoNullJSON = `{ "Details": { "Config" : null }}`

func TestSiteToSiteInfoMarshalJSON(t *testing.T) {
	// has value JSON
	var info SiteToSiteConnectionInfo
	err := json.Unmarshal([]byte(testSiteToSiteInfoJSON), &info)

	assert.NoError(t, err)
	assert.NotNil(t, info)
	assert.NotNil(t, info.Details.Config)
	assert.Len(t, info.Details.Config, 1)

	err = json.Unmarshal([]byte(testSiteToSiteInfoNullJSON), &info)
	assert.NoError(t, err)
	assert.NotNil(t, info)
	assert.Len(t, info.Details.Config, 0)
}

func TestSiteToSiteConnectionDetailMarshalJSON(t *testing.T) {

	var info SiteToSiteConnectionDetail
	err := json.Unmarshal([]byte(testSiteToSiteConnectionDetailJSON), &info)

	assert.NoError(t, err)
	assert.NotEmpty(t, info)

	assert.NotEmpty(t, info.ESP)
	assert.NotEmpty(t, info.ESP.AuthenticationProtocol)
	assert.NotEmpty(t, info.ESP.DHGroup)
	assert.NotEmpty(t, info.ESP.EncryptionProtocol)
	assert.NotEmpty(t, info.ESP.Lifetime)
	assert.NotEmpty(t, info.ESP.Mode)
	assert.NotEmpty(t, info.ESP.PerfectForwardSecrecy)

	assert.NotEmpty(t, info.IKE)
	assert.NotEmpty(t, info.IKE.AuthenticationProtocol)
	assert.NotEmpty(t, info.IKE.EncryptionProtocol)
	assert.NotEmpty(t, info.IKE.Lifetime)
	assert.NotEmpty(t, info.IKE.Mode)
	assert.NotEmpty(t, info.IKE.PerfectForwardSecrecy)

	assert.NotEmpty(t, info.Peer)
	assert.NotEmpty(t, info.Peer.ID)
	assert.NotEmpty(t, info.Peer.InsideNetworks)
	assert.NotEmpty(t, info.Peer.OutsideIPAddress)

	assert.NotEmpty(t, info.VPCRouter)
	assert.NotEmpty(t, info.VPCRouter.ID)
	assert.NotEmpty(t, info.VPCRouter.InsideNetworks)
	assert.NotEmpty(t, info.VPCRouter.OutsideIPAddress)
}
