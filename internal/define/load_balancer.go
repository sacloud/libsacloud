package define

import (
	"github.com/sacloud/libsacloud-v2/internal/schema"
	"github.com/sacloud/libsacloud-v2/internal/schema/meta"
	"github.com/sacloud/libsacloud-v2/sacloud/naked"
)

func init() {
	nakedType := meta.Static(naked.LoadBalancer{})

	loadBalancer := &schema.Model{
		Fields: []*schema.FieldDesc{
			fields.ID(),
			fields.Name(),
			fields.Description(),
			fields.Tags(),
			fields.Availability(),
			fields.Class(),
			fields.Icon(),
			fields.CreatedAt(),
			fields.ModifiedAt(),
			// instance
			fields.InstanceHostName(),
			fields.InstanceHostInfoURL(),
			fields.InstanceStatus(),
			fields.InstanceStatusChangedAt(),
			// interfaces
			fields.Interfaces(),
			// plan
			fields.AppliancePlanID(),
			// switch
			fields.ApplianceSwitchID(),
			fields.Switch(),
			// remark
			fields.RemarkDefaultRoute(),
			fields.RemarkNetworkMaskLen(),
			fields.RemarkServerIPAddress(),
			fields.RemarkZoneID(),
			fields.RemarkVRID(),
			// settings
			fields.LoadBalancerVIP(),
			fields.SettingsHash(),
		},
	}

	createParam := &schema.Model{
		Fields: []*schema.FieldDesc{
			fields.LoadBalancerClass(),
			fields.ApplianceSwitchID(),
			fields.AppliancePlanID(),
			fields.RemarkVRID(),
			fields.ApplianceIPAddresses(),
			fields.RemarkNetworkMaskLen(),
			fields.RemarkDefaultRoute(),
			fields.Name(),
			fields.Description(),
			fields.Tags(),
			fields.IconID(),
			fields.LoadBalancerVIP(),
		},
	}

	updateParam := &schema.Model{
		Fields: []*schema.FieldDesc{
			fields.Name(),
			fields.Description(),
			fields.Tags(),
			fields.IconID(),
			fields.LoadBalancerVIP(),
		},
	}

	Resources.DefineWith("LoadBalancer", func(r *schema.Resource) {
		r.Operations(
			// find
			r.DefineOperationApplianceFind(nakedType, findParameter, loadBalancer),

			// create
			r.DefineOperationApplianceCreate(nakedType, createParam, loadBalancer),

			// read
			r.DefineOperationApplianceRead(nakedType, loadBalancer),

			// update
			r.DefineOperationApplianceUpdate(nakedType, updateParam, loadBalancer),

			// delete
			r.DefineOperationDelete(),

			// config
			r.DefineOperationConfig(),

			// power management(boot/shutdown/reset)
			r.DefineOperationBoot(),
			r.DefineOperationShutdown(),
			r.DefineOperationReset(),

			// monitor
			r.DefineOperationMonitorChild("Interface", "interface",
				monitorParameter, monitors.interfaceModel()),
		)
	}).PathName("appliance")
}
