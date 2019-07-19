package test

import (
	"context"
	"testing"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/stretchr/testify/assert"
)

func TestAuthStatusOp_Read(t *testing.T) {
	Run(t, &CRUDTestCase{
		Parallel:           true,
		SetupAPICallerFunc: singletonAPICaller,
		Read: &CRUDTestFunc{
			Func: func(testContext *CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
				client := sacloud.NewAuthStatusOp(singletonAPICaller())
				authStatus, err := client.Read(context.Background())

				assert.NotNil(t, authStatus)

				return nil, err
			},
		},
	})
}
