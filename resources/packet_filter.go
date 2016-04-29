package resources

// PacketFilter type of PacketFilter
type PacketFilter struct {
	*Resource
	Name                string
	Description         string `json:",omitempty"`
	RequiredHostVersion int    `json:",omitempty"`
	//	Expression          string `json:",omitempty"`
	Notice string `json:",omitempty"`
}
