package sacloud

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	testVPCRouterSettingsJSON = `
{
    "DHCPServer": {
        "Config": [
            {
                "Interface": "eth1",
                "RangeStart": "192.168.0.240",
                "RangeStop": "192.168.0.249"
            },
            {
                "Interface": "eth2",
                "RangeStart": "192.168.1.240",
                "RangeStop": "192.168.1.249"
            }
        ],
        "Enabled": "True"
    },
    "DHCPStaticMapping": {
        "Config": [
            {
                "IPAddress": "192.168.0.250",
                "MACAddress": "aa:bb:cc:dd:ee:ff"
            },
            {
                "IPAddress": "192.168.1.251",
                "MACAddress": "aa:bb:cc:dd:ee:00"
            }
        ],
        "Enabled": "True"
    },
    "Firewall": {
        "Config": [
            {
                "Receive": [
                    {
                        "Action": "allow",
                        "DestinationNetwork": "192.168.0.10",
                        "DestinationPort": "22",
                        "Protocol": "tcp",
                        "SourceNetwork": "203.0.113.1",
                        "SourcePort": "32768-61000"
                    },
                    {
                        "Action": "allow",
                        "DestinationNetwork": null,
                        "DestinationPort": "53",
                        "Protocol": "udp",
                        "SourceNetwork": "198.51.100.0/24",
                        "SourcePort": null
                    },
                    {
                        "Action": "deny",
                        "DestinationNetwork": null,
                        "Protocol": "ip",
                        "SourceNetwork": null
                    }
                ],
                "Send": [
                    {
                        "Action": "deny",
                        "DestinationNetwork": "198.51.100.0/24",
                        "Protocol": "icmp",
                        "SourceNetwork": "192.168.0.0/24"
                    },
                    {
                        "Action": "deny",
                        "DestinationNetwork": "198.51.100.111",
                        "Protocol": "ip",
                        "SourceNetwork": null
                    }
                ]
            }
        ],
        "Enabled": "True"
    },
    "Interfaces": [
        {
            "IPAddress": [
                "192.0.2.2",
                "192.0.2.3"
            ],
            "IPAliases": [
                "192.0.2.4",
                "192.0.2.5"
            ],
            "VirtualIPAddress": "192.0.2.1"
        },
        {
            "IPAddress": [
                "192.168.0.2",
                "192.168.0.3"
            ],
            "NetworkMaskLen": 24,
            "VirtualIPAddress": "192.168.0.1"
        },
        {
            "IPAddress": [
                "192.168.1.2",
                "192.168.1.3"
            ],
            "NetworkMaskLen": 24,
            "VirtualIPAddress": "192.168.1.1"
        }
    ],
    "L2TPIPsecServer": {
        "Config": {
            "PreSharedSecret": "abcd1234",
            "RangeStart": "192.168.0.220",
            "RangeStop": "192.168.0.229"
        },
        "Enabled": "True"
    },
    "PPTPServer": {
        "Config": {
            "RangeStart": "192.168.0.230",
            "RangeStop": "192.168.0.239"
        },
        "Enabled": "True"
    },
    "PortForwarding": {
        "Config": [
            {
                "GlobalPort": "80",
                "PrivateAddress": "192.168.0.10",
                "PrivatePort": "80",
                "Protocol": "tcp"
            },
            {
                "GlobalPort": "53",
                "PrivateAddress": "192.168.0.11",
                "PrivatePort": "53",
                "Protocol": "udp"
            }
        ],
        "Enabled": "True"
    },
    "RemoteAccessUsers": {
        "Config": [
            {
                "Password": "asdf1234",
                "UserName": "micho1"
            },
            {
                "Password": "asdf1234",
                "UserName": "micho2"
            }
        ],
        "Enabled": "True"
    },
    "SiteToSiteIPsecVPN": {
        "Config": [
            {
                "LocalPrefix": [
                    "192.168.0.0/24",
                    "192.168.1.0/24"
                ],
                "Peer": "198.51.100.10",
                "PreSharedSecret": "abcd1234",
                "RemoteID": "198.51.100.10",
                "Routes": [
                    "10.0.0.0/24",
                    "10.0.1.0/24"
                ]
            },
            {
                "LocalPrefix": [
                    "192.168.0.0/24",
                    "192.168.1.0/24"
                ],
                "Peer": "203.0.113.20",
                "PreSharedSecret": "abcd1234",
                "RemoteID": "203.0.113.20",
                "Routes": [
                    "10.0.2.0/24",
                    "10.0.3.0/24"
                ]
            }
        ],
        "Enabled": "True"
    },
    "StaticNAT": {
        "Config": [
            {
                "GlobalAddress": "192.0.2.4",
                "PrivateAddress": "192.168.1.10"
            },
            {
                "GlobalAddress": "192.0.2.5",
                "PrivateAddress": "192.168.1.11"
            }
        ],
        "Enabled": "True"
    },
    "VRID": 1
}
`
)

func TestMarshalVPCRouterSettingJSON(t *testing.T) {
	var setting VPCRouterSetting
	err := json.Unmarshal([]byte(testVPCRouterSettingsJSON), &setting)

	assert.NoError(t, err)
	assert.NotEmpty(t, setting)

	assert.Equal(t, *setting.VRID, 1)
}

func TestVPCRouterStaticNatFunc(t *testing.T) {
	setting := &VPCRouterSetting{}

	setting.AddStaticNAT("1.2.3.4", "192.168.0.1", "")
	assert.NotNil(t, setting.StaticNAT)
	assert.NotNil(t, setting.StaticNAT.Config)
	assert.Equal(t, setting.StaticNAT.Enabled, "True")
	assert.Len(t, setting.StaticNAT.Config, 1)

	setting.AddStaticNAT("5.6.7.8", "192.168.0.2", "")
	assert.NotNil(t, setting.StaticNAT)
	assert.NotNil(t, setting.StaticNAT.Config)
	assert.Equal(t, setting.StaticNAT.Enabled, "True")
	assert.Len(t, setting.StaticNAT.Config, 2)

	// it is not delete
	setting.RemoveStaticNAT("5.6.7.8", "192.168.99.99")
	assert.NotNil(t, setting.StaticNAT)
	assert.NotNil(t, setting.StaticNAT.Config)
	assert.Equal(t, setting.StaticNAT.Enabled, "True")
	assert.Len(t, setting.StaticNAT.Config, 2)

	setting.RemoveStaticNAT("9.9.9.9", "192.168.0.2")
	assert.NotNil(t, setting.StaticNAT)
	assert.NotNil(t, setting.StaticNAT.Config)
	assert.Equal(t, setting.StaticNAT.Enabled, "True")
	assert.Len(t, setting.StaticNAT.Config, 2)

	// delete
	setting.RemoveStaticNAT("1.2.3.4", "192.168.0.1")
	assert.NotNil(t, setting.StaticNAT)
	assert.NotNil(t, setting.StaticNAT.Config)
	assert.Equal(t, setting.StaticNAT.Enabled, "True")
	assert.Len(t, setting.StaticNAT.Config, 1)

	setting.RemoveStaticNAT("5.6.7.8", "192.168.0.2")
	assert.NotNil(t, setting.StaticNAT)
	assert.Nil(t, setting.StaticNAT.Config)
	assert.Equal(t, setting.StaticNAT.Enabled, "False")

}
