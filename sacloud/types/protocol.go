package types

// Protocol パケットフィルタやロードバランサ、VPCルータなどで利用するプロトコル
type Protocol string

// Protocols パケットフィルタやロードバランサ、VPCルータなどで利用するプロトコル
var Protocols = &struct {
	HTTP     Protocol
	HTTPS    Protocol
	TCP      Protocol
	UDP      Protocol
	ICMP     Protocol
	Fragment Protocol
	IP       Protocol
}{
	HTTP:     "http",
	HTTPS:    "https",
	TCP:      "tcp",
	UDP:      "udp",
	ICMP:     "icmp",
	Fragment: "fragment",
	IP:       "ip",
}

// String Protocolの文字列表現
func (p Protocol) String() string {
	return string(p)
}
