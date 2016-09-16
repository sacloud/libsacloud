package sacloud

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

var testAuthStatusJSON = `
{
    "Account": {
        "Class": "account",
        "Code": "your_account_code",
        "ID": "999999999999",
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
