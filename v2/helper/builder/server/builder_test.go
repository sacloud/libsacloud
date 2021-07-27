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

package server

import (
	"bytes"
	"context"
	"errors"
	"testing"
	"time"

	"github.com/sacloud/libsacloud/v2/helper/api"
	"github.com/sacloud/libsacloud/v2/helper/builder"
	"github.com/sacloud/libsacloud/v2/helper/builder/disk"
	"github.com/sacloud/libsacloud/v2/helper/plans"
	"github.com/sacloud/libsacloud/v2/helper/power"
	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/ostype"
	"github.com/sacloud/libsacloud/v2/sacloud/testutil"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/ssh"
)

func init() {
	if !testutil.IsAccTest() {
		api.SetupFakeDefaults()
	}
}

func TestBuilder_setDefaults(t *testing.T) {
	in := &Builder{}
	in.setDefaults()

	expected := &Builder{
		CPU:             defaultCPU,
		MemoryGB:        defaultMemoryGB,
		Commitment:      defaultCommitment,
		Generation:      defaultGeneration,
		InterfaceDriver: defaultInterfaceDriver,
	}
	require.Equal(t, expected, in)
}

func TestBuilder_Validate(t *testing.T) {
	cases := []struct {
		msg string
		in  *Builder
		err error
	}{
		{
			msg: "Client is not set",
			in:  &Builder{},
			err: errors.New("client is empty"),
		},
		{
			msg: "invalid NICs",
			in: &Builder{
				NIC: nil,
				AdditionalNICs: []AdditionalNICSettingHolder{
					&DisconnectedNICSetting{},
				},
				Client: &APIClient{
					ServerPlan: &dummyPlanFinder{},
				},
			},
			err: errors.New("NIC is required when AdditionalNICs is specified"),
		},
		{
			msg: "Additional NICs over 9",
			in: &Builder{
				NIC: &SharedNICSetting{},
				AdditionalNICs: []AdditionalNICSettingHolder{
					&DisconnectedNICSetting{},
					&DisconnectedNICSetting{},
					&DisconnectedNICSetting{},
					&DisconnectedNICSetting{},
					&DisconnectedNICSetting{},
					&DisconnectedNICSetting{},
					&DisconnectedNICSetting{},
					&DisconnectedNICSetting{},
					&DisconnectedNICSetting{},
					&DisconnectedNICSetting{},
				},
				Client: &APIClient{
					ServerPlan: &dummyPlanFinder{},
				},
			},
			err: errors.New("AdditionalNICs must be less than 9"),
		},
		{
			msg: "invalid InterfaceDriver",
			in: &Builder{
				NIC:             &SharedNICSetting{},
				InterfaceDriver: types.EInterfaceDriver("invalid"),
				Client: &APIClient{
					ServerPlan: &dummyPlanFinder{},
				},
			},
			err: errors.New("invalid InterfaceDriver: invalid"),
		},
		{
			msg: "finding plan returns unexpected error",
			in: &Builder{
				Client: &APIClient{
					ServerPlan: &dummyPlanFinder{
						err: errors.New("dummy"),
					},
				},
			},
			err: errors.New("dummy"),
		},
		{
			msg: "eth0: switch not found",
			in: &Builder{
				NIC: &ConnectedNICSetting{
					SwitchID: 1111111,
				},
				Client: &APIClient{
					Switch: &dummySwitchReader{
						err: errors.New("not found"),
					},
				},
			},
			err: errors.New("invalid NIC: reading switch info(id:1111111) is failed: not found"),
		},
		{
			msg: "eth1: switch not found",
			in: &Builder{
				NIC: &SharedNICSetting{},
				AdditionalNICs: []AdditionalNICSettingHolder{
					&ConnectedNICSetting{
						SwitchID: 1111111,
					},
				},
				Client: &APIClient{
					Switch: &dummySwitchReader{
						err: errors.New("not found"),
					},
				},
			},
			err: errors.New("invalid AdditionalNICs[0]: reading switch info(id:1111111) is failed: not found"),
		},
		{
			msg: "plan not found",
			in: &Builder{
				CPU:      1000,
				MemoryGB: 1024,
				Client: &APIClient{
					ServerPlan: &dummyPlanFinder{},
				},
			},
			err: errors.New("server plan not found"),
		},
	}

	for _, tc := range cases {
		err := tc.in.Validate(context.Background(), "tk1v")
		require.Equal(t, tc.err, err, tc.msg)
	}
}

