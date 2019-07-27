package setup

import (
	"context"
	"testing"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/accessor"
	"github.com/sacloud/libsacloud/v2/sacloud/testutil"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
	"github.com/sacloud/libsacloud/v2/utils/nfs"
)

func TestRetryableSetup(t *testing.T) {
	var switchID types.ID
	testZone := testutil.TestZone()

	testutil.Run(t, &testutil.CRUDTestCase{
		Parallel:           true,
		SetupAPICallerFunc: testutil.SingletonAPICaller,

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
						nfsPlanID, err := nfs.FindNFSPlanID(ctx, sacloud.NewNoteOp(caller), types.NFSPlans.HDD, types.NFSHDDSizes.Size100GB)
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
