package sacloud

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
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
