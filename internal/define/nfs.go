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
	nfsAPI := &schema.Resource{
		Name:       "NFS",
		PathName:   "appliance",
		PathSuffix: schema.CloudAPISuffix,
	}
	nfsAPI.Operations = []*schema.Operation{
		// find
		nfsAPI.DefineOperationApplianceFind(nakedType, findParameter, nfs),

		// create
		nfsAPI.DefineOperationApplianceCreate(nakedType, createParam, nfs),

		// read
		nfsAPI.DefineOperationApplianceRead(nakedType, nfs),

		// update
		nfsAPI.DefineOperationApplianceUpdate(nakedType, updateParam, nfs),

		// delete
		nfsAPI.DefineOperationDelete(),

		// power management(boot/shutdown/reset)
		nfsAPI.DefineOperationBoot(),
		nfsAPI.DefineOperationShutdown(),
		nfsAPI.DefineOperationReset(),

		// monitor
		nfsAPI.DefineOperationMonitorChild("FreeDiskSize", "database",
			monitorParameter, monitors.freeDiskSizeModel()),
		nfsAPI.DefineOperationMonitorChild("Interface", "interface",
			monitorParameter, monitors.interfaceModel()),
	}
	Resources.Def(nfsAPI)
}
