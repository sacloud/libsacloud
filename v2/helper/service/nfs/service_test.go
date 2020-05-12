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

package nfs

import (
	"testing"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/pointer"
	"github.com/sacloud/libsacloud/v2/sacloud/testutil"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

func TestNFSService_CRUD(t *testing.T) {
	svc := New(testutil.SingletonAPICaller())
	name := testutil.ResourceName("nfs")
	zone := testutil.TestZone()
	var sw *sacloud.Switch

	testutil.RunCRUD(t, &testutil.CRUDTestCase{
		Parallel:           true,
		PreCheck:           nil,
		SetupAPICallerFunc: testutil.SingletonAPICaller,
		Setup: func(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) error {
			s, err := sacloud.NewSwitchOp(caller).Create(ctx, zone, &sacloud.SwitchCreateRequest{Name: name})
			if err != nil {
				return err
			}
			sw = s

			return err
		},
		Create: &testutil.CRUDTestFunc{
			Func: func(ctx *testutil.CRUDTestContext, _ sacloud.APICaller) (interface{}, error) {
				return svc.Create(&CreateRequest{
					Name:           name,
					Description:    "test",
					Tags:           types.Tags{"tag1", "tag2"},
					Zone:           zone,
					SwitchID:       sw.ID,
					Plan:           types.NFSPlans.SSD,
					Size:           types.NFSSSDSizes.Size100GB,
					IPAddresses:    []string{"192.168.0.11"},
					NetworkMaskLen: 24,
					DefaultRoute:   "192.168.0.1",
				})
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
						Zone:        zone,
					})
				},
			},
		},
		Delete: &testutil.CRUDTestDeleteFunc{
			Func: func(ctx *testutil.CRUDTestContext, _ sacloud.APICaller) error {
				return svc.Delete(&DeleteRequest{ID: ctx.ID, Zone: zone})
			},
		},
		Shutdown: func(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) error {
			return svc.Shutdown(&ShutdownRequest{
				Zone:          zone,
				ID:            ctx.ID,
				ForceShutdown: true,
			})
		},
		Cleanup: nil,
	})
}