func TestBuilder_Build(t *testing.T) {
	cases := []struct {
		msg string
		in  *Builder
		out *BuildResult
		err error
	}{
		{
			msg: "Validate func is called",
			in:  &Builder{},
			out: nil,
			err: errors.New("client is empty"),
		},
		{
			msg: "finding server plan API returns error",
			in: &Builder{
				Client: &APIClient{
					Switch:       &dummySwitchReader{},
					PacketFilter: &dummyPackerFilterReader{},
					ServerPlan: &dummyPlanFinder{
						err: errors.New("dummy"),
					},
				},
			},
			out: nil,
			err: errors.New("dummy"),
		},
		{
			msg: "creating server returns error",
			in: &Builder{
				Client: &APIClient{
					Switch:       &dummySwitchReader{},
					PacketFilter: &dummyPackerFilterReader{},
					ServerPlan: &dummyPlanFinder{
						plans: []*sacloud.ServerPlan{
							{
								ID: 1,
							},
						},
					},
					Server: &dummyCreateServerHandler{
						err: errors.New("dummy"),
					},
				},
			},
			out: nil,
			err: errors.New("dummy"),
		},
		{
			msg: "validating disk returns error",
			in: &Builder{
				DiskBuilders: []disk.Builder{
					&dummyDiskBuilder{
						err: errors.New("dummy"),
					},
				},
				Client: &APIClient{
					Switch:       &dummySwitchReader{},
					PacketFilter: &dummyPackerFilterReader{},
					ServerPlan: &dummyPlanFinder{
						plans: []*sacloud.ServerPlan{
							{
								ID: 1,
							},
						},
					},
					Server: &dummyCreateServerHandler{
						server: &sacloud.Server{ID: 1},
					},
				},
			},
			out: nil,
			err: errors.New("dummy"),
		},
		{
			msg: "updating NIC returns error",
			in: &Builder{
				NIC: &SharedNICSetting{
					PacketFilterID: 2,
				},
				Client: &APIClient{
					Switch:       &dummySwitchReader{},
					PacketFilter: &dummyPackerFilterReader{},
					ServerPlan: &dummyPlanFinder{
						plans: []*sacloud.ServerPlan{
							{
								ID: 1,
							},
						},
					},
					Server: &dummyCreateServerHandler{
						server: &sacloud.Server{
							ID: 1,
							Interfaces: []*sacloud.InterfaceView{
								{ID: 1},
							},
						},
					},
					Interface: &dummyInterfaceHandler{
						err: errors.New("dummy"),
					},
				},
			},
			out: &BuildResult{ServerID: 1},
			err: errors.New("dummy"),
		},
		{
			msg: "inserting CD-ROM returns error",
			in: &Builder{
				CDROMID: 1,
				Client: &APIClient{
					Switch:       &dummySwitchReader{},
					PacketFilter: &dummyPackerFilterReader{},
					ServerPlan: &dummyPlanFinder{
						plans: []*sacloud.ServerPlan{
							{
								ID: 1,
							},
						},
					},
					Server: &dummyCreateServerHandler{
						server:   &sacloud.Server{ID: 1},
						cdromErr: errors.New("dummy"),
					},
				},
			},
			out: &BuildResult{ServerID: 1},
			err: errors.New("dummy"),
		},
		{
			msg: "booting server returns error",
			in: &Builder{
				BootAfterCreate: true,
				Client: &APIClient{
					Switch:       &dummySwitchReader{},
					PacketFilter: &dummyPackerFilterReader{},
					ServerPlan: &dummyPlanFinder{
						plans: []*sacloud.ServerPlan{
							{
								ID: 1,
							},
						},
					},
					Server: &dummyCreateServerHandler{
						server:  &sacloud.Server{ID: 1},
						bootErr: errors.New("dummy"),
					},
				},
			},
			out: &BuildResult{ServerID: 1},
			err: errors.New("dummy"),
		},
	}
	for _, tc := range cases {
		res, err := tc.in.Build(context.Background(), "tk1v")
		require.Equal(t, tc.err, err, tc.msg)
		require.Equal(t, tc.out, res, tc.msg)
	}
}

