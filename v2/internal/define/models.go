// Copyright 2016-2020 The Libsacloud Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package define

import (
	"github.com/sacloud/libsacloud/v2/internal/dsl"
	"github.com/sacloud/libsacloud/v2/internal/dsl/meta"
	"github.com/sacloud/libsacloud/v2/sacloud/naked"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

type modelsDef struct{}

var models = &modelsDef{}

func (m *modelsDef) ftpServerOpenParameter() *dsl.Model {
	return &dsl.Model{
		Name: "OpenFTPRequest",
		Fields: []*dsl.FieldDesc{
			fields.FTPServerChangePassword(),
		},
	}
}

func (m *modelsDef) ftpServer() *dsl.Model {
	return &dsl.Model{
		Name:      "FTPServer",
		NakedType: meta.Static(naked.OpeningFTPServer{}),
		Fields: []*dsl.FieldDesc{
			fields.HostName(),
			fields.IPAddress(),
			fields.User(),
			fields.Password(),
		},
	}
}

func (m *modelsDef) ftpServerInfo() *dsl.Model {
	return &dsl.Model{
		Name:      "FTPServerInfo",
		NakedType: meta.Static(naked.FTPServer{}),
		Fields: []*dsl.FieldDesc{
			fields.HostName(),
			fields.IPAddress(),
		},
	}
}

func (m *modelsDef) diskEdit() *dsl.Model {
	sshKeyFields := []*dsl.FieldDesc{
		{
			Name: "ID",
			Type: meta.TypeID,
			Tags: &dsl.FieldTags{
				MapConv: ",omitempty",
				JSON:    ",omitempty",
			},
		},
		{
			Name: "PublicKey",
			Type: meta.TypeString,
			Tags: &dsl.FieldTags{
				MapConv: ",omitempty",
				JSON:    ",omitempty",
			},
		},
	}

	noteFields := []*dsl.FieldDesc{
		{
			Name: "ID",
			Type: meta.TypeID,
			Tags: &dsl.FieldTags{
				MapConv: ",omitempty",
				JSON:    ",omitempty",
			},
		},
		{
			Name: "APIKeyID",
			Type: meta.TypeID,
			Tags: &dsl.FieldTags{
				MapConv: "APIKey.ID,omitempty",
				JSON:    ",omitempty",
			},
		},
		{
			Name: "Variables",
			Type: meta.Static(map[string]interface{}{}),
			Tags: &dsl.FieldTags{
				MapConv: ",omitempty",
				JSON:    ",omitempty",
			},
		},
	}

	userSubnetFields := []*dsl.FieldDesc{
		{
			Name: "DefaultRoute",
			Type: meta.TypeString,
			Tags: &dsl.FieldTags{
				MapConv: ",omitempty",
				JSON:    ",omitempty",
			},
		},
		{
			Name: "NetworkMaskLen",
			Type: meta.TypeInt,
			Tags: &dsl.FieldTags{
				MapConv: ",omitempty",
				JSON:    ",omitempty",
			},
		},
	}

	return &dsl.Model{
		Name:      "DiskEditRequest",
		NakedType: meta.Static(naked.DiskEdit{}),
		Fields: []*dsl.FieldDesc{
			{
				Name: "Background",
				Type: meta.TypeFlag,
				Tags: &dsl.FieldTags{
					MapConv: ",omitempty",
					JSON:    ",omitempty",
				},
			},
			{
				Name: "Password",
				Type: meta.TypeString,
				Tags: &dsl.FieldTags{
					MapConv: ",omitempty",
					JSON:    ",omitempty",
				},
			},
			{
				Name: "SSHKey",
				Type: &dsl.Model{
					Name:   "DiskEditSSHKey",
					Fields: sshKeyFields,
				},
				Tags: &dsl.FieldTags{
					MapConv: ",omitempty,recursive",
					JSON:    ",omitempty",
				},
			},
			{
				Name: "SSHKeys",
				Type: &dsl.Model{
					Name:    "DiskEditSSHKey",
					IsArray: true,
					Fields:  sshKeyFields,
				},
				Tags: &dsl.FieldTags{
					MapConv: "[]SSHKeys,omitempty,recursive",
					JSON:    ",omitempty",
				},
			},
			{
				Name: "DisablePWAuth",
				Type: meta.TypeFlag,
				Tags: &dsl.FieldTags{
					MapConv: ",omitempty",
					JSON:    ",omitempty",
				},
			},
			{
				Name: "EnableDHCP",
				Type: meta.TypeFlag,
				Tags: &dsl.FieldTags{
					MapConv: ",omitempty",
					JSON:    ",omitempty",
				},
			},
			{
				Name: "ChangePartitionUUID",
				Type: meta.TypeFlag,
				Tags: &dsl.FieldTags{
					MapConv: ",omitempty",
					JSON:    ",omitempty",
				},
			},
			{
				Name: "HostName",
				Type: meta.TypeString,
				Tags: &dsl.FieldTags{
					MapConv: ",omitempty",
					JSON:    ",omitempty",
				},
			},
			{
				Name: "Notes",
				Type: &dsl.Model{
					Name:      "DiskEditNote",
					NakedType: meta.Static([]*naked.DiskEditNote{}),
					IsArray:   true,
					Fields:    noteFields,
				},
				Tags: &dsl.FieldTags{
					MapConv: "[]Notes,omitempty,recursive",
					JSON:    ",omitempty",
				},
			},
			{
				Name: "UserIPAddress",
				Type: meta.TypeString,
				Tags: &dsl.FieldTags{
					MapConv: ",omitempty",
					JSON:    ",omitempty",
				},
			},
			{
				Name: "UserSubnet",
				Type: &dsl.Model{
					Name:   "DiskEditUserSubnet",
					Fields: userSubnetFields,
				},
				Tags: &dsl.FieldTags{
					MapConv: ",omitempty",
					JSON:    ",omitempty",
				},
			},
		},
	}
}

func (m *modelsDef) interfaceModel() *dsl.Model {
	return &dsl.Model{
		Name:      "InterfaceView",
		NakedType: meta.Static(naked.Interface{}),
		IsArray:   true,
		Fields: []*dsl.FieldDesc{
			fields.ID(),
			fields.MACAddress(),
			fields.IPAddress(),
			fields.UserIPAddress(),
			fields.HostName(),
			// switch
			{
				Name: "SwitchID",
				Type: meta.TypeID,
				Tags: &dsl.FieldTags{
					MapConv: "Switch.ID",
				},
			},
			{
				Name: "SwitchName",
				Type: meta.TypeString,
				Tags: &dsl.FieldTags{
					MapConv: "Switch.Name",
				},
			},
			{
				Name: "SwitchScope",
				Type: meta.TypeScope,
				Tags: &dsl.FieldTags{
					MapConv: "Switch.Scope",
				},
			},
			{
				Name: "UserSubnetDefaultRoute",
				Type: meta.TypeString,
				Tags: &dsl.FieldTags{
					MapConv: "Switch.UserSubnet.DefaultRoute",
				},
			},
			{
				Name: "UserSubnetNetworkMaskLen",
				Type: meta.TypeInt,
				Tags: &dsl.FieldTags{
					MapConv: "Switch.UserSubnet.NetworkMaskLen",
				},
			},
			{
				Name: "SubnetDefaultRoute",
				Type: meta.TypeString,
				Tags: &dsl.FieldTags{
					MapConv: "Switch.Subnet.DefaultRoute",
				},
			},
			{
				Name: "SubnetNetworkMaskLen",
				Type: meta.TypeInt,
				Tags: &dsl.FieldTags{
					MapConv: "Switch.Subnet.NetworkMaskLen",
				},
			},
			{
				Name: "SubnetNetworkAddress",
				Type: meta.TypeString,
				Tags: &dsl.FieldTags{
					MapConv: "Switch.Subnet.NetworkAddress",
				},
			},
			{
				Name: "SubnetBandWidthMbps",
				Type: meta.TypeInt,
				Tags: &dsl.FieldTags{
					MapConv: "Switch.Subnet.Internet.BandWidthMbps",
				},
			},
			// packet filter
			{
				Name: "PacketFilterID",
				Type: meta.TypeID,
				Tags: &dsl.FieldTags{
					MapConv: "PacketFilter.ID",
				},
			},
			{
				Name: "PacketFilterName",
				Type: meta.TypeString,
				Tags: &dsl.FieldTags{
					MapConv: "PacketFilter.Name",
				},
			},
			{
				Name: "PacketFilterRequiredHostVersion",
				Type: meta.TypeStringNumber,
				Tags: &dsl.FieldTags{
					MapConv: "PacketFilter.RequiredHostVersionn",
				},
			},
			{
				Name: "UpstreamType",
				Type: meta.Static(types.EUpstreamNetworkType("")),
			},
		},
	}
}

func (m *modelsDef) vpcRouterInterfaceModel() *dsl.Model {
	ifModel := m.interfaceModel()
	ifModel.Name = "VPCRouterInterface"
	ifModel.Fields = append(ifModel.Fields, &dsl.FieldDesc{
		Name: "Index",
		Type: meta.TypeInt,
		Tags: &dsl.FieldTags{
			MapConv: ",omitempty",
		},
	})
	return ifModel
}

func (m *modelsDef) mobileGatewayInterfaceModel() *dsl.Model {
	ifModel := m.interfaceModel()
	ifModel.Name = "MobileGatewayInterface"
	ifModel.Fields = append(ifModel.Fields, &dsl.FieldDesc{
		Name: "Index",
		Type: meta.TypeInt,
		Tags: &dsl.FieldTags{
			MapConv: ",omitempty",
		},
	})
	return ifModel
}

func (m *modelsDef) bundleInfoModel() *dsl.Model {
	return &dsl.Model{
		Name:      "BundleInfo",
		NakedType: meta.Static(naked.BundleInfo{}),
		Fields: []*dsl.FieldDesc{
			fields.ID(),
			{
				Name: "HostClass",
				Type: meta.TypeString,
				Tags: &dsl.FieldTags{
					MapConv: ",omitempty",
					JSON:    ",omitempty",
				},
			},
			{
				Name: "ServiceClass",
				Type: meta.TypeString,
				Tags: &dsl.FieldTags{
					MapConv: ",omitempty",
					JSON:    ",omitempty",
				},
			},
		},
	}
}

func (m *modelsDef) storageModel() *dsl.Model {
	return &dsl.Model{
		Name:      "Storage",
		NakedType: meta.Static(naked.Storage{}),
		Fields: []*dsl.FieldDesc{
			fields.ID(),
			fields.Name(),
			{
				Name: "Class",
				Type: meta.TypeString,
				Tags: &dsl.FieldTags{
					MapConv: ",omitempty",
					JSON:    ",omitempty",
				},
			},
			{
				Name: "Generation",
				Type: meta.TypeInt,
				Tags: &dsl.FieldTags{
					MapConv: ",omitempty",
					JSON:    ",omitempty",
				},
			},
		},
	}
}

func (m *modelsDef) region() *dsl.Model {
	return &dsl.Model{
		Name:      "Region",
		NakedType: meta.Static(naked.Region{}),
		Fields: []*dsl.FieldDesc{
			fields.ID(),
			fields.Name(),
			fields.Description(),
			fields.NameServers(),
		},
	}
}

func (m *modelsDef) zoneInfoModel() *dsl.Model {
	return &dsl.Model{
		Name:      "ZoneInfo",
		NakedType: meta.Static(naked.Zone{}),
		Fields: []*dsl.FieldDesc{
			fields.ID(),
			fields.Name(),
			{
				Name: "DisplayName",
				Type: meta.TypeString,
				Tags: &dsl.FieldTags{
					MapConv: "Description,omitempty",
					JSON:    ",omitempty",
				},
			},
			{
				Name: "IsDummy",
				Type: meta.TypeFlag,
				Tags: &dsl.FieldTags{
					MapConv: ",omitempty",
					JSON:    ",omitempty",
				},
			},
			{
				Name: "VNCProxy",
				Type: m.vncProxyModel(),
				Tags: &dsl.FieldTags{
					MapConv: ",omitempty,recursive",
					JSON:    ",omitempty",
				},
			},
			{
				Name: "FTPServer",
				Type: m.ftpServerInfo(),
				Tags: &dsl.FieldTags{
					MapConv: ",omitempty,recursive",
					JSON:    ",omitempty",
				},
			},
			{
				Name: "Region",
				Type: m.region(),
				Tags: &dsl.FieldTags{
					MapConv: ",omitempty,recursive",
					JSON:    ",omitempty",
				},
			},
		},
	}
}

func (m *modelsDef) vncProxyModel() *dsl.Model {
	return &dsl.Model{
		Name:      "VNCProxy",
		NakedType: meta.Static(naked.VNCProxy{}),
		Fields: []*dsl.FieldDesc{
			{
				Name: "HostName",
				Type: meta.TypeString,
				Tags: &dsl.FieldTags{
					MapConv: ",omitempty",
					JSON:    ",omitempty",
				},
			},
			{
				Name: "IPAddress",
				Type: meta.TypeString,
				Tags: &dsl.FieldTags{
					MapConv: ",omitempty",
					JSON:    ",omitempty",
				},
			},
		},
	}
}

func (m *modelsDef) sourceArchiveInfo() *dsl.Model {
	return &dsl.Model{
		Name: "SourceArchiveInfo",
		Fields: []*dsl.FieldDesc{
			{
				Name: "ID",
				Type: meta.TypeID,
				Tags: &dsl.FieldTags{
					MapConv: "ArchiveUnderZone.ID",
				},
			},
			{
				Name: "AccountID",
				Type: meta.TypeID,
				Tags: &dsl.FieldTags{
					MapConv: "ArchiveUnderZone.Account.ID",
				},
			},
			{
				Name: "ZoneID",
				Type: meta.TypeID,
				Tags: &dsl.FieldTags{
					MapConv: "ArchiveUnderZone.Zone.ID",
				},
			},
			{
				Name: "ZoneName",
				Type: meta.TypeString,
				Tags: &dsl.FieldTags{
					MapConv: "ArchiveUnderZone.Zone.Name",
				},
			},
		},
	}
}

func (m *modelsDef) packetFilterExpressions() *dsl.Model {
	return &dsl.Model{
		Name:      "PacketFilterExpression",
		NakedType: meta.Static(naked.PacketFilterExpression{}),
		IsArray:   true,
		Fields: []*dsl.FieldDesc{
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
			{
				Name: "Description",
				Type: meta.TypeString,
			},
		},
	}
}
func (m *modelsDef) bridgeInfoModel() *dsl.Model {
	return &dsl.Model{
		Name:      "BridgeInfo",
		IsArray:   true,
		NakedType: meta.Static(naked.Switch{}),
		Fields: []*dsl.FieldDesc{
			fields.ID(),
			fields.Name(),
			fields.ZoneID(),
			fields.ZoneName(),
		},
	}
}

func (m *modelsDef) switchInZoneModel() *dsl.Model {
	return &dsl.Model{
		Name:      "BridgeSwitchInfo",
		NakedType: meta.Static(naked.BridgeSwitchInfo{}),
		Fields: []*dsl.FieldDesc{
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

func (m *modelsDef) internetModel() *dsl.Model {
	return &dsl.Model{
		Name:      "Internet",
		NakedType: meta.Static(naked.Internet{}),
		Fields: []*dsl.FieldDesc{
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
				Tags: &dsl.FieldTags{
					MapConv: ",recursive",
				},
			},
		},
	}
}

// switchIPv6NetModel InternetリソースのフィールドとしてのIPv6Net
func (m *modelsDef) switchIPv6NetModel() *dsl.Model {
	return &dsl.Model{
		Name:      "IPv6NetInfo",
		NakedType: meta.Static(naked.IPv6Net{}),
		IsArray:   false,
		Fields: []*dsl.FieldDesc{
			fields.ID(),
			fields.Def("IPv6Prefix", meta.TypeString),
			fields.Def("IPv6PrefixLen", meta.TypeInt),
		},
	}
}

// switchIPv6NetsModel InternetリソースのフィールドとしてのIPv6Net
func (m *modelsDef) switchIPv6NetsModel() *dsl.Model {
	return &dsl.Model{
		Name:      "IPv6NetInfo",
		NakedType: meta.Static(naked.IPv6Net{}),
		IsArray:   true,
		Fields: []*dsl.FieldDesc{
			fields.ID(),
			fields.Def("IPv6Prefix", meta.TypeString),
			fields.Def("IPv6PrefixLen", meta.TypeInt),
		},
	}
}

// switchInfoModel Internetリソースのフィールドとしてのswitch
//
// Subnetの情報は限定的にしか返ってこない(IPAddresses.Max/Minなどがない)ため注意
// 必要であればSwitchリソース配下のSubnetsを参照すればOK
func (m *modelsDef) switchInfoModel() *dsl.Model {
	return &dsl.Model{
		Name:      "SwitchInfo",
		NakedType: meta.Static(naked.Switch{}),
		Fields: []*dsl.FieldDesc{
			fields.ID(),
			fields.Name(),
			fields.Description(),
			fields.Tags(),
			fields.Scope(),
			{
				Name: "Subnets",
				Type: m.internetSubnet(),
				Tags: &dsl.FieldTags{
					MapConv: "[]Subnets,recursive",
				},
			},
			fields.Def("IPv6Nets", m.switchIPv6NetsModel(), mapConvTag("[]IPv6Nets,recursive,omitempty")),
		},
	}
}

func (m *modelsDef) internetSubnet() *dsl.Model {
	return &dsl.Model{
		Name:      "InternetSubnet",
		NakedType: meta.Static(naked.Subnet{}),
		IsArray:   true,
		Fields: []*dsl.FieldDesc{
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
func (m *modelsDef) internetSubnetOperationResult() *dsl.Model {
	return &dsl.Model{
		Name:      "InternetSubnetOperationResult",
		NakedType: meta.Static(naked.Subnet{}),
		Fields: []*dsl.FieldDesc{
			fields.ID(),
			fields.DefaultRoute(),
			fields.NextHop(),
			fields.StaticRoute(),
			fields.NetworkAddress(),
			fields.NetworkMaskLen(),
			{
				Name: "IPAddresses",
				Type: meta.TypeStringSlice,
				Tags: &dsl.FieldTags{
					MapConv: "[]IPAddresses.IPAddress",
				},
			},
		},
	}
}

func (m *modelsDef) switchSubnet() *dsl.Model {
	// switchSubnetはinternetSubnetにInternetとIPAddressesを追加したもの
	subnet := m.internetSubnet()
	subnet.Name = "SwitchSubnet"
	subnet.Fields = append(subnet.Fields,
		&dsl.FieldDesc{
			Name: "Internet",
			Type: m.internetModel(),
		},
		&dsl.FieldDesc{
			Name: "AssignedIPAddressMax",
			Type: meta.TypeString,
			Tags: &dsl.FieldTags{
				MapConv: "IPAddresses.Max",
			},
		},
		&dsl.FieldDesc{
			Name: "AssignedIPAddressMin",
			Type: meta.TypeString,
			Tags: &dsl.FieldTags{
				MapConv: "IPAddresses.Min",
			},
		},
	)
	subnet.Methods = []*dsl.MethodDesc{
		{
			Name:        "GetAssignedIPAddresses",
			Description: "割り当てられたIPアドレスのリスト",
			ResultTypes: []meta.Type{meta.TypeStringSlice},
		},
	}
	return subnet
}

//******************************************************************************
// VPCRouter
//******************************************************************************

func (m *modelsDef) vpcRouterSetting() *dsl.Model {
	return &dsl.Model{
		Name:      "VPCRouterSetting",
		NakedType: meta.Static(naked.VPCRouterSettings{}),
		Fields: []*dsl.FieldDesc{
			{
				Name: "VRID",
				Type: meta.TypeInt,
				Tags: &dsl.FieldTags{
					JSON:    ",omitempty",
					MapConv: "Router.VRID",
				},
			},
			{
				Name: "InternetConnectionEnabled",
				Type: meta.TypeStringFlag,
				Tags: &dsl.FieldTags{
					MapConv: "Router.InternetConnection.Enabled,omitempty",
				},
			},
			{
				Name: "Interfaces",
				Type: m.vpcRouterInterface(),
				Tags: &dsl.FieldTags{
					MapConv: "Router.[]Interfaces,omitempty,recursive",
				},
			},
			{
				Name: "StaticNAT",
				Type: m.vpcRouterStaticNAT(),
				Tags: &dsl.FieldTags{
					MapConv: "Router.StaticNAT.[]Config,omitempty,recursive",
				},
			},
			{
				Name: "PortForwarding",
				Type: m.vpcRouterPortForwarding(),
				Tags: &dsl.FieldTags{
					MapConv: "Router.PortForwarding.[]Config,omitempty,recursive",
				},
			},
			{
				Name: "Firewall",
				Type: m.vpcRouterFirewall(),
				Tags: &dsl.FieldTags{
					MapConv: "Router.Firewall.[]Config,omitempty,recursive",
				},
			},
			{
				Name: "DHCPServer",
				Type: m.vpcRouterDHCPServer(),
				Tags: &dsl.FieldTags{
					MapConv: "Router.DHCPServer.[]Config,omitempty,recursive",
				},
			},
			{
				Name: "DHCPStaticMapping",
				Type: m.vpcRouterDHCPStaticMapping(),
				Tags: &dsl.FieldTags{
					MapConv: "Router.DHCPStaticMapping.[]Config,omitempty,recursive",
				},
			},
			{
				Name: "PPTPServer",
				Type: m.vpcRouterPPTPServer(),
				Tags: &dsl.FieldTags{
					MapConv: "Router.PPTPServer.Config,omitempty,recursive",
				},
			},
			{
				Name: "PPTPServerEnabled",
				Type: meta.TypeStringFlag,
				Tags: &dsl.FieldTags{
					MapConv: "Router.PPTPServer.Enabled,omitempty",
				},
			},
			{
				Name: "L2TPIPsecServer",
				Type: m.vpcRouterL2TPIPsecServer(),
				Tags: &dsl.FieldTags{
					MapConv: "Router.L2TPIPsecServer.Config,omitempty,recursive",
				},
			},
			{
				Name: "L2TPIPsecServerEnabled",
				Type: meta.TypeStringFlag,
				Tags: &dsl.FieldTags{
					MapConv: "Router.L2TPIPsecServer.Enabled,omitempty",
				},
			},
			{
				Name: "RemoteAccessUsers",
				Type: m.vpcRouterRemoteAccessUser(),
				Tags: &dsl.FieldTags{
					MapConv: "Router.RemoteAccessUsers.[]Config,omitempty,recursive",
				},
			},
			{
				Name: "SiteToSiteIPsecVPN",
				Type: m.vpcRouterSiteToSiteIPsecVPN(),
				Tags: &dsl.FieldTags{
					MapConv: "Router.SiteToSiteIPsecVPN.[]Config,omitempty,recursive",
				},
			},
			{
				Name: "StaticRoute",
				Type: m.vpcRouterStaticRoute(),
				Tags: &dsl.FieldTags{
					MapConv: "Router.StaticRoutes.[]Config,omitempty,recursive",
				},
			},
			{
				Name: "SyslogHost",
				Type: meta.TypeString,
				Tags: &dsl.FieldTags{
					MapConv: "Router.SyslogHost",
				},
			},
		},
	}
}

func (m *modelsDef) vpcRouterInterface() *dsl.Model {
	return &dsl.Model{
		Name:      "VPCRouterInterfaceSetting",
		NakedType: meta.Static(naked.VPCRouterInterface{}),
		IsArray:   true,
		Fields: []*dsl.FieldDesc{
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

func (m *modelsDef) vpcRouterStaticNAT() *dsl.Model {
	return &dsl.Model{
		Name:      "VPCRouterStaticNAT",
		NakedType: meta.Static(naked.VPCRouterStaticNATConfig{}),
		IsArray:   true,
		Fields: []*dsl.FieldDesc{
			{
				Name: "GlobalAddress",
				Type: meta.TypeString,
				Tags: &dsl.FieldTags{
					MapConv: "GlobalAddress",
				},
			},
			{
				Name: "PrivateAddress",
				Type: meta.TypeString,
				Tags: &dsl.FieldTags{
					MapConv: "PrivateAddress",
				},
			},
			{
				Name: "Description",
				Type: meta.TypeString,
			},
		},
	}
}

func (m *modelsDef) vpcRouterPortForwarding() *dsl.Model {
	return &dsl.Model{
		Name:      "VPCRouterPortForwarding",
		NakedType: meta.Static(naked.VPCRouterPortForwardingConfig{}),
		IsArray:   true,
		Fields: []*dsl.FieldDesc{
			{
				Name: "Protocol",
				Type: meta.Static(types.EVPCRouterPortForwardingProtocol("")),
			},
			{
				Name: "GlobalPort",
				Type: meta.TypeStringNumber,
			},
			{
				Name: "PrivateAddress",
				Type: meta.TypeString,
				Tags: &dsl.FieldTags{
					MapConv: "PrivateAddress",
				},
			},
			{
				Name: "PrivatePort",
				Type: meta.TypeStringNumber,
			},
			{
				Name: "Description",
				Type: meta.TypeString,
			},
		},
	}
}

func (m *modelsDef) vpcRouterFirewall() *dsl.Model {
	return &dsl.Model{
		Name:      "VPCRouterFirewall",
		NakedType: meta.Static(naked.VPCRouterFirewallConfig{}),
		IsArray:   true,
		Fields: []*dsl.FieldDesc{
			{
				Name: "Send",
				Type: m.vpcRouterFirewallRule(),
			},
			{
				Name: "Receive",
				Type: m.vpcRouterFirewallRule(),
			},
			{
				Name: "Index",
				Type: meta.TypeInt,
			},
		},
	}
}
func (m *modelsDef) vpcRouterFirewallRule() *dsl.Model {
	return &dsl.Model{
		Name:      "VPCRouterFirewallRule",
		NakedType: meta.Static(naked.VPCRouterFirewallRule{}),
		IsArray:   true,
		Fields: []*dsl.FieldDesc{
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

func (m *modelsDef) vpcRouterDHCPServer() *dsl.Model {
	return &dsl.Model{
		Name:      "VPCRouterDHCPServer",
		NakedType: meta.Static(naked.VPCRouterDHCPServerConfig{}),
		IsArray:   true,
		Fields: []*dsl.FieldDesc{
			{
				Name: "Interface",
				Type: meta.TypeString,
			},
			{
				Name: "RangeStart",
				Type: meta.TypeString,
			},
			{
				Name: "RangeStop",
				Type: meta.TypeString,
			},
			{
				Name: "DNSServers",
				Type: meta.TypeStringSlice,
			},
		},
	}
}

func (m *modelsDef) vpcRouterDHCPStaticMapping() *dsl.Model {
	return &dsl.Model{
		Name:      "VPCRouterDHCPStaticMapping",
		NakedType: meta.Static(naked.VPCRouterDHCPStaticMappingConfig{}),
		IsArray:   true,
		Fields: []*dsl.FieldDesc{
			{
				Name: "MACAddress",
				Type: meta.TypeString,
			},
			{
				Name: "IPAddress",
				Type: meta.TypeString,
			},
		},
	}
}

func (m *modelsDef) vpcRouterPPTPServer() *dsl.Model {
	return &dsl.Model{
		Name:      "VPCRouterPPTPServer",
		NakedType: meta.Static(naked.VPCRouterPPTPServerConfig{}),
		Fields: []*dsl.FieldDesc{
			{
				Name: "RangeStart",
				Type: meta.TypeString,
			},
			{
				Name: "RangeStop",
				Type: meta.TypeString,
			},
		},
	}
}

func (m *modelsDef) vpcRouterL2TPIPsecServer() *dsl.Model {
	return &dsl.Model{
		Name:      "VPCRouterL2TPIPsecServer",
		NakedType: meta.Static(naked.VPCRouterL2TPIPsecServerConfig{}),
		Fields: []*dsl.FieldDesc{
			{
				Name: "RangeStart",
				Type: meta.TypeString,
			},
			{
				Name: "RangeStop",
				Type: meta.TypeString,
			},
			{
				Name: "PreSharedSecret",
				Type: meta.TypeString,
			},
		},
	}
}

func (m *modelsDef) vpcRouterRemoteAccessUser() *dsl.Model {
	return &dsl.Model{
		Name:      "VPCRouterRemoteAccessUser",
		NakedType: meta.Static(naked.VPCRouterRemoteAccessUserConfig{}),
		IsArray:   true,
		Fields: []*dsl.FieldDesc{
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

func (m *modelsDef) vpcRouterSiteToSiteIPsecVPN() *dsl.Model {
	return &dsl.Model{
		Name:      "VPCRouterSiteToSiteIPsecVPN",
		NakedType: meta.Static(naked.VPCRouterSiteToSiteIPsecVPNConfig{}),
		IsArray:   true,
		Fields: []*dsl.FieldDesc{
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

func (m *modelsDef) vpcRouterStaticRoute() *dsl.Model {
	return &dsl.Model{
		Name:      "VPCRouterStaticRoute",
		NakedType: meta.Static(naked.VPCRouterStaticRouteConfig{}),
		IsArray:   true,
		Fields: []*dsl.FieldDesc{
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

//******************************************************************************
// Mobile Gateway
//******************************************************************************

func (m *modelsDef) mobileGatewayInterface() *dsl.Model {
	return &dsl.Model{
		Name:      "MobileGatewayInterfaceSetting",
		NakedType: meta.Static(naked.MobileGatewayInterface{}),
		IsArray:   true,
		Fields: []*dsl.FieldDesc{
			fields.Def("IPAddress", meta.TypeStringSlice),
			fields.Def("NetworkMaskLen", meta.TypeInt),
			fields.Def("Index", meta.TypeInt),
		},
	}
}

func (m *modelsDef) mobileGatewayStaticRoute() *dsl.Model {
	return &dsl.Model{
		Name:      "MobileGatewayStaticRoute",
		NakedType: meta.Static(naked.MobileGatewayStaticRoute{}),
		IsArray:   true,
		Fields: []*dsl.FieldDesc{
			fields.Def("Prefix", meta.TypeString),
			fields.Def("NextHop", meta.TypeString),
		},
	}
}

//******************************************************************************
// SIM
//******************************************************************************

func (m *modelsDef) simInfo() *dsl.Model {
	return &dsl.Model{
		Name:      "SIMInfo",
		NakedType: meta.Static(naked.SIMInfo{}),
		Fields:    m.simInfoFields(),
	}
}

func (m *modelsDef) simInfoList() *dsl.Model {
	return &dsl.Model{
		Name:      "MobileGatewaySIMInfo",
		Alias:     "MobileGatewaySIMs",
		NakedType: meta.Static(naked.SIMInfo{}),
		Fields:    m.simInfoFields(),
	}
}

func (m *modelsDef) simInfoFields() []*dsl.FieldDesc {
	return []*dsl.FieldDesc{
		fields.Def("ICCID", meta.TypeString),
		fields.Def("IMSI", meta.TypeStringSlice),
		fields.Def("IMEI", meta.TypeString),
		fields.Def("IP", meta.TypeString),
		fields.Def("SessionStatus", meta.TypeString),
		fields.Def("IMEILock", meta.TypeFlag),
		fields.Def("Registered", meta.TypeFlag),
		fields.Def("Activated", meta.TypeFlag),
		fields.Def("ResourceID", meta.TypeString),
		fields.Def("RegisteredDate", meta.TypeTime),
		fields.Def("ActivatedDate", meta.TypeTime),
		fields.Def("DeactivatedDate", meta.TypeTime),
		fields.Def("SIMGroupID", meta.TypeString),
		{
			Name: "TrafficBytesOfCurrentMonth",
			Type: &dsl.Model{
				Name: "SIMTrafficBytes",
				Fields: []*dsl.FieldDesc{
					fields.Def("UplinkBytes", meta.TypeInt64),
					fields.Def("DownlinkBytes", meta.TypeInt64),
				},
			},
			Tags: &dsl.FieldTags{
				MapConv: ",recursive",
			},
		},
		fields.Def("ConnectedIMEI", meta.TypeString),
	}
}

func (m *modelsDef) migrationJobStatus() *dsl.Model {
	return &dsl.Model{
		Name:      "JobStatus",
		NakedType: meta.Static(naked.MigrationJobStatus{}),
		Fields: []*dsl.FieldDesc{
			fields.Def("Status", meta.TypeString),
			fields.Def("ConfigError", &dsl.Model{
				Name:      "JobConfigError",
				NakedType: meta.Static(naked.JobConfigError{}),
				Fields: []*dsl.FieldDesc{
					fields.Def("ErrorCode", meta.TypeString),
					fields.Def("ErrorMsg", meta.TypeString),
					fields.Def("Status", meta.TypeString),
				},
			}),
		},
	}
}
