package sacloud

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

var testPrivateHostPlanJSON = `
{
      "Index": 0,
      "ID": 112900526366,
      "Name": "200Core 224GB 標準",
      "Description": "200Core\n224GB\n標準",
      "Class": "dynamic",
      "CPU": 200,
      "MemoryMB": 229376,
      "ServiceClass": "cloud/host/200core-224gb-std",
      "Availability": "available"
    }
`

func TestMarshalProductPrivateHostJSON(t *testing.T) {
	var productPrivateHost ProductPrivateHost
	err := json.Unmarshal([]byte(testPrivateHostPlanJSON), &productPrivateHost)

	assert.NoError(t, err)
	assert.NotEmpty(t, productPrivateHost)
	assert.NotEmpty(t, productPrivateHost.ID)
}
