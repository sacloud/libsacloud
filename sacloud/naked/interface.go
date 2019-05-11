package naked

// Interface サーバなどに接続されているNICの情報
type Interface struct {
	HostName      string  `json:",omitempty" yaml:"host_name,omitempty" structs:",omitempty"`
	IPAddress     string  `json:",omitempty" yaml:"ip_address,omitempty" structs:",omitempty"`
	Switch        *Switch `json:",omitempty" yaml:"switch,omitempty" structs:",omitempty"`
	UserIPAddress string  `json:",omitempty" yaml:"user_ip_address,omitempty" structs:",omitempty"`
}
