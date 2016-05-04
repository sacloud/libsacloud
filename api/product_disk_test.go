package api

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProductDiskRead(t *testing.T) {
	api := client.Product.Disk
	res, err := api.Find()

	assert.NoError(t, err)
	assert.NotEmpty(t, res)
	assert.NotEmpty(t, res.DiskPlans)
	assert.NotEmpty(t, res.DiskPlans[0].ID)
}
