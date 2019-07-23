package archive

import (
	"context"
	"errors"
	"testing"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/ostype"
	"github.com/sacloud/libsacloud/v2/sacloud/testutil"
	"github.com/stretchr/testify/require"
)

type dummyFinder struct {
	archive *sacloud.ArchiveFindResult
	err     error
}

func (f *dummyFinder) Find(ctx context.Context, zone string, conditions *sacloud.FindCondition) (*sacloud.ArchiveFindResult, error) {
	return f.archive, f.err
}

func TestFindByOSType(t *testing.T) {
	t.Parallel()

	cases := []struct {
		input         ostype.ArchiveOSType
		finder        Finder
		expectedValue *sacloud.Archive
		expectedError error
	}{
		{
			input:         ostype.Custom,
			finder:        &dummyFinder{},
			expectedValue: nil,
			expectedError: errors.New("unsupported ostype.ArchiveOSType: Custom"),
		},
		{
			input: ostype.CentOS,
			finder: &dummyFinder{
				archive: &sacloud.ArchiveFindResult{}, // count: 0
			},
			expectedValue: nil,
			expectedError: errors.New("archive not found with ostype.ArchiveOSType: CentOS"),
		},
		{
			input: ostype.CentOS,
			finder: &dummyFinder{
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
		actual, err := FindByOSType(context.Background(), tc.finder, "tk1v", tc.input)
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

func TestAccFindByOSType(t *testing.T) {
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
			archive, err := FindByOSType(ctx, archiveOp, zone, os)
			require.NoError(t, err)
			t.Logf("zone: %s ostype[%s] => {ID: %d, Name: %s}", zone, os, archive.ID, archive.Name)
		}
	}
}
