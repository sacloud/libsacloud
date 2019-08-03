package types

// EGSLBHealthCheckProtocol GSLB 監視プロトコル
type EGSLBHealthCheckProtocol string

// GSLBHealthCheckProtocols GSLB 監視プロトコル
var GSLBHealthCheckProtocols = struct {
	// Unknown 不明
	Unknown EGSLBHealthCheckProtocol
	// HTTP http
	HTTP EGSLBHealthCheckProtocol
	// HTTPS https
	HTTPS EGSLBHealthCheckProtocol
	// TCP tcp
	TCP EGSLBHealthCheckProtocol
	// Ping ping
	Ping EGSLBHealthCheckProtocol
}{
	Unknown: EGSLBHealthCheckProtocol(""),
	HTTP:    EGSLBHealthCheckProtocol("http"),
	HTTPS:   EGSLBHealthCheckProtocol("https"),
	TCP:     EGSLBHealthCheckProtocol("tcp"),
	Ping:    EGSLBHealthCheckProtocol("ping"),
}
