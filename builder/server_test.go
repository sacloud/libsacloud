package builder

import (
	"testing"

	"github.com/sacloud/libsacloud/sacloud"
	"github.com/sacloud/libsacloud/sacloud/ostype"
	"github.com/stretchr/testify/assert"
)

var (
	serverBuilderTestServerName = "testServerName"
	serverBuilderTestPassword   = "testPassword01"
	serverBuilderTestNote       = `#!/bin/bash
yum -y update || exit 1
exit 0
`
	serverBuilderTestSSHKey = `ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDVbJLAQHDVpgsjLhauPl1dY5o5MeC1f+sPQW1W5D9Iug+qCdUI3VpWSq5oPSe4sx4n+l3eFbEsjA6Z2pDwboBwZ142P5Ry5npiIX1Bi8xbx3Cp8KylgILf+pJtFqkRMdpFEDPxN2cmqsSR4yPyMJ8R5sBPMFRtBOkBLiRrdfLBOoh4RwpS3tiNwqkLCc2YVirLL+NTz6/1shQu8++hO0xWDjzvrl/plIAHpVG8nuPuMr9zE+MPW3m+1O0jV9iFFh8/8vTnfP1kPY/CQCht05Lh+q53XKXp0a7tdKRzJ6TKV6l5VfySKIIcuKSVJ16ysbnbYMacc2mEsH5DAXxlPESl`
)

func TestServerBuilder_buildParams(t *testing.T) {

	builder := ServerDiskless(NewAPIClient(client), serverBuilderTestServerName).(*serverBuilder)

	// この段階では ディスクレス/ISOイメージレス/NICレス、全ての設定がデフォルト値のサーバーになる
	builder.currentBuildValue = &ServerBuildValue{}
	builder.currentBuildResult = &ServerBuildResult{}
	err := builder.buildParams()
	params := builder.currentBuildValue

	assert.NoError(t, err)
	assert.NotNil(t, params)

	// この時点ではディスクレスなため、ディスク関連のプロパティは設定できない(コンパイルエラー)
	//builder.GetPassword()

	// Unix系アーカイブからのインストールの場合、パスワードの設定などのディスクの編集ができるようになる。
	tempBuilder := ServerPublicArchiveUnix(NewAPIClient(client), ostype.CentOS, serverBuilderTestServerName, serverBuilderTestPassword)

	assert.NotNil(t, tempBuilder)
	assert.Equal(t, tempBuilder.GetPassword(), serverBuilderTestPassword)

}

func TestServerBuilder_DisklessDefaults(t *testing.T) {
	defer initServers()()

	builder := ServerDiskless(NewAPIClient(client), serverBuilderTestServerName)

	assert.Equal(t, builder.GetServerName(), serverBuilderTestServerName)        // サーバー名
	assert.Equal(t, builder.GetCore(), 1)                                        // コア数 : デフォルト1
	assert.Equal(t, builder.GetMemory(), 1)                                      // メモリ : デフォルト1GB
	assert.Equal(t, builder.GetCommitment(), sacloud.ECommitmentStandard)        // コミットメント: デフォルト standard
	assert.Equal(t, builder.GetInterfaceDriver(), sacloud.InterfaceDriverVirtIO) // 準仮想化モード(@virtio-net-pci) : デフォルト有効

}

func TestServerBuilder_DisklessDedicatedCPU(t *testing.T) {
	defer initServers()()

	builder := ServerDiskless(NewAPIClient(client), serverBuilderTestServerName)

	builder.SetCore(2)
	builder.SetMemory(4)
	builder.SetCommitment("dedicatedcpu")

	assert.Equal(t, builder.GetServerName(), serverBuilderTestServerName)        // サーバー名
	assert.Equal(t, builder.GetCore(), 2)                                        // コア数 : デフォルト1
	assert.Equal(t, builder.GetMemory(), 4)                                      // メモリ : デフォルト1GB
	assert.Equal(t, builder.GetCommitment(), sacloud.ECommitmentDedicatedCPU)    // コミットメント: デフォルト standard
	assert.Equal(t, builder.GetInterfaceDriver(), sacloud.InterfaceDriverVirtIO) // 準仮想化モード(@virtio-net-pci) : デフォルト有効

}

