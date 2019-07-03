package types

// Protocol パケットフィルタやロードバランサ、VPCルータなどで利用するプロトコル
type Protocol string

// Protocols パケットフィルタやロードバランサ、VPCルータなどで利用するプロトコル
var Protocols = &struct {
	HTTP     Protocol
	HTTPS    Protocol
	TCP      Protocol
	UDP      Protocol
	Ping     Protocol
	Fragment Protocol
	IP       Protocol
}{
	HTTP:     "http",
	HTTPS:    "https",
	TCP:      "tcp",
	UDP:      "udp",
	Ping:     "pint",
	Fragment: "fragment",
	IP:       "ip",
}
