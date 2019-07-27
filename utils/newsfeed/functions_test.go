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
