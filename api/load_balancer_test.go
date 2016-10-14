package api

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/yamamoto-febc/libsacloud/sacloud"
	"testing"
	"time"
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
	api := client.LoadBalancer

	//prerequired
	sw := client.Switch.New()
	sw.Name = testLoadBalancerName
	sw, err := client.Switch.Create(sw)

	assert.NoError(t, err)
	assert.NotEmpty(t, sw)

	//CREATE
	createLoadBalancerValues.SwitchID = fmt.Sprintf("%d", sw.ID)
	newItem, err := sacloud.CreateNewLoadBalancerSingle(createLoadBalancerValues, loadBalancerSettings)
	assert.NoError(t, err)

	item, err := api.Create(newItem)

	assert.NoError(t, err)
	assert.NotEmpty(t, item)

	id := item.ID

	err = api.SleepWhileCopying(id, 20*time.Minute, 3)
	if !assert.NoError(t, err) {
		return
	}

	err = api.SleepUntilUp(id, 10*time.Minute)
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

	err = api.SleepUntilDown(id, 120*time.Second)
	if !assert.NoError(t, err) {
		return
	}

	//Delete
	_, err = api.Delete(id)
	assert.NoError(t, err)

	_, err = client.Switch.Delete(sw.ID)
	assert.NoError(t, err)

}

func TestLoadBalancerCRUDWithoutVIP(t *testing.T) {
	api := client.LoadBalancer

	//prerequired
	sw := client.Switch.New()
	sw.Name = testLoadBalancerName
	sw, err := client.Switch.Create(sw)

	assert.NoError(t, err)
	assert.NotEmpty(t, sw)

	//CREATE
	createLoadBalancerValues.SwitchID = fmt.Sprintf("%d", sw.ID)
	newItem, err := sacloud.CreateNewLoadBalancerSingle(createLoadBalancerValues, nil)
	assert.NoError(t, err)

	item, err := api.Create(newItem)
	assert.NoError(t, err)
	assert.NotEmpty(t, item)

	id := item.ID

	err = api.SleepWhileCopying(id, 20*time.Minute, 3)
	if !assert.NoError(t, err) {
		return
	}

	err = api.SleepUntilUp(id, 10*time.Minute)
	if !assert.NoError(t, err) {
		return
	}
	//power off
	_, err = api.Stop(id)
	assert.NoError(t, err)
	err = api.SleepUntilDown(id, 120*time.Second)
	if !assert.NoError(t, err) {
		return
	}

	//Delete
	_, err = api.Delete(id)
	assert.NoError(t, err)

	_, err = client.Switch.Delete(sw.ID)
	assert.NoError(t, err)

}

func init() {
	testSetupHandlers = append(testSetupHandlers, cleanupLoadBalancer)
	testTearDownHandlers = append(testTearDownHandlers, cleanupLoadBalancer)
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
