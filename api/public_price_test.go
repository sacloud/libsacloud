package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPublicPriceFind(t *testing.T) {
	api := client.Product.Price
	res, err := api.Find()
	assert.NoError(t, err)
	assert.NotEmpty(t, res)
	assert.NotEmpty(t, res.ServiceClasses)
	assert.NotEmpty(t, res.ServiceClasses[0].DisplayName)

}
