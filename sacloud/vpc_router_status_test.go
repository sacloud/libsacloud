package sacloud

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
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

	assert.EqualValues(t, "slog1", vpcRouterStatus.FirewallSendLogs[0])
	assert.EqualValues(t, "rlog1", vpcRouterStatus.FirewallReceiveLogs[0])
	assert.EqualValues(t, "vlog1", vpcRouterStatus.VPNLogs[0])

}
