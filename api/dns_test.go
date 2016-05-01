package api

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const testDNSDomain = "docker-machine-sakuracloud.com"

func TestUpdateDnsCommonServiceItem(t *testing.T) {
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
