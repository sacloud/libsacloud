package define

import (
	"net/http"

	"github.com/sacloud/libsacloud/v2/internal/define/names"
	"github.com/sacloud/libsacloud/v2/internal/define/ops"
	"github.com/sacloud/libsacloud/v2/internal/dsl"
	"github.com/sacloud/libsacloud/v2/internal/dsl/meta"
	"github.com/sacloud/libsacloud/v2/sacloud/naked"
)

const (
	databaseAPIName     = "Database"
	databaseAPIPathName = "appliance"
)

var databaseAPI = &dsl.Resource{
	Name:       databaseAPIName,
	PathName:   databaseAPIPathName,
	PathSuffix: dsl.CloudAPISuffix,
	Operations: dsl.Operations{
		// find
		ops.FindAppliance(databaseAPIName, databaseNakedType, findParameter, databaseView),

		// create
		ops.CreateAppliance(databaseAPIName, databaseNakedType, databaseCreateParam, databaseView),

		// read
		ops.ReadAppliance(databaseAPIName, databaseNakedType, databaseView),

		// update
		ops.UpdateAppliance(databaseAPIName, databaseNakedType, databaseUpdateParam, databaseView),

		// delete
		ops.Delete(databaseAPIName),

		// config
		ops.Config(databaseAPIName),

		// power management(boot/shutdown/reset)
		ops.Boot(databaseAPIName),
		ops.Shutdown(databaseAPIName),
		ops.Reset(databaseAPIName),

		// monitor
		ops.MonitorChild(databaseAPIName, "CPU", "cpu",
			monitorParameter, monitors.cpuTimeModel()),
		ops.MonitorChild(databaseAPIName, "Disk", "disk/0",
			monitorParameter, monitors.diskModel()),
		ops.MonitorChild(databaseAPIName, "Interface", "interface",
			monitorParameter, monitors.interfaceModel()),
		ops.MonitorChild(databaseAPIName, "Database", "database",
			monitorParameter, monitors.databaseModel()),

		// status
		{
			ResourceName: databaseAPIName,
			Name:         "Status",
			PathFormat:   dsl.IDAndSuffixPathFormat("status"),
			Method:       http.MethodGet,
			Arguments: dsl.Arguments{
				dsl.ArgumentID,
			},
			ResponseEnvelope: dsl.ResponseEnvelope(&dsl.EnvelopePayloadDesc{
				Type: meta.Static(naked.DatabaseStatusResponse{}),
				Name: "Appliance",
			}),
			Results: dsl.Results{
				{
					SourceField: "Appliance",
					DestField:   databaseStatusView.Name,
					IsPlural:    false,
					Model:       databaseStatusView,
				},
			},
		},

		// TODO ログやバックアップはクライアントからの利用頻度低(usacloudのみ)なため、現状非対応とする。
	},
}

