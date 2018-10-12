package sacloud

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testFTPServerJSON = `
{
	"HostName": "xxxxxxxxxxxxxx",
	"IPAddress": "XXX.XXX.XXX.XXX",
	"User": "user",
	"Password": "password"
}
`

func TestMarshalFTPServerJSON(t *testing.T) {
	var ftp FTPServer
	err := json.Unmarshal([]byte(testFTPServerJSON), &ftp)

	assert.NoError(t, err)
	assert.NotEmpty(t, ftp)
	assert.NotEmpty(t, ftp.HostName)
}
