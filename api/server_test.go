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
	"time"

	"github.com/sacloud/libsacloud/sacloud"
	"github.com/stretchr/testify/assert"
)

const testServerName = "libsacloud-test-Server"

func TestServerCRUD(t *testing.T) {
	defer initServer()()

	api := client.Server

	//CREATE
	newItem := api.New()
	newItem.Name = testServerName
	newItem.Description = "before"
	newItem.SetServerPlanByValue(1, 1, sacloud.PlanDefault) // 1Core/1GBメモリ
	newItem.AddPublicNWConnectedParam()                     //公開セグメントに接続

	item, err := api.Create(newItem)

	assert.NoError(t, err)
	assert.NotEmpty(t, item)

	id := item.ID

	//Find
	items, err := api.WithNameLike(testServerName).Find()
	assert.NoError(t, err)
	assert.NotEmpty(t, items)
	assert.True(t, len(items.Servers) > 0)
	assert.Equal(t, id, items.Servers[0].ID)

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

func TestServerOperations(t *testing.T) {
	defer initServer()()

	currentRegion := client.Zone
	defer func() { client.Zone = currentRegion }()
	client.Zone = "is1a"

	api := client.Server

	//CREATE
	newItem := api.New()
	newItem.Name = testServerName
	newItem.Description = "before"
	newItem.SetServerPlanByID("100001001") // 1core 1GBメモリ
	newItem.AddPublicNWConnectedParam()    //公開セグメントに接続

	item, err := api.Create(newItem)

	assert.NoError(t, err)
	assert.NotEmpty(t, item)

	id := item.ID

	// change plan
	newPlan, err := client.GetProductServerAPI().GetBySpec(1, 2, sacloud.PlanDefault)
	assert.NoError(t, err)
	assert.NotEmpty(t, newPlan)

	item, err = api.ChangePlan(id, newPlan)
	assert.NoError(t, err)
	assert.True(t, id != item.ID)
	assert.Equal(t, 1, item.ServerPlan.CPU)
	assert.Equal(t, 2, item.ServerPlan.GetMemoryGB())
	assert.Equal(t, sacloud.PlanG2, item.ServerPlan.Generation)

	id = item.ID

	// boot
	res, err := api.Boot(id)
	assert.NoError(t, err)
	assert.True(t, res)
	api.SleepUntilUp(id, 180*time.Second)

	//VNC Proxy
	vncProxy, err := api.GetVNCProxy(id)
	assert.NoError(t, err)
	assert.NotEmpty(t, vncProxy)
	assert.NotEmpty(t, vncProxy.VNCFile)

	//VNC Size
	vncSize, err := api.GetVNCSize(id)
	assert.NoError(t, err)
	assert.NotEmpty(t, vncSize)
	assert.NotEmpty(t, vncSize.Height)

	//VNC SnapShot
	vncSnapRequest := api.NewVNCSnapshotRequest()
	vncSnapRequest.ScreenSaverExitTimeMS = 2000

	vncSnapResponse, err := api.GetVNCSnapshot(id, vncSnapRequest)
	assert.NoError(t, err)
	assert.NotEmpty(t, vncSnapResponse)
	assert.NotEmpty(t, vncSnapResponse.Image)

	// shutdown(force)
	// shutdown(force)
	res, err = api.Stop(id)
	assert.NoError(t, err)
	assert.True(t, res)
	api.SleepUntilDown(id, 180*time.Second)

	//Delete
	_, err = api.Delete(id)
	assert.NoError(t, err)
}

func TestSearchServer(t *testing.T) {
	assert.True(t, true)
	//
	//client.Zone = "is1a"
	//api := client.Server
	//s, err := api.Reset().withNameLike("sakura-dev").Find()
	//assert.NoError(t, err)
	//assert.NotEmpty(t, s)
	//
	//server := s.Servers[0]
	//
	//now := time.Now()
	//before10 := now.Add(-10 * time.Minute)
	//
	//// HACK : /monitor系APIはナノ秒があるとうまく範囲指定ができない
	////req := &sacloud.ResourceMonitorRequest{
	////	Start: &now,
	////	End:   &now,
	////}
	//req := sacloud.NewResourceMonitorRequest(&before10, &now)
	//
	//monitor, err := api.Monitor(server.ID, req)
	//
	//assert.NoError(t, err)
	//assert.NotEmpty(t, monitor)
	//
	//rawMonitorValue := map[string]*sacloud.MonitorValue(*monitor)
	//
	//assert.True(t, len(rawMonitorValue) >= 2)
	//assert.True(t, len(rawMonitorValue) <= 3)
	//
	//monitor, err = client.Disk.Monitor(server.Disks[0].ID, req)
	//assert.NoError(t, err)
	//assert.NotEmpty(t, monitor)
	//
	//rawMonitorValue = map[string]*sacloud.MonitorValue(*monitor)
	//assert.True(t, len(rawMonitorValue) >= 2)
	//assert.True(t, len(rawMonitorValue) <= 3)
	//
	//monitor, err = client.Interface.Monitor(server.Interfaces[0].ID, req)
	//assert.NoError(t, err)
	//assert.NotEmpty(t, monitor)
	//
	//rawMonitorValue = map[string]*sacloud.MonitorValue(*monitor)
	//assert.True(t, len(rawMonitorValue) >= 2)
	//assert.True(t, len(rawMonitorValue) <= 3)

}

func initServer() func() {
	cleanupServer()
	return cleanupServer
}

func cleanupServer() {
	items, _ := client.Server.Reset().WithNameLike(testServerName).Find()
	for _, item := range items.Servers {
		client.Server.Delete(item.ID)
	}
}
