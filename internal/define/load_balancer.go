package define

import (
	"github.com/sacloud/libsacloud/internal/schema"
	"github.com/sacloud/libsacloud/internal/schema/meta"
	"github.com/sacloud/libsacloud/sacloud/naked"
)

var loadBalancerAPI = &schema.Resource{
	Name:       "LoadBalancer",
	PathName:   "appliance",
	PathSuffix: schema.CloudAPISuffix,
	OperationsDefineFunc: func(r *schema.Resource) []*schema.Operation {
		return []*schema.Operation{
			// find
			r.DefineOperationApplianceFind(loadBalancerNakedType, findParameter, loadBalancerView),

			// create
			r.DefineOperationApplianceCreate(loadBalancerNakedType, loadBalancerCreateParam, loadBalancerView),

			// read
			r.DefineOperationApplianceRead(loadBalancerNakedType, loadBalancerView),

			// update
			r.DefineOperationApplianceUpdate(loadBalancerNakedType, loadBalancerUpdateParam, loadBalancerView),

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

			// status
			r.DefineOperationStatus(meta.Static(naked.LoadBalancerStatus{}), loadBalancerStatus),
		}
	},
}

var (
	loadBalancerNakedType = meta.Static(naked.LoadBalancer{})

	loadBalancerView = &schema.Model{
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
