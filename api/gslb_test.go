// Copyright 2016-2020 The Libsacloud Authors
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
	"testing"

	"github.com/stretchr/testify/assert"
)

const testGslbName = "test_libsakuracloud_gslb"

func TestGslbGet(t *testing.T) {
	defer initGSLB()()

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
	defer initGSLB()()

	currentRegion := client.Zone
	defer func() { client.Zone = currentRegion }()
	client.Zone = "is1a"

	item := client.GSLB.New(testGslbName)
	assert.Equal(t, item.Name, testGslbName)

	item.Settings.GSLB.SorryServer = "8.8.8.8"
	item.Settings.GSLB.HealthCheck.Host = "libsacloud.com"

	//IPを追加して保存してみる
	item.AddGSLBServer(item.CreateGSLBServer("8.8.8.8"))
	item.AddGSLBServer(item.CreateGSLBServer("8.8.4.4"))

	item, err := client.GSLB.Create(item)

	assert.NoError(t, err)
	assert.Equal(t, item.Settings.GSLB.HealthCheck.Host, "libsacloud.com")
	assert.Equal(t, item.Settings.GSLB.SorryServer, "8.8.8.8")

	assert.Equal(t, item.Settings.GSLB.Servers[0].IPAddress, "8.8.8.8")
	assert.Equal(t, item.Settings.GSLB.Servers[0].Weight, "1")
	assert.Equal(t, item.Settings.GSLB.Servers[0].Enabled, "True")
	assert.Equal(t, item.Settings.GSLB.Servers[1].IPAddress, "8.8.4.4")
	assert.Equal(t, item.Settings.GSLB.Servers[1].Weight, "1")
	assert.Equal(t, item.Settings.GSLB.Servers[1].Enabled, "True")

	client.GSLB.Delete(item.ID)
}

func TestGSLBWithEmptyServer(t *testing.T) {
	defer initGSLB()()

	currentRegion := client.Zone
	defer func() { client.Zone = currentRegion }()
	client.Zone = "is1a"

	item := client.GSLB.New(testGslbName)
	assert.Equal(t, item.Name, testGslbName)

	item.Settings.GSLB.SorryServer = "8.8.8.8"
	item.Settings.GSLB.HealthCheck.Host = "libsacloud.com"

	item, err := client.GSLB.Create(item)

	assert.NoError(t, err)
	assert.Equal(t, item.Settings.GSLB.HealthCheck.Host, "libsacloud.com")
	assert.Equal(t, item.Settings.GSLB.SorryServer, "8.8.8.8")

	assert.Len(t, item.Settings.GSLB.Servers, 0)

	client.GSLB.Delete(item.ID)
}

func initGSLB() func() {
	cleanupGSLB()
	return cleanupGSLB
}

func cleanupGSLB() {
	items, _ := client.GSLB.Reset().WithNameLike(testGslbName).Find()

	for _, item := range items.CommonServiceGSLBItems {
		client.GSLB.Delete(item.ID)
	}
}
