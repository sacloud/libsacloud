package api

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const testGslbName = "test_libsakuracloud_gslb"

func TestGslbGet(t *testing.T) {

	currentRegion := client.Zone
	defer func() { client.Zone = currentRegion }()
	client.Zone = "is1a"

	item, err := client.GSLB.findOrCreateBy(testGslbName)
	assert.NoError(t, err)
	assert.NotNil(t, item)
	assert.Equal(t, item.Name, testGslbName)

	//IPを追加して保存してみる
	item.Settings.GSLB.AddServer("8.8.8.8")
	item, err = client.GSLB.updateGSLBServers(item)
	assert.NoError(t, err)
	assert.NotNil(t, item)
	assert.Equal(t, item.Settings.GSLB.Servers[0].IPAddress, "8.8.8.8")
	assert.Equal(t, item.Settings.GSLB.Servers[0].Weight, "1")
	assert.Equal(t, item.Settings.GSLB.Servers[0].Enabled, "True")

	//IPを追加して保存してみる(2個目)
	item.Settings.GSLB.AddServer("8.8.4.4")
	item, err = client.GSLB.updateGSLBServers(item)
	assert.NoError(t, err)
	assert.NotNil(t, item)
	assert.Equal(t, item.Settings.GSLB.Servers[1].IPAddress, "8.8.4.4")
	assert.Equal(t, item.Settings.GSLB.Servers[1].Weight, "1")
	assert.Equal(t, item.Settings.GSLB.Servers[1].Enabled, "True")

	client.GSLB.Delete(item.ID)

}

func TestGSLBCreate(t *testing.T) {
	currentRegion := client.Zone
	defer func() { client.Zone = currentRegion }()
	client.Zone = "is1a"

	item := client.GSLB.New(testGslbName)
	assert.Equal(t, item.Name, testGslbName)

	item.Settings.GSLB.HealthCheck.Host = "libsacloud.com"

	//IPを追加して保存してみる
	item.AddGSLBServer(item.CreateGSLBServer("8.8.8.8"))
	item.AddGSLBServer(item.CreateGSLBServer("8.8.4.4"))

	item, err := client.GSLB.Create(item)

	assert.NoError(t, err)
	assert.NotEmpty(t, item.Settings.GSLB.HealthCheck.Host, "libsacloud.com")

	assert.Equal(t, item.Settings.GSLB.Servers[0].IPAddress, "8.8.8.8")
	assert.Equal(t, item.Settings.GSLB.Servers[0].Weight, "1")
	assert.Equal(t, item.Settings.GSLB.Servers[0].Enabled, "True")
	assert.Equal(t, item.Settings.GSLB.Servers[1].IPAddress, "8.8.4.4")
	assert.Equal(t, item.Settings.GSLB.Servers[1].Weight, "1")
	assert.Equal(t, item.Settings.GSLB.Servers[1].Enabled, "True")

	client.GSLB.Delete(item.ID)
}

func init() {
	testSetupHandlers = append(testSetupHandlers, cleanupGslbCommonServiceItem)
	testTearDownHandlers = append(testTearDownHandlers, cleanupGslbCommonServiceItem)
}

func cleanupGslbCommonServiceItem() {
	item, _ := client.GSLB.findOrCreateBy(testGslbName)

	if item.ID != "" {
		client.GSLB.Delete(item.ID)
	}
}
