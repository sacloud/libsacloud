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
	"time"

	"github.com/stretchr/testify/assert"
)

var testProxyLBJSON = `

{
    "ID": 123456789012,
    "Name": "example",
    "Description": "example",
    "Settings": {
      "ProxyLB": {
        "HealthCheck": {
          "Protocol": "http",
          "Path": "/",
          "Host": "example.com",
          "DelayLoop": 10
        },
        "SorryServer": {
          "IPAddress": "192.2.0.1",
          "Port": 80
        },
        "BindPorts": [
          {
            "ProxyMode": "https",
            "Port": 443
          }
        ],
        "Servers": [
          {
            "IPAddress": "192.2.0.11",
            "Port": 443,
            "Enabled": true
          },
          {
            "IPAddress": "192.2.0.12",
            "Port": 443,
            "Enabled": true
          }
        ]
      }
    },
    "Status": {
      "FQDN": "site-xxxxxxx.proxylbN.sakura.ne.jp",
      "ProxyNetworks": [
        "192.51.0.0/28"
      ],
      "UseVIPFailover": true
    },
    "ServiceClass": "cloud/proxylb/plain/1000",
    "Availability": "available",
    "CreatedAt": "2019-02-28T11:15:59+09:00",
    "ModifiedAt": "2019-02-28T11:15:59+09:00",
    "Provider": {
      "ID": 9100001,
      "Class": "proxylb",
      "Name": "proxylb1",
      "ServiceClass": "cloud/proxylb"
    },
    "Tags": [
      "tag1",
      "tag2"
    ]
  }
`

var testProxyLBCertificatesJSON = `
	{
		"ServerCertificate": "dummy1",
		"IntermediateCertificate": "dummy2",
    	"PrivateKey": "dummy3",
    	"CertificateEndDate": "May  4 01:37:47 2019 GMT",
		"CertificateCommonName": ""
	}
`

func TestMarshalProxyLBJSON(t *testing.T) {
	var proxyLB ProxyLB
	err := json.Unmarshal([]byte(testProxyLBJSON), &proxyLB)

	assert.NoError(t, err)
	assert.NotEmpty(t, proxyLB)

	assert.NotEmpty(t, proxyLB.ID)
	assert.NotEmpty(t, proxyLB.Status.FQDN)
	assert.NotEmpty(t, proxyLB.Status.ProxyNetworks)
	assert.True(t, proxyLB.Status.UseVIPFailover)
	assert.NotEmpty(t, proxyLB.Provider.Class)
}

func TestMarshalProxyLBCertificate(t *testing.T) {
	var certs ProxyLBCertificate
	err := json.Unmarshal([]byte(testProxyLBCertificatesJSON), &certs)

	assert.NoError(t, err)
	assert.NotEmpty(t, certs)

	assert.Equal(t, "dummy1", certs.ServerCertificate)
	assert.Equal(t, "dummy2", certs.IntermediateCertificate)
	assert.Equal(t, "dummy3", certs.PrivateKey)
	loc, _ := time.LoadLocation("GMT")
	assert.Equal(t, time.Date(2019, 5, 4, 1, 37, 47, 0, loc).Unix(), certs.CertificateEndDate.Unix())
}

func TestMarshalProxyLBCertificates(t *testing.T) {

	t.Run("AdditionalCerts is empty", func(t *testing.T) {
		data := `{	
			"AdditionalCerts": "",
			"PrimaryCert": {
				"CertificateCommonName": "",
				"CertificateEndDate": "",
				"IntermediateCertificate": "",
				"PrivateKey": "",
				"ServerCertificate": ""
			}
		}`

		res := &ProxyLBCertificates{}
		err := json.Unmarshal([]byte(data), res)
		assert.NoError(t, err)
	})

	t.Run("AdditionalCerts is array", func(t *testing.T) {
		data := `{	
			"AdditionalCerts": [
				{
					"CertificateCommonName": "bbb",
					"CertificateEndDate": "",
					"IntermediateCertificate": "",
					"PrivateKey": "",
					"ServerCertificate": ""
				},
				{
					"CertificateCommonName": "ccc",
					"CertificateEndDate": "",
					"IntermediateCertificate": "",
					"PrivateKey": "",
					"ServerCertificate": ""
				}
			],
			"PrimaryCert": {
				"CertificateCommonName": "aaa",
				"CertificateEndDate": "",
				"IntermediateCertificate": "",
				"PrivateKey": "",
				"ServerCertificate": ""
			}
		}`

		var res ProxyLBCertificates
		err := json.Unmarshal([]byte(data), &res)
		assert.NoError(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, "aaa", res.PrimaryCert.CertificateCommonName)
		assert.NotNil(t, res.AdditionalCerts)
		assert.Len(t, res.AdditionalCerts, 2)
	})
}
