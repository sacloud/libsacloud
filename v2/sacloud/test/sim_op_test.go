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
	"context"
	"os"
	"testing"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/testutil"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
	"github.com/stretchr/testify/assert"
)

func TestSIMOpCRUD(t *testing.T) {
	testutil.PreCheckEnvsFunc("SAKURACLOUD_SIM_ICCID", "SAKURACLOUD_SIM_PASSCODE")(t)

	initSIMVariables()

	testutil.RunCRUD(t, &testutil.CRUDTestCase{
		Parallel:          true,
		IgnoreStartupWait: true,

		SetupAPICallerFunc: singletonAPICaller,

		Create: &testutil.CRUDTestFunc{
			Func: testSIMCreate,
			CheckFunc: testutil.AssertEqualWithExpected(&testutil.CRUDTestExpect{
				ExpectValue:  createSIMExpected,
				IgnoreFields: ignoreSIMFields,
			}),
		},

		Read: &testutil.CRUDTestFunc{
			Func: testSIMRead,
			CheckFunc: testutil.AssertEqualWithExpected(&testutil.CRUDTestExpect{
				ExpectValue:  createSIMExpected,
				IgnoreFields: ignoreSIMFields,
			}),
		},

		Updates: []*testutil.CRUDTestFunc{
			{
				Func: testSIMUpdate,
				CheckFunc: testutil.AssertEqualWithExpected(&testutil.CRUDTestExpect{
					ExpectValue:  updateSIMExpected,
					IgnoreFields: ignoreSIMFields,
				}),
			},
			// activate
			{
				Func: func(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
					client := sacloud.NewSIMOp(caller)
					if err := client.Activate(ctx, ctx.ID); err != nil {
						return nil, err
					}
					return client.Status(ctx, ctx.ID)
				},
				CheckFunc: func(t testutil.TestT, ctx *testutil.CRUDTestContext, v interface{}) error {
					simInfo := v.(*sacloud.SIMInfo)
					return testutil.DoAsserts(
						testutil.AssertNotNilFunc(t, simInfo, "SIMInfo"),
						testutil.AssertTrueFunc(t, simInfo.Activated, "SIMInfo.Activated"),
					)
				},
				SkipExtractID: true,
			},
			// deactivate
			{
				Func: func(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
					client := sacloud.NewSIMOp(caller)
					if err := client.Deactivate(ctx, ctx.ID); err != nil {
						return nil, err
					}
					return client.Status(ctx, ctx.ID)
				},
				CheckFunc: func(t testutil.TestT, ctx *testutil.CRUDTestContext, v interface{}) error {
					simInfo := v.(*sacloud.SIMInfo)
					return testutil.DoAsserts(
						testutil.AssertNotNilFunc(t, simInfo, "SIMInfo"),
						testutil.AssertFalseFunc(t, simInfo.Activated, "SIMInfo.Activated"),
					)
				},
				SkipExtractID: true,
			},
			// IMEI lock
			{
				Func: func(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
					client := sacloud.NewSIMOp(caller)
					if err := client.IMEILock(ctx, ctx.ID, &sacloud.SIMIMEILockRequest{
						IMEI: "123456789012345",
					}); err != nil {
						return nil, err
					}
					return client.Status(ctx, ctx.ID)
				},
				CheckFunc: func(t testutil.TestT, ctx *testutil.CRUDTestContext, v interface{}) error {
					simInfo := v.(*sacloud.SIMInfo)
					return testutil.DoAsserts(
						testutil.AssertTrueFunc(t, simInfo.IMEILock, "SIMInfo.IMEILock"),
					)
				},
				SkipExtractID: true,
			},
			// IMEI unlock
			{
				Func: func(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
					client := sacloud.NewSIMOp(caller)
					if err := client.IMEIUnlock(ctx, ctx.ID); err != nil {
						return nil, err
					}
					return client.Status(ctx, ctx.ID)
				},
				CheckFunc: func(t testutil.TestT, ctx *testutil.CRUDTestContext, v interface{}) error {
					simInfo := v.(*sacloud.SIMInfo)
					return testutil.DoAsserts(
						testutil.AssertFalseFunc(t, simInfo.IMEILock, "SIMInfo.IMEILock"),
					)
				},
				SkipExtractID: true,
			},
			// network operator
			{
				Func: func(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
					client := sacloud.NewSIMOp(caller)
					configs := []*sacloud.SIMNetworkOperatorConfig{
						{
							Name:  "SoftBank",
							Allow: true,
						},
					}
					if err := client.SetNetworkOperator(ctx, ctx.ID, configs); err != nil {
						return nil, err
					}
					return client.GetNetworkOperator(ctx, ctx.ID)
				},
				CheckFunc: func(t testutil.TestT, ctx *testutil.CRUDTestContext, v interface{}) error {
					config := v.([]*sacloud.SIMNetworkOperatorConfig)
					return testutil.DoAsserts(
						testutil.AssertNotEmptyFunc(t, config, "NetworkOperatorConfig"),
					)
				},
				SkipExtractID: true,
			},
		},

		Delete: &testutil.CRUDTestDeleteFunc{
			Func: testSIMDelete,
		},
	})
}

