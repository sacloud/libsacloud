// Copyright 2016-2020 The Libsacloud Authors
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

	"github.com/stretchr/testify/require"
)

func TestDatabaseSettingSourceNetworks_UnmarshalJSON(t *testing.T) {
	cases := []struct {
		name     string
		in       string
		expect   DatabaseSettingSourceNetworks
		hasError bool
	}{
		{
			name:   "empty string",
			in:     `""`,
			expect: nil,
		},
		{
			name:   "allow all",
			in:     `["0.0.0.0/0"]`,
			expect: nil,
		},
		{
			name:   "allow a CIDR block",
			in:     `["192.168.0.0/24"]`,
			expect: DatabaseSettingSourceNetworks{"192.168.0.0/24"},
		},
		{
			name:   "allow CIDR blocks",
			in:     `["192.168.0.0/24","192.168.1.0/24"]`,
			expect: DatabaseSettingSourceNetworks{"192.168.0.0/24", "192.168.1.0/24"},
		},
		{
			name:     "invalid values",
			in:       `"dummy"`,
			expect:   nil,
			hasError: true,
		},
	}

	for _, tc := range cases {
		var sn DatabaseSettingSourceNetworks
		err := json.Unmarshal([]byte(tc.in), &sn)
		require.Equal(t, tc.expect, sn, tc.name)
		require.Equal(t, tc.hasError, err != nil, tc.name)
	}
}

func TestDatabaseSettingSourceNetworks_MarshalJSON(t *testing.T) {
	cases := []struct {
		name     string
		in       DatabaseSettingSourceNetworks
		expect   string
		hasError bool
	}{
		{
			name:     "nil",
			in:       nil,
			expect:   `["0.0.0.0/0"]`,
			hasError: false,
		},
		{
			name:     "empty list",
			in:       DatabaseSettingSourceNetworks{},
			expect:   `["0.0.0.0/0"]`,
			hasError: false,
		},
		{
			name:     "allow a CIDR block",
			in:       DatabaseSettingSourceNetworks{"192.168.0.0/24"},
			expect:   `["192.168.0.0/24"]`,
			hasError: false,
		},
		{
			name:     "allow CIDR blocks",
			in:       DatabaseSettingSourceNetworks{"192.168.0.0/24", "192.168.1.0/24"},
			expect:   `["192.168.0.0/24","192.168.1.0/24"]`,
			hasError: false,
		},
	}

	for _, tc := range cases {
		data, err := json.Marshal(tc.in)
		require.Equal(t, tc.expect, string(data), tc.name)
		require.Equal(t, tc.hasError, err != nil, tc.name)
	}
}
