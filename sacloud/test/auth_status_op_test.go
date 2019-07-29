package test

import (
	"testing"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/testutil"
	"github.com/stretchr/testify/assert"
)

func TestAuthStatusOp_Read(t *testing.T) {
	testutil.Run(t, &testutil.CRUDTestCase{
		Parallel:           true,
		SetupAPICallerFunc: singletonAPICaller,
		Read: &testutil.CRUDTestFunc{
			Func: func(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
				client := sacloud.NewAuthStatusOp(singletonAPICaller())
				authStatus, err := client.Read(ctx)

				assert.NotNil(t, authStatus)

				return nil, err
			},
		},
	})
}
