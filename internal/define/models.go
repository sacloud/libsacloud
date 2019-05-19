package define

import (
	"github.com/sacloud/libsacloud-v2/internal/schema"
	"github.com/sacloud/libsacloud-v2/internal/schema/meta"
	"github.com/sacloud/libsacloud-v2/sacloud/naked"
)

type modelsDef struct{}

var models = &modelsDef{}

func (m *modelsDef) ftpServerOpenParameter() *schema.Model {
	return &schema.Model{
		Name: "OpenFTPRequest",
		Fields: []*schema.FieldDesc{
			{
				Name: "ChangePassword",
				Type: meta.TypeFlag,
			},
		},
	}
}

func (m *modelsDef) ftpServer() *schema.Model {
	return &schema.Model{
		Name:      "FTPServer",
		NakedType: meta.Static(naked.OpeningFTPServer{}),
		Fields: []*schema.FieldDesc{
			fields.HostName(),
			fields.IPAddress(),
			fields.User(),
			fields.Password(),
		},
	}
}

func (m *modelsDef) ftpServerInfo() *schema.Model {
	return &schema.Model{
		Name:      "FTPServerInfo",
		NakedType: meta.Static(naked.FTPServer{}),
		Fields: []*schema.FieldDesc{
			fields.HostName(),
			fields.IPAddress(),
		},
	}
}

func (m *modelsDef) diskEdit() *schema.Model {

	sshKeyFields := []*schema.FieldDesc{
		{
			Name: "ID",
			Type: meta.TypeID,
			Tags: &schema.FieldTags{
				MapConv: ",omitempty",
				JSON:    ",omitempty",
			},
		},
		{
			Name: "PublicKey",
			Type: meta.TypeString,
			Tags: &schema.FieldTags{
				MapConv: ",omitempty",
				JSON:    ",omitempty",
			},
		},
	}

	noteFields := []*schema.FieldDesc{
		{
			Name: "ID",
			Type: meta.TypeID,
			Tags: &schema.FieldTags{
				MapConv: ",omitempty",
				JSON:    ",omitempty",
			},
		},
		{
			Name: "Variables",
			Type: meta.Static(map[string]interface{}{}),
			Tags: &schema.FieldTags{
				MapConv: ",omitempty",
				JSON:    ",omitempty",
			},
		},
	}

	userSubnetFdields := []*schema.FieldDesc{
		{
			Name: "DefaultRoute",
			Type: meta.TypeString,
			Tags: &schema.FieldTags{
				MapConv: ",omitempty",
				JSON:    ",omitempty",
			},
		},
		{
			Name: "NetworkMaskLen",
			Type: meta.TypeInt,
			Tags: &schema.FieldTags{
				MapConv:  ",omitempty",
				Validate: "min=0,max=32",
				JSON:     ",omitempty",
			},
		},
	}

	return &schema.Model{
		Name:      "DiskEditRequest",
		NakedType: meta.Static(naked.DiskEdit{}),
		Fields: []*schema.FieldDesc{
			{
				Name: "Password",
				Type: meta.TypeString,
				Tags: &schema.FieldTags{
					MapConv: ",omitempty",
					JSON:    ",omitempty",
				},
			},
			{
				Name: "SSHKey",
				Type: &schema.Model{
					Name:   "DiskEditSSHKey",
					Fields: sshKeyFields,
				},
				Tags: &schema.FieldTags{
					MapConv: ",omitempty,recursive",
					JSON:    ",omitempty",
				},
			},
			{
				Name: "SSHKeys",
				Type: &schema.Model{
					Name:    "DiskEditSSHKey",
					IsArray: true,
					Fields:  sshKeyFields,
				},
				Tags: &schema.FieldTags{
					MapConv: "[]SSHKeys,omitempty,recursive",
					JSON:    ",omitempty",
				},
			},
			{
				Name: "DisablePWAuth",
				Type: meta.TypeFlag,
				Tags: &schema.FieldTags{
					MapConv: ",omitempty",
					JSON:    ",omitempty",
				},
			},
			{
				Name: "EnableDHCP",
				Type: meta.TypeFlag,
				Tags: &schema.FieldTags{
					MapConv: ",omitempty",
					JSON:    ",omitempty",
				},
			},
			{
				Name: "ChangePartitionUUID",
				Type: meta.TypeFlag,
				Tags: &schema.FieldTags{
					MapConv: ",omitempty",
					JSON:    ",omitempty",
				},
			},
			{
				Name: "HostName",
				Type: meta.TypeString,
				Tags: &schema.FieldTags{
					MapConv: ",omitempty",
					JSON:    ",omitempty",
				},
			},
			{
				Name: "Notes",
				Type: &schema.Model{
					Name:    "DiskEditNote",
					IsArray: true,
					Fields:  noteFields,
				},
				Tags: &schema.FieldTags{
					MapConv: ",omitempty,recursive",
					JSON:    ",omitempty",
				},
			},
			{
				Name: "UserIPAddress",
				Type: meta.TypeString,
				Tags: &schema.FieldTags{
					MapConv: ",omitempty",
					JSON:    ",omitempty",
				},
			},
			{
				Name: "UserSubnet",
				Type: &schema.Model{
					Name:   "DiskEditUserSubnet",
					Fields: userSubnetFdields,
				},
				Tags: &schema.FieldTags{
					MapConv: ",omitempty",
					JSON:    ",omitempty",
				},
			},
		},
	}
}

