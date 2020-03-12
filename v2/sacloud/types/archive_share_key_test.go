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
	"testing"

	"github.com/stretchr/testify/require"
)

func TestArchiveShareKey(t *testing.T) {
	cases := []struct {
		in    string
		zone  string
		id    ID
		token string
		valid bool
	}{
		{
			in:    "zone:11111:token",
			zone:  "zone",
			id:    StringID("11111"),
			token: "token",
			valid: true,
		},
		{
			in:    "zone:11111",
			zone:  "zone",
			id:    StringID("11111"),
			token: "",
			valid: false,
		},
		{
			in:    "zone",
			zone:  "zone",
			id:    StringID(""),
			token: "",
			valid: false,
		},
		{
			in:    "",
			zone:  "",
			id:    StringID(""),
			token: "",
			valid: false,
		},
	}

	for _, tc := range cases {
		key := ArchiveShareKey(tc.in)
		require.Equal(t, key.Zone(), tc.zone)
		require.Equal(t, key.SourceArchiveID(), tc.id)
		require.Equal(t, key.Token(), tc.token)
		require.Equal(t, key.ValidFormat(), tc.valid)
	}
}
