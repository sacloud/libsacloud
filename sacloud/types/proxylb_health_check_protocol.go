package types

// EProxyLBHealthCheckProtocol エンハンスドロードバランサ 監視プロトコル
type EProxyLBHealthCheckProtocol string

// ProxyLBProtocols エンハンスドロードバランサ 監視プロトコル
var ProxyLBProtocols = struct {
	// Unknown 不明
	Unknown EProxyLBHealthCheckProtocol
	// HTTP http
	HTTP EProxyLBHealthCheckProtocol
	// TCP tcp
	TCP EProxyLBHealthCheckProtocol
}{
	Unknown: EProxyLBHealthCheckProtocol(""),
	HTTP:    EProxyLBHealthCheckProtocol("http"),
	TCP:     EProxyLBHealthCheckProtocol("tcp"),
}
