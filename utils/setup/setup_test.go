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

package setup

import (
	"context"
	"testing"
	"time"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/accessor"
	"github.com/sacloud/libsacloud/v2/sacloud/testutil"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
	"github.com/sacloud/libsacloud/v2/utils/query"
)

func TestRetryableSetup(t *testing.T) {
	var switchID types.ID
	testZone := testutil.TestZone()

	testutil.Run(t, &testutil.CRUDTestCase{
		Parallel: true,
		SetupAPICallerFunc: func() sacloud.APICaller {
			return testutil.SingletonAPICaller()
		},

		Setup: func(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) error {
			switchOp := sacloud.NewSwitchOp(caller)
			sw, err := switchOp.Create(ctx, testZone, &sacloud.SwitchCreateRequest{Name: "libsacloud-switch-for-util-setup"})
			if err != nil {
				return err
			}
			switchID = sw.ID
			return nil
		},

		Create: &testutil.CRUDTestFunc{
			Func: func(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
				nfsOp := sacloud.NewNFSOp(caller)
				nfsSetup := &RetryableSetup{
					Create: func(ctx context.Context, zone string) (accessor.ID, error) {
						nfsPlanID, err := query.FindNFSPlanID(ctx, sacloud.NewNoteOp(caller), types.NFSPlans.HDD, types.NFSHDDSizes.Size100GB)
						if err != nil {
							return nil, err
						}
						return nfsOp.Create(ctx, zone, &sacloud.NFSCreateRequest{
							Name:           "libsacloud-nfs-for-util-setup",
							SwitchID:       switchID,
							PlanID:         nfsPlanID,
							IPAddresses:    []string{"192.168.0.11"},
							NetworkMaskLen: 24,
							DefaultRoute:   "192.168.0.1",
						})
					},
					Read: func(ctx context.Context, zone string, id types.ID) (interface{}, error) {
						return nfsOp.Read(ctx, zone, id)
					},
					Delete: func(ctx context.Context, zone string, id types.ID) error {
						return nfsOp.Delete(ctx, zone, id)
					},
					IsWaitForCopy: true,
					IsWaitForUp:   true,
					RetryCount:    3,
				}
				if !testutil.IsAccTest() {
					nfsSetup.ProvisioningRetryInterval = time.Millisecond
					nfsSetup.DeleteRetryInterval = time.Millisecond
					nfsSetup.PollingInterval = time.Millisecond
				}

				return nfsSetup.Setup(ctx, testZone)
			},
		},
		Read: &testutil.CRUDTestFunc{
			Func: func(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
				nfsOp := sacloud.NewNFSOp(caller)
				return nfsOp.Read(ctx, testZone, ctx.ID)
			},
			CheckFunc: func(t testutil.TestT, ctx *testutil.CRUDTestContext, i interface{}) error {
				nfs := i.(*sacloud.NFS)
				return testutil.DoAsserts(
					testutil.AssertEqualFunc(t, types.Availabilities.Available, nfs.Availability, "NFS.Availability"),
				)
			},
		},

		Shutdown: func(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) error {
			nfsOp := sacloud.NewNFSOp(caller)
			return nfsOp.Shutdown(ctx, testZone, ctx.ID, &sacloud.ShutdownOption{Force: true})
		},

		Delete: &testutil.CRUDTestDeleteFunc{
			Func: func(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) error {
				nfsOp := sacloud.NewNFSOp(caller)
				if err := nfsOp.Delete(ctx, testZone, ctx.ID); err != nil {
					return err
				}

				switchOp := sacloud.NewSwitchOp(caller)
				if err := switchOp.Delete(ctx, testZone, switchID); err != nil {
					return err
				}
				return nil
			},
		},
	})
}
