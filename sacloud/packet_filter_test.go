package sacloud

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

var testPacketFilterJSON = `
{
	"ID": "123456789012",
	"Name": "\u307b\u3052\u307b\u3052\uff12",
	"RequiredHostVersion": 103
}
`

func TestMarshalPacketFilterJSON(t *testing.T) {
	var packetFilter PacketFilter
	err := json.Unmarshal([]byte(testPacketFilterJSON), &packetFilter)

	assert.NoError(t, err)
	assert.NotEmpty(t, packetFilter)

	assert.NotEmpty(t, packetFilter.ID)
}
