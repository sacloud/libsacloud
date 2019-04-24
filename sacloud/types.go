package sacloud

// EAvailability 有効状態
type EAvailability string

// Availability 有効状態
type Availability struct {
	Available EAvailability // 有効
	Uploading EAvailability // アップロード中
	Failed    EAvailability // 失敗
	Migrating EAvailability // マイグレーション中
}

// Availabilities 有効状態
var Availabilities = Availability{
	Available: EAvailability("available"),
	Uploading: EAvailability("uploading"),
	Failed:    EAvailability("failed"),
	Migrating: EAvailability("migrating"),
}

// IsAvailable 有効状態が"有効"か判定
func (e EAvailability) IsAvailable() bool {
	return e == Availabilities.Available
}

// IsUploading 有効状態が"アップロード中"か判定
func (e EAvailability) IsUploading() bool {
	return e == Availabilities.Uploading
}

// IsFailed 有効状態が"失敗"か判定
func (e EAvailability) IsFailed() bool {
	return e == Availabilities.Failed
}

// IsMigrating 有効状態が"マイグレーション中"か判定
func (e EAvailability) IsMigrating() bool {
	return e == Availabilities.Migrating
}

// EInterfaceDriver インターフェースドライバ
type EInterfaceDriver string

func (d EInterfaceDriver) String() string {
	return string(d)
}

// InterfaceDriver インターフェースドライバ
type InterfaceDriver struct {
	VirtIO EInterfaceDriver // virtio
	E1000  EInterfaceDriver // e1000
}

var (
	// InterfaceDrivers インターフェースドライバ
	InterfaceDrivers = InterfaceDriver{
		VirtIO: EInterfaceDriver("virtio"),
		E1000:  EInterfaceDriver("e1000"),
	}

	// InterfaceDriverMap インターフェースドライバと文字列表現のマップ
	InterfaceDriverMap = map[string]EInterfaceDriver{
		InterfaceDrivers.VirtIO.String(): InterfaceDrivers.VirtIO,
		InterfaceDrivers.E1000.String():  InterfaceDrivers.E1000,
	}

	// InterfaceDriverValues インターフェースドライバが取りうる有効値
	InterfaceDriverValues = []string{
		InterfaceDrivers.VirtIO.String(),
		InterfaceDrivers.E1000.String(),
	}
)

// EServerInstanceStatus サーバーインスタンスステータス
type EServerInstanceStatus string

// ServerInstanceStatus サーバーインスタンスステータス
type ServerInstanceStatus struct {
	Up   EServerInstanceStatus
	Down EServerInstanceStatus
}

// ServerInstanceStatuses サーバーインスタンスステータス
var ServerInstanceStatuses = &ServerInstanceStatus{
	Up:   EServerInstanceStatus("up"),
	Down: EServerInstanceStatus("down"),
}

// IsUp インスタンスが起動しているか判定
func (e EServerInstanceStatus) IsUp() bool {
	return e == ServerInstanceStatuses.Up
}

// IsDown インスタンスがダウンしているか確認
func (e EServerInstanceStatus) IsDown() bool {
	return e == ServerInstanceStatuses.Down
}

// EScope スコープ
type EScope string

// Scope スコープ
type Scope struct {
	Shared EScope // 共有
	User   EScope // ユーザー
}

// Scopes スコープ
var Scopes = &Scope{
	Shared: EScope("shared"),
	User:   EScope("user"),
}

// EDiskConnection ディスク接続方法
type EDiskConnection string

// String EDiskConnectionの文字列表現
func (c EDiskConnection) String() string {
	return string(c)
}

// DiskConnection ディスク接続方法
type DiskConnection struct {
	VirtIO EDiskConnection
	IDE    EDiskConnection
}

// DiskConnections ディスク接続方法
var (
	DiskConnections = DiskConnection{
		VirtIO: EDiskConnection("virtio"),
		IDE:    EDiskConnection("ide"),
	}

	DiskConnectionMap = map[string]EDiskConnection{
		DiskConnections.VirtIO.String(): DiskConnections.VirtIO,
		DiskConnections.IDE.String():    DiskConnections.IDE,
	}

	DiskConnectionValues = []string{
		DiskConnections.VirtIO.String(),
		DiskConnections.IDE.String(),
	}
)

