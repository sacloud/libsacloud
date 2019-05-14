package define

import (
	"github.com/sacloud/libsacloud-v2/internal/schema"
	"github.com/sacloud/libsacloud-v2/internal/schema/meta"
	"github.com/sacloud/libsacloud-v2/sacloud/naked"
)

type fieldsDef struct{}

var fields = &fieldsDef{}

func (f *fieldsDef) ID() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "ID",
		Type: meta.TypeID,
		ExtendAccessors: []*schema.ExtendAccessor{
			{
				Name: "StringID",
				Type: meta.TypeString,
			},
			{
				Name: "Int64ID",
				Type: meta.TypeInt64,
			},
		},
	}
}

func (f *fieldsDef) Name() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "Name",
		Type: meta.TypeString,
		Tags: &schema.FieldTags{
			Validate: "required",
		},
	}
}

func (f *fieldsDef) PlanID() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "PlanID",
		Tags: &schema.FieldTags{
			MapConv: "Plan.ID",
		},
		Type: meta.TypeID,
	}
}

func (f *fieldsDef) IconID() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "IconID",
		Tags: &schema.FieldTags{
			MapConv: "Icon.ID",
		},
		Type: meta.TypeID,
	}
}

func (f *fieldsDef) AppliancePlanID() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "PlanID",
		Tags: &schema.FieldTags{
			MapConv: "Remark.Plan.ID,Plan.ID",
		},
		Type: meta.TypeID,
	}
}

func (f *fieldsDef) ApplianceSwitchID() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "SwitchID",
		Tags: &schema.FieldTags{
			MapConv: "Remark.Switch.ID",
		},
		Type: meta.TypeID,
	}
}

func (f *fieldsDef) ApplianceIPAddress() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "IPAddress",
		Tags: &schema.FieldTags{
			MapConv: "Remark.[]Servers.IPAddress",
		},
		Type: meta.TypeString,
	}
}

func (f *fieldsDef) ApplianceIPAddresses() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "IPAddresses",
		Type: meta.TypeStringSlice,
		Tags: &schema.FieldTags{
			MapConv:  "Remark.[]Servers.IPAddress",
			Validate: "min=1,max=2,dive,ipv4",
		},
	}
}

func (f *fieldsDef) LoadBalancerVIP() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "VirtualIPAddresses",
		Type: &schema.Model{
			Name:    "LoadBalancerVirtualIPAddress",
			IsArray: true,
			Fields: []*schema.FieldDesc{
				{
					Name: "VirtualIPAddress",
					Type: meta.TypeString,
					Tags: &schema.FieldTags{
						Validate: "ipv4",
					},
				},
				{
					Name: "Port",
					Type: meta.TypeStringNumber,
				},
				{
					Name: "DelayLoop",
					Type: meta.TypeStringNumber,
					Tags: &schema.FieldTags{
						MapConv:  ",default=10",
						Validate: "min=0,max=60", // TODO 最大値確認
					},
				},
				{
					Name: "SorryServer",
					Type: meta.TypeString,
					Tags: &schema.FieldTags{
						Validate: "ipv4",
					},
				},
				f.Description(),
				{
					Name: "Servers",
					Type: &schema.Model{
						Name:    "LoadBalancerServer",
						IsArray: true,
						Fields: []*schema.FieldDesc{
							{
								Name: "IPAddress",
								Type: meta.TypeString,
								Tags: &schema.FieldTags{
									Validate: "ipv4",
								},
							},
							{
								Name: "Port",
								Type: meta.TypeStringNumber,
								Tags: &schema.FieldTags{
									Validate: "min=1,max=65535",
								},
							},
							{
								Name: "Enabled",
								Type: meta.TypeStringFlag,
								Tags: &schema.FieldTags{
									MapConv: ",default=true",
								},
							},
							{
								Name: "HealthCheckProtocol",
								Type: meta.TypeString,
								Tags: &schema.FieldTags{
									MapConv:  "HealthCheck.Protocol",
									Validate: "oneof=http https ping tcp",
								},
							},
							{
								Name: "HealthCheckPath",
								Type: meta.TypeString,
								Tags: &schema.FieldTags{
									MapConv: "HealthCheck.Path",
								},
							},
							{
								Name: "HealthCheckResponseCode",
								Type: meta.TypeStringNumber,
								Tags: &schema.FieldTags{
									MapConv: "HealthCheck.Status",
								},
							},
						},
					},
					Tags: &schema.FieldTags{
						MapConv:  ",recursive",
						Validate: "min=0,max=40",
					},
				},
			},
		},
		Tags: &schema.FieldTags{
			MapConv:  "Settings.[]LoadBalancer,recursive",
			Validate: "min=0,max=10",
		},
	}
}