func (m *modelsDef) interfaceModel() *schema.Model {
	return &schema.Model{
		Name:      "Interface",
		NakedType: meta.Static(naked.Interface{}),
		IsArray:   true,
		Fields: []*schema.FieldDesc{
			fields.ID(),
			fields.MACAddress(),
			fields.IPAddress(),
			fields.UserIPAddress(),
			fields.HostName(),
			// switch
			{
				Name: "SwitchID",
				Type: meta.TypeID,
				Tags: &schema.FieldTags{
					MapConv: "Switch.ID",
				},
			},
			{
				Name: "SwitchName",
				Type: meta.TypeString,
				Tags: &schema.FieldTags{
					MapConv: "Switch.Name",
				},
			},
			{
				Name: "SwitchScope",
				Type: meta.TypeScope,
				Tags: &schema.FieldTags{
					MapConv: "Switch.Scope",
				},
			},
			{
				Name: "UserSubnetDefaultRoute",
				Type: meta.TypeString,
				Tags: &schema.FieldTags{
					MapConv: "Switch.UserSubnet.DefaultRoute",
				},
			},
			{
				Name: "UserSubnetNetworkMaskLen",
				Type: meta.TypeInt,
				Tags: &schema.FieldTags{
					MapConv: "Switch.UserSubnet.NetworkMaskLen",
				},
			},
			{
				Name: "SubnetDefaultRoute",
				Type: meta.TypeString,
				Tags: &schema.FieldTags{
					MapConv: "Switch.Subnet.DefaultRoute",
				},
			},
			{
				Name: "SubnetNetworkMaskLen",
				Type: meta.TypeInt,
				Tags: &schema.FieldTags{
					MapConv: "Switch.Subnet.NetworkMaskLen",
				},
			},
			{
				Name: "SubnetNetworkAddress",
				Type: meta.TypeString,
				Tags: &schema.FieldTags{
					MapConv: "Switch.Subnet.NetworkAddress",
				},
			},
			{
				Name: "SubnetBandWidthMbps",
				Type: meta.TypeInt,
				Tags: &schema.FieldTags{
					MapConv: "Switch.Subnet.Internet.BandWidthMbps",
				},
			},
			// packet filter
			{
				Name: "PacketFilterID",
				Type: meta.TypeString,
				Tags: &schema.FieldTags{
					MapConv: "PacketFilter.ID",
				},
			},
			{
				Name: "PacketFilterName",
				Type: meta.TypeString,
				Tags: &schema.FieldTags{
					MapConv: "PacketFilter.Name",
				},
			},
			{
				Name: "PacketFilterRequiredHostVersion",
				Type: meta.TypeStringNumber,
				Tags: &schema.FieldTags{
					MapConv: "PacketFilter.RequiredHostVersionn",
				},
			},
		},
	}
}

func (m *modelsDef) bundleInfoModel() *schema.Model {
	return &schema.Model{
		Name:      "BundleInfo",
		NakedType: meta.Static(naked.BundleInfo{}),
		Fields: []*schema.FieldDesc{
			fields.ID(),
			{
				Name: "HostClass",
				Type: meta.TypeString,
				Tags: &schema.FieldTags{
					MapConv: ",omitempty",
					JSON:    ",omitempty",
				},
			},
			{
				Name: "ServiceClass",
				Type: meta.TypeString,
				Tags: &schema.FieldTags{
					MapConv: ",omitempty",
					JSON:    ",omitempty",
				},
			},
		},
	}
}

