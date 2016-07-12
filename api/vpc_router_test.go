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

	// hack
	newItem.InitVPCRouterSetting()
	newItem.Settings.Router.AddInterface("", []string{"192.168.11.1"}, 24)
	newItem.Settings.Router.EnableL2TPIPsecServer("preshared", "192.168.11.100", "192.168.11.200")
	newItem.Settings.Router.AddRemoteAccessUser("yamamoto", "hogehogeo")

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
	//item.Description = "after"
	//item.Settings.Router.Interfaces = nil
	//item, err = api.Update(id, item)
	//
	//assert.NoError(t, err)
	//assert.NotEqual(t, item.Description, "before")
	//
	//////connect to switch
	////sw := client.Switch.New()
	////sw.Name = testSwitchName
	////
	////sw, err = client.Switch.Create(sw)
	////assert.NoError(t, err)
	////assert.NotEmpty(t, sw)
	////
	//
	//_, err = client.VPCRouter.Config(item.ID)
	//assert.NoError(t, err)

	item, err = api.Read(id)
	assert.NoError(t, err)
	assert.NotEmpty(t, item)

	//check connected switch
	assert.Equal(t, item.Settings.Router.Interfaces[1].IPAddress[0], "192.168.11.1")
	assert.Equal(t, item.Settings.Router.Interfaces[1].NetworkMaskLen, 24)
	assert.Equal(t, item.Settings.Router.Interfaces[1].VirtualIPAddress, "")

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
	newItem.SetPremiumPlan(sw.ID, vip, ip1, ip2, 1)
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

//func TestVPCRouterCRUDWithL2TP(t *testing.T) {
//	api := client.VPCRouter
//
//	//CREATE
//	newItem := api.New()
//	newItem.SetStandardPlan()
//	newItem.Name = testVPCRouterName
//	newItem.Description = "l2tp"
//	item, err := api.Create(newItem)
//
//	assert.NoError(t, err)
//	assert.NotEmpty(t, item)
//
//	id := item.ID
//
//	//wait
//	api.SleepWhileCopying(id, 300*time.Second)
//
//	////connect to switch
//	//sw := client.Switch.New()
//	//sw.Name = testSwitchName
//	//
//	//sw, err = client.Switch.Create(sw)
//	//assert.NoError(t, err)
//	//assert.NotEmpty(t, sw)
//	//
//	//err = client.VPCRouter.AddStandardInterface(item.ID, sw.ID, "192.168.100.1", 24)
//	//assert.NoError(t, err)
//
//	item.InitVPCRouterSetting()
//	item.Settings.Router.EnableL2TPIPsecServer("secret-hogehoge", "192.168.100.100", "192.168.100.200")
//	item.Settings.Router.AddRemoteAccessUser("user01", "password")
//
//	//update
//	item, err = api.Update(id, item)
//	assert.NoError(t, err)
//	assert.NotEmpty(t, item)
//
//	// config
//	res, err := api.Config(id)
//	assert.NoError(t, err)
//	assert.True(t, res)
//
//	//Delete
//	//_, err = api.Delete(id)
//	//assert.NoError(t, err)
//}

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
	items, _ := client.Internet.Reset().WithNameLike(testVPCRouterSwitchName).Find()
	for _, item := range items.Internet {
		client.Internet.Delete(item.ID)
	}
}
