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

package search

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestKey(t *testing.T) {
	cases := []struct {
		input  FilterKey
		expect string
	}{
		{
			input: FilterKey{
				Field: "field",
				Op:    OpEqual,
			},
			expect: "field",
		},
		{
			input: FilterKey{
				Field: "field",
				Op:    OpGreaterThan,
			},
			expect: "field>",
		},
		{
			input: FilterKey{
				Field: "field",
				Op:    OpGreaterEqual,
			},
			expect: "field>=",
		},
		{
			input: FilterKey{
				Field: "field",
				Op:    OpLessThan,
			},
			expect: "field<",
		},
		{
			input: FilterKey{
				Field: "field",
				Op:    OpLessEqual,
			},
			expect: "field<=",
		},
		{
			input: FilterKey{
				Field: "another-field-name",
				Op:    OpEqual,
			},
			expect: "another-field-name",
		},
	}

	for _, tc := range cases {
		require.Equal(t, tc.expect, tc.input.String())
	}
}
