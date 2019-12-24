// Copyright 2016-2019 The Libsacloud Authors
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

package registry

import (
	"testing"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/testutil"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

func TestBuilder_Build(t *testing.T) {
	testutil.Run(t, &testutil.CRUDTestCase{
		SetupAPICallerFunc: func() sacloud.APICaller {
			return testutil.SingletonAPICaller()
		},
		Parallel:          true,
		IgnoreStartupWait: true,
		Create: &testutil.CRUDTestFunc{
			Func: func(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
				builder := &Builder{
					Name:        testutil.ResourceName("container-registry-builder"),
					Description: "description",
					Tags:        types.Tags{"tag1", "tag2"},
					AccessLevel: types.ContainerRegistryAccessLevels.None,
					NamePrefix:  testutil.RandomName(60, testutil.CharSetAlpha),
					Users: []*User{
						{
							UserName: "user1",
							Password: "password",
						},
						{
							UserName: "user2",
							Password: "password",
						},
					},
					Client: NewAPIClient(caller),
				}
				return builder.Build(ctx)
			},
		},
		Read: &testutil.CRUDTestFunc{
			Func: func(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
				regOp := sacloud.NewContainerRegistryOp(caller)
				return regOp.Read(ctx, ctx.ID)
			},
			CheckFunc: func(t testutil.TestT, ctx *testutil.CRUDTestContext, value interface{}) error {
				reg := value.(*sacloud.ContainerRegistry)
				return testutil.DoAsserts(
					testutil.AssertNotNilFunc(t, reg, "ContainerRegistry"),
				)
			},
		},
		Delete: &testutil.CRUDTestDeleteFunc{
			Func: func(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) error {
				regOp := sacloud.NewContainerRegistryOp(caller)
				return regOp.Delete(ctx, ctx.ID)
			},
		},
	})
}