func TestDisklessServerBuilder_ServerPublicArchiveUnixDefaults(t *testing.T) {
	defer initServers()()

	b := ServerPublicArchiveUnix(NewAPIClient(client), ostype.CentOS, serverBuilderTestServerName, serverBuilderTestPassword)

	assert.Equal(t, b.GetServerName(), serverBuilderTestServerName)        // サーバー名
	assert.Equal(t, b.GetCore(), 1)                                        // コア数 : デフォルト1
	assert.Equal(t, b.GetMemory(), 1)                                      // メモリ : デフォルト1GB
	assert.Equal(t, b.GetInterfaceDriver(), sacloud.InterfaceDriverVirtIO) // 準仮想化モード(@virtio-net-pci) : デフォルト有効

	assert.Equal(t, b.GetDiskSize(), 20)                                 // デフォルト 20GB
	assert.Equal(t, b.GetDiskPlanID(), sacloud.DiskPlanSSDID)            // デフォルトSSD
	assert.Equal(t, b.GetDiskConnection(), sacloud.DiskConnectionVirtio) // デフォルト　virtio
	assert.Equal(t, b.IsDisablePWAuth(), false)
}

func TestServerBuilder_Build_WithMinimum(t *testing.T) {
	defer initServers()()

	builder := ServerPublicArchiveUnix(NewAPIClient(client), ostype.CentOS, serverBuilderTestServerName, serverBuilderTestPassword)
	builder.AddPublicNWConnectedNIC()
	result, err := builder.Build()

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.NotNil(t, result.Server)
	assert.NotNil(t, result.Disks[0])

	assert.True(t, result.Server.Instance.IsUp())
}

func TestServerBuilder_Build_WithPacketFilter(t *testing.T) {
	defer initServers()()

	builder := ServerDiskless(NewAPIClient(client), serverBuilderTestServerName)

	pfReq := client.PacketFilter.New()
	pfReq.Name = "Test"
	pf, err := client.PacketFilter.Create(pfReq)

	assert.NoError(t, err)

	builder.AddPublicNWConnectedNIC()
	builder.SetPacketFilterIDs([]int64{pf.ID})
	res, err := builder.Build()

	assert.NoError(t, err)
	assert.NotNil(t, res.Server)
	assert.NotNil(t, res.Server.Interfaces[0])
	assert.Equal(t, res.Server.Interfaces[0].PacketFilter.ID, pf.ID)

}

func TestServerBuilder_Build_WithSSHKeyAndNoteEphemeral(t *testing.T) {
	defer initServers()()

	builder := ServerPublicArchiveUnix(NewAPIClient(client), ostype.CentOS, serverBuilderTestServerName, serverBuilderTestPassword)

	builder.AddPublicNWConnectedNIC()
	builder.AddNote(serverBuilderTestNote)
	builder.AddSSHKey(serverBuilderTestSSHKey)
	builder.SetDisablePWAuth(true)
	res, err := builder.Build()

	assert.NoError(t, err)
	assert.NotNil(t, res.Server)
	assert.NotNil(t, res.Disks[0].Disk)
	assert.NotNil(t, res.Disks[0].SSHKeys[0])
	assert.NotNil(t, res.Disks[0].Notes[0])

	// Ephemeralに設定していると、SSH/Noteは消えているはず
	key, err := client.SSHKey.Read(res.Disks[0].SSHKeys[0].ID)
	assert.Error(t, err)
	assert.Nil(t, key)

	note, err := client.Note.Read(res.Disks[0].Notes[0].ID)
	assert.Error(t, err)
	assert.Nil(t, note)

}

func TestServerBuilder_Build_WithExistsSwitch(t *testing.T) {
	defer initServers()()

	newSw := client.Switch.New()
	newSw.Name = serverBuilderTestServerName
	sw, err := client.Switch.Create(newSw)
	assert.NoError(t, err)

	expectAddr := "19.2.0.1"
	builder := ServerDiskless(NewAPIClient(client), serverBuilderTestServerName)
	builder.AddExistsSwitchConnectedNICWithDisplayIP(sw.GetStrID(), expectAddr)
	res, err := builder.Build()
	assert.NoError(t, err)

	assert.Equal(t, expectAddr, res.Server.Interfaces[0].GetUserIPAddress())
}

