package sacloud

import "time"

// Subnet type of SakuraCloud Subnet
type Subnet struct {
	*Resource
	DefaultRoute   string       `json:",omitempty"`
	CreatedAt      *time.Time   `json:",omitempty"`
	IPAddresses    []*IPAddress `json:",omitempty"`
	NetworkAddress string       `json:",omitempty"`
	NetworkMaskLen int          `json:",omitempty"`
	ServiceClass   string       `json:",omitempty"`
	ServiceID      int64        `json:",omitempty"`
	StaticRoute    string       `json:",omitempty"`
	Switch         *Switch      `json:",omitempty"`
	Internet       *Internet    `json:",omitempty"`
}
