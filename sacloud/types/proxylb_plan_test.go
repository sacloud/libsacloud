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
