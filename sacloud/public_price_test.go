package sacloud

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

var testPublicPriceJSON = `
{
	"DisplayName": "\u30d7\u30e9\u30f311(\u30c7\u30a3\u30b9\u30af\u306a\u3057)",
	"IsPublic": true,
	"Price": {
		"Base": 0,
		"Daily": 3270,
		"Hourly": 326,
		"Monthly": 65397,
		"Zone": "is1b"
	},
	"ServiceClassID": 50060,
	"ServiceClassName": "plan\/11",
	"ServiceClassPath": "cloud\/plan\/11"
}
`

func TestMarshalPublicPriceJSON(t *testing.T) {
	var publicPrice PublicPrice
	err := json.Unmarshal([]byte(testPublicPriceJSON), &publicPrice)

	assert.NoError(t, err)
	assert.NotEmpty(t, publicPrice)

	assert.NotEmpty(t, publicPrice.DisplayName)
	assert.NotEmpty(t, publicPrice.Price)
}
