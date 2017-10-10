package sacloud

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

var testPacketFilterJSON = `
{
	"ID": 123456789012,
	"Name": "\u307b\u3052\u307b\u3052\uff12",
	"RequiredHostVersion": 103,
	"Expression": [
                {
                    "Protocol": "tcp",
                    "SourceNetwork": null,
                    "SourcePort": null,
                    "DestinationPort": "22",
                    "Action": "allow"
                },
                {
                    "Protocol": "tcp",
                    "SourceNetwork": null,
                    "SourcePort": null,
                    "DestinationPort": "80",
                    "Action": "allow"
                },
                {
                    "Protocol": "tcp",
                    "SourceNetwork": null,
                    "SourcePort": null,
                    "DestinationPort": "443",
                    "Action": "allow"
                },
                {
                    "Protocol": "tcp",
                    "SourceNetwork": null,
                    "SourcePort": null,
                    "DestinationPort": "2376",
                    "Action": "allow"
                },
                {
                    "Protocol": "tcp",
                    "SourceNetwork": null,
                    "SourcePort": null,
                    "DestinationPort": "32768-61000",
                    "Action": "allow"
                },
                {
                    "Protocol": "udp",
                    "SourceNetwork": null,
                    "SourcePort": null,
                    "DestinationPort": "32768-61000",
                    "Action": "allow"
                },
                {
                    "Protocol": "fragment",
                    "SourceNetwork": null,
                    "Action": "allow"
                },
                {
                    "Protocol": "icmp",
                    "SourceNetwork": null,
                    "Action": "allow"
                },
                {
                    "Protocol": "ip",
                    "SourceNetwork": null,
                    "Action": "deny"
                }
            ]
}
`

func TestMarshalPacketFilterJSON(t *testing.T) {
	var packetFilter PacketFilter
	err := json.Unmarshal([]byte(testPacketFilterJSON), &packetFilter)

	assert.NoError(t, err)
	assert.NotEmpty(t, packetFilter)

	assert.NotEmpty(t, packetFilter.ID)
}

func TestPacketFilterRuleFuncs(t *testing.T) {
	pf := CreateNewPacketFilter()

	pf.AddTCPRuleAt("", "", "", "", true, 0)
	assert.Len(t, pf.Expression, 1)
	assert.Equal(t, pf.Expression[0].Protocol, "tcp")

	pf.AddUDPRule("", "", "", "", true)
	assert.Len(t, pf.Expression, 2)
	assert.Equal(t, pf.Expression[1].Protocol, "udp")

	pf.AddTCPRuleAt("", "", "", "", true, 1)
	assert.Len(t, pf.Expression, 3)
	assert.Equal(t, pf.Expression[0].Protocol, "tcp")
	assert.Equal(t, pf.Expression[1].Protocol, "tcp")
	assert.Equal(t, pf.Expression[2].Protocol, "udp")

	pf.AddTCPRuleAt("", "", "", "", true, 5)
	assert.Len(t, pf.Expression, 4)
	assert.Equal(t, pf.Expression[0].Protocol, "tcp")
	assert.Equal(t, pf.Expression[1].Protocol, "tcp")
	assert.Equal(t, pf.Expression[2].Protocol, "udp")
	assert.Equal(t, pf.Expression[3].Protocol, "tcp")

	pf.RemoveRuleAt(2)
	assert.Len(t, pf.Expression, 3)
	assert.Equal(t, pf.Expression[0].Protocol, "tcp")
	assert.Equal(t, pf.Expression[1].Protocol, "tcp")
	assert.Equal(t, pf.Expression[2].Protocol, "tcp")

	pf.RemoveRuleAt(5)
	assert.Len(t, pf.Expression, 2)
	assert.Equal(t, pf.Expression[0].Protocol, "tcp")
	assert.Equal(t, pf.Expression[1].Protocol, "tcp")
}
