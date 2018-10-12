package sacloud

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testRegionJSON = `
{
	    "ID": 310,
	    "Name": "\u77f3\u72e9",
	    "Description": "\u77f3\u72e9",
	    "NameServers": [
		"133.242.0.3",
		"133.242.0.4"
	    ]
}
`

func TestMarshalRegionJSON(t *testing.T) {
	var region Region
	err := json.Unmarshal([]byte(testRegionJSON), &region)

	assert.NoError(t, err)
	assert.NotEmpty(t, region)

	assert.NotEmpty(t, region.ID)
}
