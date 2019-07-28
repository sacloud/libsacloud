package define

import (
	"github.com/sacloud/libsacloud/v2/internal/dsl"
	"github.com/sacloud/libsacloud/v2/internal/dsl/meta"
	"github.com/sacloud/libsacloud/v2/sacloud/naked"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

func mapConvTag(tag string) *dsl.FieldTags {
	return &dsl.FieldTags{
		MapConv: tag,
	}
}

type fieldsDef struct{}

var fields = &fieldsDef{}

func (f *fieldsDef) Def(name string, t meta.Type, tag ...*dsl.FieldTags) *dsl.FieldDesc {
	desc := &dsl.FieldDesc{
		Name: name,
		Type: t,
	}
	if len(tag) > 0 {
		desc.Tags = tag[0]
	}
	return desc
}

func (f *fieldsDef) ID() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "ID",
		Type: meta.TypeID,
		Methods: []*dsl.MethodDesc{
			{
				Name: "SetStringID",
				Arguments: dsl.Arguments{
					{
						Name: "id",
						Type: meta.TypeString,
					},
				},
			},
			{
				Name:        "GetStringID",
				ResultTypes: []meta.Type{meta.TypeString},
			},
			{
				Name: "SetInt64ID",
				Arguments: dsl.Arguments{
					{
						Name: "id",
						Type: meta.TypeInt64,
					},
				},
			},
			{
				Name:        "GetInt64ID",
				ResultTypes: []meta.Type{meta.TypeInt64},
			},
		},
	}
}

func (f *fieldsDef) Name() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "Name",
		Type: meta.TypeString,
		Tags: &dsl.FieldTags{
			Validate: "required",
		},
	}
}

func (f *fieldsDef) InterfaceDriver() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name:         "InterfaceDriver",
		Type:         meta.TypeInterfaceDriver,
		DefaultValue: `types.InterfaceDrivers.VirtIO`,
	}
}

func (f *fieldsDef) BridgeInfo() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "BridgeInfo",
		Type: models.bridgeInfoModel(),
		Tags: &dsl.FieldTags{
			MapConv: "[]Switches,recursive",
		},
	}
}

func (f *fieldsDef) SwitchInZone() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "SwitchInZone",
		Type: models.switchInZoneModel(),
	}
}

func (f *fieldsDef) DiskPlanID() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "DiskPlanID",
		Tags: &dsl.FieldTags{
			MapConv: "Plan.ID",
		},
		Type: meta.TypeID,
	}
}

func (f *fieldsDef) DiskPlanName() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "DiskPlanName",
		Tags: &dsl.FieldTags{
			MapConv: "Plan.Name",
		},
		Type: meta.TypeString,
	}
}

func (f *fieldsDef) DiskPlanStorageClass() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "DiskPlanStorageClass",
		Tags: &dsl.FieldTags{
			MapConv: "Plan.StorageClass",
		},
		Type: meta.TypeString,
	}
}

func (f *fieldsDef) CPU() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "CPU",
		Type: meta.TypeInt,
	}
}

/*
func (f *fieldsDef) ServerPlanCPU() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "CPU",
		Tags: &schema.FieldTags{
			MapConv: "ServerPlan.CPU",
		},
		Type: meta.TypeInt,
	}
}
*/

func (f *fieldsDef) MemoryMB() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "MemoryMB",
		Type: meta.TypeInt,
		Methods: []*dsl.MethodDesc{
			{
				Name:        "GetMemoryGB",
				ResultTypes: []meta.Type{meta.TypeInt},
			},
			{
				Name: "SetMemoryGB",
				Arguments: dsl.Arguments{
					{
						Name: "memory",
						Type: meta.TypeInt,
					},
				},
			},
		},
	}
}

func (f *fieldsDef) Generation() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "ServerPlanGeneration",
		Type: meta.TypePlanGeneration,
	}
}

func (f *fieldsDef) Commitment() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "ServerPlanCommitment",
		Tags: &dsl.FieldTags{
			MapConv: "Commitment",
		},
		Type:         meta.TypeCommitment,
		DefaultValue: "types.Commitments.Standard",
	}
}

func (f *fieldsDef) ServerPlanID() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "ServerPlanID",
		Tags: &dsl.FieldTags{
			MapConv: "ServerPlan.ID",
		},
		Type: meta.TypeID,
	}
}

func (f *fieldsDef) ServerPlanName() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "ServerPlanName",
		Tags: &dsl.FieldTags{
			MapConv: "ServerPlan.Name",
		},
		Type: meta.TypeString,
	}
}

func (f *fieldsDef) ServerPlanCPU() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "CPU",
		Tags: &dsl.FieldTags{
			MapConv: "ServerPlan.CPU",
		},
		Type: meta.TypeInt,
	}
}

func (f *fieldsDef) ServerPlanMemoryMB() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "MemoryMB",
		Tags: &dsl.FieldTags{
			MapConv: "ServerPlan.MemoryMB",
		},
		Type: meta.TypeInt,
		Methods: []*dsl.MethodDesc{
			{
				Name:        "GetMemoryGB",
				ResultTypes: []meta.Type{meta.TypeInt},
			},
		},
	}
}

func (f *fieldsDef) ServerPlanGeneration() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "ServerPlanGeneration",
		Type: meta.TypePlanGeneration,
		Tags: &dsl.FieldTags{
			MapConv: "ServerPlan.Generation",
		},
	}
}

func (f *fieldsDef) ServerPlanCommitment() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "ServerPlanCommitment",
		Tags: &dsl.FieldTags{
			MapConv: "ServerPlan.Commitment",
		},
		Type:         meta.TypeCommitment,
		DefaultValue: "types.Commitments.Standard",
	}
}

func (f *fieldsDef) PlanID() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "PlanID",
		Tags: &dsl.FieldTags{
			MapConv: "Plan.ID",
		},
		Type: meta.TypeID,
	}
}

func (f *fieldsDef) PlanName() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "PlanName",
		Tags: &dsl.FieldTags{
			MapConv: "Plan.Name",
		},
		Type: meta.TypeString,
	}
}

func (f *fieldsDef) stringEnabled() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "Enabled",
		Tags: &dsl.FieldTags{
			MapConv: ",omitempty",
		},
		Type: meta.TypeStringFlag,
	}
}

func (f *fieldsDef) IconURL() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "URL",
		Type: meta.TypeString,
	}
}

