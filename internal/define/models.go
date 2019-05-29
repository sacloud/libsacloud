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

func (m *modelsDef) vpcRouterInterfaceModel() *schema.Model {
	ifModel := m.interfaceModel()
	ifModel.Name = "VPCRouterInterface"
	ifModel.Fields = append(ifModel.Fields, &schema.FieldDesc{
		Name: "Index",
		Type: meta.TypeInt,
		Tags: &schema.FieldTags{
			MapConv: ",omitempty",
		},
	})
	return ifModel
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
				Type: meta.TypeAction,
			},
		},
	}
}
func (m *modelsDef) bridgeInfoModel() *schema.Model {
	return &schema.Model{
		Name:      "BridgeInfo",
		IsArray:   true,
		NakedType: meta.Static(naked.Switch{}),
		Fields: []*schema.FieldDesc{
			fields.ID(),
			fields.Name(),
			fields.ZoneID(),
		},
	}
}

func (m *modelsDef) switchInZoneModel() *schema.Model {
	return &schema.Model{
		Name:      "BridgeSwitchInfo",
		NakedType: meta.Static(naked.BridgeSwitchInfo{}),
		Fields: []*schema.FieldDesc{
			fields.ID(),
			fields.Name(),
			fields.Scope(),
			{
				Name: "ServerCount",
				Type: meta.TypeInt,
			},
			{
				Name: "ApplianceCount",
				Type: meta.TypeInt,
			},
		},
	}
}

func (m *modelsDef) internetModel() *schema.Model {
	return &schema.Model{
		Name:      "Internet",
		NakedType: meta.Static(naked.Internet{}),
		Fields: []*schema.FieldDesc{
			fields.ID(),
			fields.Name(),
			fields.Description(),
			fields.Tags(),
			fields.IconID(),
			fields.CreatedAt(),
			fields.BandWidthMbps(),
			fields.NetworkMaskLen(),
			{
				Name: "Switch",
				Type: m.switchInfoModel(),
				Tags: &schema.FieldTags{
					MapConv: ",recursive",
				},
			},
		},
	}
}

// switchInfoModel Internetリソースのフィールドとしてのswitch
//
// Subnetの情報は限定的にしか返ってこない(IPAddresses.Max/Minなどがない)ため注意
// 必要であればSwitchリソース配下のSubnetsを参照すればOK
func (m *modelsDef) switchInfoModel() *schema.Model {
	return &schema.Model{
		Name:      "SwitchInfo",
		NakedType: meta.Static(naked.Switch{}),
		Fields: []*schema.FieldDesc{
			fields.ID(),
			fields.Name(),
			fields.Description(),
			fields.Tags(),
			fields.Scope(),
			{
				Name: "Subnets",
				Type: m.internetSubnet(),
				Tags: &schema.FieldTags{
					MapConv: "[]Subnets,recursive",
				},
			},
		},
	}
}

func (m *modelsDef) internetSubnet() *schema.Model {
	return &schema.Model{
		Name:      "InternetSubnet",
		NakedType: meta.Static(naked.Subnet{}),
		IsArray:   true,
		Fields: []*schema.FieldDesc{
			fields.ID(),
			fields.DefaultRoute(),
			fields.NextHop(),
			fields.StaticRoute(),
			fields.NetworkAddress(),
			fields.NetworkMaskLen(),
		},
	}
}

// internetSubnetOperationResult Internetリソースへのサブネット追加/更新時の戻り値
//
// internetSubnetに対しIPAddresses(文字列配列)を追加したもの
func (m *modelsDef) internetSubnetOperationResult() *schema.Model {
	return &schema.Model{
		Name:      "InternetSubnetOperationResult",
		NakedType: meta.Static(naked.Subnet{}),
		Fields: []*schema.FieldDesc{
			fields.ID(),
			fields.DefaultRoute(),
			fields.NextHop(),
			fields.StaticRoute(),
			fields.NetworkAddress(),
			fields.NetworkMaskLen(),
			{
				Name: "IPAddresses",
				Type: meta.TypeStringSlice,
				Tags: &schema.FieldTags{
					MapConv: "[]IPAddresses.IPAddress",
				},
			},
		},
	}
}

