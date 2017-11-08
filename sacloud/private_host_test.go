package sacloud

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

var testPrivateHostJSON = `
	{
	    "ID": 112900645140,
	    "Name": "private_host_name",
	    "Description": "private_host_description",
	    "AssignedCPU": 4,
	    "AssignedMemoryMB": 1024,
	    "CreatedAt": "2017-05-25T15:21:02+09:00",
	    "Plan": {
	      "ID": 112900526366,
	      "Name": "200Core 224GB 標準",
	      "Class": "dynamic",
	      "CPU": 200,
	      "MemoryMB": 229376,
	      "ServiceClass": "cloud/host/200core-224gb-std"
	    },
	    "Host": {
	      "Name": "sac-tk1a-sv511"
	    },
	    "Icon": null,
	    "Tags": [
	      "t1",
	      "t2",
	      "t3"
	    ]
	  }
`

func TestMarshalPrivateHostJSON(t *testing.T) {
	var privateHost *PrivateHost

	err := json.Unmarshal([]byte(testPrivateHostJSON), &privateHost)
	assert.NoError(t, err)
	assert.NotEmpty(t, privateHost)
	assert.NotEmpty(t, privateHost.ID)
	assert.EqualValues(t, privateHost.Name, "private_host_name")
	assert.EqualValues(t, privateHost.Description, "private_host_description")
	assert.EqualValues(t, privateHost.AssignedCPU, 4)
	assert.EqualValues(t, privateHost.AssignedMemoryMB, 1024)
	assert.NotNil(t, privateHost.Plan)
	assert.NotNil(t, privateHost.Host)
}
