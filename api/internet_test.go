package api

import (
	"fmt"
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

	err = api.SleepWhileCreating(id, 120*time.Second)
	if err != nil {
		assert.Fail(t, fmt.Sprintf("Timeout: Can't read /internet/%d", id))
	}

	item, err = api.Read(id)
	assert.NoError(t, err)
	assert.NotNil(t, item)
	assert.NotEmpty(t, item)

	//UPDATE
	item.Description = "after"
	item, err = api.Update(id, item)

	assert.NoError(t, err)
	assert.NotEqual(t, item.Description, "before")

	// UPDATE BandWidth
	item, err = api.UpdateBandWidth(id, 500) //IDが変わる
	assert.NoError(t, err)

	id = item.ID

	err = api.SleepWhileCreating(id, 120*time.Second)
	if err != nil {
		assert.Fail(t, fmt.Sprintf("Timeout: Can't read /internet/%d", id))
	}
	assert.NoError(t, err)

	item, err = api.Read(id)
	assert.Equal(t, item.BandWidthMbps, 500)

	// Enable/Disable IPv6
	ipv6Net, err := api.EnableIPv6(id)
	assert.NoError(t, err)
	assert.Equal(t, ipv6Net.Switch.Internet.ID, id)

	// disable
	item, err = api.Read(id)
	res, err := api.DisableIPv6(id, item.Switch.IPv6Nets[0].ID)
	assert.NoError(t, err)
	assert.True(t, res)

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
