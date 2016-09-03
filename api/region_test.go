package api

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRegion_Find(t *testing.T) {
	api := client.Facility.Region
	res, err := api.Find()
	assert.NoError(t, err)
	assert.NotEmpty(t, res)
	assert.NotEmpty(t, res.Regions)
	assert.NotEmpty(t, res.Regions[0].ID)

	id := res.Regions[0].ID

	region, err := api.Read(id)
	assert.NoError(t, err)
	assert.NotEmpty(t, region)
	assert.NotEmpty(t, region.ID)

}
