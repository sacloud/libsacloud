package sacloud

// IPAddress type of SakuraCloud IPAddress
type IPAddress struct {
	HostName  string    `json:",omitempty"`
	IPAddress string    `json:",omitempty"`
	Interface *Internet `json:",omitempty"`
	Subnet    *Subnet   `json:",omitempty"`
}
