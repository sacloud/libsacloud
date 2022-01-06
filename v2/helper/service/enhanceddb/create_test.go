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

package enhanceddb

import (
	"testing"

	"github.com/sacloud/libsacloud/v2/sacloud/testutil"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
	"github.com/stretchr/testify/require"
)

func TestEnhancedDBService_convertCreateRequest(t *testing.T) {
	name := testutil.ResourceName("container-registry-service")
	dbName := testutil.RandomName(10, testutil.CharSetAlpha)
	password := testutil.RandomName(16, testutil.CharSetAlpha)

	cases := []struct {
		in     *CreateRequest
		expect *ApplyRequest
	}{
		{
			in: &CreateRequest{
				Name:         name,
				Description:  "desc",
				Tags:         types.Tags{"tag1", "tag2"},
				IconID:       1,
				DatabaseName: dbName,
				Password:     password,
			},
			expect: &ApplyRequest{
				ID:           0,
				Name:         name,
				Description:  "desc",
				Tags:         types.Tags{"tag1", "tag2"},
				IconID:       1,
				DatabaseName: dbName,
				Password:     password,
				SettingsHash: "",
			},
		},
	}

	for _, tc := range cases {
		require.EqualValues(t, tc.expect, tc.in.ApplyRequest())
	}
}
