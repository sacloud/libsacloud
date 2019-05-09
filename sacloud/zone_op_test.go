package sacloud

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFindZone(t *testing.T) {
	if !isAccTest() {
		t.Skip("TESTACC is not set. skip")
	}

	client := NewZoneOp(singletonAPICaller())

	zones, err := client.Find(context.Background(), DefaultZone, &FindCondition{Count: 1})
	require.NoError(t, err)
	require.Len(t, zones, 1)
}
