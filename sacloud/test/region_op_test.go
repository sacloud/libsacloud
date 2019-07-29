package test

import (
	"context"
	"testing"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/testutil"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
	"github.com/stretchr/testify/assert"
)

func TestRegionOp_Find(t *testing.T) {
	t.Parallel()

	client := sacloud.NewRegionOp(singletonAPICaller())

	searched, err := client.Find(context.Background(), &sacloud.FindCondition{Count: 1})
	assert.NoError(t, err)

	err = testutil.DoAsserts(
		testutil.AssertLenFunc(t, searched.Regions, 1, "Regions"),
		testutil.AssertNotEmptyFunc(t, searched.Regions[0].ID, "Region.ID"),
		testutil.AssertNotEmptyFunc(t, searched.Regions[0].Name, "Region.Name"),
		testutil.AssertNotEmptyFunc(t, searched.Regions[0].Description, "Region.Description"),
		testutil.AssertNotEmptyFunc(t, searched.Regions[0].NameServers, "Region.NameServers"),
	)
	assert.NoError(t, err)
}

func TestRegionOp_Read(t *testing.T) {
	t.Parallel()

	client := sacloud.NewRegionOp(singletonAPICaller())

	sandboxRegionID := types.ID(290)
	region, err := client.Read(context.Background(), sandboxRegionID)
	assert.NoError(t, err)

	err = testutil.DoAsserts(
		testutil.AssertEqualFunc(t, region.ID, sandboxRegionID, "Region.ID"),
		testutil.AssertNotEmptyFunc(t, region.Name, "Region.Name"),
		testutil.AssertNotEmptyFunc(t, region.Description, "Region.Description"),
		testutil.AssertNotEmptyFunc(t, region.NameServers, "Region.NameServers"),
	)
	assert.NoError(t, err)
}
