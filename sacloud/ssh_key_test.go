// Copyright 2016-2019 The Libsacloud Authors
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

var testSSHKeyJSON = `
{
	"ID": 123456789012,
	"Name": "test_key",
	"Description": "",
	"PublicKey": "ssh-rsa xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
	"Fingerprint": "xx:xx:xx:xx:xx:xx:xx:xx:xx:xx:xx:xx:xx:xx:xx:xx",
	"CreatedAt": "2016-02-15T19:00:01+09:00"
}
`

func TestMarshalSSHKeyJSON(t *testing.T) {
	var key SSHKey
	err := json.Unmarshal([]byte(testSSHKeyJSON), &key)

	assert.NoError(t, err)
	assert.NotEmpty(t, key)

	assert.NotEmpty(t, key.ID)
}
