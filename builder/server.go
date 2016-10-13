package builder

import (
	"fmt"
	"github.com/yamamoto-febc/libsacloud/api"
	"github.com/yamamoto-febc/libsacloud/sacloud"
	"github.com/yamamoto-febc/libsacloud/sacloud/ostype"
)

/**********************************************************
  Type : ServerBuildEvents
**********************************************************/

// ServerBuildEvents サーバー構築時イベント種別
type ServerBuildEvents int

const (
	// ServerBuildOnStart サーバー構築 開始
	ServerBuildOnStart ServerBuildEvents = iota

	// ServerBuildOnSetPlanBefore サーバープラン設定 開始時
	ServerBuildOnSetPlanBefore

	// ServerBuildOnSetPlanAfter サーバープラン設定 終了時
	ServerBuildOnSetPlanAfter

	// ServerBuildOnCreateServerBefore サーバー作成 開始時
	ServerBuildOnCreateServerBefore

	// ServerBuildOnCreateServerAfter サーバー作成 終了時
	ServerBuildOnCreateServerAfter

	// ServerBuildOnConnectDiskBefore ディスク接続 開始時
	ServerBuildOnConnectDiskBefore

	// ServerBuildOnConnectDiskAfter ディスク接続 終了時
	ServerBuildOnConnectDiskAfter

	// ServerBuildOnInsertCDROMBefore ISOイメージ挿入 開始時
	ServerBuildOnInsertCDROMBefore

	// ServerBuildOnInsertCDROMAfter ISOイメージ挿入 終了時
	ServerBuildOnInsertCDROMAfter

	// ServerBuildOnBootBefore サーバー起動 開始時
	ServerBuildOnBootBefore

	// ServerBuildOnBootAfter サーバー起動 終了時
	ServerBuildOnBootAfter

	// ServerBuildOnComplete サーバー構築 完了
	ServerBuildOnComplete
)

// ServerBuildEventHandler サーバー構築時イベントハンドラ
type ServerBuildEventHandler func(value *ServerBuildValue, result *ServerBuildResult)

/**********************************************************
  Type : ServerBuilder
**********************************************************/

//serverBuilder サーバービルダー基底
type serverBuilder struct {
	*baseBuilder
	buildEventHandlers map[ServerBuildEvents]ServerBuildEventHandler
	// for server
	serverName      string
	core            int
	memory          int
	useVirtIONetPCI bool
	description     string
	iconID          int64
	*sacloud.TagsType
	bootAfterCreate bool

	// CDROM
	isoImageID int64

	// for nic
	nicConnections []string

	// for disks
	disk            *DiskBuilder
	additionalDisks []*DiskBuilder

	currentBuildValue  *ServerBuildValue
	currentBuildResult *ServerBuildResult
}

var (
	defaultCore            = 1
	defaultMemory          = 1
	defaultUseVirtIONetCPI = true
	defaultDescription     = ""
	defaultIconID          = int64(0)
	defaultBootAfterCreate = true
)

func newServerBuilder(client *api.Client, serverName string) *serverBuilder {
	return &serverBuilder{
		baseBuilder: &baseBuilder{
			client: client,
			errors: []error{},
		},
		buildEventHandlers: map[ServerBuildEvents]ServerBuildEventHandler{},
		serverName:         serverName,
		core:               defaultCore,
		memory:             defaultMemory,
		useVirtIONetPCI:    defaultUseVirtIONetCPI,
		TagsType:           &sacloud.TagsType{},
		description:        defaultDescription,
		iconID:             defaultIconID,
		bootAfterCreate:    defaultBootAfterCreate,
	}

}

/*---------------------------------------------------------
  for connect disk functions
---------------------------------------------------------*/

// FromDiskless ディスクレスサーバービルダー
func FromDiskless(client *api.Client, name string) *DisklessServerBuilder {
	b := newServerBuilder(client, name)
	return &DisklessServerBuilder{
		serverBuilder: b,
	}
}

// FromPublicArchiveUnix ディスクの編集が可能なLinux(Unix)系パブリックアーカイブを利用するビルダー
func FromPublicArchiveUnix(client *api.Client, os ostype.ArchiveOSTypes, name string, password string) *PublicArchiveUnixServerBuilder {

	b := newServerBuilder(client, name)
	b.fromPublicArchiveUnix(os, password)
	return &PublicArchiveUnixServerBuilder{
		serverBuilder: b,
	}

}

