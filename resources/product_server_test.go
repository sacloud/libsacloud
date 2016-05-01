package resources

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

var testServerPlanJSON = `
{
	"ID": 1001,
	"Name": "\u30d7\u30e9\u30f3\/1Core-1GB",
	"CPU": 1,
	"MemoryMB": 1024,
	"ServiceClass": "cloud\/plan\/1core-1gb"
}
`

func TestMarshalProductServerJSON(t *testing.T) {
	var productServer ProductServer
	err := json.Unmarshal([]byte(testServerPlanJSON), &productServer)

	assert.NoError(t, err)
	assert.NotEmpty(t, productServer)

	assert.NotEmpty(t, productServer.ID)
}
