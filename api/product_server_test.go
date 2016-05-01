package api

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetProductServer(t *testing.T) {
	api := client.Product.Server

	//READ
	res, err := api.Read(1001) // 1core 1GB
	assert.NoError(t, err)
	assert.NotEmpty(t, res)
	assert.NotEmpty(t, res.ID)

	isValid, err := api.IsValidPlan(1, 1)
	assert.NoError(t, err)
	assert.True(t, isValid)

	isValid, err = api.IsValidPlan(9999999, 99999999)
	assert.Error(t, err)
	assert.False(t, isValid)

}