type dummyDiskBuilder struct {
	result       *disk.BuildResult
	updateResult *disk.UpdateResult
	diskID       types.ID
	updateLevel  builder.UpdateLevel
	noWait       bool
	err          error
}

func (d *dummyDiskBuilder) Validate(ctx context.Context, zone string) error {
	return d.err
}

func (d *dummyDiskBuilder) Build(ctx context.Context, zone string, serverID types.ID) (*disk.BuildResult, error) {
	if d.err != nil {
		return nil, d.err
	}
	return d.result, nil
}

// Update ディスクの更新
func (d *dummyDiskBuilder) Update(ctx context.Context, zone string) (*disk.UpdateResult, error) {
	if d.err != nil {
		return nil, d.err
	}
	return d.updateResult, nil
}

func (d *dummyDiskBuilder) DiskID() types.ID {
	return d.diskID
}

func (d *dummyDiskBuilder) UpdateLevel(ctx context.Context, zone string, disk *sacloud.Disk) builder.UpdateLevel {
	return d.updateLevel
}

func (d *dummyDiskBuilder) NoWaitFlag() bool {
	return d.noWait
}

func TestBuilder_Build_BlackBox(t *testing.T) {
	var switchID types.ID
	var diskIDs []types.ID
	var blackboxBuilder *Builder
	var buildResult *BuildResult
	var testZone = testutil.TestZone()

	testutil.RunCRUD(t, &testutil.CRUDTestCase{
		SetupAPICallerFunc: func() sacloud.APICaller {
			return testutil.SingletonAPICaller()
		},
		Parallel:          true,
		IgnoreStartupWait: true,

		Setup: func(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) error {
			switchOp := sacloud.NewSwitchOp(caller)
			sw, err := switchOp.Create(ctx, testZone,
				&sacloud.SwitchCreateRequest{
					Name: "libsacloud-switch-for-builder",
				},
			)
			if err != nil {
				return err
			}
			switchID = sw.ID
			blackboxBuilder = getBlackBoxTestBuilder(switchID)
			return nil
		},

		Create: &testutil.CRUDTestFunc{
			Func: func(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
				return blackboxBuilder.Build(ctx, testZone)
			},
			SkipExtractID: true,
			CheckFunc: func(t testutil.TestT, ctx *testutil.CRUDTestContext, v interface{}) error {
				result := v.(*BuildResult)
				err := testutil.DoAsserts(
					testutil.AssertNotEmptyFunc(t, result.ServerID, "BuildResult.ServerID"),
					testutil.AssertNotEmptyFunc(t, result.GeneratedSSHPrivateKey, "BuildResult.GeneratedSSHPrivateKey"),
				)
				if err != nil {
					return err
				}
				buildResult = result
				ctx.ID = result.ServerID
				return nil
			},
		},

		Read: &testutil.CRUDTestFunc{
			Func: func(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
				serverOp := sacloud.NewServerOp(caller)
				server, err := serverOp.Read(ctx, testZone, ctx.ID)
				if err != nil {
					return nil, err
				}
				diskIDs = []types.ID{}
				for _, disk := range server.Disks {
					diskIDs = append(diskIDs, disk.ID)
				}
				return server, nil
			},
			CheckFunc: func(t testutil.TestT, ctx *testutil.CRUDTestContext, i interface{}) error {
				if testutil.IsAccTest() && testZone != "tk1v" { // サンドボックス以外
					time.Sleep(30 * time.Second) // sshd起動まで少し待つ
					server := i.(*sacloud.Server)
					ip := server.Interfaces[0].IPAddress
					return connectToServerViaSSH(t, "root", ip, []byte(buildResult.GeneratedSSHPrivateKey), []byte("libsacloud-test-passphrase"))
				}
				return nil
			},
			SkipExtractID: true,
		},
		Updates: []*testutil.CRUDTestFunc{
			{
				Func: func(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
					blackboxBuilder.AdditionalNICs = []AdditionalNICSettingHolder{blackboxBuilder.AdditionalNICs[1]}
					blackboxBuilder.DiskBuilders = append(blackboxBuilder.DiskBuilders, &disk.BlankBuilder{
						Name:        "libsacloud-disk-builder",
						SizeGB:      20,
						PlanID:      types.DiskPlans.SSD,
						Connection:  types.DiskConnections.VirtIO,
						Description: "libsacloud-disk-builder-description",
						Tags:        types.Tags{"tag1", "tag2"},
						Client:      disk.NewBuildersAPIClient(testutil.SingletonAPICaller()),
					})
					return blackboxBuilder.Update(ctx, testZone)
				},
				SkipExtractID: true,
				CheckFunc: func(t testutil.TestT, ctx *testutil.CRUDTestContext, v interface{}) error {
					result := v.(*BuildResult)
					buildResult = result
					ctx.ID = result.ServerID
					return nil
				},
			},
		},
		Shutdown: func(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) error {
			return power.ShutdownServer(ctx, sacloud.NewServerOp(caller), testZone, ctx.ID, true)
		},
		Delete: &testutil.CRUDTestDeleteFunc{
			Func: func(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) error {
				serverOp := sacloud.NewServerOp(caller)
				if err := serverOp.DeleteWithDisks(ctx, testZone, ctx.ID, &sacloud.ServerDeleteWithDisksRequest{IDs: diskIDs}); err != nil {
					return err
				}

				switchOp := sacloud.NewSwitchOp(caller)
				return switchOp.Delete(ctx, testZone, switchID)
			},
		},
	})
}

