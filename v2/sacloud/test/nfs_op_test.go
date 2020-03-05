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

package test

import (
	"testing"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/testutil"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
	"github.com/sacloud/libsacloud/v2/utils/power"
	"github.com/sacloud/libsacloud/v2/utils/query"
)

func TestNFSOp_CRUD(t *testing.T) {
	testutil.Run(t, &testutil.CRUDTestCase{
		Parallel: true,

		SetupAPICallerFunc: singletonAPICaller,
		Setup: func(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) error {
			err := setupSwitchFunc("nfs",
				createNFSParam,
				createNFSExpected,
				updateNFSExpected,
				updateNFSToMinExpected,
			)(ctx, caller)
			if err != nil {
				return err
			}

			// find plan id
			planID, err := query.FindNFSPlanID(ctx, sacloud.NewNoteOp(caller), types.NFSPlans.HDD, types.NFSHDDSizes.Size100GB)
			if err != nil {
				return err
			}
			createNFSParam.PlanID = planID
			createNFSExpected.PlanID = planID
			updateNFSExpected.PlanID = planID
			updateNFSToMinExpected.PlanID = planID
			return nil
		},

		Create: &testutil.CRUDTestFunc{
			Func: testNFSCreate,
			CheckFunc: testutil.AssertEqualWithExpected(&testutil.CRUDTestExpect{
				ExpectValue:  createNFSExpected,
				IgnoreFields: ignoreNFSFields,
			}),
		},

		Read: &testutil.CRUDTestFunc{
			Func: testNFSRead,
			CheckFunc: testutil.AssertEqualWithExpected(&testutil.CRUDTestExpect{
				ExpectValue:  createNFSExpected,
				IgnoreFields: ignoreNFSFields,
			}),
		},

		Updates: []*testutil.CRUDTestFunc{
			{
				Func: testNFSUpdate,
				CheckFunc: testutil.AssertEqualWithExpected(&testutil.CRUDTestExpect{
					ExpectValue:  updateNFSExpected,
					IgnoreFields: ignoreNFSFields,
				}),
			},
			{
				Func: testNFSUpdateToMin,
				CheckFunc: testutil.AssertEqualWithExpected(&testutil.CRUDTestExpect{
					ExpectValue:  updateNFSToMinExpected,
					IgnoreFields: ignoreNFSFields,
				}),
			},
		},

		Shutdown: func(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) error {
			client := sacloud.NewNFSOp(caller)
			return power.ShutdownNFS(ctx, client, testZone, ctx.ID, true)
		},

		Delete: &testutil.CRUDTestDeleteFunc{
			Func: testNFSDelete,
		},

		Cleanup: cleanupSwitchFunc("nfs"),
	})
}

var (
	ignoreNFSFields = []string{
		"ID",
		"Class",
		"Availability",
		"InstanceStatus",
		"InstanceHostName",
		"InstanceHostInfoURL",
		"InstanceStatusChangedAt",
		"Interfaces",
		"SwitchName",
		"ZoneID",
		"CreatedAt",
		"ModifiedAt",
	}
	createNFSParam = &sacloud.NFSCreateRequest{
		// PlanID:      type.ID(0), // プランIDはSetUpで設定する
		IPAddresses:    []string{"192.168.0.11"},
		NetworkMaskLen: 24,
		DefaultRoute:   "192.168.0.1",
		Name:           testutil.ResourceName("nfs"),
		Description:    "desc",
		Tags:           []string{"tag1", "tag2"},
	}
	createNFSExpected = &sacloud.NFS{
		Name:           createNFSParam.Name,
		Description:    createNFSParam.Description,
		Tags:           createNFSParam.Tags,
		PlanID:         createNFSParam.PlanID,
		DefaultRoute:   createNFSParam.DefaultRoute,
		NetworkMaskLen: createNFSParam.NetworkMaskLen,
		IPAddresses:    createNFSParam.IPAddresses,
	}
	updateNFSParam = &sacloud.NFSUpdateRequest{
		Name:        testutil.ResourceName("nfs-upd"),
		Tags:        []string{"tag1-upd", "tag2-upd"},
		Description: "desc-upd",
		IconID:      testIconID,
	}
	updateNFSExpected = &sacloud.NFS{
		Name:           updateNFSParam.Name,
		Description:    updateNFSParam.Description,
		Tags:           updateNFSParam.Tags,
		DefaultRoute:   createNFSParam.DefaultRoute,
		NetworkMaskLen: createNFSParam.NetworkMaskLen,
		IPAddresses:    createNFSParam.IPAddresses,
		IconID:         testIconID,
	}
	updateNFSToMinParam = &sacloud.NFSUpdateRequest{
		Name: testutil.ResourceName("nfs-to-min"),
	}
	updateNFSToMinExpected = &sacloud.NFS{
		Name:           updateNFSToMinParam.Name,
		DefaultRoute:   createNFSParam.DefaultRoute,
		NetworkMaskLen: createNFSParam.NetworkMaskLen,
		IPAddresses:    createNFSParam.IPAddresses,
	}
)

func testNFSCreate(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewNFSOp(caller)
	return client.Create(ctx, testZone, createNFSParam)
}

func testNFSRead(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewNFSOp(caller)
	return client.Read(ctx, testZone, ctx.ID)
}

func testNFSUpdate(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewNFSOp(caller)
	return client.Update(ctx, testZone, ctx.ID, updateNFSParam)
}

func testNFSUpdateToMin(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewNFSOp(caller)
	return client.Update(ctx, testZone, ctx.ID, updateNFSToMinParam)
}

func testNFSDelete(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) error {
	client := sacloud.NewNFSOp(caller)
	return client.Delete(ctx, testZone, ctx.ID)
}
