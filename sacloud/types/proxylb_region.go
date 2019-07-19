package types

// EProxyLBRegion エンハンスドロードバランサ 設置先リージョン
type EProxyLBRegion string

// ProxyLBRegions エンハンスドロードバランサ 設置先リージョン
var ProxyLBRegions = struct {
	// Unknown 不明
	Unknown EProxyLBRegion
	// TK1 東京
	TK1 EProxyLBRegion
	// IS1 石狩
	IS1 EProxyLBRegion
}{
	Unknown: EProxyLBRegion(""),
	TK1:     EProxyLBRegion("tk1"),
	IS1:     EProxyLBRegion("is1"),
}