func (m *modelsDef) switchSubnet() *schema.Model {
	// switchSubnetはinternetSubnetにInternetとIPAddressesを追加したもの
	subnet := m.internetSubnet()
	subnet.Name = "SwitchSubnet"
	subnet.Fields = append(subnet.Fields,
		&schema.FieldDesc{
			Name: "Internet",
			Type: m.internetModel(),
		},
		&schema.FieldDesc{
			Name: "AssignedIPAddressMax",
			Type: meta.TypeString,
			Tags: &schema.FieldTags{
				MapConv: "IPAddresses.Max",
			},
		},
		&schema.FieldDesc{
			Name: "AssignedIPAddressMin",
			Type: meta.TypeString,
			Tags: &schema.FieldTags{
				MapConv: "IPAddresses.Min",
			},
		},
	)
	subnet.Accessors = []*schema.Accessor{
		{
			Name:             "GetAssignedIPAddresses",
			Description:      "割り当てられたIPアドレスのリスト",
			AccessorTypeName: "AssignedIPAddress",
			ResultType:       meta.TypeStringSlice,
		},
	}
	return subnet
}

func (m *modelsDef) vpcRouterSetting() *schema.Model {
	return &schema.Model{
		Name:      "VPCRouterSetting",
		NakedType: meta.Static(naked.VPCRouterSettings{}),
		Fields: []*schema.FieldDesc{
			{
				Name: "VRID",
				Type: meta.TypeInt,
				Tags: &schema.FieldTags{
					JSON:    ",omitempty",
					MapConv: "Router.VRID",
				},
			},
			{
				Name: "InternetConnectionEnabled",
				Type: meta.TypeStringFlag,
				Tags: &schema.FieldTags{
					JSON:    ",omitempty",
					MapConv: "Router.InternetConnection.Enabled,omitempty",
				},
			},
			{
				Name: "Interfaces",
				Type: m.vpcRouterInterface(),
				Tags: &schema.FieldTags{
					JSON:    ",omitempty",
					MapConv: "Router.[]Interface,omitempty,recursive",
				},
			},
			{
				Name: "StaticNAT",
				Type: m.vpcRouterStaticNAT(),
				Tags: &schema.FieldTags{
					JSON:    ",omitempty",
					MapConv: "Router.StaticNAT.[]Config,omitempty,recursive",
				},
			},
			{
				Name: "Firewall",
				Type: m.vpcRouterFirewall(),
				Tags: &schema.FieldTags{
					JSON:    ",omitempty",
					MapConv: "Router.Firewall.[]Config,omitempty,recursive",
				},
			},
			{
				Name: "DHCPServer",
				Type: m.vpcRouterDHCPServer(),
				Tags: &schema.FieldTags{
					JSON:    ",omitempty",
					MapConv: "Router.DHCPServer.[]Config,omitempty,recursive",
				},
			},
			{
				Name: "DHCPStaticMapping",
				Type: m.vpcRouterDHCPStaticMapping(),
				Tags: &schema.FieldTags{
					JSON:    ",omitempty",
					MapConv: "Router.DHCPStaticMapping.[]Config,omitempty,recursive",
				},
			},
			{
				Name: "PPTPServer",
				Type: m.vpcRouterPPTPServer(),
				Tags: &schema.FieldTags{
					JSON:    ",omitempty",
					MapConv: "Router.PPTPServer.Config,omitempty,recursive",
				},
			},
			{
				Name: "PPTPServerEnabled",
				Type: meta.TypeStringFlag,
				Tags: &schema.FieldTags{
					JSON:    ",omitempty",
					MapConv: "Router.PPTPServer.Enabled,omitempty",
				},
			},
			{
				Name: "L2TPIPsecServer",
				Type: m.vpcRouterL2TPIPsecServer(),
				Tags: &schema.FieldTags{
					JSON:    ",omitempty",
					MapConv: "Router.L2TPIPsecServer.Config,omitempty,recursive",
				},
			},
			{
				Name: "L2TPIPsecServerEnabled",
				Type: meta.TypeStringFlag,
				Tags: &schema.FieldTags{
					JSON:    ",omitempty",
					MapConv: "Router.L2TPIPsecServer.Enabled,omitempty",
				},
			},
			{
				Name: "RemoteAccessUsers",
				Type: m.vpcRouterRemoteAccessUser(),
				Tags: &schema.FieldTags{
					JSON:    ",omitempty",
					MapConv: "Router.RemoteAccessUsers.[]Config,omitempty,recursive",
				},
			},
			{
				Name: "SiteToSiteIPsecVPN",
				Type: m.vpcRouterSiteToSiteIPsecVPN(),
				Tags: &schema.FieldTags{
					JSON:    ",omitempty",
					MapConv: "Router.SiteToSiteIPsecVPN.[]Config,omitempty,recursive",
				},
			},
			{
				Name: "StaticRoute",
				Type: m.vpcRouterStaticRoute(),
				Tags: &schema.FieldTags{
					JSON:    ",omitempty",
					MapConv: "Router.StaticRoutes.[]Config,omitempty,recursive",
				},
			},
		},
	}
}