func (f *fieldsDef) Tags() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "Tags",
		Type: meta.TypeStringSlice,
	}
}

func (f *fieldsDef) Class() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "Class",
		Type: meta.TypeString,
	}
}

func (f *fieldsDef) NFSClass() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name:     "Class",
		ReadOnly: true,
		Type:     meta.TypeString,
		Tags: &schema.FieldTags{
			MapConv: ",default=nfs",
		},
	}
}

func (f *fieldsDef) LoadBalancerClass() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name:     "Class",
		ReadOnly: true,
		Type:     meta.TypeString,
		Tags: &schema.FieldTags{
			MapConv: ",default=loadbalancer",
		},
	}
}

func (f *fieldsDef) GSLBProviderClass() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name:     "Class",
		ReadOnly: true,
		Type:     meta.TypeString,
		Tags: &schema.FieldTags{
			MapConv: "Provider.Class,default=gslb",
		},
	}
}

func (f *fieldsDef) GSLBFQDN() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name:     "FQDN",
		ReadOnly: true,
		Type:     meta.TypeString,
		Tags: &schema.FieldTags{
			MapConv: "Status.FQDN",
		},
	}
}

func (f *fieldsDef) GSLBHealthCheckProtocol() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "HealthCheckProtocol",
		Type: meta.TypeString,
		Tags: &schema.FieldTags{
			MapConv:  "Settings.GSLB.HealthCheck.Protocol",
			Validate: "oneof=http https ping tcp",
		},
	}
}

func (f *fieldsDef) GSLBHealthCheckHostHeader() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "HealthCheckHostHeader",
		Type: meta.TypeString,
		Tags: &schema.FieldTags{
			MapConv: "Settings.GSLB.HealthCheck.Host",
		},
	}
}

func (f *fieldsDef) GSLBHealthCheckPath() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "HealthCheckPath",
		Type: meta.TypeString,
		Tags: &schema.FieldTags{
			MapConv: "Settings.GSLB.HealthCheck.Path",
		},
	}
}

func (f *fieldsDef) GSLBHealthCheckResponseCode() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "HealthCheckResponseCode",
		Type: meta.TypeStringNumber,
		Tags: &schema.FieldTags{
			MapConv: "Settings.GSLB.HealthCheck.Status",
		},
	}
}

func (f *fieldsDef) GSLBHealthCheckPort() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "HealthCheckPort",
		Type: meta.TypeStringNumber,
		Tags: &schema.FieldTags{
			MapConv: "Settings.GSLB.HealthCheck.Port",
		},
	}
}

func (f *fieldsDef) GSLBDelayLoop() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "DelayLoop",
		Type: meta.TypeInt,
		Tags: &schema.FieldTags{
			Validate: "min=10,max=60",
			MapConv:  "Settings.GSLB.DelayLoop,default=10",
		},
	}
}

func (f *fieldsDef) GSLBWeighted() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "Weighted",
		Type: meta.TypeStringFlag,
		Tags: &schema.FieldTags{
			MapConv: "Settings.GSLB.Weighted",
		},
	}
}