func getBlackBoxTestBuilder(switchID types.ID) *Builder {
	return &Builder{
		Name:            "libsacloud-server-builder",
		CPU:             1,
		MemoryGB:        1,
		Description:     "libsacloud-server-builder-description",
		Tags:            types.Tags{"tag1", "tag2"},
		BootAfterCreate: true,
		NIC:             &SharedNICSetting{},
		AdditionalNICs: []AdditionalNICSettingHolder{
			&DisconnectedNICSetting{},
			&ConnectedNICSetting{SwitchID: switchID},
		},
		DiskBuilders: []disk.Builder{
			&disk.FromUnixBuilder{
				OSType:      ostype.CentOS,
				Name:        "libsacloud-disk-builder",
				SizeGB:      20,
				PlanID:      types.DiskPlans.SSD,
				Connection:  types.DiskConnections.VirtIO,
				Description: "libsacloud-disk-builder-description",
				Tags:        types.Tags{"tag1", "tag2"},
				EditParameter: &disk.UnixEditRequest{
					HostName:                  "libsacloud-disk-builder",
					Password:                  "libsacloud-test-password",
					DisablePWAuth:             true,
					EnableDHCP:                false,
					ChangePartitionUUID:       true,
					IsSSHKeysEphemeral:        true,
					GenerateSSHKeyName:        "libsacloud-sshkey-generated",
					GenerateSSHKeyDescription: "libsacloud-sshkey-generated-for-builder",
					GenerateSSHKeyPassPhrase:  "libsacloud-test-passphrase",
					//IPAddress      string
					//NetworkMaskLen int
					//DefaultRoute   string
					//SSHKeys   []string
					//SSHKeyIDs []types.ID
					IsNotesEphemeral: true,
					NoteContents: []string{
						`libsacloud-startup-script-for-builder`,
					},
					//Notes          []*sacloud.DiskEditNote{},
				},
				Client: disk.NewBuildersAPIClient(testutil.SingletonAPICaller()),
			},
		},
		Client: NewBuildersAPIClient(testutil.SingletonAPICaller()),
	}
}

