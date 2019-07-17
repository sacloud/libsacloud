package test

import (
	"context"
	"os"
	"testing"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/stretchr/testify/assert"
)

// TestIPv6AddrOp_CRUD .
//
// Note: IDが特殊(数値でなくIPv6アドレス)なため、CRUDTestCaseを利用しない
func TestIPv6AddrOp_CRUD(t *testing.T) {

	t.Parallel()

	PreCheckEnvsFunc("SAKURACLOUD_IPV6ADDRESS", "SAKURACLOUD_IPV6HOSTNAME")(t)

	client := sacloud.NewIPv6AddrOp(singletonAPICaller())
	ip := os.Getenv("SAKURACLOUD_IPV6ADDRESS")
	hostName := os.Getenv("SAKURACLOUD_IPV6HOSTNAME")

	// create
	created, err := client.Create(context.Background(), testZone, &sacloud.IPv6AddrCreateRequest{
		IPv6Addr: ip,
		HostName: hostName,
	})
	assert.NoError(t, err)
	assert.Equal(t, created.IPv6Addr, ip)
	assert.Equal(t, created.HostName, hostName)

	// read
	read, err := client.Read(context.Background(), testZone, ip)
	assert.NoError(t, err)
	assert.Equal(t, read.IPv6Addr, ip)
	assert.Equal(t, read.HostName, hostName)

	// update
	updated, err := client.Update(context.Background(), testZone, ip, &sacloud.IPv6AddrUpdateRequest{
		HostName: "",
	})
	assert.NoError(t, err)
	assert.Equal(t, updated.IPv6Addr, ip)
	assert.Equal(t, updated.HostName, "")

	// delete
	err = client.Delete(context.Background(), testZone, ip)
	assert.NoError(t, err)
}
