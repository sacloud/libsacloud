package api

import (
	"fmt"
	"testing"

	"github.com/sacloud/libsacloud/sacloud"
	"github.com/stretchr/testify/assert"
)

const testNFSName = "libsacloud_test_NFS"

var createNFSValues = &sacloud.CreateNFSValue{
	IPAddress:    "192.168.11.11",
	MaskLen:      24,
	DefaultRoute: "192.168.11.1",
	Name:         "TestNFS",
	Description:  "TestDescription",
	Tags:         []string{"tag1", "tag2", "tag3"},
}

func TestNFSCRUD(t *testing.T) {
	defer initNFS()()

	api := client.NFS

	//prerequired
	sw := client.Switch.New()
	sw.Name = testNFSName
	sw, err := client.Switch.Create(sw)

	assert.NoError(t, err)
	assert.NotEmpty(t, sw)

	//CREATE
	createNFSValues.SwitchID = fmt.Sprintf("%d", sw.ID)
	item, err := api.CreateWithPlan(createNFSValues, sacloud.NFSPlanHDD, sacloud.NFSSize100G)

	assert.NoError(t, err)
	assert.NotEmpty(t, item)

	id := item.ID

	err = api.SleepWhileCopying(id, client.DefaultTimeoutDuration, 30)
	if !assert.NoError(t, err) {
		return
	}

	err = api.SleepUntilUp(id, client.DefaultTimeoutDuration)
	if !assert.NoError(t, err) {
		return
	}

	//READ
	item, err = api.Read(id)
	assert.NoError(t, err)
	assert.NotEmpty(t, item)

	//UPDATE
	item.Description = "after"
	item, err = api.Update(id, item)

	assert.NoError(t, err)
	assert.NotEqual(t, item.Description, "before")

	_, err = api.Stop(id)
	assert.NoError(t, err)

	err = api.SleepUntilDown(id, client.DefaultTimeoutDuration)
	if !assert.NoError(t, err) {
		return
	}

	//Delete
	_, err = api.Delete(id)
	assert.NoError(t, err)

	_, err = client.Switch.Delete(sw.ID)
	assert.NoError(t, err)

}

func initNFS() func() {
	cleanupNFS()
	return cleanupNFS
}

func cleanupNFS() {
	sw, _ := client.Switch.Reset().WithNameLike(testNFSName).Find()
	for _, item := range sw.Switches {
		client.Switch.Delete(item.ID)
	}

	items, _ := client.NFS.Reset().WithNameLike(testNFSName).Find()
	for _, item := range items.NFS {
		client.NFS.Delete(item.ID)
	}

}
