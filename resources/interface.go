package resources

// Interface type of server nic
type Interface struct {
	*Resource
	MACAddress    string        `json:",omitempty"`
	IPAddress     string        `json:",omitempty"`
	UserIPAddress string        `json:",omitempty"`
	HostName      string        `json:",omitempty"`
	Switch        *Switch       `json:",omitempty"`
	PacketFilter  *PacketFilter `json:",omitempty"`
}
