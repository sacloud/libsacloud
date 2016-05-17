package api

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

const testInternetName = "libsacloud_test_internet"

func TestInternetCRUD(t *testing.T) {
	api := client.Internet

	//CREATE
	newItem := api.New()
	newItem.Name = testInternetName
	newItem.Description = "before"
	newItem.BandWidthMbps = 100
	newItem.NetworkMaskLen = 28

	item, err := api.Create(newItem)

	assert.NoError(t, err)
	assert.NotEmpty(t, item)

	id := item.ID

	timeout := 120 * time.Second
	current := 0 * time.Second
	interval := 5 * time.Second

	item = nil
	err = nil
	//READ
	for item == nil && timeout > current {
		item, err = api.Read(id)

		if err != nil {
			time.Sleep(interval)
			current = current + interval
			err = nil
		}
	}

	if err != nil || current > timeout {
		assert.Fail(t, "Timeout: Can't read /internet/"+id)
	}

	//item, err = api.Read(id)
	assert.NoError(t, err)
	assert.NotNil(t, item)
	assert.NotEmpty(t, item)

	//UPDATE
	item.Description = "after"
	item, err = api.Update(id, item)

	assert.NoError(t, err)
	assert.NotEqual(t, item.Description, "before")

	item, err = api.UpdateBandWidth(id, 500) //IDが変わる
	assert.NoError(t, err)

	id = item.ID

	item = nil
	err = nil
	current = 0 * time.Second
	//READ
	for item == nil && timeout > current {
		item, err = api.Read(id)

		if err != nil {
			time.Sleep(interval)
			current = current + interval
			err = nil
		}
	}

	if err != nil || current > timeout {
		assert.Fail(t, "Timeout: Can't read /internet/"+id)
	}

	assert.NoError(t, err)

	assert.Equal(t, item.BandWidthMbps, 500)

	//Delete
	_, err = api.Delete(id)
	assert.NoError(t, err)
}

func init() {
	testSetupHandlers = append(testSetupHandlers, cleanupInternet)
	testTearDownHandlers = append(testTearDownHandlers, cleanupInternet)
}

func cleanupInternet() {
	items, _ := client.Internet.Reset().WithNameLike(testInternetName).Find()
	for _, item := range items.Internet {
		client.Internet.Delete(item.ID)
	}
}
