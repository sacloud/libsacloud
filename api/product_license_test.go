package api

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProductLicenseRead(t *testing.T) {
	api := client.Product.License
	res, err := api.Find()
	assert.NoError(t, err)
	assert.NotEmpty(t, res.LicenseInfo)
}
