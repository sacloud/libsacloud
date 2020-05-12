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

package archive

import (
	"testing"

	"github.com/sacloud/libsacloud/v2/helper/query"
	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/ostype"
	"github.com/sacloud/libsacloud/v2/sacloud/testutil"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

func TestStandardArchiveBuilder_Build(t *testing.T) {
	testZone := testutil.TestZone()
	var sourceArchive *sacloud.Archive

	testutil.Run(t, &testutil.CRUDTestCase{
		SetupAPICallerFunc: func() sacloud.APICaller {
			return testutil.SingletonAPICaller()
		},
		Parallel:          true,
		IgnoreStartupWait: true,
		Setup: func(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) error {
			archiveOp := sacloud.NewArchiveOp(caller)
			source, err := query.FindArchiveByOSType(ctx, archiveOp, testZone, ostype.CentOS)
			if err != nil {
				return err
			}
			sourceArchive = source
			return nil
		},
		Create: &testutil.CRUDTestFunc{
			Func: func(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
				builder := &StandardArchiveBuilder{
					Name:            testutil.ResourceName("standard-archive-builder"),
					Description:     "description",
					Tags:            types.Tags{"tag1", "tag2"},
					SourceArchiveID: sourceArchive.ID,
					Client:          NewAPIClient(caller),
				}
				return builder.Build(ctx, testZone)
			},
		},
		Read: &testutil.CRUDTestFunc{
			Func: func(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
				return sacloud.NewArchiveOp(caller).Read(ctx, testZone, ctx.ID)
			},
			CheckFunc: func(t testutil.TestT, ctx *testutil.CRUDTestContext, value interface{}) error {
				archive := value.(*sacloud.Archive)
				return testutil.DoAsserts(
					testutil.AssertNotNilFunc(t, archive, "Archive"),
				)
			},
		},
		Delete: &testutil.CRUDTestDeleteFunc{
			Func: func(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) error {
				archiveOp := sacloud.NewArchiveOp(caller)

				_, err := sacloud.WaiterForReady(func() (interface{}, error) {
					return archiveOp.Read(ctx, testZone, ctx.ID)
				}).WaitForState(ctx)
				if err != nil {
					return err
				}
				return archiveOp.Delete(ctx, testZone, ctx.ID)
			},
		},
	})
}
