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
