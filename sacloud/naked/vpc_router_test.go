package naked

import (
	"encoding/json"
	"testing"

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
