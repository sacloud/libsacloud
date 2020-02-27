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

	"github.com/sacloud/libsacloud/sacloud"
	"github.com/stretchr/testify/assert"
)

const testLoadBalancerName = "libsacloud_test_LoadBalancer"

var (
	createLoadBalancerValues = &sacloud.CreateLoadBalancerValue{
		VRID:         1,
		Plan:         sacloud.LoadBalancerPlanStandard,
		IPAddress1:   "192.168.11.11",
		MaskLen:      24,
		DefaultRoute: "192.168.11.1",
		Name:         "TestLoadBalancer",
		Description:  "TestDescription",
		Tags:         []string{"tag1", "tag2", "tag3"},
	}
	loadBalancerSettings = []*sacloud.LoadBalancerSetting{
		{
			VirtualIPAddress: "192.168.11.101",
			Port:             "8080",
			DelayLoop:        "30",
			SorryServer:      "192.168.11.201",
			Servers: []*sacloud.LoadBalancerServer{
				{
					IPAddress: "192.168.11.51",
					Port:      "8080",
					HealthCheck: &sacloud.LoadBalancerHealthCheck{
						Protocol: "http",
						Path:     "/",
						Status:   "200",
					},
				},
				{
					IPAddress: "192.168.11.52",
					Port:      "8080",
					HealthCheck: &sacloud.LoadBalancerHealthCheck{
						Protocol: "http",
						Path:     "/",
						Status:   "200",
					},
				},
			},
		},
	}
)

func TestLoadBalancerCRUD(t *testing.T) {
	defer initLoadBalancer()()

	api := client.LoadBalancer

	//prerequired
	sw := client.Switch.New()
	sw.Name = testLoadBalancerName
	sw, err := client.Switch.Create(sw)

	assert.NoError(t, err)
	assert.NotEmpty(t, sw)

	//CREATE
	createLoadBalancerValues.SwitchID = sw.ID
	newItem, err := sacloud.CreateNewLoadBalancerSingle(createLoadBalancerValues, loadBalancerSettings)
	assert.NoError(t, err)

	item, err := api.Create(newItem)

	assert.NoError(t, err)
	assert.NotEmpty(t, item)

	id := item.ID

	err = api.SleepWhileCopying(id, client.DefaultTimeoutDuration, 30)
	if !assert.NoError(t, err) {
		return
	}

	err = api.SleepUntilUp(id, client.DefaultTimeoutDuration)
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

	_, err = client.Switch.Delete(sw.ID)
	assert.NoError(t, err)

}

func _TestLoadBalancerCRUDWithoutVIP(t *testing.T) {
	defer initLoadBalancer()()

	api := client.LoadBalancer

	//prerequired
	sw := client.Switch.New()
	sw.Name = testLoadBalancerName
	sw, err := client.Switch.Create(sw)

	assert.NoError(t, err)
	assert.NotEmpty(t, sw)

	//CREATE
	createLoadBalancerValues.SwitchID = sw.ID
	newItem, err := sacloud.CreateNewLoadBalancerSingle(createLoadBalancerValues, nil)
	assert.NoError(t, err)

	item, err := api.Create(newItem)
	assert.NoError(t, err)
	assert.NotEmpty(t, item)

	id := item.ID

	err = api.SleepWhileCopying(id, client.DefaultTimeoutDuration, 30)
	if !assert.NoError(t, err) {
		return
	}

	err = api.SleepUntilUp(id, client.DefaultTimeoutDuration)
	if !assert.NoError(t, err) {
		return
	}
	//power off
	_, err = api.Stop(id)
	assert.NoError(t, err)
	err = api.SleepUntilDown(id, client.DefaultTimeoutDuration)
	if !assert.NoError(t, err) {
		return
	}

	//Delete
	_, err = api.Delete(id)
	assert.NoError(t, err)

	_, err = client.Switch.Delete(sw.ID)
	assert.NoError(t, err)

}

func initLoadBalancer() func() {
	cleanupLoadBalancer()
	return cleanupLoadBalancer
}

func cleanupLoadBalancer() {
	sw, _ := client.Switch.Reset().WithNameLike(testLoadBalancerName).Find()
	for _, item := range sw.Switches {
		client.Switch.Delete(item.ID)
	}

	items, _ := client.LoadBalancer.Reset().WithNameLike(testLoadBalancerName).Find()
	for _, item := range items.LoadBalancers {
		client.LoadBalancer.Delete(item.ID)
	}

}
