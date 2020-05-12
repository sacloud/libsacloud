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

package disk

import (
	"context"
	"errors"
	"testing"

	"github.com/sacloud/libsacloud/v2/helper/wait"
	"github.com/sacloud/libsacloud/v2/pkg/size"
	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/testutil"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

func TestService_connectAndDisconnect(t *testing.T) {
	prefix := testutil.RandomPrefix()
	name := prefix + "disk-service"
	zone := testutil.TestZone()

	var server *sacloud.Server
	var disk *sacloud.Disk

	testutil.RunResource(t, &testutil.ResourceTestCase{
		PreCheck:           nil,
		SetupAPICallerFunc: testutil.SingletonAPICaller,
		Setup: func(ctx context.Context, caller sacloud.APICaller) error {
			s, d, err := setupServerAndDisk(ctx, caller, zone, name)
			if err != nil {
				return err
			}

			server = s
			disk = d
			return nil
		},
		Tests: []testutil.ResourceTestFunc{
			// connect
			func(ctx context.Context, caller sacloud.APICaller) error {
				return New(caller).ConnectToServer(&ConnectToServerRequest{
					Zone:     zone,
					ID:       disk.ID,
					ServerID: server.ID,
				})
			},
			// check
			func(ctx context.Context, caller sacloud.APICaller) error {
				read, err := sacloud.NewDiskOp(caller).Read(ctx, zone, disk.ID)
				if err != nil {
					return err
				}
				if read.ServerID.IsEmpty() {
					return errors.New("disk is not connected to the server")
				}
				if read.ServerID != server.ID {
					return errors.New("disk is connected to invalid server")
				}
				return nil
			},
			// disconnect
			func(ctx context.Context, caller sacloud.APICaller) error {
				return New(caller).DisconnectFromServer(&DisconnectFromServerRequest{
					Zone: zone,
					ID:   disk.ID,
				})
			},
			// check
			func(ctx context.Context, caller sacloud.APICaller) error {
				read, err := sacloud.NewDiskOp(caller).Read(ctx, zone, disk.ID)
				if err != nil {
					return err
				}
				if !read.ServerID.IsEmpty() {
					return errors.New("disk is still connected to the server")
				}
				return nil
			},
		},
		Cleanup: testutil.ComposeCleanupResourceFunc(prefix,
			testutil.CleanupTargets.Server,
			testutil.CleanupTargets.Disk,
		),
		Parallel: true,
	})
}

func setupServerAndDisk(ctx context.Context, caller sacloud.APICaller, zone string, name string) (*sacloud.Server, *sacloud.Disk, error) {
	server, err := setupServer(ctx, caller, zone, name)
	if err != nil {
		return nil, nil, err
	}
	disk, err := setupDisk(ctx, caller, zone, name)
	if err != nil {
		return nil, nil, err
	}
	return server, disk, nil
}

func setupServer(ctx context.Context, caller sacloud.APICaller, zone string, name string) (*sacloud.Server, error) {
	return sacloud.NewServerOp(caller).Create(ctx, zone, &sacloud.ServerCreateRequest{
		CPU:                  1,
		MemoryMB:             1 * size.GiB,
		ServerPlanCommitment: types.Commitments.Standard,
		ServerPlanGeneration: types.PlanGenerations.Default,
		Name:                 name,
	})
}

func setupDisk(ctx context.Context, caller sacloud.APICaller, zone string, name string) (*sacloud.Disk, error) {
	diskOp := sacloud.NewDiskOp(caller)
	disk, err := diskOp.Create(ctx, zone, &sacloud.DiskCreateRequest{
		DiskPlanID: types.DiskPlans.SSD,
		Connection: types.DiskConnections.VirtIO,
		SizeMB:     20 * size.GiB,
		Name:       name,
	}, nil)
	if err != nil {
		return nil, err
	}
	return wait.UntilDiskIsReady(ctx, diskOp, zone, disk.ID)
}
