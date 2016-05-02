package api

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

const testVPCRouterName = "libsacloud_test_VPCRouter"
const testVPCRouterSwitchName = "libsacloud_test_vpc_and_internet"

func TestVPCRouterCRUD(t *testing.T) {
	api := client.VPCRouter

	//CREATE
	newItem := api.New()
	newItem.SetStandardPlan()
	newItem.Name = testVPCRouterName
	newItem.Description = "before"
	item, err := api.Create(newItem)

	assert.NoError(t, err)
	assert.NotEmpty(t, item)

	id := item.ID

	//wait
	api.SleepWhileCopying(id, 300*time.Second)

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

func TestVPCRouterPremiumCRUD(t *testing.T) {
	api := client.VPCRouter

	//create internet(switch + router)
	inet := client.Internet.New()
	inet.Name = testVPCRouterSwitchName
	inet.BandWidthMbps = 100
	inet.NetworkMaskLen = 28
	internet, err := client.Internet.Create(inet)
	inetID := internet.ID
	timeout := 180 * time.Second
	current := 0 * time.Second
	interval := 5 * time.Second

	internet = nil
	err = nil
	//READ
	for internet == nil && timeout > current {
		internet, err = client.Internet.Read(inetID)

		if err != nil {
			time.Sleep(interval)
			current = current + interval
			err = nil
		}
	}

	if err != nil || current > timeout {
		assert.Fail(t, "Timeout: Can't read /internet/"+inetID)
	}

	assert.NotNil(t, internet)
	assert.NotNil(t, internet.Switch)
	assert.NoError(t, err)

	sw, err := client.Switch.Read(internet.Switch.ID)
	assert.NoError(t, err)

	vip, ip1, ip2, err := sw.GetDefaultIPAddressesForVPCRouter()

	assert.NoError(t, err)

	//CREATE
	newItem := api.New()
	newItem.SetPremiumPlan(sw.ID, vip, ip1, ip2, 1, internet.NetworkMaskLen)
	newItem.Name = testVPCRouterName
	newItem.Description = "before"
	item, err := api.Create(newItem)

	assert.NoError(t, err)
	assert.NotEmpty(t, item)

	id := item.ID

	//wait
	api.SleepWhileCopying(id, 300*time.Second)

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
	testSetupHandlers = append(testSetupHandlers, cleanupVPCRouter)
	testTearDownHandlers = append(testTearDownHandlers, cleanupVPCRouter)

	testSetupHandlers = append(testSetupHandlers, cleanupInternetForVPCRouter)
	testTearDownHandlers = append(testTearDownHandlers, cleanupInternetForVPCRouter)
}

func cleanupVPCRouter() {
	items, _ := client.VPCRouter.Reset().WithNameLike(testVPCRouterName).Find()
	for _, item := range items.VPCRouters {
		client.VPCRouter.Delete(item.ID)
	}
}

func cleanupInternetForVPCRouter() {
	items, _ := client.Switch.Reset().WithNameLike(testVPCRouterSwitchName).Find()
	for _, item := range items.Switches {
		client.Switch.Delete(item.ID)
	}
}
