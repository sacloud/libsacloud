package sacloud

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testDNSJSON = `
 {
	"ID": 123456789012,
	"Name": "libsacloud-test.com",
	"Description": "",
	"Settings": {
		"DNS": {
			"ResourceRecordSets": [
				{
				    "Name": "libsacloud-cname-record",
				    "Type": "CNAME",
				    "RData": "site-112800377587.gslb3.sakura.ne.jp.",
				    "TTL": 10
				},
				{
				    "Name": "libsacloud-a-record",
				    "Type": "A",
				    "RData": "133.242.224.254"
				}
			]
		}
	},
	"Status": {
		"Zone": "libsacloud-test.com",
		"NS": [
		    "ns1.gslb2.sakura.ne.jp",
		    "ns2.gslb2.sakura.ne.jp"
		]
	},
	"ServiceClass": "cloud\/dns",
	"CreatedAt": "2016-01-22T11:48:17+09:00",
	"ModifiedAt": "2016-01-22T11:48:17+09:00",
	"Provider": {
		"ID": 2000001,
		"Class": "dns",
		"Name": "gslb2.sakura.ne.jp",
		"ServiceClass": "cloud\/dns"
	},
	"Icon": {
		"ID": 112300511382,
		"URL": "https:\/\/secure.sakura.ad.jp\/cloud\/zone\/is1b\/api\/cloud\/1.1\/icon\/112300511382.png",
		"Name": "DNS",
		"Scope": "shared"
	},
	"Tags": []
}

`

func TestMarshalDNSJSON(t *testing.T) {
	var dns DNS
	err := json.Unmarshal([]byte(testDNSJSON), &dns)

	assert.NoError(t, err)
	assert.NotEmpty(t, dns)

	assert.NotEmpty(t, dns.ID)
	assert.NotEmpty(t, dns.Status.Zone)
	assert.NotEmpty(t, dns.Provider.Class)
	assert.NotEmpty(t, dns.Settings.DNS.ResourceRecordSets[0].Name)
}

func TestDnsRecordSets(t *testing.T) {
	records := DNSRecordSets{}
	assert.True(t, len(records.ResourceRecordSets) == 0)

	records.AddDNSRecordSet("test1", "192.168.0.1")
	assert.True(t, len(records.ResourceRecordSets) == 1)
	assert.Equal(t, records.ResourceRecordSets[0].RData, "192.168.0.1")
	//t.Logf("records:%#v", records)

	records.AddDNSRecordSet("test2", "192.168.0.2")
	assert.True(t, len(records.ResourceRecordSets) == 2)
	assert.Equal(t, records.ResourceRecordSets[1].RData, "192.168.0.2")
	//t.Logf("records:%#v", records)

	records.AddDNSRecordSet("test1", "192.168.0.3")
	assert.True(t, len(records.ResourceRecordSets) == 2)
	assert.Equal(t, records.ResourceRecordSets[0].RData, "192.168.0.3")
	//t.Logf("records:%#v", records)

	records.DeleteDNSRecordSet("test1", "192.168.0.1")
	assert.True(t, len(records.ResourceRecordSets) == 2)
	assert.Equal(t, records.ResourceRecordSets[0].RData, "192.168.0.3")

	records.DeleteDNSRecordSet("test1", "192.168.0.3")
	assert.True(t, len(records.ResourceRecordSets) == 1)
	assert.Equal(t, records.ResourceRecordSets[0].RData, "192.168.0.2")

}
