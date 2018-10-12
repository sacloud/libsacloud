package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProductInternetRead(t *testing.T) {
	api := client.Product.Internet
	res, err := api.Find()

	assert.NoError(t, err)
	assert.NotEmpty(t, res)
	assert.NotEmpty(t, res.InternetPlans)
	assert.NotEmpty(t, res.InternetPlans[0].ID)
}
