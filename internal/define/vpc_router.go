package define

import (
	"net/http"

	"github.com/sacloud/libsacloud-v2/internal/schema"
	"github.com/sacloud/libsacloud-v2/internal/schema/meta"
	"github.com/sacloud/libsacloud-v2/sacloud/naked"
)

func init() {
	nakedType := meta.Static(naked.VPCRouter{})

	vpcRouter := &schema.Model{
		Fields: []*schema.FieldDesc{
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
				Tags: &schema.FieldTags{
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

	createParam := &schema.Model{
		Fields: []*schema.FieldDesc{
			fields.VPCRouterClass(),
			fields.Name(),
			fields.Description(),
			fields.Tags(),
			fields.IconID(),
			fields.PlanID(),

			// nic
			{
				Name: "Switch",
				Type: &schema.Model{
					Name: "ApplianceConnectedSwitch",
					Fields: []*schema.FieldDesc{
						fields.ID(),
						fields.Scope(),
					},
					NakedType: meta.Static(naked.ConnectedSwitch{}),
				},
				Tags: &schema.FieldTags{
					JSON:    ",omitempty",
					MapConv: "Remark.Switch,recursive",
				},
			},

			// TODO remarkとsettings.Interfaces両方に設定する必要がある。うまい方法が思いつかないため当面は利用者側で両方に設定する方法としておく
			fields.ApplianceIPAddresses(),

			{
				Name: "Settings",
				Type: models.vpcRouterSetting(),
				Tags: &schema.FieldTags{
					MapConv: ",omitempty,recursive",
				},
			},
		},
	}

	updateParam := &schema.Model{
		Fields: []*schema.FieldDesc{
			fields.Name(),
			fields.Description(),
			fields.Tags(),
			fields.IconID(),
			{
				Name: "Settings",
				Type: models.vpcRouterSetting(),
				Tags: &schema.FieldTags{
					MapConv: ",omitempty,recursive",
				},
			},
		},
	}

	api := &schema.Resource{
		Name:       "VPCRouter",
		PathName:   "appliance",
		PathSuffix: schema.CloudAPISuffix,
	}

	api.Operations = []*schema.Operation{
		// find
		api.DefineOperationApplianceFind(nakedType, findParameter, vpcRouter),

		// create
		api.DefineOperationApplianceCreate(nakedType, createParam, vpcRouter),

		// read
		api.DefineOperationApplianceRead(nakedType, vpcRouter),

		// update
		api.DefineOperationApplianceUpdate(nakedType, updateParam, vpcRouter),

		// delete
		api.DefineOperationDelete(),

		// config
		api.DefineOperationConfig(),

		// power management(boot/shutdown/reset)
		api.DefineOperationBoot(),
		api.DefineOperationShutdown(),
		api.DefineOperationReset(),

		// connect to switch
		api.DefineSimpleOperation("ConnectToSwitch", http.MethodPut, "interface/{{.nicIndex}}/to/switch/{{.switchID}}",
			&schema.Argument{
				Name: "nicIndex",
				Type: meta.TypeInt,
			},
			&schema.Argument{
				Name: "switchID",
				Type: meta.TypeID,
			},
		),

		// disconnect from switch
		api.DefineSimpleOperation("DisconnectFromSwitch", http.MethodDelete, "interface/{{.nicIndex}}/to/switch",
			&schema.Argument{
				Name: "nicIndex",
				Type: meta.TypeInt,
			},
		),

		// monitor
		api.DefineOperationMonitorChildBy("Interface", "interface",
			monitorParameter, monitors.interfaceModel()),
	}
	Resources.Def(api)
}