func (m *modelsDef) vpcRouterInterface() *schema.Model {
	return &schema.Model{
		Name:      "VPCRouterInterfaceSetting",
		NakedType: meta.Static(naked.VPCRouterInterface{}),
		IsArray:   true,
		Fields: []*schema.FieldDesc{
			fields.stringEnabled(),
			{
				Name: "IPAddress",
				Type: meta.TypeStringSlice,
			},
			{
				Name: "VirtualIPAddress",
				Type: meta.TypeString,
			},
			{
				Name: "IPAliases",
				Type: meta.TypeStringSlice,
			},
			{
				Name: "NetworkMaskLen",
				Type: meta.TypeInt,
			},
			{
				Name: "Index",
				Type: meta.TypeInt,
			},
		},
	}
}

func (m *modelsDef) vpcRouterStaticNAT() *schema.Model {
	return &schema.Model{
		Name:      "VPCRouterStaticNAT",
		NakedType: meta.Static(naked.VPCRouterStaticNATConfig{}),
		IsArray:   true,
		Fields: []*schema.FieldDesc{
			{
				Name: "GlobalAddress",
				Type: meta.TypeString,
				Tags: &schema.FieldTags{
					MapConv:  "GlobalAddress",
					Validate: "ipv4",
				},
			},
			{
				Name: "PrivateAddress",
				Type: meta.TypeString,
				Tags: &schema.FieldTags{
					MapConv:  "PrivateAddress",
					Validate: "ipv4",
				},
			},
			{
				Name: "Description",
				Type: meta.TypeString,
			},
		},
	}
}

func (m *modelsDef) vpcRouterFirewall() *schema.Model {
	return &schema.Model{
		Name:      "VPCRouterFirewall",
		NakedType: meta.Static(naked.VPCRouterFirewallConfig{}),
		IsArray:   true,
		Fields: []*schema.FieldDesc{
			{
				Name: "Send",
				Type: m.vpcRouterFirewallRule(),
			},
			{
				Name: "Receive",
				Type: m.vpcRouterFirewallRule(),
			},
		},
	}
}
func (m *modelsDef) vpcRouterFirewallRule() *schema.Model {
	return &schema.Model{
		Name:      "VPCRouterFirewallRule",
		NakedType: meta.Static(naked.VPCRouterFirewallRule{}),
		IsArray:   true,
		Fields: []*schema.FieldDesc{
			{
				Name: "Protocol",
				Type: meta.TypeProtocol,
			},
			{
				Name: "SourceNetwork",
				Type: meta.TypeVPCFirewallNetwork,
			},
			{
				Name: "SourcePort",
				Type: meta.TypeVPCFirewallPort,
			},
			{
				Name: "DestinationNetwork",
				Type: meta.TypeVPCFirewallNetwork,
			},
			{
				Name: "DestinationPort",
				Type: meta.TypeVPCFirewallPort,
			},
			{
				Name: "Action",
				Type: meta.TypeAction,
			},
			{
				Name: "Logging",
				Type: meta.TypeStringFlag,
			},
			{
				Name: "Description",
				Type: meta.TypeString,
			},
		},
	}
}

