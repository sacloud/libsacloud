package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const testSwitchName = "libsacloud_test_Switch"

func TestSwitchCRUD(t *testing.T) {
	defer initSwitch()()

	api := client.Switch

	//CREATE
	newItem := api.New()
	newItem.Name = testSwitchName
	newItem.Description = "before"

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

func initSwitch() func() {
	cleanupSwitch()
	return cleanupSwitch
}

func cleanupSwitch() {
	items, _ := client.Switch.Reset().WithNameLike(testSwitchName).Find()
	for _, item := range items.Switches {
		client.Switch.Delete(item.ID)
	}
}
