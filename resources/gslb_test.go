package resources

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

var testGSLBJSON = `
        {
            "ID": "123456789012",
            "Name": "hogeGSLB",
            "Description": "GSLB\u306e\u307b\u3052\u3067\u3059",
            "Settings": {
                "GSLB": {
                    "DelayLoop": 10,
                    "HealthCheck": {
                        "Protocol": "ping"
                    },
                    "Weighted": "True",
                    "Servers": [
                        {
                            "IPAddress": "133.242.224.11",
                            "Enabled": "True",
                            "Weight": "20"
                        },
                        {
                            "IPAddress": "133.242.224.12",
                            "Enabled": "True",
                            "Weight": "30"
                        }
                    ]
                }
            },
            "Status": {
                "FQDN": "site-123456789012.gslb3.sakura.ne.jp"
            },
            "ServiceClass": "cloud\/gslb",
            "CreatedAt": "2016-04-29T22:38:50+09:00",
            "ModifiedAt": "2016-04-29T22:38:50+09:00",
            "Provider": {
                "ID": 1000002,
                "Class": "gslb",
                "Name": "gslb3.sakura.ne.jp",
                "ServiceClass": "cloud\/gslb"
            },
            "Icon": {
                "ID": "112800442534",
                "URL": "https:\/\/secure.sakura.ad.jp\/cloud\/zone\/is1b\/api\/cloud\/1.1\/icon\/112800442534.png",
                "Name": "\u30c6\u30b9\u30c8\u30a2\u30a4\u30b3\u30f3",
                "Scope": "user"
            },
            "Tags": [
                "hoge",
                "hoge2"
            ]
        }

`

func TestMarshalGSLBJSON(t *testing.T) {
	var gslb CommonServiceGSLBItem
	err := json.Unmarshal([]byte(testGSLBJSON), &gslb)

	assert.NoError(t, err)
	assert.NotEmpty(t, gslb)

	assert.NotEmpty(t, gslb.ID)
	assert.NotEmpty(t, gslb.Status.FQDN)
	assert.NotEmpty(t, gslb.Provider.Class)
	assert.NotEmpty(t, gslb.Settings.GSLB.Servers[0].IPAddress)
}
