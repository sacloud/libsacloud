package sacloud

type VPCRouter struct {
	*appliance
	Remark   VPCRouterRemark    `json:",omitempty"`
	Settings *VPCRouterSettings `json:",omitempty"`
}

type VPCRouterRemark struct {
	*applianceRemarkBase
	Zone *Resource
}

type VPCRouterSettings struct {
	Router *VPCRouterSetting `json:",omitempty"`
}

type VPCRouterSetting struct {
	Interfaces []*VPCRouterInterface `json:",omitempty"`
	VRID       int                   `json:",omitempty"`
}
type VPCRouterInterface struct {
	IPAddress        []string `json:",omitempty"`
	NetworkMaskLen   int      `json:",omitempty"`
	VirtualIPAddress string   `json:",omitempty"`
}