func (f *fieldsDef) IconImage() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "Image",
		Type: meta.TypeString,
	}
}

func (f *fieldsDef) PublicKey() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "PublicKey",
		Type: meta.TypeString,
	}
}

func (f *fieldsDef) PrivateKey() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "PrivateKey",
		Type: meta.TypeString,
	}
}

func (f *fieldsDef) Fingerprint() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "Fingerprint",
		Type: meta.TypeString,
	}
}

func (f *fieldsDef) PassPhrase() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "PassPhrase",
		Type: meta.TypeString,
	}
}

func (f *fieldsDef) SourceDiskID() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "SourceDiskID",
		Tags: &dsl.FieldTags{
			MapConv: "SourceDisk.ID,omitempty",
		},
		Type: meta.TypeID,
	}
}

func (f *fieldsDef) SourceDiskAvailability() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "SourceDiskAvailability",
		Tags: &dsl.FieldTags{
			MapConv: "SourceDisk.Availability,omitempty",
		},
		Type: meta.TypeAvailability,
	}
}

func (f *fieldsDef) SourceArchiveID() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "SourceArchiveID",
		Tags: &dsl.FieldTags{
			MapConv: "SourceArchive.ID,omitempty",
		},
		Type: meta.TypeID,
	}
}

func (f *fieldsDef) SourceArchiveAvailability() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "SourceArchiveAvailability",
		Tags: &dsl.FieldTags{
			MapConv: "SourceArchive.Availability,omitempty",
		},
		Type: meta.TypeAvailability,
	}
}

func (f *fieldsDef) OriginalArchiveID() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "OriginalArchiveID",
		Tags: &dsl.FieldTags{
			MapConv: "OriginalArchive.ID,omitempty",
		},
		Type: meta.TypeID,
	}
}

func (f *fieldsDef) BridgeID() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "BridgeID",
		Type: meta.TypeID,
		Tags: &dsl.FieldTags{
			MapConv: "Bridge.ID,omitempty",
		},
	}
}

func (f *fieldsDef) SwitchID() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "SwitchID",
		Type: meta.TypeID,
		Tags: &dsl.FieldTags{
			MapConv: "Switch.ID,omitempty",
		},
	}
}

func (f *fieldsDef) PacketFilterID() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "PacketFilterID",
		Type: meta.TypeID,
		Tags: &dsl.FieldTags{
			MapConv: "PacketFilter.ID,omitempty",
		},
	}
}

func (f *fieldsDef) ServerID() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "ServerID",
		Tags: &dsl.FieldTags{
			MapConv: "Server.ID,omitempty",
		},
		Type: meta.TypeID,
	}
}

func (f *fieldsDef) PrivateHostHostName() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "HostName",
		Type: meta.TypeString,
		Tags: &dsl.FieldTags{
			MapConv: "Host.Name",
		},
	}
}

func (f *fieldsDef) PrivateHostPlanID() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "PlanID",
		Tags: &dsl.FieldTags{
			MapConv: "Plan.ID,omitempty",
		},
		Type: meta.TypeID,
	}
}

func (f *fieldsDef) IconID() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "IconID",
		Tags: &dsl.FieldTags{
			MapConv: "Icon.ID",
		},
		Type: meta.TypeID,
	}
}

func (f *fieldsDef) LicenseInfoID() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "LicenseInfoID",
		Tags: &dsl.FieldTags{
			MapConv: "LicenseInfo.ID",
		},
		Type: meta.TypeID,
	}
}

func (f *fieldsDef) LicenseInfoName() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "LicenseInfoName",
		Tags: &dsl.FieldTags{
			MapConv: "LicenseInfo.Name",
		},
		Type: meta.TypeString,
	}
}

func (f *fieldsDef) ZoneID() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "ZoneID",
		Type: meta.TypeID,
		Tags: &dsl.FieldTags{
			MapConv: "Zone.ID",
		},
	}
}

func (f *fieldsDef) CDROMID() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "CDROMID",
		Type: meta.TypeID,
		Tags: &dsl.FieldTags{
			MapConv: "Instance.CDROM.ID",
		},
	}
}

func (f *fieldsDef) PrivateHostID() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "PrivateHostID",
		Type: meta.TypeID,
		Tags: &dsl.FieldTags{
			MapConv: "PrivateHost.ID",
		},
	}
}

func (f *fieldsDef) PrivateHostName() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "PrivateHostName",
		Type: meta.TypeString,
		Tags: &dsl.FieldTags{
			MapConv: "PrivateHost.Name",
		},
	}
}

func (f *fieldsDef) AppliancePlanID() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "PlanID",
		Tags: &dsl.FieldTags{
			MapConv: "Remark.Plan.ID/Plan.ID",
		},
		Type: meta.TypeID,
	}
}

func (f *fieldsDef) ApplianceSwitchID() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "SwitchID",
		Tags: &dsl.FieldTags{
			MapConv: "Remark.Switch.ID",
		},
		Type: meta.TypeID,
	}
}

func (f *fieldsDef) ApplianceIPAddress() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "IPAddress",
		Tags: &dsl.FieldTags{
			MapConv: "Remark.[]Servers.IPAddress",
		},
		Type: meta.TypeString,
	}
}

func (f *fieldsDef) ApplianceIPAddresses() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "IPAddresses",
		Type: meta.TypeStringSlice,
		Tags: &dsl.FieldTags{
			MapConv:  "Remark.[]Servers.IPAddress",
			Validate: "min=1,max=2,dive,ipv4",
		},
	}
}

