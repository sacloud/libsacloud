package test

import (
	"testing"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/stretchr/testify/assert"
)

func TestAuthStatusOp_Read(t *testing.T) {
	Run(t, &CRUDTestCase{
		Parallel:           true,
		SetupAPICallerFunc: singletonAPICaller,
		Read: &CRUDTestFunc{
			Func: func(ctx *CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
				client := sacloud.NewAuthStatusOp(singletonAPICaller())
				authStatus, err := client.Read(ctx)

				assert.NotNil(t, authStatus)

				return nil, err
			},
		},
	})
}
