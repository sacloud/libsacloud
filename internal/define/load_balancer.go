package define

import (
	"github.com/sacloud/libsacloud/v2/internal/define/names"
	"github.com/sacloud/libsacloud/v2/internal/define/ops"
	"github.com/sacloud/libsacloud/v2/internal/dsl"
	"github.com/sacloud/libsacloud/v2/internal/dsl/meta"
	"github.com/sacloud/libsacloud/v2/sacloud/naked"
)

const (
	loadBalancerAPIName     = "LoadBalancer"
	loadBalancerAPIPathName = "appliance"
)

var loadBalancerAPI = &dsl.Resource{
	Name:       loadBalancerAPIName,
	PathName:   loadBalancerAPIPathName,
	PathSuffix: dsl.CloudAPISuffix,
	Operations: dsl.Operations{
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

	loadBalancerView = &dsl.Model{
		Name:      loadBalancerAPIName,
		NakedType: loadBalancerNakedType,
		Fields: []*dsl.FieldDesc{
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

	loadBalancerCreateParam = &dsl.Model{
		Name:      names.CreateParameterName(loadBalancerAPIName),
		NakedType: loadBalancerNakedType,
		ConstFields: []*dsl.ConstFieldDesc{
			{
				Name:  "Class",
				Type:  meta.TypeString,
				Value: `"loadbalancer"`,
			},
		},
		Fields: []*dsl.FieldDesc{
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

	loadBalancerUpdateParam = &dsl.Model{
		Name:      names.UpdateParameterName(loadBalancerAPIName),
		NakedType: loadBalancerNakedType,
		Fields: []*dsl.FieldDesc{
			// common fields
			fields.Name(),
			fields.Description(),
			fields.Tags(),
			fields.IconID(),

			// settings
			fields.LoadBalancerVIP(),
			// settings hash
			fields.SettingsHash(),
		},
	}

	loadBalancerStatus = &dsl.Model{
		Name: "LoadBalancerStatus",
		Fields: []*dsl.FieldDesc{
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
				Type: &dsl.Model{
					Name:    "LoadBalancerServerStatus",
					IsArray: true,
					Fields: []*dsl.FieldDesc{
						fields.Def("ActiveConn", meta.TypeStringNumber),
						fields.Def("Status", meta.TypeInstanceStatus),
						fields.Def("IPAddress", meta.TypeString),
						fields.Def("Port", meta.TypeStringNumber),
						fields.Def("CPS", meta.TypeStringNumber),
					},
				},
				Tags: &dsl.FieldTags{
					MapConv: ",recursive",
				},
			},
		},
		NakedType: meta.Static(naked.LoadBalancerStatus{}),
	}
)
