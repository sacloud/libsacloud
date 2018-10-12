package sacloud

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCDROMJSON = `
{
	"ID": 123456789012,
	"DisplayOrder": 100904201011,
	"StorageClass": "iscsi1204",
	"Name": "openSUSE 42.1 64bit",
	"Description": "openSUSE-Leap-42.1-DVD-x86_64.iso",
	"SizeMB": 5120,
	"Scope": "shared",
	"Availability": "available",
	"ServiceClass": null,
	"CreatedAt": "2015-12-24T18:14:57+09:00",
	"Icon": null,
	"Storage": ` + testStorageJSON + `
}
`

func TestMarshalCDROMSON(t *testing.T) {
	var cd CDROM
	err := json.Unmarshal([]byte(testCDROMJSON), &cd)

	assert.NoError(t, err)
	assert.NotEmpty(t, cd)

	assert.NotEmpty(t, cd.ID)
	assert.NotEmpty(t, cd.Storage)
}
