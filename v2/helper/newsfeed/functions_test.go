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

package newsfeed

import (
	"testing"

	"github.com/sacloud/libsacloud/v2/sacloud/testutil"
	"github.com/stretchr/testify/require"
)

func TestGet(t *testing.T) {
	if !testutil.IsAccTest() {
		t.Skip("newsfeed.TestGet is only exec at Acceptance Test")
	}

	items, err := Get()
	require.NoError(t, err)
	require.True(t, len(items) > 0)
	fetched := items[0]

	// by URL
	item, err := GetByURL(fetched.URL)
	require.NoError(t, err)
	require.Equal(t, fetched, item)
}
