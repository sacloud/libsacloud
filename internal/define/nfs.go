package define

import (
	"github.com/sacloud/libsacloud/v2/internal/define/ops"
	"github.com/sacloud/libsacloud/v2/internal/schema"
	"github.com/sacloud/libsacloud/v2/internal/schema/meta"
	"github.com/sacloud/libsacloud/v2/sacloud/naked"
)

const (
	nfsAPIName     = "NFS"
	nfsAPIPathName = "appliance"
)

var nfsAPI = &schema.Resource{
	Name:       nfsAPIName,
	PathName:   nfsAPIPathName,
	PathSuffix: schema.CloudAPISuffix,
	Operations: schema.Operations{
		// find
		ops.FindAppliance(nfsAPIName, nfsNakedType, findParameter, nfsView),

		// create
		ops.CreateAppliance(nfsAPIName, nfsNakedType, nfsCreateParam, nfsView),

		// read
		ops.ReadAppliance(nfsAPIName, nfsNakedType, nfsView),

		// update
		ops.UpdateAppliance(nfsAPIName, nfsNakedType, nfsUpdateParam, nfsView),

		// delete
		ops.Delete(nfsAPIName),

		// power management(boot/shutdown/reset)
		ops.Boot(nfsAPIName),
		ops.Shutdown(nfsAPIName),
		ops.Reset(nfsAPIName),

		// monitor
		ops.MonitorChild(nfsAPIName, "FreeDiskSize", "database",
			monitorParameter, monitors.freeDiskSizeModel()),
		ops.MonitorChild(nfsAPIName, "Interface", "interface",
			monitorParameter, monitors.interfaceModel()),
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
