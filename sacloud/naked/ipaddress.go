package naked

// IPAddress IPアドレス(IPv4)
type IPAddress struct {
	HostName  string     `yaml:"host_name"`
	IPAddress string     `json:",omitempty" yaml:"ip_address,omitempty" structs:",omitempty"`
	Interface *Interface `json:",omitempty" yaml:"interface,omitempty" structs:",omitempty"`
	Subnet    *Subnet    `json:",omitempty" yaml:"subnet,omitempty" structs:",omitempty"`
}
