package test

import (
	"context"
	"testing"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/stretchr/testify/assert"
)

func TestIPv6NetOp_List_Read(t *testing.T) {
	t.Parallel()

	internetOp := sacloud.NewInternetOp(singletonAPICaller())
	ipv6NetOp := sacloud.NewIPv6NetOp(singletonAPICaller())
	ctx := context.Background()

	internet, err := internetOp.Create(ctx, testZone, &sacloud.InternetCreateRequest{
		Name:           "libsacloud-internet-ipv6",
		NetworkMaskLen: 28,
		BandWidthMbps:  100,
	})
	assert.NoError(t, err)

	// wait
	internet, err = readInternet(internet.ID, singletonAPICaller())
	assert.NoError(t, err)

	// Enable IPv6
	ipv6Net, err := internetOp.EnableIPv6(ctx, testZone, internet.ID)
	assert.NoError(t, err)

	err = DoAsserts(
		AssertNotEmptyFunc(t, ipv6Net.ID, "IPv6Net.ID"),
		AssertNotEmptyFunc(t, ipv6Net.IPv6Prefix, "IPv6Net.IPv6Prefix"),
		AssertNotEmptyFunc(t, ipv6Net.IPv6PrefixLen, "IPv6Net.IPv6PrefixLen"),
	)
	assert.NoError(t, err)

	// find
	searched, err := ipv6NetOp.List(ctx, testZone)
	assert.NoError(t, err)
	assert.True(t, len(searched.IPv6Nets) > 0)

	// read
	read, err := ipv6NetOp.Read(ctx, testZone, ipv6Net.ID)
	assert.NoError(t, err)
	err = DoAsserts(
		AssertEqualFunc(t, ipv6Net.ID, read.ID, "IPv6Net.ID"),
		AssertEqualFunc(t, ipv6Net.IPv6Prefix, read.IPv6Prefix, "IPv6Net.IPv6Prefix"),
		AssertEqualFunc(t, ipv6Net.IPv6PrefixLen, read.IPv6PrefixLen, "IPv6Net.IPv6PrefixLen"),
	)
	assert.NoError(t, err)

	// Disable IPv6
	err = internetOp.DisableIPv6(ctx, testZone, internet.ID, ipv6Net.ID)
	assert.NoError(t, err)

	// cleanup
	err = internetOp.Delete(ctx, testZone, internet.ID)
	assert.NoError(t, err)
}
