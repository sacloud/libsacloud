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

package bill

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestListRequest_validate(t *testing.T) {
	cases := []struct {
		in       *ListRequest
		hasError bool
	}{
		{
			in:       &ListRequest{},
			hasError: false,
		},
		{
			in: &ListRequest{
				Month: 1, // without Year
			},
			hasError: true,
		},
		{
			in: &ListRequest{
				Year:  2020,
				Month: 0,
			},
			hasError: false,
		},
		{
			in: &ListRequest{
				Year:  2020,
				Month: 13, // invalid month
			},
			hasError: true,
		},
		{
			in: &ListRequest{
				Year:  2020,
				Month: 1,
			},
			hasError: false,
		},
	}

	for _, tc := range cases {
		err := tc.in.Validate()
		require.Equal(t, tc.hasError, err != nil, "with: %#+v", tc.in)
	}
}
