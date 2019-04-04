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
    	"CertificateEndDate": "May  4 01:37:47 2019 GMT"
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

func TestMarshalProxyLBCertificates(t *testing.T) {
	var certs ProxyLBCertificate
	err := json.Unmarshal([]byte(testProxyLBCertificatesJSON), &certs)

	assert.NoError(t, err)
	assert.NotEmpty(t, certs)

	assert.Equal(t, "dummy1", certs.ServerCertificate)
	assert.Equal(t, "dummy2", certs.IntermediateCertificate)
	assert.Equal(t, "dummy3", certs.PrivateKey)
	loc, _ := time.LoadLocation("GMT")
	assert.Equal(t, time.Date(2019, 5, 4, 1, 37, 47, 0, loc), certs.CertificateEndDate)
}
