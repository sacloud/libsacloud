package api

//import (
//	"github.com/stretchr/testify/assert"
//	"testing"
//)
//
//var (
//	targetIPAddress = "2001:e42:100:100::5"
//	targetHostName  = "test.libsacloud.com"
//)
//
//func TestIPv6AddrCRUD(t *testing.T) {
//	api := client.IPv6Addr
//
//	// create
//	ip, err := api.Create(targetIPAddress, targetHostName)
//	assert.NoError(t, err)
//	assert.NotNil(t, ip)
//	assert.Equal(t, ip.IPv6Addr, targetIPAddress)
//	assert.Equal(t, ip.HostName, targetHostName)
//
//	// read
//	ip, err = api.Read(targetIPAddress)
//	assert.NoError(t, err)
//	assert.NotNil(t, ip)
//	assert.Equal(t, ip.IPv6Addr, targetIPAddress)
//	assert.Equal(t, ip.HostName, targetHostName)
//
//	// update
//	ip, err = api.Update(targetIPAddress, "")
//	assert.NoError(t, err)
//	assert.Equal(t, ip.HostName, "")
//
//	// delete
//	ip, err = api.Delete(targetIPAddress)
//	assert.NoError(t, err)
//	assert.Equal(t, ip.IPv6Addr, targetIPAddress)
//
//}
