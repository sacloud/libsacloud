package api

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const testDNSDomain = "test.domain.libsacloud.com"

func TestUpdateDnsCommonServiceItem(t *testing.T) {
	currentRegion := client.Zone
	defer func() { client.Zone = currentRegion }()
	client.Zone = "is1a"

	item, err := client.DNS.findOrCreateBy(testDNSDomain) //存在しないため新たに作る
	assert.NoError(t, err)
	assert.NotNil(t, item)
	assert.Equal(t, item.Name, testDNSDomain)

	//IPを追加して保存してみる
	item.Settings.DNS.AddDNSRecordSet("test1", "192.168.0.1")

	item, err = client.DNS.updateDNSRecord(item)
	assert.NoError(t, err)
	assert.NotNil(t, item)
	assert.Equal(t, item.Settings.DNS.ResourceRecordSets[0].Name, "test1")
	assert.Equal(t, item.Settings.DNS.ResourceRecordSets[0].RData, "192.168.0.1")
	assert.Equal(t, item.Settings.DNS.ResourceRecordSets[0].Type, "A")

	//IPを追加して保存してみる(２個目)
	item.Settings.DNS.AddDNSRecordSet("test2", "192.168.0.2")

	item, err = client.DNS.updateDNSRecord(item)
	assert.NoError(t, err)
	assert.NotNil(t, item)
	assert.Equal(t, item.Settings.DNS.ResourceRecordSets[1].Name, "test2")
	assert.Equal(t, item.Settings.DNS.ResourceRecordSets[1].RData, "192.168.0.2")
	assert.Equal(t, item.Settings.DNS.ResourceRecordSets[1].Type, "A")

	client.DNS.Delete(item.ID)

}

func TestCreateDNSRecords(t *testing.T) {
	currentRegion := client.Zone
	defer func() { client.Zone = currentRegion }()
	client.Zone = "is1a"

	item := client.DNS.New(testDNSDomain)

	assert.Equal(t, item.Name, testDNSDomain)
	assert.Equal(t, item.Status.Zone, testDNSDomain)

	//IPを追加して保存してみる
	item.AddRecord(item.CreateNewRecord("test1", "A", "192.168.0.1", 3600))
	item.AddRecord(item.CreateNewRecord("test2", "A", "192.168.0.2", 3600))

	item, err := client.DNS.Create(item)
	assert.NoError(t, err)
	assert.NotNil(t, item)
	assert.Equal(t, item.Settings.DNS.ResourceRecordSets[0].Name, "test1")
	assert.Equal(t, item.Settings.DNS.ResourceRecordSets[0].RData, "192.168.0.1")
	assert.Equal(t, item.Settings.DNS.ResourceRecordSets[0].Type, "A")
	assert.Equal(t, item.Settings.DNS.ResourceRecordSets[1].Name, "test2")
	assert.Equal(t, item.Settings.DNS.ResourceRecordSets[1].RData, "192.168.0.2")
	assert.Equal(t, item.Settings.DNS.ResourceRecordSets[1].Type, "A")

	client.DNS.Delete(item.ID)
}

func init() {
	testSetupHandlers = append(testSetupHandlers, cleanupDNSCommonServiceItem)
	testTearDownHandlers = append(testTearDownHandlers, cleanupDNSCommonServiceItem)
}

func cleanupDNSCommonServiceItem() {
	item, _ := client.DNS.findOrCreateBy(testDNSDomain)

	if item.ID != "" {
		client.DNS.Delete(item.ID)
	}
}
