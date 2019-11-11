// Copyright 2016-2019 The Libsacloud Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package api

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

const testDNSDomain = "test.domain.libsacloud.com"

func TestUpdateDnsCommonServiceItem(t *testing.T) {
	defer initDNS()()

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
	defer initDNS()()

	currentRegion := client.Zone
	defer func() { client.Zone = currentRegion }()
	client.Zone = "is1a"

	item := client.DNS.New(testDNSDomain)

	assert.Equal(t, item.Name, testDNSDomain)
	assert.Equal(t, item.Status.Zone, testDNSDomain)

	//IPを追加して保存してみる
	item.AddRecord(item.CreateNewRecord("test1", "A", "192.168.0.1", 3600))
	item.AddRecord(item.CreateNewRecord("test1", "A", "192.168.0.2", 3600))
	//item.AddRecord(item.CreateNewSRVRecord("_sip._tls", testDNSDomain+".", 3600, 100, 1, 443))

	item, err := client.DNS.Create(item)
	assert.NoError(t, err)
	assert.NotNil(t, item)
	assert.Equal(t, len(item.Settings.DNS.ResourceRecordSets), 2)
	assert.Equal(t, item.Settings.DNS.ResourceRecordSets[0].Name, "test1")
	assert.Equal(t, item.Settings.DNS.ResourceRecordSets[0].RData, "192.168.0.1")
	assert.Equal(t, item.Settings.DNS.ResourceRecordSets[0].Type, "A")
	assert.Equal(t, item.Settings.DNS.ResourceRecordSets[1].Name, "test1")
	assert.Equal(t, item.Settings.DNS.ResourceRecordSets[1].RData, "192.168.0.2")
	assert.Equal(t, item.Settings.DNS.ResourceRecordSets[1].Type, "A")

	item.AddRecord(item.CreateNewSRVRecord("_sip._tls", "test1."+testDNSDomain+".", 3600, 100, 1, 443))
	item, err = client.DNS.Update(item.ID, item)
	assert.NoError(t, err)

	assert.Equal(t, len(item.Settings.DNS.ResourceRecordSets), 3)
	assert.Equal(t, item.Settings.DNS.ResourceRecordSets[2].Name, "_sip._tls")
	assert.Equal(t, item.Settings.DNS.ResourceRecordSets[2].RData, fmt.Sprintf("%d %d %d %s", 100, 1, 443, "test1."+testDNSDomain+"."))
	assert.Equal(t, item.Settings.DNS.ResourceRecordSets[2].Type, "SRV")

	client.DNS.Delete(item.ID)
}

func initDNS() func() {
	cleanupDNS()
	return cleanupDNS
}

func cleanupDNS() {
	item, _ := client.DNS.findOrCreateBy(testDNSDomain)

	if item.ID > 0 {
		client.DNS.Delete(item.ID)
	}
}
