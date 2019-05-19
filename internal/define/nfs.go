package define

import (
	"github.com/sacloud/libsacloud-v2/internal/schema"
	"github.com/sacloud/libsacloud-v2/internal/schema/meta"
	"github.com/sacloud/libsacloud-v2/sacloud/naked"
)

func init() {
	nakedType := meta.Static(naked.NFS{})

	nfs := &schema.Model{
		Fields: []*schema.FieldDesc{
			fields.ID(),
			fields.Name(),
			fields.Description(),
			fields.Tags(),
			fields.Availability(),
			fields.Class(),
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
			//fields.Switch(),
			// remark
			fields.RemarkDefaultRoute(),
			fields.RemarkNetworkMaskLen(),
			fields.RemarkServerIPAddress(),
			fields.RemarkZoneID(),
			// other
			fields.IconID(),
			fields.CreatedAt(),
			fields.ModifiedAt(),
		},
	}

	createParam := &schema.Model{
		Fields: []*schema.FieldDesc{
			fields.NFSClass(),
			fields.ApplianceSwitchID(),
			fields.AppliancePlanID(),
			fields.ApplianceIPAddresses(),
			fields.RemarkNetworkMaskLen(),
			fields.RemarkDefaultRoute(),
			fields.Name(),
			fields.Description(),
			fields.Tags(),
			fields.IconID(),
		},
	}

	updateParam := &schema.Model{
		Fields: []*schema.FieldDesc{
			fields.Name(),
			fields.Description(),
			fields.Tags(),
			fields.IconID(),
		},
	}

	Resources.DefineWith("NFS", func(r *schema.Resource) {
		r.Operations(
			// find
			r.DefineOperationApplianceFind(nakedType, findParameter, nfs),

			// create
			r.DefineOperationApplianceCreate(nakedType, createParam, nfs),

			// read
			r.DefineOperationApplianceRead(nakedType, nfs),

			// update
			r.DefineOperationApplianceUpdate(nakedType, updateParam, nfs),

			// delete
			r.DefineOperationDelete(),

			// power management(boot/shutdown/reset)
			r.DefineOperationBoot(),
			r.DefineOperationShutdown(),
			r.DefineOperationReset(),

			// monitor
			r.DefineOperationMonitorChild("FreeDiskSize", "database",
				monitorParameter, monitors.freeDiskSizeModel()),
			r.DefineOperationMonitorChild("Interface", "interface",
				monitorParameter, monitors.interfaceModel()),
		)
	}).PathName("appliance")
}
