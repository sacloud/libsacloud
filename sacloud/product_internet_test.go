package sacloud

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testInternetPlanJSON = `
{
	"Index": 0,
	"ID": 100,
	"Name": "100Mbps\u5171\u6709",
	"BandWidthMbps": 100,
	"ServiceClass": "cloud\/internet\/router\/100m",
	"Availability": "available"
}
`

func TestMarshalProductInternetJSON(t *testing.T) {
	var productInternet ProductInternet
	err := json.Unmarshal([]byte(testInternetPlanJSON), &productInternet)

	assert.NoError(t, err)
	assert.NotEmpty(t, productInternet)

	assert.NotEmpty(t, productInternet.ID)
}
