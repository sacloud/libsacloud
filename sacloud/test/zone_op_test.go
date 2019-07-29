package test

import (
	"context"
	"testing"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/stretchr/testify/assert"
)

func TestZoneOp_Find(t *testing.T) {
	t.Parallel()

	client := sacloud.NewZoneOp(singletonAPICaller())

	zoneFindResult, err := client.Find(context.Background(), &sacloud.FindCondition{Count: 1})
	assert.NoError(t, err)
	assert.Len(t, zoneFindResult.Zones, 1)
}
