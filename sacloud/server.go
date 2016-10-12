package sacloud

import (
	"time"
)

// Server type of create server request values
type Server struct {
	*Resource
	Name        string
	Index       int    `json:",omitempty"`
	HostName    string `json:",omitempty"`
	Description string `json:",omitempty"`
	*EAvailability
	ServiceClass string         `json:",omitempty"`
	CreatedAt    *time.Time     `json:",omitempty"`
	Icon         *Resource      `json:",omitempty"`
	ServerPlan   *ProductServer `json:",omitempty"`
	Zone         *Zone          `json:",omitempty"`
	*TagsType
	//Tags              []string       //`json:",omitempty"`
	ConnectedSwitches []interface{} `json:",omitempty" libsacloud:"requestOnly"`
	//InterfaceNum      int            `json:",omitempty" libsacloud:"requestOnly"` !Not support! ConnectedSwitchesで代替
	Disks      []Disk      `json:",omitempty"`
	Interfaces []Interface `json:",omitempty"`
	Instance   *Instance   `json:",omitempty"`
}

func (s *Server) SetServerPlanByID(planID string) {
	if s.ServerPlan == nil {
		s.ServerPlan = &ProductServer{Resource: NewResourceByStringID(planID)}
	}
}

func (s *Server) ClearConnectedSwitches() {
	s.ConnectedSwitches = []interface{}{}
}

// AddPublicNWConnectedParam add "Scope:shared" to Server#ConnectedSwitches
func (s *Server) AddPublicNWConnectedParam() {
	if s.ConnectedSwitches == nil {
		s.ClearConnectedSwitches()
	}
	s.ConnectedSwitches = append(s.ConnectedSwitches, map[string]interface{}{"Scope": "shared"})
}

// AddExistsSwitchConnectedParam add "ID:[switchID]" to Server#ConnectedSwitches
func (s *Server) AddExistsSwitchConnectedParam(switchID string) {
	if s.ConnectedSwitches == nil {
		s.ClearConnectedSwitches()
	}
	s.ConnectedSwitches = append(s.ConnectedSwitches, map[string]interface{}{"ID": switchID})
}

// AddEmptyConnectedParam  add "null" to Server#ConnectedSwitches
func (s *Server) AddEmptyConnectedParam() {
	if s.ConnectedSwitches == nil {
		s.ClearConnectedSwitches()
	}
	s.ConnectedSwitches = append(s.ConnectedSwitches, nil)
}

// GetDiskIDs ディスクID配列を返す
func (s *Server) GetDiskIDs() []int64 {

	ids := []int64{}
	for _, disk := range s.Disks {
		ids = append(ids, disk.ID)
	}
	return ids

}

// KeyboardRequest type of send-key request
type KeyboardRequest struct {
	Keys []string `json:",omitempty"`
	Key  string   `json:",omitempty"`
}

// MouseRequest type of send-mouse request
type MouseRequest struct {
	X       *int                 `json:",omitempty"`
	Y       *int                 `json:",omitempty"`
	Z       *int                 `json:",omitempty"`
	Buttons *MouseRequestButtons `json:",omitempty"`
}

// VNCSnapshotRequest type of VNC snapshot request
type VNCSnapshotRequest struct {
	ScreenSaverExitTimeMS int `json:",omitempty"`
}

// MouseRequestButtons type of send-mouse request buttons
type MouseRequestButtons struct {
	L bool `json:",omitempty"`
	R bool `json:",omitempty"`
	M bool `json:",omitempty"`
}

// VNCProxyResponse type of VNC Proxy response from server
type VNCProxyResponse struct {
	*ResultFlagValue
	Status   string `json:",omitempty"`
	Host     string `json:",omitempty"`
	Port     string `json:",omitempty"`
	Password string `json:",omitempty"`
	VNCFile  string `json:",omitempty"`
}

// VNCSizeResponse type of VNC display size response from server
type VNCSizeResponse struct {
	Width  int `json:",string,omitempty"`
	Height int `json:",string,omitempty"`
}

// VNCSnapshotResponse type of VNC snapshot response
type VNCSnapshotResponse struct {
	Image string `json:",omitempty"`
}
