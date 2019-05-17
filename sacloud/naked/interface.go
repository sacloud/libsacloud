package naked

import "github.com/sacloud/libsacloud-v2/sacloud/types"

// Interface サーバなどに接続されているNICの情報
type Interface struct {
	ID            types.ID      `json:",omitempty" yaml:"id,omitempty" structs:",omitempty"`
	MACAddress    string        `json:",omitempty" yaml:"mac_address,omitempty" structs:",omitempty"`
	IPAddress     string        `json:",omitempty" yaml:"ip_address,omitempty" structs:",omitempty"`
	UserIPAddress string        `json:",omitempty" yaml:"user_ip_address,omitempty" structs:",omitempty"`
	HostName      string        `json:",omitempty" yaml:"host_name,omitempty" structs:",omitempty"`
	Switch        *Switch       `json:",omitempty" yaml:"switch,omitempty" structs:",omitempty"`
	PacketFilter  *PacketFilter `json:",omitempty" yaml:"packet_filter,omitempty" structs:",omitempty"`
	Server        *Server       `json:",omitempty" yaml:"server,omitempty" structs:",omitempty"`
}
