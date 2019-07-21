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
	vpcRouterAPIName     = "VPCRouter"
	vpcRouterAPIPathName = "appliance"
)

var vpcRouterAPI = &dsl.Resource{
	Name:       vpcRouterAPIName,
	PathName:   vpcRouterAPIPathName,
	PathSuffix: dsl.CloudAPISuffix,
	Operations: dsl.Operations{
		// find
		ops.FindAppliance(vpcRouterAPIName, vpcRouterNakedType, findParameter, vpcRouterView, nil),

		// create
		ops.CreateAppliance(vpcRouterAPIName, vpcRouterNakedType, vpcRouterCreateParam, vpcRouterView),

		// read
		ops.ReadAppliance(vpcRouterAPIName, vpcRouterNakedType, vpcRouterView),

		// update
		ops.UpdateAppliance(vpcRouterAPIName, vpcRouterNakedType, vpcRouterUpdateParam, vpcRouterView),

		// delete
		ops.Delete(vpcRouterAPIName),

		// config
		ops.Config(vpcRouterAPIName),

		// power management(boot/shutdown/reset)
		ops.Boot(vpcRouterAPIName),
		ops.Shutdown(vpcRouterAPIName),
		ops.Reset(vpcRouterAPIName),

		// connect to switch
		ops.WithIDAction(
			vpcRouterAPIName, "ConnectToSwitch", http.MethodPut, "interface/{{.nicIndex}}/to/switch/{{.switchID}}",
			&dsl.Argument{
				Name: "nicIndex",
				Type: meta.TypeInt,
			},
			&dsl.Argument{
				Name: "switchID",
				Type: meta.TypeID,
			},
		),

		// disconnect from switch
		ops.WithIDAction(
			vpcRouterAPIName, "DisconnectFromSwitch", http.MethodDelete, "interface/{{.nicIndex}}/to/switch",
			&dsl.Argument{
				Name: "nicIndex",
				Type: meta.TypeInt,
			},
		),

		// monitor
		ops.MonitorChildBy(vpcRouterAPIName, "Interface", "interface",
			monitorParameter, monitors.interfaceModel()),
	},
}

var (
	vpcRouterNakedType = meta.Static(naked.VPCRouter{})

	vpcRouterView = &dsl.Model{
		Name:      vpcRouterAPIName,
		NakedType: vpcRouterNakedType,
		Fields: []*dsl.FieldDesc{
			fields.ID(),
			fields.Name(),
			fields.Description(),
			fields.Tags(),
			fields.Availability(),
			fields.Class(),
			fields.IconID(),
			fields.CreatedAt(),
			// plan
			fields.AppliancePlanID(),
			// settings
			fields.SettingsHash(),
			{
				Name: "Settings",
				Type: models.vpcRouterSetting(),
				Tags: &dsl.FieldTags{
					MapConv: ",omitempty,recursive",
				},
			},
			// instance
			fields.InstanceHostName(),
			fields.InstanceHostInfoURL(),
			fields.InstanceStatus(),
			fields.InstanceStatusChangedAt(),
			// interfaces
			fields.VPCRouterInterfaces(),
			// switch
			fields.ApplianceSwitchID(),
			// remark
			fields.RemarkServerIPAddress(),
			fields.RemarkZoneID(),
		},
	}

	vpcRouterCreateParam = &dsl.Model{
		Name:      names.CreateParameterName(vpcRouterAPIName),
		NakedType: vpcRouterNakedType,
		ConstFields: []*dsl.ConstFieldDesc{
			{
				Name:  "Class",
				Type:  meta.TypeString,
				Value: `"vpcrouter"`,
			},
		},
		Fields: []*dsl.FieldDesc{
			fields.Name(),
			fields.Description(),
			fields.Tags(),
			fields.IconID(),
			fields.PlanID(),

			// nic
			{
				Name: "Switch",
				Type: &dsl.Model{
					Name: "ApplianceConnectedSwitch",
					Fields: []*dsl.FieldDesc{
						fields.ID(),
						fields.Scope(),
					},
					NakedType: meta.Static(naked.ConnectedSwitch{}),
				},
				Tags: &dsl.FieldTags{
					JSON:    ",omitempty",
					MapConv: "Remark.Switch,recursive",
				},
			},

			// TODO remarkとsettings.Interfaces両方に設定する必要がある。うまい方法が思いつかないため当面は利用者側で両方に設定する方法としておく
			fields.ApplianceIPAddresses(),

			{
				Name: "Settings",
				Type: models.vpcRouterSetting(),
				Tags: &dsl.FieldTags{
					MapConv: ",omitempty,recursive",
				},
			},
		},
	}

	vpcRouterUpdateParam = &dsl.Model{
		Name:      names.UpdateParameterName(vpcRouterAPIName),
		NakedType: vpcRouterNakedType,
		Fields: []*dsl.FieldDesc{
			fields.Name(),
			fields.Description(),
			fields.Tags(),
			fields.IconID(),
			{
				Name: "Settings",
				Type: models.vpcRouterSetting(),
				Tags: &dsl.FieldTags{
					MapConv: ",omitempty,recursive",
				},
			},
		},
	}
)
