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

package query

import (
	"context"
	"errors"
	"testing"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/ostype"
	"github.com/sacloud/libsacloud/v2/sacloud/testutil"
	"github.com/stretchr/testify/require"
)

func TestFindArchiveByOSType(t *testing.T) {
	t.Parallel()

	cases := []struct {
		input         ostype.ArchiveOSType
		finder        ArchiveFinder
		expectedValue *sacloud.Archive
		expectedError error
	}{
		{
			input:         ostype.Custom,
			finder:        &dummyArchiveFinder{},
			expectedValue: nil,
			expectedError: errors.New("unsupported ostype.ArchiveOSType: Custom"),
		},
		{
			input: ostype.CentOS,
			finder: &dummyArchiveFinder{
				archive: &sacloud.ArchiveFindResult{}, // count: 0
			},
			expectedValue: nil,
			expectedError: errors.New("archive not found with ostype.ArchiveOSType: CentOS"),
		},
		{
			input: ostype.CentOS,
			finder: &dummyArchiveFinder{
				archive: &sacloud.ArchiveFindResult{
					Count: 2,
					Total: 2,
					Archives: []*sacloud.Archive{
						{
							ID: 1,
						},
						{
							ID: 2,
						},
					},
				},
			},
			expectedValue: &sacloud.Archive{ID: 1},
			expectedError: nil,
		},
	}

	for _, tc := range cases {
		actual, err := FindArchiveByOSType(context.Background(), tc.finder, "tk1v", tc.input)
		if tc.expectedError != nil {
			require.Equal(t, tc.expectedError, err)
		} else {
			require.NoError(t, err)
		}

		if tc.expectedValue != nil {
			require.Equal(t, tc.expectedValue, actual)
		} else {
			require.Nil(t, actual)
		}
	}
}

func TestAccFindArchiveByOSType(t *testing.T) {
	if !testutil.IsAccTest() {
		t.Skip("TestAccFindByOSType only exec at Acceptance Test")
	}

	t.Parallel()

	caller := testutil.SingletonAPICaller()
	archiveOp := sacloud.NewArchiveOp(caller)
	ctx := context.Background()

	zones := []string{"is1a", "is1b", "tk1a", "tk1v"}

	for _, zone := range zones {
		for _, os := range ostype.ArchiveOSTypes {
			archive, err := FindArchiveByOSType(ctx, archiveOp, zone, os)
			require.NoError(t, err)
			t.Logf("zone: %s ostype[%s] => {ID: %d, Name: %s}", zone, os, archive.ID, archive.Name)
		}
	}
}
