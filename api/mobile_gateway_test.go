package api

import (
	"os"
	"testing"
	"time"

	"github.com/sacloud/libsacloud/sacloud"
	"github.com/stretchr/testify/assert"
)

const testMGWName = "libsacloud_test_mgw"

var (
	createMGWValues = &sacloud.CreateMobileGatewayValue{
		Name:        testMGWName,
		Description: "TestDescription",
		Tags:        []string{"tag1", "tag2", "tag3"},
	}
	mgwSetting = &sacloud.MobileGatewaySetting{
		InternetConnection: &sacloud.MGWInternetConnection{
			Enabled: "True",
		},
		Interfaces: []*sacloud.MGWInterface{
			nil,
			{
				IPAddress: []string{
					"192.168.0.1",
				},
				NetworkMaskLen: 24,
			},
		},
		StaticRoutes: []*sacloud.MGWStaticRoute{
			{
				Prefix:  "172.10.0.0/16",
				NextHop: "192.168.0.2",
			},
			{
				Prefix:  "172.0.0.0/8",
				NextHop: "192.168.0.3",
			},
		},
	}
)

func TestMobileGatwayCRUD(t *testing.T) {
	defer initMobileGateway()()

	api := client.MobileGateway

	//CREATE
	newItem, err := sacloud.CreateNewMobileGateway(createMGWValues, mgwSetting)
	assert.NoError(t, err)

	item, err := api.Create(newItem)

	assert.NoError(t, err)
	assert.NotEmpty(t, item)

	id := item.ID

	err = api.SleepWhileCopying(id, client.DefaultTimeoutDuration, 30)
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

}

func TestMobileGatewayWithSIM(t *testing.T) {
	iccID := os.Getenv(testSIMEnvICCID)
	passcode := os.Getenv(testSIMEnvPASSCODE)

	if iccID == "" || passcode == "" {
		t.Skipf("%s and %s is required. skip", testSIMEnvICCID, testSIMEnvPASSCODE)
	}

	defer initSIM()()

	// create sim
	simAPI := client.SIM
	sim, err := simAPI.Create(simAPI.New(testSIMName, iccID, passcode))
	assert.NoError(t, err)
	assert.NotNil(t, sim)

	//CREATE
	api := client.MobileGateway
	newItem, err := sacloud.CreateNewMobileGateway(createMGWValues, mgwSetting)
	assert.NoError(t, err)

	item, err := api.Create(newItem)

	assert.NoError(t, err)
	assert.NotEmpty(t, item)

	id := item.ID
	err = api.SleepWhileCopying(id, client.DefaultTimeoutDuration, 30)
	if !assert.NoError(t, err) {
		return
	}

	// sleep(HACK)
	time.Sleep(30 * time.Second)

	// Set DNS
	_, err = api.SetDNS(id, &sacloud.MobileGatewayResolver{
		SimGroup: &sacloud.MobileGatewaySIMGroup{
			DNS1: "8.8.8.8",
			DNS2: "8.8.4.4",
		},
	})
	assert.NoError(t, err)

	// add SIM to MGW
	_, err = api.AddSIM(id, sim.ID)
	assert.NoError(t, err)

	// List SIM
	sims, err := api.ListSIM(id, nil)
	assert.NoError(t, err)
	assert.Len(t, sims, 1)

	// Set SIM Route
	simRoutes := &sacloud.MobileGatewaySIMRoutes{
		SIMRoutes: []*sacloud.MobileGatewaySIMRoute{
			{
				ResourceID: sim.GetStrID(),
				Prefix:     "192.168.10.0/24",
			},
		},
	}
	_, err = api.SetSIMRoutes(id, simRoutes)
	assert.NoError(t, err)

	// List SIMRoute
	routes, err := api.GetSIMRoutes(id)
	assert.NotNil(t, simRoutes)
	assert.NoError(t, err)
	assert.Len(t, routes, 1)
	assert.Equal(t, sim.GetStrID(), routes[0].ResourceID)
	assert.Equal(t, "192.168.10.0/24", routes[0].Prefix)

	// Delete(all) SIMRoute
	_, err = api.DeleteSIMRoutes(id)
	assert.NoError(t, err)

	// add SIM Route
	added, err := api.AddSIMRoute(id, sim.ID, "192.168.10.0/24")
	assert.True(t, added)
	assert.NoError(t, err)
	// List SIMRoute
	routes, err = api.GetSIMRoutes(id)
	assert.NotNil(t, simRoutes)
	assert.NoError(t, err)
	assert.Len(t, routes, 1)
	assert.Equal(t, sim.GetStrID(), routes[0].ResourceID)
	assert.Equal(t, "192.168.10.0/24", routes[0].Prefix)

	// Delete(by value) SIMRoute
	deleted, err := api.DeleteSIMRoute(id, sim.ID, "192.168.10.0/24")
	assert.True(t, deleted)
	assert.NoError(t, err)

	// Delete SIM
	_, err = api.DeleteSIM(id, sim.ID)
	assert.NoError(t, err)

	// List SIM(after delete)
	sims, err = api.ListSIM(id, nil)
	assert.NoError(t, err)
	assert.Len(t, sims, 0)

	// Del MGW
	_, err = api.Delete(id)
	assert.NoError(t, err)

}

func initMobileGateway() func() {
	cleanupMobileGateway()
	return cleanupMobileGateway
}

func cleanupMobileGateway() {

	items, _ := client.MobileGateway.Reset().WithNameLike(testMGWName).Find()
	for _, item := range items.MobileGateways {
		client.MobileGateway.Delete(item.ID)
	}

}
