package define

import (
	"github.com/sacloud/libsacloud/v2/internal/schema"
	"github.com/sacloud/libsacloud/v2/internal/schema/meta"
	"github.com/sacloud/libsacloud/v2/sacloud/naked"
)

var nfsAPI = &schema.Resource{
	Name:       "NFS",
	PathName:   "appliance",
	PathSuffix: schema.CloudAPISuffix,
	OperationsDefineFunc: func(r *schema.Resource) []*schema.Operation {
		return []*schema.Operation{
			// find
			r.DefineOperationApplianceFind(nfsNakedType, findParameter, nfsView),

			// create
			r.DefineOperationApplianceCreate(nfsNakedType, nfsCreateParam, nfsView),

			// read
			r.DefineOperationApplianceRead(nfsNakedType, nfsView),

			// update
			r.DefineOperationApplianceUpdate(nfsNakedType, nfsUpdateParam, nfsView),

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
		}
	},
}

var (
	nfsNakedType = meta.Static(naked.NFS{})

	nfsView = &schema.Model{
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

	nfsCreateParam = &schema.Model{
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

	nfsUpdateParam = &schema.Model{
		Fields: []*schema.FieldDesc{
			fields.Name(),
			fields.Description(),
			fields.Tags(),
			fields.IconID(),
		},
	}
)
