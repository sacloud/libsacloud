package sacloud

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testBridgeJSON = `
{
            "Index": 0,
            "ID": 123456789012,
            "Name": "sakura_hyb",
            "Description": "\u30cf\u30a4\u30d6\u30ea\u30c3\u30c9\u63a5\u7d9a\u307b\u3052\u307b\u3052",
            "Info": {
                "Switches": [
                    {
                        "ID": 123456789012,
                        "Name": "\u3059\u3046\u3043\u3063\u3061",
                        "Zone": {
                            "ID": 31002,
                            "Name": "is1b"
                        }
                    }
                ]
            },
            "ServiceClass": "cloud\/bridge\/default",
            "CreatedAt": "2016-05-02T10:44:22+09:00",
            "Region": ` + testRegionJSON + `,
            "SwitchInZone": {
                "ID": 123456789012,
                "Name": "\u3059\u3046\u3043\u3063\u3061",
                "ServerCount": 1,
                "ApplianceCount": 2,
                "Scope": "user"
            }
        }
`

func TestMarshalBridgeJSON(t *testing.T) {
	var br Bridge
	err := json.Unmarshal([]byte(testBridgeJSON), &br)

	assert.NoError(t, err)
	assert.NotEmpty(t, br)

	assert.NotEmpty(t, br.ID)
	assert.NotEmpty(t, br.Info)
	assert.NotEmpty(t, br.Region)
	assert.NotEmpty(t, br.SwitchInZone)

}