// EUpstreamNetworkType 上流ネットワーク種別
type EUpstreamNetworkType string

// String EUpstreamNetworkTypeの文字列表現
func (t EUpstreamNetworkType) String() string {
	return string(t)
}

// UpstreamNetworkType 上流ネットワーク種別
type UpstreamNetworkType struct {
	Unknown EUpstreamNetworkType
	Shared  EUpstreamNetworkType
	Switch  EUpstreamNetworkType
	Router  EUpstreamNetworkType
	None    EUpstreamNetworkType
}

var (
	// UpstreamNetworkTypes 上流ネットワーク種別
	UpstreamNetworkTypes = UpstreamNetworkType{
		Unknown: EUpstreamNetworkType("unknown"),
		Shared:  EUpstreamNetworkType("shared"),
		Switch:  EUpstreamNetworkType("switch"),
		Router:  EUpstreamNetworkType("router"),
		None:    EUpstreamNetworkType("none"),
	}

	// UpstreamNetworkTypeMap 文字列とEUpstreamNetworkTypeのマッピング
	UpstreamNetworkTypeMap = map[string]EUpstreamNetworkType{
		"unknown": UpstreamNetworkTypes.Unknown,
		"shared":  UpstreamNetworkTypes.Shared,
		"switch":  UpstreamNetworkTypes.Switch,
		"router":  UpstreamNetworkTypes.Router,
		"none":    UpstreamNetworkTypes.None,
	}
)

// EPlanGeneration サーバプラン世代
type EPlanGeneration int

// PlanGeneration サーバプラン世代
type PlanGeneration struct {
	Default EPlanGeneration // デフォルト(自動選択)
	G100    EPlanGeneration // 第1世代
	G200    EPlanGeneration // 第2世代
}

var (
	// PlanGenerationValues 有効なサーバプラン世代の値
	PlanGenerationValues = []int{
		int(PlanGenerations.Default),
		int(PlanGenerations.G100),
		int(PlanGenerations.G200),
	}

	// PlanGenerations サーバプラン世代
	PlanGenerations = PlanGeneration{
		Default: EPlanGeneration(0),
		G100:    EPlanGeneration(100),
		G200:    EPlanGeneration(200),
	}
)

// SpecialTag 特殊タグ
type SpecialTag string

// SpecialTagValue 特殊タグ
type SpecialTagValue struct {
	// GroupA サーバをグループ化し起動ホストを分離します(グループA)
	GroupA SpecialTag
	// GroupB サーバをグループ化し起動ホストを分離します(グループB)
	GroupB SpecialTag
	// GroupC サーバをグループ化し起動ホストを分離します(グループC)
	GroupC SpecialTag
	// GroupD サーバをグループ化し起動ホストを分離します(グループD)
	GroupD SpecialTag
	// AutoReboot サーバ停止時に自動起動します
	AutoReboot SpecialTag
	// KeyboardUS リモートスクリーン画面でUSキーボード入力します
	KeyboardUS SpecialTag
	// BootCDROM 優先ブートデバイスをCD-ROMに設定します
	BootCDROM SpecialTag
	// BootNetwork 優先ブートデバイスをPXE bootに設定します
	BootNetwork SpecialTag
}

// SpecialTags 特殊タグ一覧
var SpecialTags = &SpecialTagValue{
	GroupA:      SpecialTag("@group=a"),
	GroupB:      SpecialTag("@group=b"),
	GroupC:      SpecialTag("@group=c"),
	GroupD:      SpecialTag("@group=d"),
	AutoReboot:  SpecialTag("@auto-reboot"),
	KeyboardUS:  SpecialTag("@keyboard-us"),
	BootCDROM:   SpecialTag("@boot-cdrom"),
	BootNetwork: SpecialTag("@boot-network"),
}
