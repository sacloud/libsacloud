// Copyright 2016-2022 The Libsacloud Authors
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

func TestStringFlag(t *testing.T) {
	expects := []struct {
		input  string
		expect StringFlag
	}{
		{input: `"True"`, expect: StringTrue},
		{input: `"true"`, expect: StringTrue},
		{input: `"False"`, expect: StringFalse},
		{input: `"false"`, expect: StringFalse},
		{input: `true`, expect: StringTrue},
		{input: `false`, expect: StringFalse},
		{input: `"On"`, expect: StringTrue},
		{input: `"on"`, expect: StringTrue},
		{input: `"Off"`, expect: StringFalse},
		{input: `"off"`, expect: StringFalse},
		{input: `"1"`, expect: StringTrue},
		{input: `"0"`, expect: StringFalse},
		{input: `1`, expect: StringTrue},
		{input: `0`, expect: StringFalse},
		{input: `""`, expect: StringFalse},
		{input: `"hoge"`, expect: StringFalse},
	}

	for _, tc := range expects {
		var f StringFlag
		err := json.Unmarshal([]byte(tc.input), &f)

		require.NotNil(t, f)
		require.NoError(t, err, "expect: %#v", tc)
		require.Equal(t, tc.expect, f)
	}
}
