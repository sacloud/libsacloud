// Copyright 2016-2021 The Libsacloud Authors
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

func TestResult_UnmarshalJSON(t *testing.T) {
	expects := []struct {
		input    string
		expect   APIResult
		hasError bool
	}{
		{
			input:    "true",
			expect:   ResultSuccess,
			hasError: false,
		},
		{
			input:    "false",
			expect:   ResultFailed,
			hasError: false,
		},
		{
			input:    `"Accepted"`,
			expect:   ResultAccepted,
			hasError: false,
		},
		{
			input:    "2",
			expect:   ResultUnknown,
			hasError: true,
		},
		{
			input:    `"Deleted"`,
			expect:   ResultUnknown,
			hasError: false,
		},
	}

	for _, tc := range expects {
		var res APIResult
		err := json.Unmarshal([]byte(tc.input), &res)

		if tc.hasError {
			require.Error(t, err)
		} else {
			require.NoError(t, err)
		}
		require.Equal(t, tc.expect, res)
	}
}
