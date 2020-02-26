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

var testWebAccelSiteJSON = `
{
      "ID": "123456789012",
      "Name": "サイト1",
      "DomainType": "own_domain",
      "Domain": "cdn1.example.com",
      "Subdomain": "0f5cty4g.user.webaccel.jp",
      "ASCIIDomain": "cdn1.example.com",
      "Origin": "198.51.100.1",
      "HostHeader": "cdn2.example.com",
      "Status": "enabled",
      "CreatedAt": "2015-11-13T02:56:01+09:00",
      "HasCertificate": true,
      "HasOldCertificate": true,
      "GibSentInLastWeek": 80,
      "CertValidNotBefore": 1457568000000,
      "CertValidNotAfter": 1526558400000
}
`

var testWebAccelCertResponseValid = `
{
  "Certificate": {
    "Current":{
      "ID": "1",
      "SiteID": "123456789012",
      "CertificateChain": "-----BEGIN CERTIFICATE-----・・・・・",
      "Key": "-----BEGIN RSA PRIVATE KEY-----・・・・・",
      "CreatedAt": "2015-11-13T02:57:01+09:00",
      "UpdatedAt": "2015-11-14T02:57:01+09:00"
    },
    "Old": [
      {
        "ID": "1",
        "SiteID": "123456789012",
        "CertificateChain": "-----BEGIN CERTIFICATE-----・・・・・",
        "CreatedAt": "2015-11-13T02:57:01+09:00",
        "UpdatedAt": "2015-11-14T02:57:01+09:00"
      }
    ]
  },
  "Success": true,
  "is_ok": true
}
`
var testWebAccelCertResponseNotExists = `
{
  "Certificate": [],
  "Success": true,
  "is_ok": true
}
`

func TestMarshalWebAccelSiteJSON(t *testing.T) {
	var site WebAccelSite

	err := json.Unmarshal([]byte(testWebAccelSiteJSON), &site)
	assert.NoError(t, err)
	assert.NotEmpty(t, site.ID)
	assert.NotEmpty(t, site.Name)
	assert.NotEmpty(t, site.DomainType)
	assert.NotEmpty(t, site.Domain)
	assert.NotEmpty(t, site.Subdomain)
	assert.NotEmpty(t, site.ASCIIDomain)
	assert.NotEmpty(t, site.Origin)
	assert.NotEmpty(t, site.HostHeader)
	assert.NotEmpty(t, site.Status)
	assert.NotEmpty(t, site.CreatedAt)
	assert.True(t, site.HasCertificate)
	assert.True(t, site.HasOldCertificate)
	assert.NotEmpty(t, site.GibSentInLastWeek)
	assert.NotEmpty(t, site.CertValidNotBefore)
	assert.NotEmpty(t, site.CertValidNotAfter)

}

func TestMarshalWebAccelCertResponseJSON(t *testing.T) {
	t.Run("Has cert response", func(t *testing.T) {
		var res WebAccelCertResponse
		err := json.Unmarshal([]byte(testWebAccelCertResponseValid), &res)

		assert.NoError(t, err)
		assert.NotNil(t, res.Certificate)
		assert.NotNil(t, res.Certificate.Current)
		assert.True(t, len(res.Certificate.Old) > 0)
	})
	t.Run("Not exists response", func(t *testing.T) {
		var res WebAccelCertResponse
		err := json.Unmarshal([]byte(testWebAccelCertResponseNotExists), &res)

		assert.NoError(t, err)
		assert.Nil(t, res.Certificate)
	})
}
