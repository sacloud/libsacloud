package naked

import (
	"encoding/json"
	"testing"

	"github.com/sacloud/libsacloud/v2/sacloud/types"
	"github.com/stretchr/testify/require"
)

func TestInterface_UnmarshalJSON(t *testing.T) {

	cases := []struct {
		in  string
		out types.EUpstreamNetworkType
	}{
		{
			in:  `{}`,
			out: types.UpstreamNetworkTypes.None,
		},
		{
			in:  `{"Switch":{}}`,
			out: types.UpstreamNetworkTypes.Switch,
		},
		{
			in:  `{"Switch":{"Scope":"shared","Subnet":{}}}`,
			out: types.UpstreamNetworkTypes.Shared,
		},
		{
			in:  `{"Switch":{"Scope":"user","Subnet":{}}}`,
			out: types.UpstreamNetworkTypes.Router,
		},
	}

	for _, tc := range cases {
		var iface Interface
		err := json.Unmarshal([]byte(tc.in), &iface)
		require.NoError(t, err)
		require.Equal(t, tc.out.String(), iface.UpstreamType.String())
	}

}