func (f *fieldsDef) LoadBalancerVIP() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "VirtualIPAddresses",
		Type: &dsl.Model{
			Name:    "LoadBalancerVirtualIPAddress",
			IsArray: true,
			Fields: []*dsl.FieldDesc{
				{
					Name: "VirtualIPAddress",
					Type: meta.TypeString,
					Tags: &dsl.FieldTags{
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
					Tags: &dsl.FieldTags{
						Validate: "min=0,max=60", // TODO 最大値確認
					},
					DefaultValue: "10",
				},
				{
					Name: "SorryServer",
					Type: meta.TypeString,
					Tags: &dsl.FieldTags{
						Validate: "ipv4",
					},
				},
				f.Description(),
				{
					Name: "Servers",
					Type: &dsl.Model{
						Name:    "LoadBalancerServer",
						IsArray: true,
						Fields: []*dsl.FieldDesc{
							{
								Name: "IPAddress",
								Type: meta.TypeString,
								Tags: &dsl.FieldTags{
									Validate: "ipv4",
								},
							},
							{
								Name: "Port",
								Type: meta.TypeStringNumber,
								Tags: &dsl.FieldTags{
									Validate: "min=1,max=65535",
								},
							},
							{
								Name: "Enabled",
								Type: meta.TypeStringFlag,
							},
							{
								Name: "HealthCheckProtocol",
								Type: meta.TypeProtocol,
								Tags: &dsl.FieldTags{
									MapConv:  "HealthCheck.Protocol",
									Validate: "oneof=http https ping tcp",
								},
							},
							{
								Name: "HealthCheckPath",
								Type: meta.TypeString,
								Tags: &dsl.FieldTags{
									MapConv: "HealthCheck.Path",
								},
							},
							{
								Name: "HealthCheckResponseCode",
								Type: meta.TypeStringNumber,
								Tags: &dsl.FieldTags{
									MapConv: "HealthCheck.Status",
								},
							},
						},
					},
					Tags: &dsl.FieldTags{
						MapConv:  ",recursive",
						Validate: "min=0,max=40",
					},
				},
			},
		},
		Tags: &dsl.FieldTags{
			MapConv:  "Settings.[]LoadBalancer,recursive",
			Validate: "min=0,max=10",
		},
	}
}

func (f *fieldsDef) Tags() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "Tags",
		Type: meta.Static(types.Tags{}),
		Methods: []*dsl.MethodDesc{
			{
				Name:        "HasTag",
				Description: "指定のタグが存在する場合trueを返す",
				Arguments: dsl.Arguments{
					{
						Name: "tag",
						Type: meta.TypeString,
					},
				},
				ResultTypes: []meta.Type{meta.TypeFlag},
			},
			{
				Name:        "AppendTag",
				Description: "指定のタグを追加",
				Arguments: dsl.Arguments{
					{
						Name: "tag",
						Type: meta.TypeString,
					},
				},
			},
			{
				Name:        "RemoveTag",
				Description: "指定のタグを削除",
				Arguments: dsl.Arguments{
					{
						Name: "tag",
						Type: meta.TypeString,
					},
				},
			},
			{
				Name:        "ClearTags",
				Description: "タグを全クリア",
			},
		},
	}
}

func (f *fieldsDef) Class() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "Class",
		Type: meta.TypeString,
	}
}

func (f *fieldsDef) SIMICCID() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "ICCID",
		Type: meta.TypeString,
		Tags: &dsl.FieldTags{
			MapConv:  "Status.ICCID",
			Validate: "numeric", // TODO 数値のみ15桁固定
		},
	}
}

func (f *fieldsDef) SIMPassCode() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "PassCode",
		Type: meta.TypeString,
		Tags: &dsl.FieldTags{
			MapConv: "Remark.PassCode",
		},
	}
}

func (f *fieldsDef) GSLBFQDN() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "FQDN",
		Type: meta.TypeString,
		Tags: &dsl.FieldTags{
			MapConv: "Status.FQDN",
		},
	}
}

func (f *fieldsDef) GSLBHealthCheck() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "HealthCheck",
		Type: &dsl.Model{
			Name:      "GSLBHealthCheck",
			NakedType: meta.Static(naked.HealthCheck{}),
			Fields: []*dsl.FieldDesc{
				{
					Name: "Protocol",
					Type: meta.TypeProtocol,
					Tags: &dsl.FieldTags{
						Validate: "oneof=http https ping tcp",
					},
				},
				{
					Name: "HostHeader",
					Type: meta.TypeString,
					Tags: &dsl.FieldTags{
						MapConv: "Host",
					},
				},
				{
					Name: "Path",
					Type: meta.TypeString,
					Tags: &dsl.FieldTags{
						MapConv: "Path",
					},
				},
				{
					Name: "ResponseCode",
					Type: meta.TypeStringNumber,
					Tags: &dsl.FieldTags{
						MapConv: "Status",
					},
				},
				{
					Name: "Port",
					Type: meta.TypeStringNumber,
					Tags: &dsl.FieldTags{
						MapConv: "Port",
					},
				},
			},
		},
		Tags: &dsl.FieldTags{
			MapConv: "Settings.GSLB.HealthCheck,recursive",
		},
	}
}

func (f *fieldsDef) GSLBDelayLoop() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "DelayLoop",
		Type: meta.TypeInt,
		Tags: &dsl.FieldTags{
			Validate: "min=10,max=60",
			MapConv:  "Settings.GSLB.DelayLoop",
		},
		DefaultValue: "10",
	}
}

func (f *fieldsDef) GSLBWeighted() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "Weighted",
		Type: meta.TypeStringFlag,
		Tags: &dsl.FieldTags{
			MapConv: "Settings.GSLB.Weighted",
		},
	}
}

func (f *fieldsDef) GSLBDestinationServers() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "DestinationServers",
		Type: &dsl.Model{
			Name:    "GSLBServer",
			IsArray: true,
			Fields: []*dsl.FieldDesc{
				{
					Name: "IPAddress",
					Type: meta.TypeString,
					Tags: &dsl.FieldTags{
						Validate: "ipv4",
					},
				},
				{
					Name: "Enabled",
					Type: meta.TypeStringFlag,
				},
				{
					Name: "Weight",
					Type: meta.TypeStringNumber,
				},
			},
		},
		Tags: &dsl.FieldTags{
			MapConv:  "Settings.GSLB.[]Servers,recursive",
			Validate: "min=0,max=12",
		},
	}
}

func (f *fieldsDef) GSLBSorryServer() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "SorryServer",
		Type: meta.TypeString,
		Tags: &dsl.FieldTags{
			MapConv: "Settings.GSLB.SorryServer",
		},
	}
}

func (f *fieldsDef) AutoBackupBackupSpanWeekDays() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "BackupSpanWeekdays",
		Type: meta.TypeBackupSpanWeekdays,
		Tags: &dsl.FieldTags{
			MapConv: "Settings.Autobackup.BackupSpanWeekdays",
		},
	}
}

func (f *fieldsDef) AutoBackupMaximumNumberOfArchives() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "MaximumNumberOfArchives",
		Type: meta.TypeInt,
		Tags: &dsl.FieldTags{
			MapConv: "Settings.Autobackup.MaximumNumberOfArchives",
		},
	}
}

