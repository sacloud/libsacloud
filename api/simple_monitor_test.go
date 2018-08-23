package api

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const testSimpleMonitorName = "8.8.8.8"

func TestSimpleMonitorCRUD(t *testing.T) {
	defer initSimpleMonitor()()

	currentRegion := client.Zone
	defer func() { client.Zone = currentRegion }()
	client.Zone = "is1b"

	api := client.SimpleMonitor

	//CREATE
	newItem := api.New(testSimpleMonitorName)
	newItem.Description = "before"
	newItem.SetDelayLoop(60)
	newItem.EnableNotifyEmail(true)
	newItem.SetHealthCheckHTTP("80", "/index.html", "200", "www.libsacloud.com", "", "")

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

func initSimpleMonitor() func() {
	cleanupSimpleMonitor()
	return cleanupSimpleMonitor
}

func cleanupSimpleMonitor() {
	items, _ := client.SimpleMonitor.Reset().WithNameLike(testSimpleMonitorName).Find()
	for _, item := range items.SimpleMonitors {
		client.SimpleMonitor.Delete(item.ID)
	}
}
