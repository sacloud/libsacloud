package define

import (
	"github.com/sacloud/libsacloud/v2/internal/define/names"
	"github.com/sacloud/libsacloud/v2/internal/define/ops"
	"github.com/sacloud/libsacloud/v2/internal/schema"
	"github.com/sacloud/libsacloud/v2/internal/schema/meta"
	"github.com/sacloud/libsacloud/v2/sacloud/naked"
)

const (
	loadBalancerAPIName     = "LoadBalancer"
	loadBalancerAPIPathName = "appliance"
)

var loadBalancerAPI = &schema.Resource{
	Name:       loadBalancerAPIName,
	PathName:   loadBalancerAPIPathName,
	PathSuffix: schema.CloudAPISuffix,
	Operations: schema.Operations{
		// find
		ops.FindAppliance(loadBalancerAPIName, loadBalancerNakedType, findParameter, loadBalancerView),

		// create
		ops.CreateAppliance(loadBalancerAPIName, loadBalancerNakedType, loadBalancerCreateParam, loadBalancerView),

		// read
		ops.ReadAppliance(loadBalancerAPIName, loadBalancerNakedType, loadBalancerView),

		// update
		ops.UpdateAppliance(loadBalancerAPIName, loadBalancerNakedType, loadBalancerUpdateParam, loadBalancerView),

		// delete
		ops.Delete(loadBalancerAPIName),

		// config
		ops.Config(loadBalancerAPIName),

		// power management(boot/shutdown/reset)
		ops.Boot(loadBalancerAPIName),
		ops.Shutdown(loadBalancerAPIName),
		ops.Reset(loadBalancerAPIName),

		// monitor
		ops.MonitorChild(loadBalancerAPIName, "Interface", "interface",
			monitorParameter, monitors.interfaceModel()),

		// status
		ops.Status(loadBalancerAPIName, meta.Static(naked.LoadBalancerStatus{}), loadBalancerStatus),
	},
}

var (
	loadBalancerNakedType = meta.Static(naked.LoadBalancer{})

	loadBalancerView = &schema.Model{
		Name:      loadBalancerAPIName,
		NakedType: loadBalancerNakedType,
		Fields: []*schema.FieldDesc{
			fields.ID(),
			fields.Name(),
			fields.Description(),
			fields.Tags(),
			fields.Availability(),
			fields.Class(),
			fields.IconID(),
			fields.CreatedAt(),
			fields.ModifiedAt(),
			// instance
			fields.InstanceHostName(),
			fields.InstanceHostInfoURL(),
			fields.InstanceStatus(),
			fields.InstanceStatusChangedAt(),
			// interfaces
			//fields.Interfaces(),
			// plan
			fields.AppliancePlanID(),
			// switch
			fields.ApplianceSwitchID(),
			// remark
			fields.RemarkDefaultRoute(),
			fields.RemarkNetworkMaskLen(),
			fields.RemarkServerIPAddress(),
			fields.RemarkZoneID(),
			fields.RemarkVRID(),
			// settings
			fields.LoadBalancerVIP(),
			fields.SettingsHash(),
			// interfaces
			fields.Interfaces(),
		},
	}

	loadBalancerCreateParam = &schema.Model{
		Name:      names.CreateParameterName(loadBalancerAPIName),
		NakedType: loadBalancerNakedType,
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

	loadBalancerUpdateParam = &schema.Model{
		Name:      names.UpdateParameterName(loadBalancerAPIName),
		NakedType: loadBalancerNakedType,
		Fields: []*schema.FieldDesc{
			fields.Name(),
			fields.Description(),
			fields.Tags(),
			fields.IconID(),
			fields.LoadBalancerVIP(),
		},
	}

	loadBalancerStatus = &schema.Model{
		Name: "LoadBalancerStatus",
		Fields: []*schema.FieldDesc{
			{
				Name: "VirtualIPAddress",
				Type: meta.TypeString,
			},
			{
				Name: "Port",
				Type: meta.TypeStringNumber,
			},
			{
				Name: "CPS",
				Type: meta.TypeStringNumber,
			},
			{
				Name: "Servers",
				Type: &schema.Model{
					Name:    "LoadBalancerServerStatus",
					IsArray: true,
					Fields: []*schema.FieldDesc{
						{
							Name: "ActiveConn",
							Type: meta.TypeStringNumber,
						},
						{
							Name: "Status",
							Type: meta.TypeInstanceStatus,
						},
						{
							Name: "IPAddress",
							Type: meta.TypeString,
						},
						{
							Name: "Port",
							Type: meta.TypeStringNumber,
						},
						{
							Name: "CPS",
							Type: meta.TypeStringNumber,
						},
					},
				},
				Tags: &schema.FieldTags{
					MapConv: ",recursive",
				},
			},
		},
		NakedType: meta.Static(naked.LoadBalancerStatus{}),
	}
)
