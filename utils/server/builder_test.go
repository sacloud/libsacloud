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

package server

import (
	"bytes"
	"context"
	"errors"
	"testing"
	"time"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/testutil"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
	"github.com/sacloud/libsacloud/v2/utils/server/ostype"
	"golang.org/x/crypto/ssh"

	"github.com/stretchr/testify/require"
)

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
		msg    string
		in     *Builder
		client *BuildersAPIClient
		err    error
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
			},
			client: &BuildersAPIClient{
				ServerPlan: &dummyPlanFinder{},
			},
			err: errors.New("NIC is required when AdditionalNICs is specified"),
		},
		{
			msg: "Additional NICs over 4",
			in: &Builder{
				NIC: &SharedNICSetting{},
				AdditionalNICs: []AdditionalNICSettingHolder{
					&DisconnectedNICSetting{},
					&DisconnectedNICSetting{},
					&DisconnectedNICSetting{},
					&DisconnectedNICSetting{},
				},
			},
			client: &BuildersAPIClient{
				ServerPlan: &dummyPlanFinder{},
			},
			err: errors.New("AdditionalNICs must be less than 4"),
		},
		{
			msg: "invalid InterfaceDriver",
			in: &Builder{
				NIC:             &SharedNICSetting{},
				InterfaceDriver: types.EInterfaceDriver("invalid"),
			},
			client: &BuildersAPIClient{
				ServerPlan: &dummyPlanFinder{},
			},
			err: errors.New("invalid InterfaceDriver: invalid"),
		},
		{
			msg: "finding plan returns unexpected error",
			in:  &Builder{},
			client: &BuildersAPIClient{
				ServerPlan: &dummyPlanFinder{
					err: errors.New("dummy"),
				},
			},
			err: errors.New("dummy"),
		},
		{
			msg: "plan not found",
			in: &Builder{
				CPU:      1000,
				MemoryGB: 1024,
			},
			client: &BuildersAPIClient{
				ServerPlan: &dummyPlanFinder{},
			},
			err: errors.New("server plan not found"),
		},
	}

	for _, tc := range cases {
		err := tc.in.Validate(context.Background(), tc.client, "tk1v")
		require.Equal(t, tc.err, err, tc.msg)
	}
}

func TestBuilder_Build(t *testing.T) {
	cases := []struct {
		msg    string
		in     *Builder
		client *BuildersAPIClient
		out    *BuildResult
		err    error
	}{
		{
			msg: "Validate func is called",
			in:  &Builder{},
			out: nil,
			err: errors.New("client is empty"),
		},
		{
			msg: "finding server plan API returns error",
			in:  &Builder{},
			client: &BuildersAPIClient{
				ServerPlan: &dummyPlanFinder{
					err: errors.New("dummy"),
				},
			},
			out: nil,
			err: errors.New("dummy"),
		},
		{
			msg: "creating server returns error",
			in:  &Builder{},
			client: &BuildersAPIClient{
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
			out: nil,
			err: errors.New("dummy"),
		},
		{
			msg: "building disk returns error",
			in: &Builder{
				DiskBuilders: []DiskBuilder{
					&dummyDiskBuilder{
						err: errors.New("dummy"),
					},
				},
			},
			client: &BuildersAPIClient{
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
			out: nil,
			err: errors.New("dummy"),
		},
		{
			msg: "updating NIC returns error",
			in: &Builder{
				NIC: &SharedNICSetting{
					PacketFilterID: 2,
				},
			},
			client: &BuildersAPIClient{
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
			out: nil,
			err: errors.New("dummy"),
		},
		{
			msg: "inserting CD-ROM returns error",
			in: &Builder{
				CDROMID: 1,
			},
			client: &BuildersAPIClient{
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
			out: nil,
			err: errors.New("dummy"),
		},
		{
			msg: "booting server returns error",
			in: &Builder{
				BootAfterCreate: true,
			},
			client: &BuildersAPIClient{
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
			out: nil,
			err: errors.New("dummy"),
		},
	}
	for _, tc := range cases {
		res, err := tc.in.Build(context.Background(), tc.client, "tk1v")
		require.Equal(t, tc.err, err, tc.msg)
		require.Equal(t, tc.out, res, tc.msg)
	}
}

type dummyDiskBuilder struct {
	result *BuildDiskResult
	err    error
}

func (d *dummyDiskBuilder) Validate(ctx context.Context, client *BuildersAPIClient, zone string) error {
	return d.err
}

func (d *dummyDiskBuilder) BuildDisk(ctx context.Context, client *BuildersAPIClient, zone string, serverID types.ID) (*BuildDiskResult, error) {
	if d.err != nil {
		return nil, d.err
	}
	return d.result, nil
}

func TestBuilder_Build_BlackBox(t *testing.T) {
	var switchID types.ID
	var diskIDs []types.ID
	var buildResult *BuildResult
	var testZone = testutil.TestZone()

	testutil.Run(t, &testutil.CRUDTestCase{
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
			return nil
		},

		Create: &testutil.CRUDTestFunc{
			Func: func(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
				client := NewBuildersAPIClient(caller)
				return getBalckBoxTestBuilder(switchID).Build(ctx, client, testZone)
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

		Shutdown: func(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) error {
			serverOp := sacloud.NewServerOp(caller)
			return serverOp.Shutdown(ctx, testZone, ctx.ID, &sacloud.ShutdownOption{Force: true})
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

func getBalckBoxTestBuilder(switchID types.ID) *Builder {
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
		DiskBuilders: []DiskBuilder{
			&FromUnixDiskBuilder{
				OSType:      ostype.CentOS,
				Name:        "libsacloud-disk-builder",
				SizeGB:      20,
				PlanID:      types.DiskPlans.SSD,
				Connection:  types.DiskConnections.VirtIO,
				Description: "libsacloud-disk-builder-description",
				Tags:        types.Tags{"tag1", "tag2"},
				EditParameter: &UnixDiskEditRequest{
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
					Notes: []string{
						`libsacloud-startup-script-for-builder`,
					},
					//NoteIDs          []types.ID
				},
			},
		},
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