func connectToServerViaSSH(t testutil.TestT, user, ip string, privateKey []byte, passPhrase []byte) error {
	signer, err := ssh.ParsePrivateKeyWithPassphrase(privateKey, passPhrase)
	if err != nil {
		return err
	}

	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         10 * time.Second,
	}

	client, err := ssh.Dial("tcp", ip+":22", config)
	if err != nil {
		return err
	}
	defer client.Close()

	session, err := client.NewSession()
	if err != nil {
		return err
	}
	defer session.Close()

	var b bytes.Buffer
	session.Stdout = &b
	if err := session.Run("/usr/bin/whoami"); err != nil {
		return err
	}
	t.Logf("Connect to the Server via SSH: `whoami`: %s\n", b.String())
	return nil
}

func TestBuilder_IsNeedShutdown(t *testing.T) {
	cases := []struct {
		msg    string
		in     *Builder
		expect bool
		err    error
	}{
		{
			msg:    "server id is empty",
			expect: false,
			err:    errors.New("server id required"),
			in:     &Builder{},
		},
		{
			msg:    "in-place update",
			expect: false,
			err:    nil,
			in: &Builder{
				ServerID:    types.ID(1),
				Name:        "update",
				Description: "update",
				Tags:        types.Tags{"update1", "update2"},
				Client: &APIClient{
					Server: &dummyCreateServerHandler{
						server: &sacloud.Server{
							ID:          types.ID(1),
							Name:        "update",
							Description: "update",
							Tags:        types.Tags{"update1", "update2-upd"},
						},
					},
				},
			},
		},
		{
			msg:    "changed: PrivateHostID",
			expect: true,
			err:    nil,
			in: &Builder{
				ServerID:      types.ID(1),
				PrivateHostID: types.ID(2),
				Client: &APIClient{
					Server: &dummyCreateServerHandler{
						server: &sacloud.Server{
							ID:            types.ID(1),
							PrivateHostID: types.ID(3),
						},
					},
				},
			},
		},
		{
			msg:    "changed: InterfaceDriver",
			expect: true,
			err:    nil,
			in: &Builder{
				ServerID:        types.ID(1),
				InterfaceDriver: types.InterfaceDrivers.E1000,
				Client: &APIClient{
					Server: &dummyCreateServerHandler{
						server: &sacloud.Server{
							ID:              types.ID(1),
							InterfaceDriver: types.InterfaceDrivers.VirtIO,
						},
					},
				},
			},
		},
		{
			msg:    "changed: Memory Size",
			expect: true,
			err:    nil,
			in: &Builder{
				ServerID: types.ID(1),
				MemoryGB: 1,
				Client: &APIClient{
					Server: &dummyCreateServerHandler{
						server: &sacloud.Server{
							ID:       types.ID(1),
							MemoryMB: 2,
						},
					},
				},
			},
		},
		{
			msg:    "changed: CPU",
			expect: true,
			err:    nil,
			in: &Builder{
				ServerID: types.ID(1),
				CPU:      1,
				Client: &APIClient{
					Server: &dummyCreateServerHandler{
						server: &sacloud.Server{
							ID:  types.ID(1),
							CPU: 2,
						},
					},
				},
			},
		},
		{
			msg:    "changed: Commitment",
			expect: true,
			err:    nil,
			in: &Builder{
				ServerID: types.ID(1),
				CPU:      1,
				Client: &APIClient{
					Server: &dummyCreateServerHandler{
						server: &sacloud.Server{
							ID:                   types.ID(1),
							CPU:                  1,
							ServerPlanCommitment: types.Commitments.DedicatedCPU,
						},
					},
				},
			},
		},
		{
			msg:    "changed: add NIC",
			expect: true,
			err:    nil,
			in: &Builder{
				ServerID: types.ID(1),
				NIC:      &SharedNICSetting{},
				Client: &APIClient{
					Server: &dummyCreateServerHandler{
						server: &sacloud.Server{
							ID: types.ID(1),
						},
					},
				},
			},
		},
		{
			msg:    "changed: delete NIC",
			expect: true,
			err:    nil,
			in: &Builder{
				ServerID: types.ID(1),
				Client: &APIClient{
					Server: &dummyCreateServerHandler{
						server: &sacloud.Server{
							ID: types.ID(1),
							Interfaces: []*sacloud.InterfaceView{
								{
									ID:             types.ID(2),
									SwitchScope:    types.Scopes.Shared,
									PacketFilterID: 0,
									UpstreamType:   types.UpstreamNetworkTypes.Shared,
								},
							},
						},
					},
				},
			},
		},
		{
			msg:    "changed: Packet Filter ID",
			expect: false,
			err:    nil,
			in: &Builder{
				ServerID: types.ID(1),
				NIC: &SharedNICSetting{
					PacketFilterID: types.ID(10),
				},
				Client: &APIClient{
					Server: &dummyCreateServerHandler{
						server: &sacloud.Server{
							ID: types.ID(1),
							Interfaces: []*sacloud.InterfaceView{
								{
									ID:             types.ID(2),
									SwitchScope:    types.Scopes.Shared,
									PacketFilterID: types.ID(11),
									UpstreamType:   types.UpstreamNetworkTypes.Shared,
								},
							},
						},
					},
				},
			},
		},
	}

	for _, tc := range cases {
		got, err := tc.in.IsNeedShutdown(context.Background(), testutil.TestZone())
		require.Equal(t, tc.err, err, tc.msg)
		require.Equal(t, tc.expect, got, tc.msg)
	}
}

