package api

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestZone_Find(t *testing.T) {
	api := client.Facility.Zone
	res, err := api.Find()
	assert.NoError(t, err)
	assert.NotEmpty(t, res)
	assert.NotEmpty(t, res.Zones)
	assert.NotEmpty(t, res.Zones[0].ID)

	id := res.Zones[0].ID

	zone, err := api.Read(id.String())
	assert.NoError(t, err)
	assert.NotEmpty(t, zone)
	assert.NotEmpty(t, zone.ID)
}
