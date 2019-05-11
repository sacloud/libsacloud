package types

// EUpstreamNetworkType 上流ネットワーク種別
type EUpstreamNetworkType string

// String EUpstreamNetworkTypeの文字列表現
func (t EUpstreamNetworkType) String() string {
	return string(t)
}

var (
	// UpstreamNetworkTypes 上流ネットワーク種別
	UpstreamNetworkTypes = struct {
		// Unknown 不明(invalid)
		Unknown EUpstreamNetworkType
		// Shared 共有セグメント
		Shared EUpstreamNetworkType
		// Switch スイッチ
		Switch EUpstreamNetworkType
		// Router ルータ
		Router EUpstreamNetworkType
		// None 接続なし
		None EUpstreamNetworkType
	}{
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
