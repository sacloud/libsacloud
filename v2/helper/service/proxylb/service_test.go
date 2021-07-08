// Copyright 2016-2021 The Libsacloud Authors
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

package proxylb

import (
	"os"
	"testing"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/pointer"
	"github.com/sacloud/libsacloud/v2/sacloud/testutil"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

func TestProxyLBService_CRUD(t *testing.T) {
	svc := New(testutil.SingletonAPICaller())
	name := testutil.ResourceName("proxylb")

	testutil.RunCRUD(t, &testutil.CRUDTestCase{
		Parallel:           true,
		PreCheck:           testutil.PreCheckEnvsFunc("SAKURACLOUD_PROXYLB_SERVER0"),
		SetupAPICallerFunc: testutil.SingletonAPICaller,
		Setup:              nil,
		IgnoreStartupWait:  true,
		Create: &testutil.CRUDTestFunc{
			Func: func(ctx *testutil.CRUDTestContext, _ sacloud.APICaller) (interface{}, error) {
				return svc.Create(&CreateRequest{
					Name:        name,
					Description: "test",
					Tags:        types.Tags{"tag1", "tag2"},
					Plan:        types.ProxyLBPlans.CPS100,
					SorryServer: &sacloud.ProxyLBSorryServer{
						IPAddress: os.Getenv("SAKURACLOUD_PROXYLB_SERVER0"),
						Port:      80,
					},
					HealthCheck: &sacloud.ProxyLBHealthCheck{
						Protocol:  types.ProxyLBProtocols.TCP,
						DelayLoop: 10,
					},
					StickySession: &sacloud.ProxyLBStickySession{
						Enabled: true,
						Method:  "cookie",
					},
					UseVIPFailover: false,
					Syslog: &sacloud.ProxyLBSyslog{
						Server: "",
						Port:   514,
					},
					Timeout: &sacloud.ProxyLBTimeout{
						InactiveSec: 10,
					},
					Region: types.ProxyLBRegions.IS1,
				})
			},
		},
		Read: &testutil.CRUDTestFunc{
			Func: func(ctx *testutil.CRUDTestContext, _ sacloud.APICaller) (interface{}, error) {
				return svc.Read(&ReadRequest{ID: ctx.ID})
			},
		},
		Updates: []*testutil.CRUDTestFunc{
			{
				Func: func(ctx *testutil.CRUDTestContext, _ sacloud.APICaller) (interface{}, error) {
					return svc.Update(&UpdateRequest{
						ID:          ctx.ID,
						Name:        pointer.NewString(name + "-upd"),
						Description: pointer.NewString("test-upd"),
					})
				},
			},
		},
		Delete: &testutil.CRUDTestDeleteFunc{
			Func: func(ctx *testutil.CRUDTestContext, _ sacloud.APICaller) error {
				return svc.Delete(&DeleteRequest{ID: ctx.ID})
			},
		},
		Cleanup: nil,
	})
}