func TestSIMOp_Logs(t *testing.T) {
	if !isAccTest() {
		t.Skip("TestSIMOp_Logs only exec at Acceptance Test")
	}
	testutil.PreCheckEnvsFunc("SAKURACLOUD_SIM_ID")(t)
	id := types.StringID(os.Getenv("SAKURACLOUD_SIM_ID"))

	client := sacloud.NewSIMOp(singletonAPICaller())
	logs, err := client.Logs(context.Background(), id)
	assert.NoError(t, err)
	assert.NotEmpty(t, logs)
}

func initSIMVariables() {
	iccid := os.Getenv("SAKURACLOUD_SIM_ICCID")
	passcode := os.Getenv("SAKURACLOUD_SIM_PASSCODE")

	createSIMParam = &sacloud.SIMCreateRequest{
		Name:        testutil.ResourceName("sim"),
		Description: "desc",
		Tags:        []string{"tag1", "tag2"},
		ICCID:       iccid,
		PassCode:    passcode,
	}
	createSIMExpected = &sacloud.SIM{
		Name:         createSIMParam.Name,
		Description:  createSIMParam.Description,
		Tags:         createSIMParam.Tags,
		Availability: types.Availabilities.Available,
		ICCID:        createSIMParam.ICCID,
	}
	updateSIMParam = &sacloud.SIMUpdateRequest{
		Name:        testutil.ResourceName("sim-upd"),
		Description: "desc-upd",
		Tags:        []string{"tag1-upd", "tag2-upd"},
	}
	updateSIMExpected = &sacloud.SIM{
		Name:         updateSIMParam.Name,
		Description:  updateSIMParam.Description,
		Tags:         updateSIMParam.Tags,
		Availability: types.Availabilities.Available,
		ICCID:        createSIMParam.ICCID,
	}
}

var (
	ignoreSIMFields = []string{
		"ID",
		"Class",
		"IconID",
		"Info",
		"CreatedAt",
		"ModifiedAt",
	}
	createSIMParam    *sacloud.SIMCreateRequest
	createSIMExpected *sacloud.SIM
	updateSIMParam    *sacloud.SIMUpdateRequest
	updateSIMExpected *sacloud.SIM
)

func testSIMCreate(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewSIMOp(caller)
	return client.Create(ctx, createSIMParam)
}

func testSIMRead(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewSIMOp(caller)
	return client.Read(ctx, ctx.ID)
}

func testSIMUpdate(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewSIMOp(caller)
	return client.Update(ctx, ctx.ID, updateSIMParam)
}

func testSIMDelete(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) error {
	client := sacloud.NewSIMOp(caller)
	return client.Delete(ctx, ctx.ID)
}
