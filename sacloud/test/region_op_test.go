package test

import (
	"context"
	"testing"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
	"github.com/stretchr/testify/assert"
)

func TestRegionOp_Find(t *testing.T) {
	t.Parallel()

	client := sacloud.NewRegionOp(singletonAPICaller())

	searched, err := client.Find(context.Background(), sacloud.APIDefaultZone, &sacloud.FindCondition{Count: 1})
	assert.NoError(t, err)

	err = DoAsserts(
		AssertLenFunc(t, searched.Regions, 1, "Regions"),
		AssertNotEmptyFunc(t, searched.Regions[0].ID, "Region.ID"),
		AssertNotEmptyFunc(t, searched.Regions[0].Name, "Region.Name"),
		AssertNotEmptyFunc(t, searched.Regions[0].Description, "Region.Description"),
		AssertNotEmptyFunc(t, searched.Regions[0].NameServers, "Region.NameServers"),
	)
	assert.NoError(t, err)
}

func TestRegionOp_Read(t *testing.T) {
	t.Parallel()

	client := sacloud.NewRegionOp(singletonAPICaller())

	sandboxRegionID := types.ID(290)
	region, err := client.Read(context.Background(), sacloud.APIDefaultZone, sandboxRegionID)
	assert.NoError(t, err)

	err = DoAsserts(
		AssertEqualFunc(t, region.ID, sandboxRegionID, "Region.ID"),
		AssertNotEmptyFunc(t, region.Name, "Region.Name"),
		AssertNotEmptyFunc(t, region.Description, "Region.Description"),
		AssertNotEmptyFunc(t, region.NameServers, "Region.NameServers"),
	)
	assert.NoError(t, err)
}
