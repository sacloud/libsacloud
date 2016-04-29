package resources

import "time"

// Server type of create server request values
type Server struct {
	*Resource
	Name              string
	Index             int                 `json:",omitempty"`
	HostName          string              `json:",omitempty"`
	Description       string              `json:",omitempty"`
	Availability      string              `json:",omitempty"`
	ServiceClass      string              `json:",omitempty`
	CreatedAt         time.Time           `json:",omitempty"`
	Icon              NumberResource      `json:",omitempty"`
	ServerPlan        *ServerPlan         `json:",omitempty"`
	Zone              *Zone               `json:",omitempty`
	Tags              []string            `json:",omitempty"`
	ConnectedSwitches []map[string]string `json:",omitempty"`
	Disks             []Disk              `json:",omitempty"`
	Interfaces        []Interface         `json:",omitempty"`
	Instance          *Instance           `json:",omitempty"`
}
