package api

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const testSimpleMonitorName = "8.8.8.8"

func TestSimpleMonitorCRUD(t *testing.T) {

	currentRegion := client.Zone
	defer func() { client.Zone = currentRegion }()
	client.Zone = "is1a"

	api := client.SimpleMonitor

	//CREATE
	newItem := api.New(testSimpleMonitorName)
	newItem.Description = "before"
	newItem.SetDelayLoop(60)
	newItem.EnableNotifyEmail()
	newItem.SetHealthCheckPing()

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

func init() {
	testSetupHandlers = append(testSetupHandlers, cleanupSimpleMonitor)
	testTearDownHandlers = append(testTearDownHandlers, cleanupSimpleMonitor)
}

func cleanupSimpleMonitor() {
	items, _ := client.SimpleMonitor.Reset().WithNameLike(testSimpleMonitorName).Find()
	for _, item := range items.SimpleMonitors {
		client.SimpleMonitor.Delete(item.ID)
	}
}
