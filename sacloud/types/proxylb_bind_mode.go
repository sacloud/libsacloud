package types

// EProxyLBProxyMode エンハンスドロードバランサでのプロキシ方式
type EProxyLBProxyMode string

// ProxyLBProxyModes エンハンスドロードバランサでのプロキシ方式
var ProxyLBProxyModes = struct {
	// Unknown 不明
	Unknown EProxyLBProxyMode
	// HTTP .
	HTTP EProxyLBProxyMode
	// HTTPS .
	HTTPS EProxyLBProxyMode
}{
	Unknown: EProxyLBProxyMode(""),
	HTTP:    EProxyLBProxyMode("http"),
	HTTPS:   EProxyLBProxyMode("https"),
}
