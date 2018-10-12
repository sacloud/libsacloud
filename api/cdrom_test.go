package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const testCDROMName = "libsacloud_test_iso_image"

func TestCRUDCDROM(t *testing.T) {

	defer initCDROM()()

	api := client.CDROM

	//CREATE
	newCD := api.New()
	newCD.Name = testCDROMName
	newCD.Description = "hoge"
	newCD.SizeMB = 5120

	cd, _, err := api.Create(newCD)

	assert.NoError(t, err)
	assert.NotEmpty(t, cd)
	id := cd.ID

	//READ
	cd, err = api.Read(id)
	assert.NoError(t, err)
	assert.NotEmpty(t, cd)

	//Delete
	_, err = api.Delete(id)
	assert.NoError(t, err)
}

func initCDROM() func() {
	cleanupCDROM()
	return cleanupCDROM
}

func cleanupCDROM() {
	items, _ := client.CDROM.Reset().WithNameLike(testCDROMName).Find()
	for _, item := range items.CDROMs {
		client.CDROM.Delete(item.ID)
	}
}