func (f *fieldsDef) AutoBackupDiskID() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "DiskID",
		Type: meta.TypeID,
		Tags: &dsl.FieldTags{
			MapConv: "Status.DiskId",
		},
	}
}

func (f *fieldsDef) AutoBackupAccountID() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "AccountID",
		Type: meta.TypeID,
		Tags: &dsl.FieldTags{
			MapConv: "Status.AccountId",
		},
	}
}

func (f *fieldsDef) AutoBackupZoneID() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "ZoneID",
		Type: meta.TypeID,
		Tags: &dsl.FieldTags{
			MapConv: "Status.ZoneId",
		},
	}
}

func (f *fieldsDef) AutoBackupZoneName() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "ZoneName",
		Type: meta.TypeString,
		Tags: &dsl.FieldTags{
			MapConv: "Status.ZoneName",
		},
	}
}

func (f *fieldsDef) DNSRecords() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "Records",
		Type: &dsl.Model{
			Name:    "DNSRecord",
			IsArray: true,
			Fields: []*dsl.FieldDesc{
				{
					Name: "Name",
					Type: meta.TypeString,
				},
				{
					Name: "Type",
					Type: meta.TypeDNSRecordType,
				},
				{
					Name: "RData",
					Type: meta.TypeString,
				},
				{
					Name: "TTL",
					Type: meta.TypeInt,
				},
			},
		},
		Tags: &dsl.FieldTags{
			MapConv:  "Settings.DNS.[]ResourceRecordSets,recursive",
			Validate: "min=0,max=1000",
		},
	}
}

func (f *fieldsDef) DNSZone() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "DNSZone",
		Type: meta.TypeString,
		Tags: &dsl.FieldTags{
			MapConv: "Status.Zone",
		},
	}
}

func (f *fieldsDef) DNSNameServers() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "DNSNameServers",
		Type: meta.TypeStringSlice,
		Tags: &dsl.FieldTags{
			MapConv: "Status.NS",
		},
	}
}

func (f *fieldsDef) SimpleMonitorTarget() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "Target",
		Type: meta.TypeString,
		Tags: &dsl.FieldTags{
			MapConv: "Status.Target",
		},
	}
}

func (f *fieldsDef) SimpleMonitorDelayLoop() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "DelayLoop",
		Type: meta.TypeInt,
		Tags: &dsl.FieldTags{
			Validate: "min=60,max=3600",
			MapConv:  "Settings.SimpleMonitor.DelayLoop",
		},
		DefaultValue: "60",
	}
}

func (f *fieldsDef) SimpleMonitorEnabled() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "Enabled",
		Type: meta.TypeStringFlag,
		Tags: &dsl.FieldTags{
			MapConv: "Settings.SimpleMonitor.Enabled",
		},
	}
}

func (f *fieldsDef) SimpleMonitorNotifyEmailEnabled() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "NotifyEmailEnabled",
		Type: meta.TypeStringFlag,
		Tags: &dsl.FieldTags{
			MapConv: "Settings.SimpleMonitor.NotifyEmail.Enabled",
		},
	}
}

func (f *fieldsDef) SimpleMonitorNotifyEmailHTML() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "NotifyEmailHTML",
		Type: meta.TypeStringFlag,
		Tags: &dsl.FieldTags{
			MapConv: "Settings.SimpleMonitor.NotifyEmail.HTML",
		},
	}
}

func (f *fieldsDef) SimpleMonitorNotifySlackEnabled() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "NotifySlackEnabled",
		Type: meta.TypeStringFlag,
		Tags: &dsl.FieldTags{
			MapConv: "Settings.SimpleMonitor.NotifySlack.Enabled",
		},
	}
}

func (f *fieldsDef) SimpleMonitorSlackWebhooksURL() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "SlackWebhooksURL",
		Type: meta.TypeString,
		Tags: &dsl.FieldTags{
			MapConv: "Settings.SimpleMonitor.NotifySlack.IncomingWebhooksURL",
		},
	}
}

func (f *fieldsDef) SimpleMonitorHealthCheck() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "HealthCheck",
		Type: &dsl.Model{
			Name: "SimpleMonitorHealthCheck",
			Fields: []*dsl.FieldDesc{
				{
					Name: "Protocol",
					Type: meta.TypeSimpleMonitorHealthCheckProtocol,
				},
				{
					Name: "Port",
					Type: meta.TypeStringNumber,
				},
				{
					Name: "Path",
					Type: meta.TypeString,
				},
				{
					Name: "Status",
					Type: meta.TypeStringNumber,
				},
				{
					Name: "SNI",
					Type: meta.TypeStringFlag,
				},
				{
					Name: "Host",
					Type: meta.TypeString,
				},
				{
					Name: "BasicAuthUsername",
					Type: meta.TypeString,
				},
				{
					Name: "BasicAuthPassword",
					Type: meta.TypeString,
				},
				{
					Name: "QName",
					Type: meta.TypeString,
				},
				{
					Name: "ExpectedData",
					Type: meta.TypeString,
				},
				{
					Name: "Community",
					Type: meta.TypeString,
				},
				{
					Name: "SNMPVersion",
					Type: meta.TypeString,
				},
				{
					Name: "OID",
					Type: meta.TypeString,
				},
				{
					Name: "RemainingDays",
					Type: meta.TypeInt,
				},
			},
		},
		Tags: &dsl.FieldTags{
			MapConv: "Settings.SimpleMonitor.HealthCheck,recursive",
		},
	}
}

func (f *fieldsDef) ProxyLBPlan() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "Plan",
		Type: meta.Static(types.EProxyLBPlan(0)),
		Tags: &dsl.FieldTags{
			MapConv: "ServiceClass",
		},
	}
}

func (f *fieldsDef) ProxyLBUseVIPFailover() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "UseVIPFailover",
		Type: meta.TypeFlag,
		Tags: &dsl.FieldTags{
			MapConv: "Status.UseVIPFailover",
		},
	}
}

func (f *fieldsDef) ProxyLBRegion() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "Region",
		Type: meta.Static(types.EProxyLBRegion("")),
		Tags: &dsl.FieldTags{
			MapConv: "Status.Region",
		},
	}
}

func (f *fieldsDef) ProxyLBProxyNetworks() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "ProxyNetworks",
		Type: meta.TypeStringSlice,
		Tags: &dsl.FieldTags{
			MapConv: "Status.ProxyNetworks",
		},
	}
}

