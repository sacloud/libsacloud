package api

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetUbuntuArchiveID(t *testing.T) {
	archiveAPI := client.Archive
	id, err := archiveAPI.GetUbuntuArchiveID()
	assert.NoError(t, err)
	assert.NotEmpty(t, id)
	t.Logf("ubuntu archive ID : %s", id)
}
