package sacloud

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

var testIconJSON = `
{
	"ID": "123456789012",
	"URL": "https:\/\/secure.sakura.ad.jp\/cloud\/zone\/is1b\/api\/cloud\/1.1\/icon\/123456789012.png",
	"Name": "TEST",
	"Scope": "shared"
}
`

func TestMarshalIconJSON(t *testing.T) {
	var icon Icon
	err := json.Unmarshal([]byte(testIconJSON), &icon)

	assert.NoError(t, err)
	assert.NotEmpty(t, icon)

	assert.NotEmpty(t, icon.ID)
}