func (f *fieldsDef) ProxyLBFQDN() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "FQDN",
		Type: meta.TypeString,
		Tags: &dsl.FieldTags{
			MapConv: "Status.FQDN",
		},
	}
}

func (f *fieldsDef) ProxyLBHealthCheck() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "HealthCheck",
		Type: &dsl.Model{
			Name: "ProxyLBHealthCheck",
			Fields: []*dsl.FieldDesc{
				{
					Name: "Protocol",
					Type: meta.Static(types.EProxyLBHealthCheckProtocol("")),
				},
				{
					Name: "Path",
					Type: meta.TypeString,
				},
				{
					Name: "DelayLoop",
					Type: meta.TypeInt,
				},
			},
		},
		Tags: &dsl.FieldTags{
			MapConv: "Settings.ProxyLB.HealthCheck,recursive",
		},
	}
}

func (f *fieldsDef) ProxyLBSorryServer() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "SorryServer",
		Type: &dsl.Model{
			Name: "ProxyLBSorryServer",
			Fields: []*dsl.FieldDesc{
				{
					Name: "IPAddress",
					Type: meta.TypeString,
					Tags: &dsl.FieldTags{
						Validate: "ipv4",
					},
				},
				{
					Name: "Port",
					Type: meta.TypeInt,
					Tags: &dsl.FieldTags{
						Validate: "min=0,max=65535",
					},
				},
			},
		},
		Tags: &dsl.FieldTags{
			MapConv: "Settings.ProxyLB.SorryServer,recursive",
		},
	}
}

func (f *fieldsDef) ProxyLBBindPorts() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "BindPorts",
		Type: &dsl.Model{
			Name:    "ProxyLBBindPort",
			IsArray: true,
			Fields: []*dsl.FieldDesc{
				{
					Name: "ProxyMode",
					Type: meta.Static(types.EProxyLBProxyMode("")),
				},
				{
					Name: "Port",
					Type: meta.TypeInt,
					Tags: &dsl.FieldTags{
						Validate: "min=0,max=65535",
					},
				},
				{
					Name: "RedirectToHTTPS",
					Type: meta.TypeFlag,
				},
				{
					Name: "SupportHTTP2",
					Type: meta.TypeFlag,
				},
			},
		},
		Tags: &dsl.FieldTags{
			MapConv: "Settings.ProxyLB.[]BindPorts,recursive",
		},
	}
}

func (f *fieldsDef) ProxyLBServers() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "Servers",
		Type: &dsl.Model{
			Name:    "ProxyLBServer",
			IsArray: true,
			Fields: []*dsl.FieldDesc{
				{
					Name: "IPAddress",
					Type: meta.TypeString,
					Tags: &dsl.FieldTags{
						Validate: "ipv4",
					},
				},
				{
					Name: "Port",
					Type: meta.TypeInt,
					Tags: &dsl.FieldTags{
						Validate: "min=0,max=65535",
					},
				},
				{
					Name: "Enabled",
					Type: meta.TypeFlag,
				},
			},
		},
		Tags: &dsl.FieldTags{
			MapConv: "Settings.ProxyLB.[]Servers,recursive",
		},
	}
}

func (f *fieldsDef) ProxyLBLetsEncrypt() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "LetsEncrypt",
		Type: &dsl.Model{
			Name: "ProxyLBACMESetting",
			Fields: []*dsl.FieldDesc{
				{
					Name: "CommonName",
					Type: meta.TypeString,
				},
				{
					Name: "Enabled",
					Type: meta.TypeFlag,
				},
			},
		},
		Tags: &dsl.FieldTags{
			MapConv: "Settings.ProxyLB.LetsEncrypt,recursive",
		},
	}
}

func (f *fieldsDef) ProxyLBStickySession() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "StickySession",
		Type: &dsl.Model{
			Name: "ProxyLBStickySession",
			Fields: []*dsl.FieldDesc{
				{
					Name: "Method",
					Type: meta.TypeString,
				},
				{
					Name: "Enabled",
					Type: meta.TypeFlag,
				},
			},
		},
		Tags: &dsl.FieldTags{
			MapConv: "Settings.ProxyLB.StickySession,recursive",
		},
	}
}

func (f *fieldsDef) SettingsHash() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "SettingsHash",
		Type: meta.TypeString,
	}
}

func (f *fieldsDef) InstanceHostName() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "InstanceHostName",
		Type: meta.TypeString,
		Tags: &dsl.FieldTags{
			MapConv: "Instance.Host.Name",
		},
	}
}

func (f *fieldsDef) InstanceHostInfoURL() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "InstanceHostInfoURL",
		Type: meta.TypeString,
		Tags: &dsl.FieldTags{
			MapConv: "Instance.Host.InfoURL",
		},
	}
}

func (f *fieldsDef) InstanceStatus() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "InstanceStatus",
		Type: meta.TypeInstanceStatus,
		Tags: &dsl.FieldTags{
			MapConv: "Instance.Status",
		},
	}
}

func (f *fieldsDef) InstanceBeforeStatus() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "InstanceBeforeStatus",
		Type: meta.TypeInstanceStatus,
		Tags: &dsl.FieldTags{
			MapConv: "Instance.BeforeStatus",
		},
	}
}

func (f *fieldsDef) InstanceStatusChangedAt() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "InstanceStatusChangedAt",
		Type: meta.TypeTime,
		Tags: &dsl.FieldTags{
			MapConv: "Instance.StatusChangedAt",
		},
	}
}

func (f *fieldsDef) InstanceWarnings() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "InstanceWarnings",
		Type: meta.TypeString,
		Tags: &dsl.FieldTags{
			MapConv: "Instance.Warnings",
		},
	}
}

func (f *fieldsDef) InstanceWarningsValue() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "InstanceWarningsValue",
		Type: meta.TypeInt,
		Tags: &dsl.FieldTags{
			MapConv: "Instance.WarningsValue",
		},
	}
}

func (f *fieldsDef) SubnetID() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "SubnetID",
		Type: meta.TypeID,
		Tags: &dsl.FieldTags{
			MapConv: "Subnet.ID,omitempty",
		},
	}
}

func (f *fieldsDef) InterfaceID() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "InterfaceID",
		Type: meta.TypeID,
		Tags: &dsl.FieldTags{
			MapConv: "Interface.ID,omitempty",
		},
	}
}

