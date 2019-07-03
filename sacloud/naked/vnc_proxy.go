package naked

// VNCProxy VNCプロキシ
type VNCProxy struct {
	HostName  string `json:",omitempty" yaml:"host_name,omitempty" structs:",omitempty"`
	IPAddress string `json:",omitempty" yaml:"ip_address,omitempty" structs:",omitempty"`
}
