package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const testLicenseName = "libsacloud_test_License"

func TestLicenseCRUD(t *testing.T) {
	defer initLicense()()

	currentRegion := client.Zone
	defer func() { client.Zone = currentRegion }()
	client.Zone = "is1a"

	api := client.License

	resLicenseInfo, _ := client.Product.License.include("ID").Find()

	//CREATE
	newItem := api.New()
	newItem.Name = testLicenseName
	newItem.Description = "before"
	newItem.LicenseInfo = &resLicenseInfo.LicenseInfo[0]

	item, err := api.Create(newItem)

	assert.NoError(t, err)
	assert.NotEmpty(t, item)

	id := item.ID

	//READ
	item, err = api.Read(id)
	assert.NoError(t, err)
	assert.NotEmpty(t, item)

	//UPDATE
	item.Description = "after"
	item, err = api.Update(id, item)

	assert.NoError(t, err)
	assert.NotEqual(t, item.Description, "before")

	//Delete
	_, err = api.Delete(id)
	assert.NoError(t, err)
}

func initLicense() func() {
	cleanupLicense()
	return cleanupLicense
}

func cleanupLicense() {
	items, _ := client.License.Reset().WithNameLike(testLicenseName).Find()
	for _, item := range items.Licenses {
		client.License.Delete(item.ID)
	}
}