func (f *fieldsDef) Interfaces() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "Interfaces",
		Type: models.interfaceModel(),
		Tags: &dsl.FieldTags{
			JSON:    ",omitempty",
			MapConv: "[]Interfaces,recursive,omitempty",
		},
	}
}

func (f *fieldsDef) VPCRouterInterfaces() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "Interfaces",
		Type: models.vpcRouterInterfaceModel(),
		Tags: &dsl.FieldTags{
			JSON:    ",omitempty",
			MapConv: "[]Interfaces,recursive,omitempty",
		},
	}
}

func (f *fieldsDef) MobileGatewayInterfaces() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "Interfaces",
		Type: models.mobileGatewayInterfaceModel(),
		Tags: &dsl.FieldTags{
			JSON:    ",omitempty",
			MapConv: "[]Interfaces,recursive,omitempty",
		},
	}
}

func (f *fieldsDef) NoteClass() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "Class",
		Type: meta.TypeString,
	}
}

func (f *fieldsDef) NoteContent() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "Content",
		Type: meta.TypeString,
	}
}

func (f *fieldsDef) Description() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "Description",
		Type: meta.TypeString,
		Tags: &dsl.FieldTags{
			Validate: "min=0,max=512",
		},
	}
}

func (f *fieldsDef) Availability() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "Availability",
		Type: meta.TypeAvailability,
	}
}

func (f *fieldsDef) Scope() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "Scope",
		Type: meta.TypeScope,
	}
}

func (f *fieldsDef) BandWidthMbps() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "BandWidthMbps",
		Type: meta.TypeInt,
	}
}

func (f *fieldsDef) DiskConnection() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "Connection",
		Type: meta.TypeDiskConnection,
		Tags: &dsl.FieldTags{
			JSON:    ",omitempty",
			MapConv: ",omitempty",
		},
	}
}

func (f *fieldsDef) DiskConnectionOrder() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "ConnectionOrder",
		Type: meta.TypeInt,
	}
}

func (f *fieldsDef) DiskReinstallCount() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "ReinstallCount",
		Type: meta.TypeInt,
	}
}

func (f *fieldsDef) SizeMB() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "SizeMB",
		Type: meta.TypeInt,
		Methods: []*dsl.MethodDesc{
			{
				Name:        "GetSizeGB",
				ResultTypes: []meta.Type{meta.TypeInt},
			},
			{
				Name: "SetSizeGB",
				Arguments: dsl.Arguments{
					{
						Name: "size",
						Type: meta.TypeInt,
					},
				},
			},
		},
	}
}

func (f *fieldsDef) MigratedMB() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "MigratedMB",
		Type: meta.TypeInt,
		Methods: []*dsl.MethodDesc{
			{
				Name:        "GetMigratedGB",
				ResultTypes: []meta.Type{meta.TypeInt},
			},
		},
	}
}

func (f *fieldsDef) DefaultRoute() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "DefaultRoute",
		Type: meta.TypeString,
		Tags: &dsl.FieldTags{
			Validate: "ipv4",
		},
	}
}

func (f *fieldsDef) NextHop() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "NextHop",
		Type: meta.TypeString,
		Description: `
			スイッチ+ルータでの追加IPアドレスブロックを示すSubnetの中でのみ設定される項目。
			この場合DefaultRouteの値は設定されないためNextHopを代用する。
			StaticRouteと同じ値が設定される。`,
		Tags: &dsl.FieldTags{
			Validate: "ipv4",
		},
	}
}

func (f *fieldsDef) StaticRoute() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "StaticRoute",
		Type: meta.TypeString,
		Description: `
			スイッチ+ルータでの追加IPアドレスブロックを示すSubnetの中でのみ設定される項目。
			この場合DefaultRouteの値は設定されないためNextHopを代用する。
			NextHopと同じ値が設定される。`,
		Tags: &dsl.FieldTags{
			Validate: "ipv4",
		},
	}
}

func (f *fieldsDef) NetworkMaskLen() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "NetworkMaskLen",
		Type: meta.TypeInt,
		Tags: &dsl.FieldTags{
			Validate: "min=24,max=28",
		},
	}
}

func (f *fieldsDef) NetworkAddress() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "NetworkAddress",
		Type: meta.TypeString,
		Tags: &dsl.FieldTags{
			Validate: "ipv4",
		},
	}
}

func (f *fieldsDef) UserSubnetNetworkMaskLen() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "NetworkMaskLen",
		Type: meta.TypeInt,
		Tags: &dsl.FieldTags{
			Validate: "min=1,max=32", // TODO
			MapConv:  "UserSubnet.NetworkMaskLen",
		},
	}
}

func (f *fieldsDef) UserSubnetDefaultRoute() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "DefaultRoute",
		Type: meta.TypeString,
		Tags: &dsl.FieldTags{
			Validate: "ipv4", // TODO
			MapConv:  "UserSubnet.DefaultRoute",
		},
	}
}

func (f *fieldsDef) RemarkNetworkMaskLen() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "NetworkMaskLen",
		Type: meta.TypeInt,
		Tags: &dsl.FieldTags{
			Validate: "min=1,max=32", // TODO
			MapConv:  "Remark.Network.NetworkMaskLen",
		},
	}
}

func (f *fieldsDef) RemarkZoneID() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "ZoneID",
		Type: meta.TypeID,
		Tags: &dsl.FieldTags{
			MapConv: "Remark.Zone.ID",
		},
	}
}

func (f *fieldsDef) RemarkDefaultRoute() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "DefaultRoute",
		Type: meta.TypeString,
		Tags: &dsl.FieldTags{
			Validate: "ipv4", // TODO
			MapConv:  "Remark.Network.DefaultRoute",
		},
	}
}

func (f *fieldsDef) RemarkServerIPAddress() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "IPAddresses",
		Type: meta.TypeStringSlice,
		Tags: &dsl.FieldTags{
			MapConv: "Remark.[]Servers.IPAddress",
		},
	}
}

func (f *fieldsDef) RemarkVRID() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "VRID",
		Type: meta.TypeInt,
		Tags: &dsl.FieldTags{
			MapConv: "Remark.VRRP.VRID",
		},
	}
}