// FromPublicArchiveWindows Windows系パブリックアーカイブを利用するビルダー
func FromPublicArchiveWindows(client *api.Client, name string, archiveID int64) *PublicArchiveWindowsServerBuilder {

	b := newServerBuilder(client, name)
	b.fromPublicArchiveWindows(archiveID)
	return &PublicArchiveWindowsServerBuilder{
		serverBuilder: b,
	}

}

//FromBlankDisk 空のディスクを利用するビルダー
func FromBlankDisk(client *api.Client, name string) *BlankDiskServerBuilder {

	b := newServerBuilder(client, name)
	return &BlankDiskServerBuilder{
		serverBuilder: b,
	}

}

// FromDisk 既存ディスクを利用するビルダー
func FromDisk(client *api.Client, name string, sourceDiskID int64) *CommonServerBuilder {
	b := newServerBuilder(client, name)

	b.fromDisk(sourceDiskID)
	return &CommonServerBuilder{
		serverBuilder: b,
	}

}

// FromArchive 既存アーカイブを利用するビルダー
func FromArchive(client *api.Client, name string, sourceArchiveID int64) *CommonServerBuilder {
	b := newServerBuilder(client, name)

	b.fromArchive(sourceArchiveID)
	return &CommonServerBuilder{
		serverBuilder: b,
	}

}

/*---------------------------------------------------------
  Inner functions
---------------------------------------------------------*/

func (b *serverBuilder) fromPublicArchiveUnix(os ostype.ArchiveOSTypes, password string) {
	archive, err := b.client.Archive.FindByOSType(os)
	if err != nil {
		b.errors = append(b.errors, err)
	}

	b.disk = NewDiskBuilder(b.client, b.serverName)
	b.disk.sourceArchiveID = archive.ID
	b.disk.password = password

}

func (b *serverBuilder) fromPublicArchiveWindows(archiveID int64) {
	b.disk = NewDiskBuilder(b.client, b.serverName)
	b.disk.sourceArchiveID = archiveID
	b.disk.sourceDiskID = 0
}

func (b *serverBuilder) fromDisk(sourceDiskID int64) {
	b.disk = NewDiskBuilder(b.client, b.serverName)
	b.disk.sourceArchiveID = 0
	b.disk.sourceDiskID = sourceDiskID
}

func (b *serverBuilder) fromArchive(sourceArchiveID int64) {

	b.disk = NewDiskBuilder(b.client, b.serverName)
	b.disk.sourceArchiveID = sourceArchiveID
	b.disk.sourceDiskID = 0
}

// Build サーバーの構築
func (b *serverBuilder) Build() (*ServerBuildResult, error) {

	// start
	b.callEventHandlerIfExists(ServerBuildOnStart)
	b.currentBuildValue = &ServerBuildValue{}
	b.currentBuildResult = &ServerBuildResult{}

	// build parameter
	if err := b.buildParams(); err != nil {
		return b.currentBuildResult, err
	}

	// create disks
	if err := b.createDisks(); err != nil {
		return b.currentBuildResult, err
	}

	// create server
	b.callEventHandlerIfExists(ServerBuildOnCreateServerBefore)
	if err := b.createServer(); err != nil {
		return b.currentBuildResult, err
	}
	b.callEventHandlerIfExists(ServerBuildOnCreateServerAfter)

	// connect disks
	b.callEventHandlerIfExists(ServerBuildOnConnectDiskBefore)
	if err := b.connectDisks(); err != nil {
		return b.currentBuildResult, err
	}
	b.callEventHandlerIfExists(ServerBuildOnConnectDiskAfter)

	// insert cdrom
	if b.isoImageID > 0 {
		b.callEventHandlerIfExists(ServerBuildOnInsertCDROMBefore)
		if err := b.insertCDROM(); err != nil {
			return b.currentBuildResult, err
		}
		b.callEventHandlerIfExists(ServerBuildOnInsertCDROMAfter)
	}

	// boot server
	if b.bootAfterCreate {
		b.callEventHandlerIfExists(ServerBuildOnBootBefore)
		if err := b.bootServer(); err != nil {
			return b.currentBuildResult, err
		}
		b.callEventHandlerIfExists(ServerBuildOnBootAfter)
	}

	// complete
	b.callEventHandlerIfExists(ServerBuildOnComplete)
	return b.currentBuildResult, nil
}

