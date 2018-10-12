package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const testPrivateHostName = "libsacloud_test_privatehost"

func TestPrivateHostCRUD(t *testing.T) {
	defer initPrivateHost()

	currentZone := client.Zone
	defer func() { client.Zone = currentZone }()
	client.Zone = "tk1a"

	api := client.PrivateHost

	// find plan
	plans, err := client.Product.GetProductPrivateHostAPI().Find()
	if !assert.NoError(t, err) {
		return
	}
	plan := plans.PrivateHostPlans[0]

	//CREATE
	p := api.New()
	p.Name = testPrivateHostName
	p.Description = "before"
	p.SetPrivateHostPlanByID(plan.ID)

	ph, err := api.Create(p)

	assert.NoError(t, err)
	assert.NotEmpty(t, ph)

	id := ph.ID

	//READ
	ph, err = api.Read(id)
	assert.NoError(t, err)
	assert.NotEmpty(t, ph)

	//UPDATE
	ph.Description = "after"
	ph, err = api.Update(id, ph)

	assert.NoError(t, err)
	assert.NotEqual(t, ph.Description, "before")

	//Delete
	_, err = api.Delete(id)
	assert.NoError(t, err)
}

func initPrivateHost() func() {
	cleanupPrivateHost()
	return cleanupPrivateHost
}

func cleanupPrivateHost() {
	currentZone := client.Zone
	defer func() { client.Zone = currentZone }()
	client.Zone = "tk1a"

	items, err := client.PrivateHost.Reset().WithNameLike(testPrivateHostName).Find()
	if err != nil {
		panic(err)
	}
	for _, item := range items.PrivateHosts {
		client.PrivateHost.Delete(item.ID)
	}
}
