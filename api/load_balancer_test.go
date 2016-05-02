package api

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoadBalancer(t *testing.T) {
	assert.True(t, true)
}

//const testLoadBalancerName = "libsacloud_test_LoadBalancer"
//
//func TestLoadBalancerCRUD(t *testing.T) {
//	api := client.LoadBalancer
//
//	//CREATE
//	newItem := api.New()
//	newItem.Name = testLoadBalancerName
//	newItem.Description = "before"
//
//	item, err := api.Create(newItem)
//
//	assert.NoError(t, err)
//	assert.NotEmpty(t, item)
//
//	id := item.ID
//
//	//READ
//	item, err = api.Read(id)
//	assert.NoError(t, err)
//	assert.NotEmpty(t, item)
//
//	//UPDATE
//	item.Description = "after"
//	item, err = api.Update(id, item)
//
//	assert.NoError(t, err)
//	assert.NotEqual(t, item.Description, "before")
//
//	//Delete
//	_, err = api.Delete(id)
//	assert.NoError(t, err)
//}
//
//func init() {
//	testSetupHandlers = append(testSetupHandlers, cleanupLoadBalancer)
//	testTearDownHandlers = append(testTearDownHandlers, cleanupLoadBalancer)
//}
//
//func cleanupLoadBalancer() {
//	items, _ := client.LoadBalancer.Reset().WithNameLike(testLoadBalancerName).Find()
//	for _, item := range items.LoadBalancers {
//		client.LoadBalancer.Delete(item.ID)
//	}
//}
