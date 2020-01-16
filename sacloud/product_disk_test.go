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

package sacloud

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
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
