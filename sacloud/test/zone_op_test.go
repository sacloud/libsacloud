package test

import (
	"context"
	"testing"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/stretchr/testify/require"
)

func TestZoneOp_Find(t *testing.T) {
	t.Parallel()

	client := sacloud.NewZoneOp(singletonAPICaller())

	zoneFindResult, err := client.Find(context.Background(), sacloud.APIDefaultZone, &sacloud.FindCondition{Count: 1})
	require.NoError(t, err)
	require.Len(t, zoneFindResult.Zones, 1)
}
