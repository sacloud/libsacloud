package sacloud

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	testIPv6NetJSON = `
	{
		"CreatedAt": "2016-10-04T13:02:51+09:00",
		"ID": 999,
		"IPv6Prefix": "2001:e42:100:100::",
		"IPv6PrefixTail": "2001:e42:100:100::",
		"IPv6PrefixLen": 64,
		"IPv6Table": {
			"ID": 1
		},
		"NamedIPv6AddrCount": 1,
		"ServiceClass": "cloud/global-ipaddress-v6/64",
		"ServiceID": 123456789012,
		"Switch": {
			"Availability": "available",
			"ID": 123456789012,
			"Internet": {
			    "BandWidthMbps": 100,
			    "ID": 123456789012,
			    "Name": "libsacloud_test_vpc_and_internet",
			    "Scope": "user",
			    "ServiceClass": "cloud/internet/router/100m"
			},
			"Name": "libsacloud_test_vpc_and_internet"
		}
        }

	`
	testIPv6AddrJSON = `
	{
            "HostName": "testhost.libsacloud.com",
            "IPv6Addr": "2001:e42:100:100::5",
            "IPv6Net": {
                "ID": 330,
                "Switch": {
                    "ID": 123456789012
                }
            },
            "Interface": null
        }

	`
)

func TestMarshalIPv6JSON(t *testing.T) {
	var net IPv6Net
	err := json.Unmarshal([]byte(testIPv6NetJSON), &net)

	assert.NoError(t, err)
	assert.NotEmpty(t, net.IPv6Prefix)

	var addr IPv6Addr
	err = json.Unmarshal([]byte(testIPv6AddrJSON), &addr)

	assert.NoError(t, err)
	assert.NotEmpty(t, addr.IPv6Addr)

}
