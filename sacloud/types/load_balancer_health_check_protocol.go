package types

// ELoadBalancerHealthCheckProtocol ロードバランサ 監視プロトコル
type ELoadBalancerHealthCheckProtocol string

// LoadBalancerHealthCheckProtocols ロードバランサ 監視プロトコル
var LoadBalancerHealthCheckProtocols = struct {
	// Unknown 不明
	Unknown ELoadBalancerHealthCheckProtocol
	// HTTP http
	HTTP ELoadBalancerHealthCheckProtocol
	// HTTPS https
	HTTPS ELoadBalancerHealthCheckProtocol
	// TCP tcp
	TCP ELoadBalancerHealthCheckProtocol
	// Ping ping
	Ping ELoadBalancerHealthCheckProtocol
}{
	Unknown: ELoadBalancerHealthCheckProtocol(""),
	HTTP:    ELoadBalancerHealthCheckProtocol("http"),
	HTTPS:   ELoadBalancerHealthCheckProtocol("https"),
	TCP:     ELoadBalancerHealthCheckProtocol("tcp"),
	Ping:    ELoadBalancerHealthCheckProtocol("ping"),
}
