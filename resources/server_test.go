package resources

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

var testServerJSON = `
{
    "ID": "999999999999",
    "Name": "lisacloud-test-server-name",
    "HostName": "libsacloud-test-server-name.com",
    "Description": "Description",
    "Availability": "available",
    "ServiceClass": "cloud\/plan\/1core-1gb",
    "CreatedAt": "2016-04-25T21:36:47+09:00",
    "Icon": null,
    "ServerPlan": ` + testServerPlanJSON + `,
    "Zone": ` + testZoneJSON + `,
    "Instance": ` + testInstanceJSON + `,
    "Disks": [ ` + testDiskJSON + `
    ],
    "Interfaces": [ ` + testPublicInterfaceJSON + `
    	,` + testPrivateInterfaceJSON + `
    ],
    "Appliance": null,
    "Tags": [
	"@virtio-net-pci"
    ]
}

`

func TestMarshalServerJSON(t *testing.T) {
	var server Server
	err := json.Unmarshal([]byte(testServerJSON), &server)

	assert.NoError(t, err)
	assert.NotEmpty(t, server)

	assert.NotEmpty(t, server.ID)
	assert.NotEmpty(t, server.ServerPlan)
	assert.NotEmpty(t, server.Zone)
	assert.NotEmpty(t, server.Disks)
	assert.NotEmpty(t, server.Interfaces)
	assert.NotEmpty(t, server.Instance)
	assert.NotEmpty(t, server.Tags)

}
