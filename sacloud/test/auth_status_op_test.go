package test

import (
	"context"
	"testing"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/stretchr/testify/require"
)

func TestAuthStatusOp_Read(t *testing.T) {
	client := sacloud.NewAuthStatusOp(singletonAPICaller())

	authStatus, err := client.Read(context.Background(), sacloud.APIDefaultZone)
	require.NoError(t, err)
	require.NotNil(t, authStatus)
}
