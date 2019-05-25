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

	Resources.DefineWith("VPCRouter", func(r *schema.Resource) {
		r.Operations(
			// find
			r.DefineOperationApplianceFind(nakedType, findParameter, vpcRouter),

			// create
			r.DefineOperationApplianceCreate(nakedType, createParam, vpcRouter),

			// read
			r.DefineOperationApplianceRead(nakedType, vpcRouter),

			// update
			r.DefineOperationApplianceUpdate(nakedType, updateParam, vpcRouter),

			// delete
			r.DefineOperationDelete(),

			// config
			r.DefineOperationConfig(),

			// power management(boot/shutdown/reset)
			r.DefineOperationBoot(),
			r.DefineOperationShutdown(),
			r.DefineOperationReset(),

			// connect to switch
			r.DefineOperation("ConnectToSwitch").
				Method(http.MethodPut).
				PathFormat(schema.IDAndSuffixPathFormat("interface/{{.nicIndex}}/to/switch/{{.switchID}}")).
				Argument(schema.ArgumentZone).
				Argument(schema.ArgumentID).
				Argument(&schema.SimpleArgument{
					Name: "nicIndex",
					Type: meta.TypeInt,
				}).
				Argument(&schema.SimpleArgument{
					Name: "switchID",
					Type: meta.TypeID,
				}),

			// disconnect from switch
			r.DefineOperation("DisconnectFromSwitch").
				Method(http.MethodDelete).
				PathFormat(schema.IDAndSuffixPathFormat("interface/{{.nicIndex}}/to/switch")).
				Argument(schema.ArgumentZone).
				Argument(schema.ArgumentID).
				Argument(&schema.SimpleArgument{
					Name: "nicIndex",
					Type: meta.TypeInt,
				}),

			// monitor
			r.DefineOperationMonitorChildBy("Interface", "interface",
				monitorParameter, monitors.interfaceModel()),
		)
	}).PathName("appliance")
}
