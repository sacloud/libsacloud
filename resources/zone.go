package resources

// Zone type of zone
type Zone struct {
	*NumberResource
	DisplayOrder int    `json:",omitempty"`
	Name         string `json:",omitempty"`
	Description  string `json:",omitempty"`
	IsDummy      bool   `json:",omitempty"`
	VNCProxy     struct {
		HostName  string `json:",omitempty"`
		IPAddress string `json:",omitempty"`
	} `json:",omitempty"`
	FTPServer struct {
		HostName  string `json:",omitempty"`
		IPAddress string `json:",omitempty"`
	} `json:",omitempty"`
	//Settings struct {
	//}
	Region struct {
		*NumberResource
		Name        string   `json:",omitempty"`
		Description string   `json:",omitempty"`
		NameServers []string `json:",omitempty"`
	} `json:",omitempty"`
}
