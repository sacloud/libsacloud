package sacloud

type Region struct {
	*Resource
	Name        string   `json:",omitempty"`
	Description string   `json:",omitempty"`
	NameServers []string `json:",omitempty"`
}
