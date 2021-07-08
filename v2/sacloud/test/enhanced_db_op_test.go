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

package test

import (
	"testing"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/testutil"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

func TestEnhancedDBOp_CRUD(t *testing.T) {
	testutil.RunCRUD(t, &testutil.CRUDTestCase{
		Parallel:           true,
		SetupAPICallerFunc: singletonAPICaller,
		Create: &testutil.CRUDTestFunc{
			Func: testEnhancedDBCreate,
			CheckFunc: testutil.AssertEqualWithExpected(&testutil.CRUDTestExpect{
				ExpectValue:  createEnhancedDBExpected,
				IgnoreFields: ignoreEnhancedDBFields,
			}),
		},

		Read: &testutil.CRUDTestFunc{
			Func: testEnhancedDBRead,
			CheckFunc: testutil.AssertEqualWithExpected(&testutil.CRUDTestExpect{
				ExpectValue:  createEnhancedDBExpected,
				IgnoreFields: ignoreEnhancedDBFields,
			}),
		},

		Updates: []*testutil.CRUDTestFunc{
			{
				Func: testEnhancedDBUpdate,
				CheckFunc: testutil.AssertEqualWithExpected(&testutil.CRUDTestExpect{
					ExpectValue:  updateEnhancedDBExpected,
					IgnoreFields: ignoreEnhancedDBFields,
				}),
			},
			{
				Func: func(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
					edbOp := sacloud.NewEnhancedDBOp(caller)
					return nil, edbOp.SetPassword(ctx, ctx.ID, &sacloud.EnhancedDBSetPasswordRequest{
						Password: "password",
					})
				},
				SkipExtractID: true,
			},
		},

		Delete: &testutil.CRUDTestDeleteFunc{
			Func: testEnhancedDBDelete,
		},
	})
}

var (
	ignoreEnhancedDBFields = []string{
		"ID",
		"Class",
		"SettingsHash",
		"CreatedAt",
		"ModifiedAt",
	}
	createEnhancedDBParam = &sacloud.EnhancedDBCreateRequest{
		Name:         testutil.ResourceName("enhanced-db"),
		Description:  "desc",
		Tags:         []string{"tag1", "tag2"},
		DatabaseName: testutil.RandomName(10, testutil.CharSetAlpha),
	}
	createEnhancedDBExpected = &sacloud.EnhancedDB{
		Name:           createEnhancedDBParam.Name,
		Description:    createEnhancedDBParam.Description,
		Tags:           createEnhancedDBParam.Tags,
		Availability:   types.Availabilities.Available,
		DatabaseName:   createEnhancedDBParam.DatabaseName,
		DatabaseType:   "tidb",
		Region:         "is1",
		HostName:       createEnhancedDBParam.DatabaseName + ".tidb-is1.db.sakurausercontent.com",
		Port:           3306,
		MaxConnections: 50,
	}
	updateEnhancedDBParam = &sacloud.EnhancedDBUpdateRequest{
		Name:        testutil.ResourceName("enhanced-db-upd"),
		Description: "desc-upd",
		Tags:        []string{"tag1-upd", "tag2-upd"},
		IconID:      testIconID,
	}
	updateEnhancedDBExpected = &sacloud.EnhancedDB{
		Name:           updateEnhancedDBParam.Name,
		Description:    updateEnhancedDBParam.Description,
		Tags:           updateEnhancedDBParam.Tags,
		Availability:   types.Availabilities.Available,
		IconID:         testIconID,
		DatabaseName:   createEnhancedDBParam.DatabaseName,
		DatabaseType:   "tidb",
		Region:         "is1",
		HostName:       createEnhancedDBParam.DatabaseName + ".tidb-is1.db.sakurausercontent.com",
		Port:           3306,
		MaxConnections: 50,
	}
)

func testEnhancedDBCreate(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewEnhancedDBOp(caller)
	return client.Create(ctx, createEnhancedDBParam)
}

func testEnhancedDBRead(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewEnhancedDBOp(caller)
	return client.Read(ctx, ctx.ID)
}

func testEnhancedDBUpdate(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewEnhancedDBOp(caller)
	return client.Update(ctx, ctx.ID, updateEnhancedDBParam)
}

func testEnhancedDBDelete(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) error {
	client := sacloud.NewEnhancedDBOp(caller)
	return client.Delete(ctx, ctx.ID)
}
