package types

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestProxyLBPlan(t *testing.T) {

	expects := []struct {
		strPlan    string
		actualPlan EProxyLBPlan
	}{
		{strPlan: `""`, actualPlan: EProxyLBPlan(0)},
		{strPlan: `"cloud/proxylb/plain/1000"`, actualPlan: EProxyLBPlan(1000)},
	}

	for _, tc := range expects {
		var n EProxyLBPlan
		err := json.Unmarshal([]byte(tc.strPlan), &n)

		require.NotNil(t, n)
		require.NoError(t, err, "expect: %#v", tc)
		require.Equal(t, tc.actualPlan, n, "expect: %#v", tc)

		// reverse
		data, err := json.Marshal(&tc.actualPlan)
		require.NoError(t, err, "expect: %#v", tc)
		require.Equal(t, tc.strPlan, string(data), "expect: %#v", tc)
	}

}
