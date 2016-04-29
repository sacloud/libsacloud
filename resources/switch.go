package resources

type Switch struct {
	*Resource
	Name       string  `json:",omitempty"`
	Scope      string  `json:",omitempty"`
	Subnet     *Subnet `json:",omitempty"`
	UserSubnet *Subnet `json:",omitempty"`
}

type Subnet struct {
	*Resource
	NetworkAddress string `json:",omitempty"`
	NetworkMaskLen int    `json:",omitempty"`
	DefaultRoute   string `json:",omitempty"`
	Internet       struct {
		BandWidthMbps int `json:",omitempty"`
	} `json:",omitempty"`
}
