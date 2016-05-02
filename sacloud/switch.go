package sacloud

import "time"

// Switch type of switch
type Switch struct {
	*Resource
	Name           string  `json:",omitempty"`
	Description    string  `json:",omitempty"`
	ServerCount    int     `json:",omitempty"`
	ApplianceCount int     `json:",omitempty"`
	Scope          EScope  `json:",omitempty"`
	Subnet         *Subnet `json:",omitempty"`
	UserSubnet     *Subnet `json:",omitempty"`
	//HybridConnection
	ServerClass string    `json:",omitempty"`
	CreatedAt   time.Time `json:",omitempty"`
	Icon        *Icon     `json:",omitempty"`
	Tags        []string  `json:",omitempty"`
	Subnets     []Subnet  `json:",omitempty"`
	IPv6Nets    []IPv6Net `json:",omitempty"`
	Internet    *Internet `json:",omitempty"`
	Bridge      *Bridge   `json:",omitempty"`
}

// Subnet type of Subnet
type Subnet struct {
	*NumberResource
	NetworkAddress string `json:",omitempty"`
	NetworkMaskLen int    `json:",omitempty"`
	DefaultRoute   string `json:",omitempty"`
	//NextHop ???
	//StaticRoute ???
	ServiceClass string `json:",omitempty"`
	IPAddresses  struct {
		Min string `json:",omitempty"`
		Max string `json:",omitempty"`
	}
	Internet *Internet `json:",omitempty"`
}

type Internet struct {
	*Resource
	Name          string `json:",omitempty"`
	BandWidthMbps int    `json:",omitempty"`
}

type IPv6Net struct {
	*NumberResource
	IPv6Prefix    string `json:",omitempty"`
	IPv6PrefixLen int    `json:",omitempty"`
	Scope         string `json:",omitempty"`
	ServiceClass  string `json:",omitempty"`
}
