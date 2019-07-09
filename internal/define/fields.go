package define

import (
	"github.com/sacloud/libsacloud/v2/internal/dsl"
	"github.com/sacloud/libsacloud/v2/internal/dsl/meta"
)

type fieldsDef struct{}

var fields = &fieldsDef{}

func (f *fieldsDef) New(name string, t meta.Type) *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: name,
		Type: t,
	}
}

func (f *fieldsDef) ID() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "ID",
		Type: meta.TypeID,
		ExtendAccessors: []*dsl.ExtendAccessor{
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

func (f *fieldsDef) Name() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "Name",
		Type: meta.TypeString,
		Tags: &dsl.FieldTags{
			Validate: "required",
		},
	}
}

func (f *fieldsDef) AccountID() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "AccountID",
		Type: meta.TypeID,
		Tags: &dsl.FieldTags{
			MapConv: "Account.ID",
		},
	}
}

func (f *fieldsDef) AccountName() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "AccountName",
		Type: meta.TypeString,
		Tags: &dsl.FieldTags{
			MapConv: "Account.Name",
		},
	}
}

func (f *fieldsDef) AccountCode() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "AccountCode",
		Type: meta.TypeString,
		Tags: &dsl.FieldTags{
			MapConv: "Account.Code",
		},
	}
}
func (f *fieldsDef) AccountClass() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "AccountClass",
		Type: meta.TypeString,
		Tags: &dsl.FieldTags{
			MapConv: "Account.Class",
		},
	}
}

func (f *fieldsDef) MemberCode() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "MemberCode",
		Type: meta.TypeString,
		Tags: &dsl.FieldTags{
			MapConv: "Member.Code",
		},
	}
}
func (f *fieldsDef) MemberClass() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "MemberClass",
		Type: meta.TypeString,
		Tags: &dsl.FieldTags{
			MapConv: "Member.Class",
		},
	}
}

func (f *fieldsDef) InterfaceDriver() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "InterfaceDriver",
		Type: meta.TypeInterfaceDriver,
		Tags: &dsl.FieldTags{
			MapConv: ",default=virtio",
		},
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

// TODO CPUとServerPlanCPUのようにmapconvのタグだけ違う値をどう扱うか

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
		ExtendAccessors: []*dsl.ExtendAccessor{
			{
				Name: "MemoryGB",
				Type: meta.TypeInt,
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
			MapConv: "Commitment,default=standard",
		},
		Type: meta.TypeCommitment,
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
		ExtendAccessors: []*dsl.ExtendAccessor{
			{
				Name: "MemoryGB",
				Type: meta.TypeInt,
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
			MapConv: "ServerPlan.Commitment,default=standard",
		},
		Type: meta.TypeCommitment,
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

func (f *fieldsDef) stringEnabled() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "Enabled",
		Tags: &dsl.FieldTags{
			MapConv: ",omitempty",
		},
		Type: meta.TypeStringFlag,
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

func (f *fieldsDef) IconID() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "IconID",
		Tags: &dsl.FieldTags{
			MapConv: "Icon.ID",
		},
		Type: meta.TypeID,
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
			MapConv: "CDROM.ID",
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
						MapConv:  ",default=10",
						Validate: "min=0,max=60", // TODO 最大値確認
					},
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
		Type: meta.TypeStringSlice,
	}
}

func (f *fieldsDef) Class() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "Class",
		Type: meta.TypeString,
	}
}

func (f *fieldsDef) NFSClass() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "Class",
		Type: meta.TypeString,
		Tags: &dsl.FieldTags{
			MapConv: ",default=nfs",
		},
	}
}

func (f *fieldsDef) LoadBalancerClass() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "Class",
		Type: meta.TypeString,
		Tags: &dsl.FieldTags{
			MapConv: ",default=loadbalancer",
		},
	}
}

func (f *fieldsDef) VPCRouterClass() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "Class",
		Type: meta.TypeString,
		Tags: &dsl.FieldTags{
			MapConv: ",default=vpcrouter",
		},
	}
}

func (f *fieldsDef) SIMProviderClass() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "Class",
		Type: meta.TypeString,
		Tags: &dsl.FieldTags{
			MapConv: "Provider.Class,default=sim",
		},
	}
}

func (f *fieldsDef) SIMICCID() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "ICCID",
		Type: meta.TypeStringNumber,
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

func (f *fieldsDef) GSLBProviderClass() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "Class",
		Type: meta.TypeString,
		Tags: &dsl.FieldTags{
			MapConv: "Provider.Class,default=gslb",
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

func (f *fieldsDef) GSLBHealthCheckProtocol() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "HealthCheckProtocol",
		Type: meta.TypeProtocol,
		Tags: &dsl.FieldTags{
			MapConv:  "Settings.GSLB.HealthCheck.Protocol",
			Validate: "oneof=http https ping tcp",
		},
	}
}

func (f *fieldsDef) GSLBHealthCheckHostHeader() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "HealthCheckHostHeader",
		Type: meta.TypeString,
		Tags: &dsl.FieldTags{
			MapConv: "Settings.GSLB.HealthCheck.Host",
		},
	}
}

func (f *fieldsDef) GSLBHealthCheckPath() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "HealthCheckPath",
		Type: meta.TypeString,
		Tags: &dsl.FieldTags{
			MapConv: "Settings.GSLB.HealthCheck.Path",
		},
	}
}

func (f *fieldsDef) GSLBHealthCheckResponseCode() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "HealthCheckResponseCode",
		Type: meta.TypeStringNumber,
		Tags: &dsl.FieldTags{
			MapConv: "Settings.GSLB.HealthCheck.Status",
		},
	}
}

func (f *fieldsDef) GSLBHealthCheckPort() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "HealthCheckPort",
		Type: meta.TypeStringNumber,
		Tags: &dsl.FieldTags{
			MapConv: "Settings.GSLB.HealthCheck.Port",
		},
	}
}

func (f *fieldsDef) GSLBDelayLoop() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "DelayLoop",
		Type: meta.TypeInt,
		Tags: &dsl.FieldTags{
			Validate: "min=10,max=60",
			MapConv:  "Settings.GSLB.DelayLoop,default=10",
		},
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
					Tags: &dsl.FieldTags{
						MapConv: ",default=1",
					},
				},
			},
		},
		Tags: &dsl.FieldTags{
			MapConv:  "Settings.GSLB.[]Servers,recursive",
			Validate: "min=0,max=6",
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

func (f *fieldsDef) AutoBackupProviderClass() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "Class",
		Type: meta.TypeString,
		Tags: &dsl.FieldTags{
			MapConv: "Provider.Class,default=autobackup",
		},
	}
}

func (f *fieldsDef) AutoBackupBackupSpanType() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "BackupSpanType",
		Type: meta.TypeBackupSpanType,
		Tags: &dsl.FieldTags{
			MapConv: "Settings.Autobackup.BackupSpanType,default=weekdays",
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

func (f *fieldsDef) DNSProviderClass() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "Class",
		Type: meta.TypeString,
		Tags: &dsl.FieldTags{
			MapConv: "Provider.Class,default=dns",
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
		ExtendAccessors: []*dsl.ExtendAccessor{
			{Name: "SizeGB"},
		},
	}
}

func (f *fieldsDef) MigratedMB() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "MigratedMB",
		Type: meta.TypeInt,
		ExtendAccessors: []*dsl.ExtendAccessor{
			{Name: "MigratedGB"},
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
		Type: meta.TypeInt,
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
