package sacloud

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

var testDiskPlanJSON = `
{
            "ID": 4,
            "StorageClass": "iscsi1204",
            "DisplayOrder": 400,
            "Name": "SSD\u30d7\u30e9\u30f3",
            "Description": "",
            "Availability": "available",
            "Size": [
                {
                    "SizeMB": 20480,
                    "DisplaySize": 20,
                    "DisplaySuffix": "GB",
                    "Availability": "available",
                    "ServiceClass": "cloud\/disk\/ssd\/20g"
                }
            ],
            "is_ok": true
}
`

func TestMarshalProductDiskJSON(t *testing.T) {
	var productDisk ProductDisk
	err := json.Unmarshal([]byte(testDiskPlanJSON), &productDisk)

	assert.NoError(t, err)
	assert.NotEmpty(t, productDisk)

	assert.NotEmpty(t, productDisk.ID)
	assert.NotEmpty(t, productDisk.Size)
	assert.NotEmpty(t, productDisk.Size[0].SizeMB)
}