func (f *fieldsDef) RemarkDBConf() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "Conf",
		Type: &dsl.Model{
			Name:      "DatabaseRemarkDBConfCommon",
			NakedType: meta.Static(naked.ApplianceRemarkDBConfCommon{}),
			Fields: []*dsl.FieldDesc{
				fields.Def("DatabaseName", meta.TypeString),
				fields.Def("DatabaseVersion", meta.TypeString),
				fields.Def("DatabaseRevision", meta.TypeString),
				fields.Def("DefaultUser", meta.TypeString),
				fields.Def("UserPassword", meta.TypeString),
			},
		},
		Tags: &dsl.FieldTags{
			MapConv: "Remark.DBConf.Common,recursive",
		},
	}
}

func (f *fieldsDef) RemarkSourceAppliance() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "SourceID",
		Type: meta.TypeID,
		Tags: &dsl.FieldTags{
			MapConv: "Remark.SourceAppliance.ID",
		},
	}
}

func (f *fieldsDef) DatabaseSettingsCommon() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "CommonSetting",
		Type: &dsl.Model{
			Name:      "DatabaseSettingCommon",
			NakedType: meta.Static(naked.DatabaseSettingCommon{}),
			Fields: []*dsl.FieldDesc{
				fields.Def("WebUI", meta.TypeFlag),
				fields.Def("ServicePort", meta.TypeInt),
				fields.Def("SourceNetwork", meta.TypeStringSlice),
				fields.Def("DefaultUser", meta.TypeString),
				fields.Def("UserPassword", meta.TypeString),
			},
		},
		Tags: &dsl.FieldTags{
			MapConv: "Settings.DBConf.Common,recursive",
		},
	}
}

func (f *fieldsDef) DatabaseSettingsCommonUpdate() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "CommonSetting",
		Type: &dsl.Model{
			Name:      "DatabaseSettingCommonUpdate",
			NakedType: meta.Static(naked.DatabaseSettingCommon{}),
			Fields: []*dsl.FieldDesc{
				fields.Def("WebUI", meta.TypeFlag), // TODO trueにした場合、文字列が返ってくるが未対応
				fields.Def("ServicePort", meta.TypeInt),
				fields.Def("SourceNetwork", meta.TypeStringSlice),
				fields.Def("DefaultUser", meta.TypeString),
				fields.Def("UserPassword", meta.TypeString),
			},
		},
		Tags: &dsl.FieldTags{
			MapConv: "Settings.DBConf.Common,recursive",
		},
	}
}

func (f *fieldsDef) DatabaseSettingsBackup() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "BackupSetting",
		Type: &dsl.Model{
			Name:      "DatabaseSettingBackup",
			NakedType: meta.Static(naked.DatabaseSettingBackup{}),
			Fields: []*dsl.FieldDesc{
				fields.Def("Rotate", meta.TypeInt),
				fields.Def("Time", meta.TypeString),
				fields.Def("DayOfWeek", meta.Static([]types.EBackupSpanWeekday{})),
			},
		},
		Tags: &dsl.FieldTags{
			MapConv: "Settings.DBConf.Backup,recursive",
		},
	}
}

func (f *fieldsDef) DatabaseSettingsReplication() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "ReplicationSetting",
		Type: &dsl.Model{
			Name:      "DatabaseReplicationSetting",
			NakedType: meta.Static(naked.DatabaseSettingReplication{}),
			Fields: []*dsl.FieldDesc{
				fields.Def("Model", meta.Static(types.EDatabaseReplicationModel(""))),
				fields.Def("IPAddress", meta.TypeString),
				fields.Def("Port", meta.TypeInt),
				fields.Def("User", meta.TypeString),
				fields.Def("Password", meta.TypeString),
				{
					Name: "ApplianceID",
					Type: meta.TypeID,
					Tags: &dsl.FieldTags{
						MapConv: "Appliance.ID",
					},
				},
			},
		},
		Tags: &dsl.FieldTags{
			MapConv: "Settings.DBConf.Replication,recursive",
		},
	}
}

func (f *fieldsDef) RequiredHostVersion() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "RequiredHostVersion",
		Type: meta.TypeStringNumber,
	}
}

func (f *fieldsDef) PacketFilterExpressions() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "Expression",
		Type: models.packetFilterExpressions(),
		Tags: &dsl.FieldTags{
			MapConv: "[]Expression,recursive",
		},
	}
}

func (f *fieldsDef) ExpressionHash() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "ExpressionHash",
		Type: meta.TypeString,
	}
}

func (f *fieldsDef) DisplayOrder() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "DisplayOrder",
		Type: meta.TypeInt64,
	}
}

func (f *fieldsDef) IsDummy() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "IsDummy",
		Type: meta.TypeFlag,
	}
}

func (f *fieldsDef) HostName() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "HostName",
		Type: meta.TypeString,
	}
}

func (f *fieldsDef) IPAddress() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "IPAddress",
		Type: meta.TypeString,
	}
}

func (f *fieldsDef) UserIPAddress() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "UserIPAddress",
		Type: meta.TypeString,
	}
}

func (f *fieldsDef) MACAddress() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "MACAddress",
		Type: meta.TypeString,
	}
}

func (f *fieldsDef) User() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "User",
		Type: meta.TypeString,
	}
}
func (f *fieldsDef) Password() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "Password",
		Type: meta.TypeString,
	}
}

func (f *fieldsDef) SourceInfo() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "SourceInfo",
		Type: models.sourceArchiveInfo(),
		Tags: &dsl.FieldTags{
			MapConv: ",omitempty,recursive",
		},
	}
}

func (f *fieldsDef) VNCProxy() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "VNCProxy",
		Type: models.vncProxyModel(),
		Tags: &dsl.FieldTags{
			JSON: ",omitempty",
		},
	}
}

func (f *fieldsDef) FTPServer() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "FTPServer",
		Type: models.ftpServerInfo(),
		Tags: &dsl.FieldTags{
			JSON: ",omitempty",
		},
	}
}

func (f *fieldsDef) Region() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "Region",
		Type: models.region(),
		Tags: &dsl.FieldTags{
			JSON: ",omitempty",
		},
	}
}

func (f *fieldsDef) NameServers() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "NameServers",
		Type: meta.TypeStringSlice,
		Tags: &dsl.FieldTags{
			MapConv: ",omitempty",
			JSON:    ",omitempty",
		},
	}
}

func (f *fieldsDef) Zone() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "Zone",
		Type: models.zoneInfoModel(),
		Tags: &dsl.FieldTags{
			MapConv: ",omitempty,recursive",
			JSON:    ",omitempty",
		},
	}
}

