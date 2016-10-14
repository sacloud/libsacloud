package api

//import (
//	"github.com/stretchr/testify/assert"
//	"testing"
//)
//
//var (
//	targetIPAddress = "xxx.xxx.xxx.xxx"
//	targetHostName  = "target.host.name"
//)
//
//func TestIPAddressReadAndUpdate(t *testing.T) {
//	api := client.IPAddress
//
//	address, err := api.Find()
//	assert.NoError(t, err)
//	assert.True(t, len(address.IPAddress) > 0)
//
//	ip, err := api.Read(targetIPAddress)
//	assert.NoError(t, err)
//	assert.NotNil(t, ip)
//
//	//upd
//	ip, err = api.Update(targetIPAddress, targetHostName)
//	assert.NoError(t, err)
//	assert.Equal(t, ip.HostName, targetHostName)
//
//	ip, err = api.Update(targetIPAddress, "")
//	assert.NoError(t, err)
//	assert.Equal(t, ip.HostName, "")
//}
