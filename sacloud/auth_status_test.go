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

var testAuthStatusJSON = `
{
    "Account": {
        "Class": "account",
        "Code": "your_account_code",
        "ID": 999999999999,
        "Name": "your_account_name"
    },
    "AuthClass": "account",
    "AuthMethod": "apikey",
    "ExternalPermission": "none",
    "IsAPIKey": true,
    "Member": {
        "Class": "xxxxxx",
        "Code": "xxx99999",
        "Errors": []
    },
    "OperationPenalty": "none",
    "Permission": "create",
    "RESTFilter": null,
    "User": null,
    "is_ok": true
}
`

func TestMarshalAuthStatusJSON(t *testing.T) {
	var authStatus AuthStatus
	err := json.Unmarshal([]byte(testAuthStatusJSON), &authStatus)

	assert.NoError(t, err)
	assert.NotEmpty(t, authStatus)

	assert.NotNil(t, authStatus.Account)
}