func (b *serverBuilder) buildParams() error {

	v := b.currentBuildValue
	v.Server = b.client.Server.New()
	if err := b.buildServerParams(); err != nil {
		return err
	}
	return nil
}

func (b *serverBuilder) buildServerParams() error {

	v := b.currentBuildValue
	b.callEventHandlerIfExists(ServerBuildOnSetPlanBefore)

	// plan
	plan, err := b.client.Product.Server.GetBySpec(b.core, b.memory)
	if err != nil {
		err = fmt.Errorf("Error building server parameters : setting plan / [%s]", err)
		return err
	}

	b.callEventHandlerIfExists(ServerBuildOnSetPlanAfter)

	s := v.Server
	s.Name = b.serverName
	s.SetServerPlanByID(plan.GetStrID())
	s.Description = b.description
	// tags
	if b.useVirtIONetPCI {
		s.AppendTag(sacloud.TagVirtIONetPCI)
	}
	for _, tag := range b.Tags {
		if !s.HasTag(tag) {
			s.AppendTag(tag)
		}
	}
	if b.iconID > 0 {
		s.Icon = sacloud.NewResource(b.iconID)
	}

	// NIC
	for _, nic := range b.nicConnections {
		switch nic {
		case "shared":
			s.AddPublicNWConnectedParam()
			break
		case "":
			s.AddEmptyConnectedParam()
			break
		default:
			s.AddExistsSwitchConnectedParam(nic)
		}
	}

	return nil
}

func (b *serverBuilder) createDisks() error {
	// build disk
	diskBuildResult, err := b.disk.Build()
	if err != nil {
		return err
	}
	b.currentBuildResult.addDisk(diskBuildResult)
	// build additional disks
	if len(b.additionalDisks) > 0 {
		for _, diskBuilder := range b.additionalDisks {
			res, err := diskBuilder.Build()
			if err != nil {
				return err
			}
			b.currentBuildResult.addDisk(res)
		}
	}
	return nil
}

func (b *serverBuilder) createServer() error {
	server, err := b.client.Server.Create(b.currentBuildValue.Server)
	if err != nil {
		return err
	}
	b.currentBuildResult.Server = server
	return nil
}

func (b *serverBuilder) connectDisks() error {
	server := b.currentBuildResult.Server
	for _, disk := range b.currentBuildResult.Disks {
		_, err := b.client.Disk.ConnectToServer(disk.Disk.ID, server.ID)
		if err != nil {
			return err
		}
	}
	return nil
}

func (b *serverBuilder) insertCDROM() error {
	server := b.currentBuildResult.Server
	_, err := b.client.Server.InsertCDROM(server.ID, b.isoImageID)
	if err != nil {
		return err
	}
	return nil
}

func (b *serverBuilder) bootServer() error {
	server := b.currentBuildResult.Server
	_, err := b.client.Server.Boot(server.ID)
	if err != nil {
		return err
	}

	if err := b.client.Server.SleepUntilUp(server.ID, b.client.DefaultTimeoutDuration); err != nil {
		return err
	}

	// refresh CurrentBildResult.Server
	s, err := b.client.Server.Read(server.ID)
	if err != nil {
		return err
	}
	b.currentBuildResult.Server = s

	return nil
}

func (b *serverBuilder) callEventHandlerIfExists(event ServerBuildEvents) {
	if handler, ok := b.buildEventHandlers[event]; ok {
		handler(b.currentBuildValue, b.currentBuildResult)
	}
}

/**********************************************************
  Type : ServerBuildValue
**********************************************************/

// ServerBuildValue サーバー構築用パラメータ
type ServerBuildValue struct {
	// Server サーバー作成用パラメータ
	Server *sacloud.Server
}

/**********************************************************
  Type : ServerBuildResult
**********************************************************/

// ServerBuildResult サーバー構築結果
type ServerBuildResult struct {
	// Server サーバー
	Server *sacloud.Server
	// Disks ディスク構築結果
	Disks []*DiskBuildResult
}

func (s *ServerBuildResult) addDisk(disk *DiskBuildResult) {
	s.Disks = append(s.Disks, disk)
}
