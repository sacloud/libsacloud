package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProductLicenseRead(t *testing.T) {
	api := client.Product.License
	res, err := api.Find()
	assert.NoError(t, err)
	assert.NotEmpty(t, res.LicenseInfo)
}
