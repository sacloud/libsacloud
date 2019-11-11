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

package test

import (
	"context"
	"os"
	"testing"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/testutil"
	"github.com/stretchr/testify/assert"
)

func TestIPAddressOp_List_Read(t *testing.T) {
	t.Parallel()

	client := sacloud.NewIPAddressOp(singletonAPICaller())
	searched, err := client.List(context.Background(), testZone)
	assert.NoError(t, err)

	if searched.Count == 0 {
		t.Skip("IPAddress is not found")
	}
	ip := searched.IPAddress[0]
	err = testutil.DoAsserts(
		testutil.AssertNotEmptyFunc(t, ip.IPAddress, "IPAddress"),
	)
	assert.NoError(t, err)

	// read
	read, err := client.Read(context.Background(), testZone, ip.IPAddress)
	assert.NoError(t, err)
	assert.Equal(t, ip, read)
}

func TestIPAddressOp_UpdateHostName(t *testing.T) {
	t.Parallel()

	testutil.PreCheckEnvsFunc("SAKURACLOUD_IPADDRESS", "SAKURACLOUD_HOSTNAME")(t)

	client := sacloud.NewIPAddressOp(singletonAPICaller())
	ip := os.Getenv("SAKURACLOUD_IPADDRESS")
	hostName := os.Getenv("SAKURACLOUD_HOSTNAME")

	updated, err := client.UpdateHostName(context.Background(), testZone, ip, hostName)
	assert.NoError(t, err)
	assert.Equal(t, updated.HostName, hostName)

	updated, err = client.UpdateHostName(context.Background(), testZone, ip, "")
	assert.NoError(t, err)
	assert.Empty(t, updated.HostName, hostName)

}