func (m *modelsDef) vpcRouterDHCPServer() *schema.Model {
	return &schema.Model{
		Name:      "VPCRouterDHCPServer",
		NakedType: meta.Static(naked.VPCRouterDHCPServerConfig{}),
		IsArray:   true,
		Fields: []*schema.FieldDesc{
			{
				Name: "Interface",
				Type: meta.TypeString,
			},
			{
				Name: "RangeStart",
				Type: meta.TypeString,
				Tags: &schema.FieldTags{
					Validate: "ipv4",
				},
			},
			{
				Name: "RangeStop",
				Type: meta.TypeString,
				Tags: &schema.FieldTags{
					Validate: "ipv4",
				},
			},
			{
				Name: "DNSServers",
				Type: meta.TypeStringSlice,
				Tags: &schema.FieldTags{
					Validate: "dive,ipv4",
				},
			},
		},
	}
}

func (m *modelsDef) vpcRouterDHCPStaticMapping() *schema.Model {
	return &schema.Model{
		Name:      "VPCRouterDHCPStaticMapping",
		NakedType: meta.Static(naked.VPCRouterDHCPStaticMappingConfig{}),
		IsArray:   true,
		Fields: []*schema.FieldDesc{
			{
				Name: "MACAddress", // TODO typesに独自型作っておくべきか?
				Type: meta.TypeString,
			},
			{
				Name: "IPAddress",
				Type: meta.TypeString,
				Tags: &schema.FieldTags{
					Validate: "ipv4",
				},
			},
		},
	}
}

func (m *modelsDef) vpcRouterPPTPServer() *schema.Model {
	return &schema.Model{
		Name:      "VPCRouterPPTPServer",
		NakedType: meta.Static(naked.VPCRouterPPTPServerConfig{}),
		Fields: []*schema.FieldDesc{
			{
				Name: "RangeStart",
				Type: meta.TypeString,
				Tags: &schema.FieldTags{
					Validate: "ipv4",
				},
			},
			{
				Name: "RangeStop",
				Type: meta.TypeString,
				Tags: &schema.FieldTags{
					Validate: "ipv4",
				},
			},
		},
	}
}

func (m *modelsDef) vpcRouterL2TPIPsecServer() *schema.Model {
	return &schema.Model{
		Name:      "VPCRouterL2TPIPsecServer",
		NakedType: meta.Static(naked.VPCRouterL2TPIPsecServerConfig{}),
		Fields: []*schema.FieldDesc{
			{
				Name: "RangeStart",
				Type: meta.TypeString,
				Tags: &schema.FieldTags{
					Validate: "ipv4",
				},
			},
			{
				Name: "RangeStop",
				Type: meta.TypeString,
				Tags: &schema.FieldTags{
					Validate: "ipv4",
				},
			},
			{
				Name: "PreSharedSecret",
				Type: meta.TypeString,
			},
		},
	}
}

func (m *modelsDef) vpcRouterRemoteAccessUser() *schema.Model {
	return &schema.Model{
		Name:      "VPCRouterRemoteAccessUser",
		NakedType: meta.Static(naked.VPCRouterRemoteAccessUserConfig{}),
		IsArray:   true,
		Fields: []*schema.FieldDesc{
			{
				Name: "UserName",
				Type: meta.TypeString,
			},
			{
				Name: "Password",
				Type: meta.TypeString,
			},
		},
	}
}

func (m *modelsDef) vpcRouterSiteToSiteIPsecVPN() *schema.Model {
	return &schema.Model{
		Name:      "VPCRouterSiteToSiteIPsecVPN",
		NakedType: meta.Static(naked.VPCRouterSiteToSiteIPsecVPNConfig{}),
		IsArray:   true,
		Fields: []*schema.FieldDesc{
			{
				Name: "Peer",
				Type: meta.TypeString,
			},
			{
				Name: "PreSharedSecret",
				Type: meta.TypeString,
			},
			{
				Name: "RemoteID",
				Type: meta.TypeString,
			},
			{
				Name: "Routes",
				Type: meta.TypeStringSlice,
			},
			{
				Name: "LocalPrefix",
				Type: meta.TypeStringSlice,
			},
		},
	}
}

func (m *modelsDef) vpcRouterStaticRoute() *schema.Model {
	return &schema.Model{
		Name:      "VPCRouterStaticRoute",
		NakedType: meta.Static(naked.VPCRouterStaticRouteConfig{}),
		IsArray:   true,
		Fields: []*schema.FieldDesc{
			{
				Name: "Prefix",
				Type: meta.TypeString,
			},
			{
				Name: "NextHop",
				Type: meta.TypeString,
			},
		},
	}
}
