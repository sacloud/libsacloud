package sacloud

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
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