func (f *fieldsDef) Storage() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "Storage",
		Type: models.storageModel(),
		Tags: &dsl.FieldTags{
			MapConv: ",omitempty,recursive",
			JSON:    ",omitempty",
		},
	}
}

func (f *fieldsDef) BundleInfo() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "BundleInfo",
		Type: models.bundleInfoModel(),
		Tags: &dsl.FieldTags{
			MapConv: ",omitempty,recursive",
			JSON:    ",omitempty",
		},
	}
}

func (f *fieldsDef) CreatedAt() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "CreatedAt",
		Type: meta.TypeTime,
	}
}

func (f *fieldsDef) ModifiedAt() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "ModifiedAt",
		Type: meta.TypeTime,
	}
}

/*
 for monitor
*/
func (f *fieldsDef) MonitorTime() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "Time",
		Type: meta.TypeTime,
		Tags: &dsl.FieldTags{
			MapConv: ",omitempty",
			JSON:    ",omitempty",
		},
	}
}

func (f *fieldsDef) MonitorCPUTime() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "CPUTime",
		Type: meta.TypeFloat64,
		Tags: &dsl.FieldTags{
			MapConv: ",omitempty",
			JSON:    ",omitempty",
		},
	}
}

func (f *fieldsDef) MonitorDiskRead() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "Read",
		Type: meta.TypeFloat64,
		Tags: &dsl.FieldTags{
			MapConv: ",omitempty",
			JSON:    ",omitempty",
		},
	}
}

func (f *fieldsDef) MonitorDiskWrite() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "Write",
		Type: meta.TypeFloat64,
		Tags: &dsl.FieldTags{
			MapConv: ",omitempty",
			JSON:    ",omitempty",
		},
	}
}

func (f *fieldsDef) MonitorRouterIn() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "In",
		Type: meta.TypeFloat64,
		Tags: &dsl.FieldTags{
			MapConv: ",omitempty",
			JSON:    ",omitempty",
		},
	}
}

func (f *fieldsDef) MonitorRouterOut() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "Out",
		Type: meta.TypeFloat64,
		Tags: &dsl.FieldTags{
			MapConv: ",omitempty",
			JSON:    ",omitempty",
		},
	}
}

func (f *fieldsDef) MonitorInterfaceSend() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "Send",
		Type: meta.TypeFloat64,
		Tags: &dsl.FieldTags{
			MapConv: ",omitempty",
			JSON:    ",omitempty",
		},
	}
}

func (f *fieldsDef) MonitorInterfaceReceive() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "Receive",
		Type: meta.TypeFloat64,
		Tags: &dsl.FieldTags{
			MapConv: ",omitempty",
			JSON:    ",omitempty",
		},
	}
}

func (f *fieldsDef) MonitorFreeDiskSize() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "FreeDiskSize",
		Type: meta.TypeFloat64,
		Tags: &dsl.FieldTags{
			MapConv: ",omitempty",
			JSON:    ",omitempty",
		},
	}
}

func (f *fieldsDef) MonitorDatabaseTotalMemorySize() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "TotalMemorySize ",
		Type: meta.TypeFloat64,
		Tags: &dsl.FieldTags{
			MapConv: ",omitempty",
			JSON:    ",omitempty",
		},
	}
}

func (f *fieldsDef) MonitorDatabaseUsedMemorySize() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "UsedMemorySize",
		Type: meta.TypeFloat64,
		Tags: &dsl.FieldTags{
			MapConv: ",omitempty",
			JSON:    ",omitempty",
		},
	}
}

func (f *fieldsDef) MonitorDatabaseTotalDisk1Size() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "TotalDisk1Size",
		Type: meta.TypeFloat64,
		Tags: &dsl.FieldTags{
			MapConv: ",omitempty",
			JSON:    ",omitempty",
		},
	}
}

func (f *fieldsDef) MonitorDatabaseUsedDisk1Size() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "UsedDisk1Size",
		Type: meta.TypeFloat64,
		Tags: &dsl.FieldTags{
			MapConv: ",omitempty",
			JSON:    ",omitempty",
		},
	}
}

func (f *fieldsDef) MonitorDatabaseTotalDisk2Size() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "TotalDisk2Size",
		Type: meta.TypeFloat64,
		Tags: &dsl.FieldTags{
			MapConv: ",omitempty",
			JSON:    ",omitempty",
		},
	}
}

func (f *fieldsDef) MonitorDatabaseUsedDisk2Size() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "UsedDisk2Size",
		Type: meta.TypeFloat64,
		Tags: &dsl.FieldTags{
			MapConv: ",omitempty",
			JSON:    ",omitempty",
		},
	}
}

func (f *fieldsDef) MonitorDatabaseBinlogUsedSizeKiB() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "BinlogUsedSizeKiB",
		Type: meta.TypeFloat64,
		Tags: &dsl.FieldTags{
			MapConv: ",omitempty",
			JSON:    ",omitempty",
		},
	}
}

func (f *fieldsDef) MonitorDatabaseDelayTimeSec() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "DelayTimeSec",
		Type: meta.TypeFloat64,
		Tags: &dsl.FieldTags{
			MapConv: ",omitempty",
			JSON:    ",omitempty",
		},
	}
}

func (f *fieldsDef) MonitorResponseTimeSec() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "ResponseTimeSec",
		Type: meta.TypeFloat64,
		Tags: &dsl.FieldTags{
			MapConv: ",omitempty",
			JSON:    ",omitempty",
		},
	}
}

func (f *fieldsDef) MonitorUplinkBPS() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "UplinkBPS",
		Type: meta.TypeFloat64,
		Tags: &dsl.FieldTags{
			MapConv: ",omitempty",
			JSON:    ",omitempty",
		},
	}
}

func (f *fieldsDef) MonitorDownlinkBPS() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "DownlinkBPS",
		Type: meta.TypeFloat64,
		Tags: &dsl.FieldTags{
			MapConv: ",omitempty",
			JSON:    ",omitempty",
		},
	}
}

func (f *fieldsDef) MonitorActiveConnections() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "ActiveConnections",
		Type: meta.TypeFloat64,
		Tags: &dsl.FieldTags{
			MapConv: ",omitempty",
			JSON:    ",omitempty",
		},
	}
}

func (f *fieldsDef) MonitorConnectionsPerSec() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "ConnectionsPerSec",
		Type: meta.TypeFloat64,
		Tags: &dsl.FieldTags{
			MapConv: ",omitempty",
			JSON:    ",omitempty",
		},
	}
}
