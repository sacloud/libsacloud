package types

// EVPCRouterPortForwardingProtocol ポートフォワーディング プロトコル
type EVPCRouterPortForwardingProtocol string

// VPCRouterPortForwardingProtocols ポートフォワーディング プロトコル
var VPCRouterPortForwardingProtocols = struct {
	// Unknown 不明
	Unknown EVPCRouterPortForwardingProtocol
	// TCP tcp
	TCP EVPCRouterPortForwardingProtocol
	// UDP http
	UDP EVPCRouterPortForwardingProtocol
}{
	Unknown: EVPCRouterPortForwardingProtocol(""),
	TCP:     EVPCRouterPortForwardingProtocol("tcp"),
	UDP:     EVPCRouterPortForwardingProtocol("udp"),
}
