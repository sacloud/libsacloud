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
