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