func (m *modelsDef) storageModel() *schema.Model {
	return &schema.Model{
		Name:      "Storage",
		NakedType: meta.Static(naked.Storage{}),
		Fields: []*schema.FieldDesc{
			fields.ID(),
			fields.Name(),
			{
				Name: "Class",
				Type: meta.TypeString,
				Tags: &schema.FieldTags{
					MapConv: ",omitempty",
					JSON:    ",omitempty",
				},
			},
			{
				Name: "Generation",
				Type: meta.TypeInt,
				Tags: &schema.FieldTags{
					MapConv: ",omitempty",
					JSON:    ",omitempty",
				},
			},
		},
	}
}

func (m *modelsDef) region() *schema.Model {
	return &schema.Model{
		Name:      "Region",
		NakedType: meta.Static(naked.Region{}),
		Fields: []*schema.FieldDesc{
			fields.ID(),
			fields.Name(),
			fields.Description(),
			{
				Name: "NameServers",
				Type: meta.TypeStringSlice,
				Tags: &schema.FieldTags{
					MapConv: ",omitempty",
					JSON:    ",omitempty",
				},
			},
		},
	}
}

func (m *modelsDef) zoneInfoModel() *schema.Model {
	return &schema.Model{
		Name:      "ZoneInfo",
		NakedType: meta.Static(naked.Zone{}),
		Fields: []*schema.FieldDesc{
			fields.ID(),
			fields.Name(),
			{
				Name: "DisplayName",
				Type: meta.TypeString,
				Tags: &schema.FieldTags{
					MapConv: "Description,omitempty",
					JSON:    ",omitempty",
				},
			},
			{
				Name: "IsDummy",
				Type: meta.TypeFlag,
				Tags: &schema.FieldTags{
					MapConv: ",omitempty",
					JSON:    ",omitempty",
				},
			},
			{
				Name: "VNCProxy",
				Type: m.vncProxyModel(),
				Tags: &schema.FieldTags{
					MapConv: ",omitempty,recursive",
					JSON:    ",omitempty",
				},
			},
			{
				Name: "FTPServer",
				Type: m.ftpServerInfo(),
				Tags: &schema.FieldTags{
					MapConv: ",omitempty,recursive",
					JSON:    ",omitempty",
				},
			},
			{
				Name: "Region",
				Type: m.region(),
				Tags: &schema.FieldTags{
					MapConv: ",omitempty,recursive",
					JSON:    ",omitempty",
				},
			},
		},
	}
}

func (m *modelsDef) vncProxyModel() *schema.Model {
	return &schema.Model{
		Name:      "VNCProxy",
		NakedType: meta.Static(naked.VNCProxy{}),
		Fields: []*schema.FieldDesc{
			{
				Name: "HostName",
				Type: meta.TypeString,
				Tags: &schema.FieldTags{
					MapConv: ",omitempty",
					JSON:    ",omitempty",
				},
			},
			{
				Name: "IPAddress",
				Type: meta.TypeString,
				Tags: &schema.FieldTags{
					MapConv: ",omitempty",
					JSON:    ",omitempty",
				},
			},
		},
	}
}

func (m *modelsDef) sourceArchiveInfo() *schema.Model {
	return &schema.Model{
		Name: "SourceArchiveInfo",
		Fields: []*schema.FieldDesc{
			{
				Name: "ID",
				Type: meta.TypeID,
				Tags: &schema.FieldTags{
					MapConv: "ArchiveUnderZone.ID",
				},
			},
			{
				Name: "AccountID",
				Type: meta.TypeID,
				Tags: &schema.FieldTags{
					MapConv: "ArchiveUnderZone.Account.ID",
				},
			},
			{
				Name: "ZoneID",
				Type: meta.TypeID,
				Tags: &schema.FieldTags{
					MapConv: "ArchiveUnderZone.Zone.ID",
				},
			},
			{
				Name: "ZoneName",
				Type: meta.TypeString,
				Tags: &schema.FieldTags{
					MapConv: "ArchiveUnderZone.Zone.Name",
				},
			},
		},
	}
}

func (m *modelsDef) packetFilterExpressions() *schema.Model {
	return &schema.Model{
		Name:      "PacketFilterExpression",
		NakedType: meta.Static(naked.PacketFilterExpression{}),
		IsArray:   true,
		Fields: []*schema.FieldDesc{
			{
				Name: "Protocol",
				Type: meta.TypeProtocol,
			},
			{
				Name: "SourceNetwork",
				Type: meta.TypePacketFilterNetwork,
			},
			{
				Name: "SourcePort",
				Type: meta.TypePacketFilterPort,
			},
			{
				Name: "DestinationPort",
				Type: meta.TypePacketFilterPort,
			},
			{
				Name: "Action",
				Type: meta.TypePacketFilterAction,
			},
		},
	}
}
