// Copyright 2016-2020 The Libsacloud Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package mobilegateway

import (
	"github.com/sacloud/libsacloud/v2/helper/wait"
	"testing"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/pointer"
	"github.com/sacloud/libsacloud/v2/sacloud/testutil"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

func TestMobileGatewayService_CRUD(t *testing.T) {
	svc := New(testutil.SingletonAPICaller())
	name := testutil.ResourceName("local-router")
	zone := testutil.TestZone()

	testutil.RunCRUD(t, &testutil.CRUDTestCase{
		Parallel:           true,
		PreCheck:           nil,
		SetupAPICallerFunc: testutil.SingletonAPICaller,
		Setup:              nil,
		IgnoreStartupWait: true,
		Create: &testutil.CRUDTestFunc{
			Func: func(ctx *testutil.CRUDTestContext, _ sacloud.APICaller) (interface{}, error) {
				created, err := svc.Create(&CreateRequest{
					Name:        name,
					Description: "test",
					Tags:        types.Tags{"tag1", "tag2"},
					Zone: zone,
				})
				if err != nil {
					return nil , err
				}
				return wait.UntilMobileGatewayIsReady(ctx, sacloud.NewMobileGatewayOp(svc.caller), zone, created.ID)
			},
		},
		Read: &testutil.CRUDTestFunc{
			Func: func(ctx *testutil.CRUDTestContext, _ sacloud.APICaller) (interface{}, error) {
				return svc.Read(&ReadRequest{ID: ctx.ID, Zone: zone})
			},
		},
		Updates: []*testutil.CRUDTestFunc{
			{
				Func: func(ctx *testutil.CRUDTestContext, _ sacloud.APICaller) (interface{}, error) {
					return svc.Update(&UpdateRequest{
						ID:          ctx.ID,
						Name:        pointer.NewString(name + "-upd"),
						Description: pointer.NewString("test-upd"),
						Zone: zone,
					})
				},
			},
		},
		Delete: &testutil.CRUDTestDeleteFunc{
			Func: func(ctx *testutil.CRUDTestContext, _ sacloud.APICaller) error {
				return svc.Delete(&DeleteRequest{ID: ctx.ID, Zone: zone})
			},
		},
		Cleanup: nil,
	})
}