var (
	databaseNakedType       = meta.Static(naked.Database{})
	databaseStatusNakedType = meta.Static(naked.DatabaseStatus{})

	databaseView = &dsl.Model{
		Name:      databaseAPIName,
		NakedType: databaseNakedType,
		Fields: []*dsl.FieldDesc{
			fields.ID(),
			fields.Class(),
			fields.Name(),
			fields.Description(),
			fields.Tags(),
			fields.Availability(),
			fields.IconID(),
			fields.CreatedAt(),
			fields.ModifiedAt(),
			// settings
			fields.DatabaseSettingsCommon(),
			fields.DatabaseSettingsBackup(),
			fields.DatabaseSettingsReplication(),
			fields.SettingsHash(),

			// instance
			fields.InstanceHostName(),
			fields.InstanceHostInfoURL(),
			fields.InstanceStatus(),
			fields.InstanceStatusChangedAt(),
			// plan
			fields.AppliancePlanID(),
			// switch
			fields.ApplianceSwitchID(),
			// remark
			fields.RemarkDBConf(),
			fields.RemarkDefaultRoute(),
			fields.RemarkNetworkMaskLen(),
			fields.RemarkServerIPAddress(),
			fields.RemarkZoneID(),
			// interfaces
			fields.Interfaces(),
		},
	}

	databaseCreateParam = &dsl.Model{
		Name:      names.CreateParameterName(databaseAPIName),
		NakedType: databaseNakedType,
		ConstFields: []*dsl.ConstFieldDesc{
			{
				Name:  "Class",
				Type:  meta.TypeString,
				Value: `"database"`,
			},
		},
		Fields: []*dsl.FieldDesc{
			fields.AppliancePlanID(),
			fields.ApplianceSwitchID(),
			fields.ApplianceIPAddresses(),
			fields.RemarkNetworkMaskLen(),
			fields.RemarkDefaultRoute(),
			fields.RemarkDBConf(),
			fields.RemarkSourceAppliance(),

			fields.DatabaseSettingsCommon(),
			fields.DatabaseSettingsBackup(),
			fields.DatabaseSettingsReplication(),

			fields.Name(),
			fields.Description(),
			fields.Tags(),
			fields.IconID(),
		},
	}

	databaseUpdateParam = &dsl.Model{
		Name:      names.UpdateParameterName(databaseAPIName),
		NakedType: databaseNakedType,
		Fields: []*dsl.FieldDesc{
			fields.DatabaseSettingsCommonUpdate(),
			fields.DatabaseSettingsBackup(),
			fields.DatabaseSettingsReplication(),

			fields.Name(),
			fields.Description(),
			fields.Tags(),
			fields.IconID(),
		},
	}

	databaseStatusView = &dsl.Model{
		Name:      "DatabaseStatus",
		NakedType: databaseStatusNakedType,
		Fields: []*dsl.FieldDesc{
			{
				Name: "Status",
				Type: meta.TypeInstanceStatus,
				Tags: &dsl.FieldTags{
					MapConv: "SettingsResponse.Status",
				},
			},
			{
				Name: "IsFatal",
				Type: meta.TypeFlag,
				Tags: &dsl.FieldTags{
					MapConv: "SettingsResponse.IsFatal",
				},
			},
			{
				Name: "Version",
				Type: databaseStatusVersionView,
				Tags: &dsl.FieldTags{
					MapConv: "SettingsResponse.DBConf.Version,recursive",
				},
			},
			{
				Name: "Logs",
				Type: databaseStatusLogView,
				Tags: &dsl.FieldTags{
					MapConv: "SettingsResponse.DBConf.[]Log,recursive",
				},
			},
			{
				Name: "Backups",
				Type: databaseStatusBackupHistoryView,
				Tags: &dsl.FieldTags{
					MapConv: "SettingsResponse.DBConf.Backup.[]History,recursive",
				},
			},
		},
	}

	databaseStatusVersionView = &dsl.Model{
		Name:      "DatabaseVersionInfo",
		NakedType: meta.Static(naked.DatabaseStatusVersion{}),
		Fields: []*dsl.FieldDesc{
			fields.Def("LastModified", meta.TypeString),
			fields.Def("CommitHash", meta.TypeString),
			fields.Def("Status", meta.TypeString),
			fields.Def("Tag", meta.TypeString),
			fields.Def("Expire", meta.TypeString),
		},
	}

	databaseStatusLogView = &dsl.Model{
		Name:      "DatabaseLog",
		NakedType: meta.Static(naked.DatabaseLog{}),
		IsArray:   true,
		Fields: []*dsl.FieldDesc{
			fields.Def("Name", meta.TypeString),
			fields.Def("Data", meta.TypeString),
			fields.Def("Size", meta.TypeInt),
		},
	}
	databaseStatusBackupHistoryView = &dsl.Model{
		Name:      "DatabaseBackupHistory",
		NakedType: meta.Static(naked.DatabaseBackupHistory{}),
		IsArray:   true,
		Fields: []*dsl.FieldDesc{
			fields.Def("CreatedAt", meta.TypeTime),
			fields.Def("Availability", meta.TypeString),
			fields.Def("RecoveredAt", meta.TypeTime),
			fields.Def("Size", meta.TypeInt64),
		},
	}
)
