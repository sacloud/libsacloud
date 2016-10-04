package sacloud

type IPv6Addr struct {
	HostName  string    `json:",omitempty"`
	IPv6Addr  string    `json:",omitempty"`
	Interface *Internet `json:",omitempty"`
	IPv6Net   *IPv6Net  `json:",omitempty"`
}

func CreateNewIPv6Addr() *IPv6Addr {
	return &IPv6Addr{
		IPv6Net: &IPv6Net{
			Resource: &Resource{},
		},
	}
}