func (f *fieldsDef) GSLBDestinationServers() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "DestinationServers",
		Type: &schema.Model{
			Name:    "GSLBServer",
			IsArray: true,
			Fields: []*schema.FieldDesc{
				{
					Name: "IPAddress",
					Type: meta.TypeString,
					Tags: &schema.FieldTags{
						Validate: "ipv4",
					},
				},
				{
					Name: "Enabled",
					Type: meta.TypeStringFlag,
					Tags: &schema.FieldTags{
						MapConv: ",default=true",
					},
				},
				{
					Name: "Weight",
					Type: meta.TypeStringNumber,
					Tags: &schema.FieldTags{
						MapConv: ",default=1",
					},
				},
			},
		},
		Tags: &schema.FieldTags{
			MapConv:  "Settings.GSLB.[]Servers,recursive",
			Validate: "min=0,max=6",
		},
	}
}

func (f *fieldsDef) GSLBSorryServer() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "SorryServer",
		Type: meta.TypeString,
		Tags: &schema.FieldTags{
			MapConv: "Settings.GSLB.SorryServer",
		},
	}
}

func (f *fieldsDef) SettingsHash() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name:     "SettingsHash",
		ReadOnly: true,
		Type:     meta.TypeString,
	}
}

func (f *fieldsDef) InstanceHostName() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name:     "InstanceHostName",
		Type:     meta.TypeString,
		ReadOnly: true,
		Tags: &schema.FieldTags{
			MapConv: "Instance.Host.Name",
		},
	}
}

func (f *fieldsDef) InstanceHostInfoURL() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name:     "InstanceHostInfoURL",
		Type:     meta.TypeString,
		ReadOnly: true,
		Tags: &schema.FieldTags{
			MapConv: "Instance.Host.InfoURL",
		},
	}
}

func (f *fieldsDef) InstanceStatus() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name:     "InstanceStatus",
		Type:     meta.TypeInstanceStatus,
		ReadOnly: true,
		Tags: &schema.FieldTags{
			MapConv: "Instance.Status",
		},
	}
}

func (f *fieldsDef) InstanceStatusChangedAt() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name:     "InstanceStatusChangedAt",
		Type:     meta.TypeTime,
		ReadOnly: true,
		Tags: &schema.FieldTags{
			MapConv: "Instance.StatusChangedAt",
		},
	}
}

func (f *fieldsDef) Interfaces() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name:     "Interfaces",
		Type:     meta.Static([]naked.Interface{}),
		ReadOnly: true,
	}
}

func (f *fieldsDef) NoteClass() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "Class",
		Type: meta.TypeString,
	}
}

func (f *fieldsDef) NoteContent() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "Content",
		Type: meta.TypeString,
	}
}

func (f *fieldsDef) Description() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "Description",
		Type: meta.TypeString,
		Tags: &schema.FieldTags{
			Validate: "min=0,max=512",
		},
	}
}

func (f *fieldsDef) Availability() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "Availability",
		Type: meta.TypeAvailability,
	}
}

func (f *fieldsDef) Scope() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "Scope",
		Type: meta.TypeScope,
	}
}

func (f *fieldsDef) SizeMB() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "SizeMB",
		Type: meta.TypeInt,
		ExtendAccessors: []*schema.ExtendAccessor{
			{Name: "SizeGB"},
		},
	}
}

func (f *fieldsDef) UserSubnetNetworkMaskLen() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "NetworkMaskLen",
		Type: meta.TypeInt,
		Tags: &schema.FieldTags{
			Validate: "min=1,max=32", // TODO
			MapConv:  "UserSubnet.NetworkMaskLen",
		},
	}
}

func (f *fieldsDef) UserSubnetDefaultRoute() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "DefaultRoute",
		Type: meta.TypeString,
		Tags: &schema.FieldTags{
			Validate: "ipv4", // TODO
			MapConv:  "UserSubnet.DefaultRoute",
		},
	}
}

