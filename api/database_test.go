package api

import (
	"fmt"
	"github.com/sacloud/libsacloud/sacloud"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

const testDatabaseName = "libsacloud_test_Database"

func TestDatabaseCRUD(t *testing.T) {
	api := client.Database
	client.Zone = "tk1a"

	//prerequired
	sw := client.Switch.New()
	sw.Name = testDatabaseName
	sw, err := client.Switch.Create(sw)

	assert.NoError(t, err)
	assert.NotEmpty(t, sw)

	//CREATE
	v := sacloud.NewCreatePostgreSQLDatabaseValue()

	v.Plan = sacloud.DatabasePlan30G
	v.DefaultUser = "defuser"
	v.UserPassword = "defuserPassword01"
	v.SourceNetwork = []string{"192.168.0.1", "192.168.1.1"}
	v.ServicePort = "54321"
	v.BackupTime = "00:30"
	v.SwitchID = fmt.Sprintf("%d", sw.ID)
	v.IPAddress1 = "192.168.11.100"
	v.MaskLen = 24
	v.DefaultRoute = "192.168.11.1"
	v.Name = testDatabaseName

	newItem := api.New(v)
	//newItem.Remark.Zone = &sacloud.Resource{ID: 21001}

	item, err := api.Create(newItem)

	assert.NoError(t, err)
	assert.NotEmpty(t, item)

	id := item.ID

	err = api.SleepWhileCopying(id, client.DefaultTimeoutDuration, 30)
	if !assert.NoError(t, err) {
		return
	}

	api.SleepUntilUp(id, client.DefaultTimeoutDuration)

	err = api.SleepUntilDatabaseRunning(id, client.DefaultTimeoutDuration, 30)
	assert.NoError(t, err)

	//READ
	item, err = api.Read(id)

	assert.NoError(t, err)
	assert.NotEmpty(t, item)

	//FIND
	items, err := api.Reset().Find()
	assert.NoError(t, err)
	assert.NotEmpty(t, items)

	//UPDATE
	item.Description = "after"
	item, err = api.Update(id, item)

	assert.NoError(t, err)
	assert.NotEqual(t, item.Description, "before")

	//power off
	for i := 0; i < 30; i++ {
		_, err = api.Stop(id)
		assert.NoError(t, err)

		err = api.SleepUntilDown(id, 10*time.Second)
		if err == nil {
			break
		}
	}
	assert.NoError(t, err)

	sourceID := id

	// clone
	clone := sacloud.NewCloneDatabaseValue(item)
	clone.Plan = sacloud.DatabasePlan10G
	clone.SourceNetwork = []string{"192.168.0.1", "192.168.1.1"}
	clone.ServicePort = "33061"
	clone.BackupTime = "13:30"
	clone.SwitchID = fmt.Sprintf("%d", sw.ID)
	clone.IPAddress1 = "192.168.11.100"
	clone.MaskLen = 24
	clone.DefaultRoute = "192.168.11.1"
	clone.Name = testDatabaseName + "_clone"

	newItem = api.New(v)
	item, err = api.Create(newItem)

	assert.NoError(t, err)
	assert.NotEmpty(t, item)
	id = item.ID
	err = api.SleepWhileCopying(id, client.DefaultTimeoutDuration, 30)
	if !assert.NoError(t, err) {
		return
	}
	api.SleepUntilUp(id, client.DefaultTimeoutDuration)
	err = api.SleepUntilDatabaseRunning(id, client.DefaultTimeoutDuration, 30)
	assert.NoError(t, err)
	//power off
	for i := 0; i < 30; i++ {
		_, err = api.Stop(id)
		assert.NoError(t, err)

		err = api.SleepUntilDown(id, 10*time.Second)
		if err == nil {
			break
		}
	}
	assert.NoError(t, err)

	//Delete
	_, err = api.Delete(sourceID)
	assert.NoError(t, err)

	_, err = api.Delete(item.ID)

	_, err = client.Switch.Delete(sw.ID)
	assert.NoError(t, err)

}

func TestDatabaseMariaDBCRUD(t *testing.T) {
	api := client.Database
	client.Zone = "tk1a"

	//prerequired
	sw := client.Switch.New()
	sw.Name = testDatabaseName
	sw, err := client.Switch.Create(sw)

	assert.NoError(t, err)
	assert.NotEmpty(t, sw)

	//CREATE
	v := sacloud.NewCreateMariaDBDatabaseValue()

	v.Plan = sacloud.DatabasePlanMini
	v.DefaultUser = "defuser"
	v.UserPassword = "defuserPassword01"
	v.SourceNetwork = []string{"192.168.0.1", "192.168.1.1"}
	v.ServicePort = "33061"
	v.BackupTime = "13:30"
	v.SwitchID = fmt.Sprintf("%d", sw.ID)
	v.IPAddress1 = "192.168.11.100"
	v.MaskLen = 24
	v.DefaultRoute = "192.168.11.1"
	v.Name = testDatabaseName

	newItem := api.New(v)
	//newItem.Remark.Zone = &sacloud.Resource{ID: 21001}

	item, err := api.Create(newItem)

	assert.NoError(t, err)
	assert.NotEmpty(t, item)

	id := item.ID

	err = api.SleepWhileCopying(id, client.DefaultTimeoutDuration, 30)
	if !assert.NoError(t, err) {
		return
	}

	api.SleepUntilUp(id, client.DefaultTimeoutDuration)

	err = api.SleepUntilDatabaseRunning(id, client.DefaultTimeoutDuration, 30)
	assert.NoError(t, err)

	//READ
	item, err = api.Read(id)

	assert.NoError(t, err)
	assert.NotEmpty(t, item)

	//FIND
	items, err := api.Reset().Find()
	assert.NoError(t, err)
	assert.NotEmpty(t, items)

	//UPDATE
	item.Description = "after"
	item, err = api.Update(id, item)

	assert.NoError(t, err)
	assert.NotEqual(t, item.Description, "before")

	//power off
	for i := 0; i < 30; i++ {
		_, err = api.Stop(id)
		assert.NoError(t, err)

		err = api.SleepUntilDown(id, 10*time.Second)
		if err == nil {
			break
		}
	}
	assert.NoError(t, err)

	//Delete
	_, err = api.Delete(id)
	assert.NoError(t, err)

	_, err = client.Switch.Delete(sw.ID)
	assert.NoError(t, err)

}

func TestDatabaseWaitForCopy(t *testing.T) {
	api := client.Database
	client.Zone = "tk1a"

	//prerequired
	sw := client.Switch.New()
	sw.Name = testDatabaseName
	sw, err := client.Switch.Create(sw)

	assert.NoError(t, err)
	assert.NotEmpty(t, sw)

	//CREATE
	v := sacloud.NewCreateMariaDBDatabaseValue()

	v.Plan = sacloud.DatabasePlanMini
	v.DefaultUser = "defuser"
	v.UserPassword = "defuserPassword01"
	v.SourceNetwork = []string{"192.168.0.1", "192.168.1.1"}
	v.ServicePort = "33061"
	v.BackupTime = "13:30"
	v.SwitchID = fmt.Sprintf("%d", sw.ID)
	v.IPAddress1 = "192.168.11.100"
	v.MaskLen = 24
	v.DefaultRoute = "192.168.11.1"
	v.Name = testDatabaseName

	newItem := api.New(v)
	//newItem.Remark.Zone = &sacloud.Resource{ID: 21001}

	item, err := api.Create(newItem)

	assert.NoError(t, err)
	assert.NotEmpty(t, item)

	id := item.ID

	defer func() {
		//power off
		for i := 0; i < 30; i++ {
			_, err = api.Stop(id)
			assert.NoError(t, err)

			err = api.SleepUntilDown(id, 10*time.Second)
			if err == nil {
				break
			}
		}
		api.Delete(id)
	}()

	complete, progress, errChan := api.AsyncSleepWhileCopying(id, client.DefaultTimeoutDuration, 10)

	for {
		select {
		case d := <-progress:
			t.Logf("Database %s... ", d.Availability)
		case <-complete:
			t.Logf("Done. Now waiting for up.")
			api.SleepUntilUp(id, client.DefaultTimeoutDuration)
			t.Logf("Done.")
			return
		case e := <-errChan:
			assert.Fail(t, e.Error(), nil)
			return
		case <-time.After(20 * time.Minute):
			assert.Fail(t, "Timeout: AsyncSleepWhileCopying: Database -> %d", id)
			return
		}
	}
}

func init() {
	testSetupHandlers = append(testSetupHandlers, cleanupDatabase)
	testTearDownHandlers = append(testTearDownHandlers, cleanupDatabase)
}

func cleanupDatabase() {
	client.Zone = "tk1a"

	sw, _ := client.Switch.Reset().WithNameLike(testDatabaseName).Find()
	for _, item := range sw.Switches {
		client.Switch.Delete(item.ID)
	}

	items, _ := client.Database.Reset().WithNameLike(testDatabaseName).Find()
	for _, item := range items.Databases {
		client.Database.Delete(item.ID)
	}

}
