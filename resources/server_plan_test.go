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

func TestMarshalServerPlanJSON(t *testing.T) {
	var serverPlan ServerPlan
	err := json.Unmarshal([]byte(testServerPlanJSON), &serverPlan)

	assert.NoError(t, err)
	assert.NotEmpty(t, serverPlan)

	assert.NotEmpty(t, serverPlan.ID)
}
