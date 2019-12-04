// Copyright 2016-2019 The Libsacloud Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package naked

import (
	"encoding/json"
	"testing"

	"github.com/sacloud/libsacloud/v2/sacloud/types"
	"github.com/stretchr/testify/require"
)

const vpcRouterMultipleInterfaceJSON = `
    [
      {
        "IPAddress": "100.65.17.75",
        "UserIPAddress": null,
        "HostName": null,
        "Switch": {
          "ID": "112600699555",
          "Name": "スイッチ",
          "Scope": "shared",
          "Subnet": {
            "NetworkAddress": "100.65.0.0",
            "NetworkMaskLen": 18,
            "DefaultRoute": "100.65.0.1",
            "Internet": {
              "BandWidthMbps": 100
            }
          },
          "UserSubnet": {
            "DefaultRoute": "100.65.0.1",
            "NetworkMaskLen": 18
          }
        }
      },
      null,
      {
        "IPAddress": null,
        "UserIPAddress": null,
        "HostName": null,
        "Switch": {
          "ID": "113100846288",
          "Name": "name",
          "Scope": "user",
          "Subnet": null,
          "UserSubnet": null
        }
      },
      null,
      null,
      null,
      null,
      null
    ]
`

func TestVPCRouterUnmarshalInterfaceJSON(t *testing.T) {
	var ifs Interfaces
	err := json.Unmarshal([]byte(vpcRouterMultipleInterfaceJSON), &ifs)
	require.NoError(t, err)
	require.Len(t, ifs, 2)

	require.Equal(t, 0, ifs[0].Index)
	require.Equal(t, 2, ifs[1].Index)
}

const (
	vpcRouterRemarkServersEmptyJSON    = `[""]`
	vpcRouterRemarkServersNotEmptyJSON = `[{"IPAddress":"192.168.0.1"}]`
)

func TestVPCRouterRemarkServers_UnmarshalJSON(t *testing.T) {
	var remarkServers ApplianceRemarkServers
	err := json.Unmarshal([]byte(vpcRouterRemarkServersEmptyJSON), &remarkServers)
	require.NoError(t, err)
	require.Len(t, remarkServers, 0)

	err = json.Unmarshal([]byte(vpcRouterRemarkServersNotEmptyJSON), &remarkServers)
	require.NoError(t, err)
	require.Len(t, remarkServers, 1)
}

const vpcRouterMultipleFirewallJSON = `
{
	"Config": [
      {
		"Receive": [],
		"Send": []
      },
      {
		"Receive": [],
		"Send": []
      },
      {
		"Receive": [],
		"Send": [
			{
                  "Protocol": "ip",
                  "SourceNetwork": null,
                  "SourcePort": null,
                  "DestinationNetwork": null,
                  "DestinationPort": null,
                  "Action": "deny",
                  "Logging": "False",
                  "Description": ""
            }
		]
      },
      {
		"Receive": [],
		"Send": []
      },
      {
		"Receive": [],
		"Send": []
      },
      {
		"Receive": [],
		"Send": []
      },
      {
		"Receive": [],
		"Send": []
      },
      {
		"Receive": [],
		"Send": []
      }
	],
	"Enabled": "True"
}
`

var vpcRouterFirewallMarshaled = `{"Config":[{"Receive":[],"Send":[],"Index":0},{"Receive":[],"Send":[],"Index":1},{"Receive":[],"Send":[{"Protocol":"ip","Action":"deny","Description":""}],"Index":2},{"Receive":[],"Send":[],"Index":3},{"Receive":[],"Send":[],"Index":4},{"Receive":[],"Send":[],"Index":5},{"Receive":[],"Send":[],"Index":6},{"Receive":[],"Send":[],"Index":7}],"Enabled":"True"}`

func TestVPCRouterFirewall_UnmarshalJSON(t *testing.T) {
	var firewallConfig VPCRouterFirewall
	err := json.Unmarshal([]byte(vpcRouterMultipleFirewallJSON), &firewallConfig)
	require.NoError(t, err)
	for i, v := range firewallConfig.Config {
		require.Equal(t, i, v.Index)
	}

}

func TestVPCRouterFirewall_MarshalJSON(t *testing.T) {
	firewallConfig := &VPCRouterFirewall{
		Config: VPCRouterFirewallConfigs{
			{
				Send: []*VPCRouterFirewallRule{
					{
						Protocol: "ip",
						Action:   types.Actions.Deny,
					},
				},
				Index: 2,
			},
		},
		Enabled: types.StringFlag(true),
	}

	data, err := json.Marshal(firewallConfig)
	require.NoError(t, err)
	require.Equal(t, vpcRouterFirewallMarshaled, string(data))
}
