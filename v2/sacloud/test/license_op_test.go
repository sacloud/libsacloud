// Copyright 2016-2022 The Libsacloud Authors
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

func TestLicenseOpCRUD(t *testing.T) {
	testutil.RunCRUD(t, &testutil.CRUDTestCase{
		Parallel:           true,
		IgnoreStartupWait:  true,
		SetupAPICallerFunc: singletonAPICaller,
		Create: &testutil.CRUDTestFunc{
			Func: testLicenseCreate,
			CheckFunc: testutil.AssertEqualWithExpected(&testutil.CRUDTestExpect{
				ExpectValue:  createLicenseExpected,
				IgnoreFields: ignoreLicenseFields,
			}),
		},
		Read: &testutil.CRUDTestFunc{
			Func: testLicenseRead,
			CheckFunc: testutil.AssertEqualWithExpected(&testutil.CRUDTestExpect{
				ExpectValue:  createLicenseExpected,
				IgnoreFields: ignoreLicenseFields,
			}),
		},
		Updates: []*testutil.CRUDTestFunc{
			{
				Func: testLicenseUpdate,
				CheckFunc: testutil.AssertEqualWithExpected(&testutil.CRUDTestExpect{
					ExpectValue:  updateLicenseExpected,
					IgnoreFields: ignoreLicenseFields,
				}),
			},
		},
		Delete: &testutil.CRUDTestDeleteFunc{
			Func: testLicenseDelete,
		},
	})
}

var (
	ignoreLicenseFields = []string{
		"ID",
		"CreatedAt",
		"ModifiedAt",
	}

	createLicenseParam = &sacloud.LicenseCreateRequest{
		Name:          testutil.ResourceName("license"),
		LicenseInfoID: types.ID(10001),
	}
	createLicenseExpected = &sacloud.License{
		Name:            createLicenseParam.Name,
		LicenseInfoID:   createLicenseParam.LicenseInfoID,
		LicenseInfoName: "Windows RDS SAL",
	}
	updateLicenseParam = &sacloud.LicenseUpdateRequest{
		Name: testutil.ResourceName("license-upd"),
	}
	updateLicenseExpected = &sacloud.License{
		Name:            updateLicenseParam.Name,
		LicenseInfoID:   createLicenseParam.LicenseInfoID,
		LicenseInfoName: "Windows RDS SAL",
	}
)

func testLicenseCreate(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewLicenseOp(caller)
	return client.Create(ctx, createLicenseParam)
}

func testLicenseRead(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewLicenseOp(caller)
	return client.Read(ctx, ctx.ID)
}

func testLicenseUpdate(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewLicenseOp(caller)
	return client.Update(ctx, ctx.ID, updateLicenseParam)
}

func testLicenseDelete(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) error {
	client := sacloud.NewLicenseOp(caller)
	return client.Delete(ctx, ctx.ID)
}
