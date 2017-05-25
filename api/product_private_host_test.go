package api

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetProductPrivateHost(t *testing.T) {
	api := client.Product.PrivateHost

	currentZone := client.Zone
	client.Zone = "tk1a"
	defer func() {
		client.Zone = currentZone
	}()

	//READ
	res, err := api.Find()
	assert.NoError(t, err)

	assert.NotEmpty(t, res.PrivateHostPlans)
	assert.NotEmpty(t, res.PrivateHostPlans[0].ID)

}
