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

package search

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestEqualExpression(t *testing.T) {
	targetTime := time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
	targetTimeString := targetTime.Format(time.RFC3339)

	cases := []struct {
		input  *EqualExpression
		expect string
	}{
		{
			input:  AndEqual("c1", "c2"),
			expect: `"c1%20c2"`,
		},
		{
			input:  AndEqual("c1", "", "c2"),
			expect: `"c1%20%20c2"`,
		},
		{
			input:  OrEqual("c1", "c2"),
			expect: `["c1","c2"]`,
		},
		{
			input:  OrEqual("c1", nil, "c2"),
			expect: `["c1","c2"]`,
		},
		{
			input:  OrEqual(1, 2),
			expect: `["1","2"]`,
		},
		{
			input:  OrEqual(targetTime),
			expect: fmt.Sprintf(`["%s"]`, targetTimeString),
		},
	}

	for _, tc := range cases {
		data, err := json.Marshal(tc.input)
		require.NoError(t, err)
		require.Equal(t, tc.expect, string(data))
	}
}
