package server

import (
	"context"
	"errors"
	"testing"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/testutil"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
	"github.com/sacloud/libsacloud/v2/utils/server/ostype"

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
		msg string
		in  *Builder
		err error
	}{
		{
			msg: "Client is not set",
			in:  &Builder{},
			err: errors.New("field 'Client' is not set"),
		},
		{
			msg: "invalid NICs",
			in: &Builder{
				Client: &BuildersAPIClient{
					ServerPlan: &dummyPlanFinder{},
				},
				NIC: nil,
				AdditionalNICs: []AdditionalNICSettingHolder{
					&DisconnectedNICRequest{},
				},
			},
			err: errors.New("NIC is required when AdditionalNICs is specified"),
		},
		{
			msg: "Additional NICs over 4",
			in: &Builder{
				Client: &BuildersAPIClient{
					ServerPlan: &dummyPlanFinder{},
				},
				NIC: &SharedNICRequest{},
				AdditionalNICs: []AdditionalNICSettingHolder{
					&DisconnectedNICRequest{},
					&DisconnectedNICRequest{},
					&DisconnectedNICRequest{},
					&DisconnectedNICRequest{},
				},
			},
			err: errors.New("AdditionalNICs must be less than 4"),
		},
		{
			msg: "invalid InterfaceDriver",
			in: &Builder{
				Client: &BuildersAPIClient{
					ServerPlan: &dummyPlanFinder{},
				},
				NIC:             &SharedNICRequest{},
				InterfaceDriver: types.EInterfaceDriver("invalid"),
			},
			err: errors.New("invalid InterfaceDriver: invalid"),
		},
		{
			msg: "finding plan returns unexpected error",
			in: &Builder{
				Client: &BuildersAPIClient{
					ServerPlan: &dummyPlanFinder{
						err: errors.New("dummy"),
					},
				},
			},
			err: errors.New("dummy"),
		},
		{
			msg: "plan not found",
			in: &Builder{
				CPU:      1000,
				MemoryGB: 1024,
				Client: &BuildersAPIClient{
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
			err: errors.New("field 'Client' is not set"),
		},
		// TODO we should add cases
	}

	for _, tc := range cases {
		res, err := tc.in.Build(context.Background(), "tk1v")
		require.Equal(t, tc.err, err, tc.msg)
		require.Equal(t, tc.out, res, tc.msg)
	}

}

func TestAccServerBuild(t *testing.T) {
	if !testutil.IsAccTest() {
		t.Skip("TestAccServerBuild only exec at Acceptance Test")
	}

	client := NewBuildersAPIClient(testutil.SingletonAPICaller())
	ctx := context.Background()
	testZone := testutil.TestZone()

	// prepare switch
	switchOp := sacloud.NewSwitchOp(testutil.SingletonAPICaller())
	sw, err := switchOp.Create(ctx, testZone,
		&sacloud.SwitchCreateRequest{
			Name: "libsacloud-switch-for-builder",
		},
	)
	require.NoError(t, err)

	builder := &Builder{
		Client:          client,
		Name:            "libsacloud-server-builder",
		CPU:             1,
		MemoryGB:        1,
		Description:     "libsacloud-server-builder-description",
		Tags:            types.Tags{"tag1", "tag2"},
		BootAfterCreate: true,
		NIC:             &SharedNICRequest{},
		AdditionalNICs: []AdditionalNICSettingHolder{
			&DisconnectedNICRequest{},
			&ConnectedNICRequest{SwitchID: sw.ID},
		},
		DiskBuilders: []DiskBuilder{
			&DiskFromUnixRequest{
				OSType:      ostype.CentOS,
				Name:        "libsacloud-disk-builder",
				SizeGB:      20,
				PlanID:      types.DiskPlans.SSD,
				Connection:  types.DiskConnections.VirtIO,
				Description: "libsacloud-disk-builder-description",
				Tags:        types.Tags{"tag1", "tag2"},
				EditParameter: &UnixDiskEditRequest{
					HostName:            "libsacloud-disk-builder",
					Password:            "libsacloud-test-password",
					DisablePWAuth:       true,
					EnableDHCP:          false,
					ChangePartitionUUID: true,
					IsSSHKeysEphemeral:  true,
					GenerateSSHKeyName:  "libsacloud-sshkey-generated",
					//IPAddress      string
					//NetworkMaskLen int
					//DefaultRoute   string
					//SSHKeys   []string
					//SSHKeyIDs []types.ID
					//IsNotesEphemeral bool
					//Notes            []string
					//NoteIDs          []types.ID
				},
			},
		},
	}

	result, err := builder.Build(ctx, testZone)
	require.NoError(t, err)
	t.Logf("ServerID: %d, private-key: %s", result.ServerID, result.GeneratedSSHPrivateKey)

	//serverOp := sacloud.NewServerOp(testutil.SingletonAPICaller())
	//serverOp.Delete(ctx, testZone, result.ServerID)
	//switchOp.Delete(ctx, testZone, sw.ID)
}