func (f *fieldsDef) RemarkNetworkMaskLen() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "NetworkMaskLen",
		Type: meta.TypeInt,
		Tags: &schema.FieldTags{
			Validate: "min=1,max=32", // TODO
			MapConv:  "Remark.Network.NetworkMaskLen",
		},
	}
}

func (f *fieldsDef) RemarkZoneID() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name:     "ZoneID",
		Type:     meta.TypeID,
		ReadOnly: true,
		Tags: &schema.FieldTags{
			MapConv: "Remark.Zone.ID",
		},
	}
}

func (f *fieldsDef) RemarkDefaultRoute() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "DefaultRoute",
		Type: meta.TypeString,
		Tags: &schema.FieldTags{
			Validate: "ipv4", // TODO
			MapConv:  "Remark.Network.DefaultRoute",
		},
	}
}

func (f *fieldsDef) RemarkServerIPAddress() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "IPAddresses",
		Type: meta.TypeStringSlice,
		Tags: &schema.FieldTags{
			MapConv: "Remark.[]Servers.IPAddress",
		},
	}
}

func (f *fieldsDef) RemarkVRID() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "VRID",
		Type: meta.TypeInt,
		Tags: &schema.FieldTags{
			MapConv: "Remark.VRRP.VRID",
		},
	}
}

func (f *fieldsDef) StorageClass() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "StorageClass",
		Type: meta.TypeString,
	}
}

func (f *fieldsDef) DisplayOrder() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "DisplayOrder",
		Type: meta.TypeInt,
	}
}

func (f *fieldsDef) IsDummy() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "IsDummy",
		Type: meta.TypeFlag,
	}
}

func (f *fieldsDef) HostName() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "HostName",
		Type: meta.TypeString,
	}
}
func (f *fieldsDef) IPAddress() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "IPAddress",
		Type: meta.TypeString,
	}
}
func (f *fieldsDef) User() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "User",
		Type: meta.TypeString,
	}
}
func (f *fieldsDef) Password() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "Password",
		Type: meta.TypeString,
	}
}

func (f *fieldsDef) VNCProxy() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "VNCProxy",
		Type: meta.Static(naked.VNCProxy{}),
		Tags: &schema.FieldTags{
			JSON: ",omitempty",
		},
	}
}

func (f *fieldsDef) FTPServer() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "FTPServer",
		Type: meta.Static(naked.FTPServer{}),
		Tags: &schema.FieldTags{
			JSON: ",omitempty",
		},
	}
}

func (f *fieldsDef) Region() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "Region",
		Type: meta.Static(naked.Region{}),
		Tags: &schema.FieldTags{
			JSON: ",omitempty",
		},
	}
}

func (f *fieldsDef) Zone() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "Zone",
		Type: meta.Static(naked.Zone{}),
		Tags: &schema.FieldTags{
			JSON: ",omitempty",
		},
	}
}

func (f *fieldsDef) UserSubnet() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "UserSubnet",
		Type: meta.Static(naked.UserSubnet{}),
		Tags: &schema.FieldTags{
			JSON: ",omitempty",
		},
	}
}

func (f *fieldsDef) Storage() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "Storage",
		Type: meta.Static(naked.Storage{}),
		Tags: &schema.FieldTags{
			JSON: ",omitempty",
		},
	}
}

func (f *fieldsDef) Icon() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "Icon",
		Type: meta.Static(naked.Icon{}),
		Tags: &schema.FieldTags{
			JSON: ",omitempty",
		},
	}
}

func (f *fieldsDef) Switch() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "Switch",
		Type: meta.Static(naked.Switch{}),
		Tags: &schema.FieldTags{
			JSON: ",omitempty",
		},
	}
}

func (f *fieldsDef) CreatedAt() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "CreatedAt",
		Type: meta.TypeTime,
	}
}

func (f *fieldsDef) ModifiedAt() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "ModifiedAt",
		Type: meta.TypeTime,
	}
}
