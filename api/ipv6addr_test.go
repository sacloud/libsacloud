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
