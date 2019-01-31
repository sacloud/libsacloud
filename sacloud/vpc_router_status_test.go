package sacloud

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

var vpcRouterStatusJSON = `
{
    "FirewallSendLogs": [
    	"slog1",
    	"slog2"
    ],
    "FirewallReceiveLogs": [
    	"rlog1",
    	"rlog2"
    ],
    "VPNLogs": [
    	"vlog1",
    	"vlog2"
   ],
    "SessionCount": 7,
    "DHCPServerLeases": [
      {
        "IPAddress": "192.168.1.210",
        "MACAddress": "xx:xx:xx:xx:xx:xx"
      },
      {
        "IPAddress": "192.168.1.211",
        "MACAddress": "xx:xx:xx:xx:xx:xx"
      }
    ],
    "L2TPIPsecServerSessions": [
      {
        "User": "micho",
        "IPAddress": "192.168.2.200",
        "TimeSec": 300
      },
      {
        "User": "micho2",
        "IPAddress": "192.168.2.201",
        "TimeSec": 6000
      }
    ],
    "PPTPServerSessions": [
      {
        "User": "micho3",
        "IPAddress": "192.168.2.202",
        "TimeSec": 500
      },
      {
        "User": "micho4",
        "IPAddress": "192.168.2.203",
        "TimeSec": 7000
      }
    ],
    "SiteToSiteIPsecVPNPeers": [
      {
        "Status": "DOWN",
        "Peer": "1.1.1.1"
      },
      {
        "Status": "UP",
        "Peer": "1.1.1.2"
      }
    ]
}`

func TestMarshalVPCRouterStatusJSON(t *testing.T) {
	// ping
	var vpcRouterStatus VPCRouterStatus
	err := json.Unmarshal([]byte(vpcRouterStatusJSON), &vpcRouterStatus)

	assert.NoError(t, err)
	assert.Len(t, vpcRouterStatus.FirewallSendLogs, 2)
	assert.Len(t, vpcRouterStatus.FirewallReceiveLogs, 2)
	assert.Len(t, vpcRouterStatus.VPNLogs, 2)
	assert.Len(t, vpcRouterStatus.DHCPServerLeases, 2)
	assert.Len(t, vpcRouterStatus.L2TPIPsecServerSessions, 2)
	assert.Len(t, vpcRouterStatus.PPTPServerSessions, 2)
	assert.Len(t, vpcRouterStatus.SiteToSiteIPsecVPNPeers, 2)

	assert.EqualValues(t, "slog1", vpcRouterStatus.FirewallSendLogs[0])
	assert.EqualValues(t, "rlog1", vpcRouterStatus.FirewallReceiveLogs[0])
	assert.EqualValues(t, "vlog1", vpcRouterStatus.VPNLogs[0])

	assert.EqualValues(t, "192.168.1.210", vpcRouterStatus.DHCPServerLeases[0].IPAddress)
	assert.EqualValues(t, "xx:xx:xx:xx:xx:xx", vpcRouterStatus.DHCPServerLeases[0].MACAddress)

	assert.EqualValues(t, "micho", vpcRouterStatus.L2TPIPsecServerSessions[0].User)
	assert.EqualValues(t, "192.168.2.200", vpcRouterStatus.L2TPIPsecServerSessions[0].IPAddress)
	assert.EqualValues(t, 300, vpcRouterStatus.L2TPIPsecServerSessions[0].TimeSec)

	assert.EqualValues(t, "micho3", vpcRouterStatus.PPTPServerSessions[0].User)
	assert.EqualValues(t, "192.168.2.202", vpcRouterStatus.PPTPServerSessions[0].IPAddress)
	assert.EqualValues(t, 500, vpcRouterStatus.PPTPServerSessions[0].TimeSec)

	assert.EqualValues(t, "DOWN", vpcRouterStatus.SiteToSiteIPsecVPNPeers[0].Status)
	assert.EqualValues(t, "1.1.1.1", vpcRouterStatus.SiteToSiteIPsecVPNPeers[0].Peer)
}
