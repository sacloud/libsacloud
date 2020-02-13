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
)

func TestContainerRegistryOp_CRUD(t *testing.T) {
	testutil.Run(t, &testutil.CRUDTestCase{
		Parallel:           true,
		SetupAPICallerFunc: singletonAPICaller,
		Create: &testutil.CRUDTestFunc{
			Func: testContainerRegistryCreate,
			CheckFunc: testutil.AssertEqualWithExpected(&testutil.CRUDTestExpect{
				ExpectValue:  createContainerRegistryExpected,
				IgnoreFields: ignoreContainerRegistryFields,
			}),
		},

		Read: &testutil.CRUDTestFunc{
			Func: testContainerRegistryRead,
			CheckFunc: testutil.AssertEqualWithExpected(&testutil.CRUDTestExpect{
				ExpectValue:  createContainerRegistryExpected,
				IgnoreFields: ignoreContainerRegistryFields,
			}),
		},

		Updates: []*testutil.CRUDTestFunc{
			{
				Func: testContainerRegistryUpdate,
				CheckFunc: testutil.AssertEqualWithExpected(&testutil.CRUDTestExpect{
					ExpectValue:  updateContainerRegistryExpected,
					IgnoreFields: ignoreContainerRegistryFields,
				}),
			},
			{
				Func: testContainerRegistryUpdateSettings,
				CheckFunc: testutil.AssertEqualWithExpected(&testutil.CRUDTestExpect{
					ExpectValue:  updateContainerRegistrySettingsExpected,
					IgnoreFields: ignoreContainerRegistryFields,
				}),
			},
			{
				Func: func(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
					registryOp := sacloud.NewContainerRegistryOp(caller)
					err := registryOp.AddUser(ctx, ctx.ID, &sacloud.ContainerRegistryUserCreateRequest{
						UserName:   "user1",
						Password:   "password",
						Permission: types.ContainerRegistryAccessLevels.ReadWrite,
					})
					if err != nil {
						return nil, err
					}
					err = registryOp.AddUser(ctx, ctx.ID, &sacloud.ContainerRegistryUserCreateRequest{
						UserName:   "user2",
						Password:   "password",
						Permission: types.ContainerRegistryAccessLevels.ReadOnly,
					})
					if err != nil {
						return nil, err
					}
					return registryOp.ListUsers(ctx, ctx.ID)
				},
				CheckFunc: func(t testutil.TestT, ctx *testutil.CRUDTestContext, value interface{}) error {
					users := value.(*sacloud.ContainerRegistryUsers).Users
					return testutil.DoAsserts(
						testutil.AssertLenFunc(t, users, 2, "ContainerRegistry.Users"),
						testutil.AssertEqualFunc(t, "user1", users[0].UserName, "ContainerRegistry.Users"),
						testutil.AssertEqualFunc(t, types.ContainerRegistryAccessLevels.ReadWrite, users[0].Permission, "ContainerRegistry.Permission"),
						testutil.AssertEqualFunc(t, "user2", users[1].UserName, "ContainerRegistry.Users"),
						testutil.AssertEqualFunc(t, types.ContainerRegistryAccessLevels.ReadOnly, users[1].Permission, "ContainerRegistry.Permission"),
					)
				},
				SkipExtractID: true,
			},
			{
				Func: func(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
					registryOp := sacloud.NewContainerRegistryOp(caller)
					if err := registryOp.DeleteUser(ctx, ctx.ID, "user1"); err != nil {
						return nil, err
					}
					return registryOp.ListUsers(ctx, ctx.ID)
				},
				CheckFunc: func(t testutil.TestT, ctx *testutil.CRUDTestContext, value interface{}) error {
					users := value.(*sacloud.ContainerRegistryUsers).Users
					return testutil.DoAsserts(
						testutil.AssertLenFunc(t, users, 1, "ContainerRegistry.Users"),
					)
				},
				SkipExtractID: true,
			},
		},

		Delete: &testutil.CRUDTestDeleteFunc{
			Func: testContainerRegistryDelete,
		},
	})
}

var (
	ignoreContainerRegistryFields = []string{
		"ID",
		"Class",
		"SettingsHash",
		"CreatedAt",
		"ModifiedAt",
	}
	createContainerRegistryParam = &sacloud.ContainerRegistryCreateRequest{
		Name:           testutil.ResourceName("container-registry"),
		Description:    "desc",
		Tags:           []string{"tag1", "tag2"},
		AccessLevel:    types.ContainerRegistryAccessLevels.ReadWrite,
		SubDomainLabel: testutil.RandomName(60, testutil.CharSetAlpha),
	}
	createContainerRegistryExpected = &sacloud.ContainerRegistry{
		Name:           createContainerRegistryParam.Name,
		Description:    createContainerRegistryParam.Description,
		Tags:           createContainerRegistryParam.Tags,
		Availability:   types.Availabilities.Available,
		AccessLevel:    createContainerRegistryParam.AccessLevel,
		SubDomainLabel: createContainerRegistryParam.SubDomainLabel,
		FQDN:           createContainerRegistryParam.SubDomainLabel + ".sakuracr.jp",
	}
	updateContainerRegistryParam = &sacloud.ContainerRegistryUpdateRequest{
		Name:        testutil.ResourceName("container-registry-upd"),
		Description: "desc-upd",
		Tags:        []string{"tag1-upd", "tag2-upd"},
		IconID:      testIconID,
		AccessLevel: types.ContainerRegistryAccessLevels.ReadOnly,
	}
	updateContainerRegistryExpected = &sacloud.ContainerRegistry{
		Name:           updateContainerRegistryParam.Name,
		Description:    updateContainerRegistryParam.Description,
		Tags:           updateContainerRegistryParam.Tags,
		Availability:   types.Availabilities.Available,
		IconID:         testIconID,
		AccessLevel:    updateContainerRegistryParam.AccessLevel,
		SubDomainLabel: createContainerRegistryParam.SubDomainLabel,
		FQDN:           createContainerRegistryParam.SubDomainLabel + ".sakuracr.jp",
	}

	updateContainerRegistrySettingsParam = &sacloud.ContainerRegistryUpdateSettingsRequest{
		AccessLevel: types.ContainerRegistryAccessLevels.None,
	}
	updateContainerRegistrySettingsExpected = &sacloud.ContainerRegistry{
		Name:           updateContainerRegistryParam.Name,
		Description:    updateContainerRegistryParam.Description,
		Tags:           updateContainerRegistryParam.Tags,
		Availability:   types.Availabilities.Available,
		IconID:         testIconID,
		AccessLevel:    updateContainerRegistrySettingsParam.AccessLevel,
		SubDomainLabel: createContainerRegistryParam.SubDomainLabel,
		FQDN:           createContainerRegistryParam.SubDomainLabel + ".sakuracr.jp",
	}
)

func testContainerRegistryCreate(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewContainerRegistryOp(caller)
	return client.Create(ctx, createContainerRegistryParam)
}

func testContainerRegistryRead(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewContainerRegistryOp(caller)
	return client.Read(ctx, ctx.ID)
}

func testContainerRegistryUpdate(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewContainerRegistryOp(caller)
	return client.Update(ctx, ctx.ID, updateContainerRegistryParam)
}

func testContainerRegistryUpdateSettings(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewContainerRegistryOp(caller)
	return client.UpdateSettings(ctx, ctx.ID, updateContainerRegistrySettingsParam)
}

func testContainerRegistryDelete(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) error {
	client := sacloud.NewContainerRegistryOp(caller)
	return client.Delete(ctx, ctx.ID)
}