func TestServerBuilder_Build_WithSSHKeyAndNote(t *testing.T) {
	defer initServers()()

	builder := ServerPublicArchiveUnix(NewAPIClient(client), ostype.CentOS, serverBuilderTestServerName, serverBuilderTestPassword)

	builder.AddPublicNWConnectedNIC()
	builder.AddNote(serverBuilderTestNote)
	builder.AddSSHKey(serverBuilderTestSSHKey)
	builder.SetDisablePWAuth(true)
	builder.SetNotesEphemeral(false)
	builder.SetSSHKeysEphemeral(false)
	builder.SetGenerateSSHKeyName("Test")
	builder.SetGenerateSSHKeyPassPhrase("12345678") // min:8,max:64
	builder.SetGenerateSSHKeyDescription("Test")
	res, err := builder.Build()

	assert.NoError(t, err)
	assert.NotNil(t, res.Server)
	assert.NotNil(t, res.Disks[0].Disk)
	assert.NotNil(t, res.Disks[0].SSHKeys[0])
	assert.NotNil(t, res.Disks[0].Notes[0])
	assert.NotNil(t, res.Disks[0].GeneratedSSHKey)
	assert.NotEmpty(t, res.Disks[0].GeneratedSSHKey.PrivateKey)

	// Ephemeral=falseに設定していると、SSH/Noteは消えていないはず
	key, err := client.SSHKey.Read(res.Disks[0].SSHKeys[0].ID)
	assert.NoError(t, err)
	assert.NotNil(t, key)

	note, err := client.Note.Read(res.Disks[0].Notes[0].ID)
	assert.NoError(t, err)
	assert.NotNil(t, note)

	generated, err := client.SSHKey.Read(res.Disks[0].GeneratedSSHKey.ID)
	assert.NoError(t, err)
	assert.NotNil(t, generated)

	client.SSHKey.Delete(res.Disks[0].SSHKeys[0].ID)
	client.Note.Delete((res.Disks[0].Notes[0].ID))

}

func TestServerBuilder_Build_WithEventHandler(t *testing.T) {
	defer initServers()()

	builder := ServerPublicArchiveUnix(NewAPIClient(client), ostype.CentOS, serverBuilderTestServerName, serverBuilderTestPassword)

	serverEvents := []ServerBuildEvents{
		ServerBuildOnStart,
		ServerBuildOnSetPlanBefore,
		ServerBuildOnSetPlanAfter,
		ServerBuildOnCreateServerBefore,
		ServerBuildOnCreateServerAfter,
		ServerBuildOnInsertCDROMBefore,
		ServerBuildOnInsertCDROMAfter,
		ServerBuildOnBootBefore,
		ServerBuildOnBootAfter,
		ServerBuildOnComplete,
	}
	serverEventCalled := map[ServerBuildEvents]bool{}
	for _, ev := range serverEvents {
		e := ev
		serverEventCalled[e] = false

		builder.SetEventHandler(e, func(value *ServerBuildValue, result *ServerBuildResult) {
			serverEventCalled[e] = true
		})
	}

	diskEvents := []DiskBuildEvents{
		DiskBuildOnStart,
		DiskBuildOnCreateSSHKeyBefore,
		DiskBuildOnCreateSSHKeyAfter,
		DiskBuildOnCreateNoteBefore,
		DiskBuildOnCreateNoteAfter,
		DiskBuildOnCreateDiskBefore,
		DiskBuildOnCreateDiskAfter,
		DiskBuildOnCleanupSSHKeyBefore,
		DiskBuildOnCleanupSSHKeyAfter,
		DiskBuildOnCleanupNoteBefore,
		DiskBuildOnCleanupNoteAfter,
		DiskBuildOnComplete,
	}
	diskEventCalled := map[DiskBuildEvents]bool{}
	for _, ev := range diskEvents {
		e := ev
		diskEventCalled[e] = false
		builder.SetDiskEventHandler(e, func(value *DiskBuildValue, result *DiskBuildResult) {
			diskEventCalled[e] = true
		})
	}

	searchISO, err := client.CDROM.Find()
	if !assert.NoError(t, err) {
		return
	}
	isoImageID := searchISO.CDROMs[0].ID

	builder.AddPublicNWConnectedNIC()
	builder.AddNote(serverBuilderTestNote)
	builder.AddSSHKey(serverBuilderTestSSHKey)
	builder.SetDisablePWAuth(true)
	builder.SetISOImageID(isoImageID)
	res, err := builder.Build()

	assert.NoError(t, err)
	assert.NotNil(t, res)

	// is callbacked?
	for ev, called := range serverEventCalled {
		assert.True(t, called, "EventHandler %s is not called", ev)
	}
	for ev, called := range diskEventCalled {
		assert.True(t, called, "EventHandler %s is not called", ev)
	}

}

func initServers() func() {
	cleanupServers()
	return cleanupServers
}

func cleanupServers() {
	res, _ := client.Server.Reset().WithNameLike(serverBuilderTestServerName).Find()
	for _, server := range res.Servers {
		deleteServer(&server)
	}
	res, _ = client.Switch.Reset().WithNameLike(serverBuilderTestServerName).Find()
	for _, sw := range res.Switches {
		client.Switch.Delete(sw.ID)
	}
}

func deleteServer(server *sacloud.Server) {
	client.Server.Stop(server.ID)
	client.Server.SleepUntilDown(server.ID, client.DefaultTimeoutDuration)
	if len(server.Disks) > 0 {
		client.Server.DeleteWithDisk(server.ID, server.GetDiskIDs())
	} else {
		client.Server.Delete(server.ID)
	}

}
