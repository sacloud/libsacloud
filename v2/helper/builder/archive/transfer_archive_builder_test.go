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

func TestTransferArchiveBuilder_Build(t *testing.T) {
	zoneFrom := "is1a"
	zoneTo := "is1b"
	var sourceArchive *sacloud.Archive

	testutil.Run(t, &testutil.CRUDTestCase{
		SetupAPICallerFunc: func() sacloud.APICaller {
			return testutil.SingletonAPICaller()
		},
		Parallel:          true,
		IgnoreStartupWait: true,
		Setup: func(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) error {
			archiveOp := sacloud.NewArchiveOp(caller)
			source, err := query.FindArchiveByOSType(ctx, archiveOp, zoneFrom, ostype.CentOS)
			if err != nil {
				return err
			}

			created, err := archiveOp.Create(ctx, zoneFrom, &sacloud.ArchiveCreateRequest{
				SourceArchiveID: source.ID,
				Name:            testutil.ResourceName("source-archive-for-transfer"),
			})
			if err != nil {
				return err
			}
			sourceArchive = created
			_, err = sacloud.WaiterForReady(func() (interface{}, error) {
				return archiveOp.Read(ctx, zoneFrom, sourceArchive.ID)
			}).WaitForState(ctx)

			return err
		},
		Create: &testutil.CRUDTestFunc{
			Func: func(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
				builder := &TransferArchiveBuilder{
					Name:              testutil.ResourceName("archive-from-other-zone"),
					Description:       "description",
					Tags:              types.Tags{"tag1", "tag2"},
					SourceArchiveID:   sourceArchive.ID,
					SourceArchiveZone: zoneFrom,
					Client:            NewAPIClient(caller),
				}
				return builder.Build(ctx, zoneTo)
			},
		},
		Read: &testutil.CRUDTestFunc{
			Func: func(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
				return sacloud.NewArchiveOp(caller).Read(ctx, zoneTo, ctx.ID)
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
					return archiveOp.Read(ctx, zoneTo, ctx.ID)
				}).WaitForState(ctx)
				if err != nil {
					return err
				}
				if err := archiveOp.Delete(ctx, zoneTo, ctx.ID); err != nil {
					return err
				}

				if sourceArchive != nil {
					if err := archiveOp.Delete(ctx, zoneFrom, sourceArchive.ID); err != nil {
						return err
					}
				}
				return nil
			},
		},
	})
}