func TestBuilder_UpdateWithPreviousID(t *testing.T) {
	ctx := context.Background()
	builder := &Builder{
		Name:            testutil.ResourceName("server-builder"),
		CPU:             1,
		MemoryGB:        1,
		Commitment:      types.Commitments.Standard,
		Generation:      types.PlanGenerations.Default,
		Tags:            types.Tags{"tag1", "tag2"},
		BootAfterCreate: false,
		Client:          NewBuildersAPIClient(testutil.SingletonAPICaller()),
		ForceShutdown:   true,
	}
	createResult, err := builder.Build(ctx, testutil.TestZone())
	if err != nil {
		t.Fatal(err)
	}

	serverOp := sacloud.NewServerOp(testutil.SingletonAPICaller())
	server, err := serverOp.Read(ctx, testutil.TestZone(), createResult.ServerID)
	if err != nil {
		t.Fatal(err)
	}

	require.EqualValues(t, builder.Tags, server.Tags)

	// プラン変更
	builder.ServerID = server.ID
	builder.CPU = 2
	builder.MemoryGB = 4

	updateResult, err := builder.Update(ctx, testutil.TestZone())
	if err != nil {
		t.Fatal(err)
	}

	// IDが変更されているはず
	require.True(t, createResult.ServerID != updateResult.ServerID)
	updated, err := serverOp.Read(ctx, testutil.TestZone(), updateResult.ServerID)
	if err != nil {
		t.Fatal(err)
	}

	require.EqualValues(t, plans.AppendPreviousIDTagIfAbsent(server.Tags, server.ID), updated.Tags)

	// cleanup
	if err := serverOp.Delete(ctx, testutil.TestZone(), updated.ID); err != nil {
		t.Fatal(err)
	}
}
