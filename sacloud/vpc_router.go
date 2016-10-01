package sacloud

type VPCRouter struct {
	*Appliance
	Remark   *VPCRouterRemark   `json:",omitempty"`
	Settings *VPCRouterSettings `json:",omitempty"`
}

type VPCRouterRemark struct {
	*ApplianceRemarkBase
	// TODO Zone
	//Zone *Resource
}

type VPCRouterSettings struct {
	Router *VPCRouterSetting `json:",omitempty"`
}

func CreateNewVPCRouter() *VPCRouter {
	return &VPCRouter{
		Appliance: &Appliance{
			Class:    "vpcrouter",
			Plan:     &Resource{},
			TagsType: &TagsType{},
		},
		Remark: &VPCRouterRemark{
			ApplianceRemarkBase: &ApplianceRemarkBase{
				Servers: []interface{}{""},
				Switch:  &ApplianceRemarkSwitch{},
			},
		},
		Settings: &VPCRouterSettings{
			Router: &VPCRouterSetting{},
		},
	}
}

func (v *VPCRouter) InitVPCRouterSetting() {
	settings := &VPCRouterSettings{
		Router: &VPCRouterSetting{},
	}

	if v.Settings != nil && v.Settings.Router != nil && v.Settings.Router.Interfaces != nil {
		settings.Router.Interfaces = v.Settings.Router.Interfaces
	}
	if v.Settings != nil && v.Settings.Router != nil && v.Settings.Router.VRID != nil {
		settings.Router.VRID = v.Settings.Router.VRID
	}

	v.Settings = settings
}

func (v *VPCRouter) IsStandardPlan() bool {
	return v.Plan.ID == 1
}
func (v *VPCRouter) IsPremiumPlan() bool {
	return v.Plan.ID == 2
}
func (v *VPCRouter) IsHighSpecPlan() bool {
	return v.Plan.ID == 3
}

func (v *VPCRouter) SetStandardPlan() {
	v.Plan.SetID(1)
	v.Remark.Switch = &ApplianceRemarkSwitch{
		Scope: "shared",
	}
	v.Settings = nil
}

func (v *VPCRouter) SetPremiumPlan(switchID string, virtualIPAddress string, ipAddress1 string, ipAddress2 string, vrid int, ipAliases []string) {
	v.Plan.SetID(2)
	v.setPremiumServices(switchID, virtualIPAddress, ipAddress1, ipAddress2, vrid, ipAliases)
}

func (v *VPCRouter) SetHighSpecPlan(switchID string, virtualIPAddress string, ipAddress1 string, ipAddress2 string, vrid int, ipAliases []string) {
	v.Plan.SetID(3)
	v.setPremiumServices(switchID, virtualIPAddress, ipAddress1, ipAddress2, vrid, ipAliases)
}

func (v *VPCRouter) setPremiumServices(switchID string, virtualIPAddress string, ipAddress1 string, ipAddress2 string, vrid int, ipAliases []string) {
	v.Remark.Switch = &ApplianceRemarkSwitch{
		ID: switchID,
	}
	v.Remark.Servers = []interface{}{
		map[string]string{"IPAddress": ipAddress1},
		map[string]string{"IPAddress": ipAddress2},
	}

	v.Settings = &VPCRouterSettings{
		Router: &VPCRouterSetting{
			Interfaces: []*VPCRouterInterface{
				{
					IPAddress: []string{
						ipAddress1,
						ipAddress2,
					},
					VirtualIPAddress: virtualIPAddress,
					IPAliases:        ipAliases,
				},
			},
			VRID: &vrid,
		},
	}

}
