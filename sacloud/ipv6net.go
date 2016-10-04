package sacloud

import "time"

// IPv6Net type of SakuraCloud IPv6 Network
type IPv6Net struct {
	*Resource

	IPv6Prefix         string     `json:",omitempty"`
	IPv6PrefixLen      int        `json:",omitempty"`
	IPv6PrefixTail     string     `json:",omitempty"`
	IPv6Table          *Resource  `json:",omitempty"`
	NamedIPv6AddrCount int        `json:",omitempty"`
	ServiceID          int64      `json:",omitempty"`
	ServiceClass       string     `json:",omitempty"`
	Scope              string     `json:",omitempty"`
	CreatedAt          *time.Time `json:",omitempty"`
	Switch             *Switch    `json:",omitempty"`
}
