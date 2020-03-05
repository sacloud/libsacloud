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

package types

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStringNumber(t *testing.T) {
	expects := []struct {
		input  string
		expect StringNumber
	}{
		{input: `""`, expect: StringNumber(0)},
		{input: `"1"`, expect: StringNumber(1)},
		{input: `0`, expect: StringNumber(0)},
		{input: `1`, expect: StringNumber(1)},
	}

	for _, tc := range expects {
		var n StringNumber
		err := json.Unmarshal([]byte(tc.input), &n)

		require.NotNil(t, n)
		require.NoError(t, err, "expect: %#v", tc)
		require.Equal(t, tc.expect, n)
	}
}
