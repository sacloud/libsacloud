package api

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReadAuthStatus(t *testing.T) {
	api := client.AuthStatus

	res, err := api.Read()

	assert.NoError(t, err)
	assert.NotEmpty(t, res)
	assert.NotNil(t, res.Account)
}
