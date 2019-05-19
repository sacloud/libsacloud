package define

import (
	"github.com/sacloud/libsacloud-v2/internal/schema"
	"github.com/sacloud/libsacloud-v2/internal/schema/meta"
	"github.com/sacloud/libsacloud-v2/sacloud/naked"
)

type monitorsDef struct{}

var monitors = &monitorsDef{}

func (m *monitorsDef) cpuTimeModel() *schema.Model {
	return &schema.Model{
		Name: "CPUTimeActivity",
		Fields: []*schema.FieldDesc{
			{
				Name: "Values",
				Type: &schema.Model{
					Name:      "MonitorCPUTimeValue",
					NakedType: meta.Static(naked.MonitorCPUTimeValue{}),
					IsArray:   true,
					Fields: []*schema.FieldDesc{
						fields.MonitorTime(),
						fields.MonitorCPUTime(),
					},
				},
				Tags: &schema.FieldTags{
					MapConv: "[]CPU",
				},
			},
		},
		NakedType: meta.Static(naked.MonitorValues{}),
	}
}

func (m *monitorsDef) diskModel() *schema.Model {
	return &schema.Model{
		Name: "DiskActivity",
		Fields: []*schema.FieldDesc{
			{
				Name: "Values",
				Type: &schema.Model{
					Name:      "MonitorDiskValue",
					NakedType: meta.Static(naked.MonitorDiskValue{}),
					IsArray:   true,
					Fields: []*schema.FieldDesc{
						fields.MonitorTime(),
						fields.MonitorDiskRead(),
						fields.MonitorDiskWrite(),
					},
				},
				Tags: &schema.FieldTags{
					MapConv: "[]Disk",
				},
			},
		},
		NakedType: meta.Static(naked.MonitorValues{}),
	}
}

func (m *monitorsDef) interfaceModel() *schema.Model {
	return &schema.Model{
		Name: "InterfaceActivity",
		Fields: []*schema.FieldDesc{
			{
				Name: "Values",
				Type: &schema.Model{
					Name:      "MonitorInterfaceValue",
					NakedType: meta.Static(naked.MonitorInterfaceValue{}),
					IsArray:   true,
					Fields: []*schema.FieldDesc{
						fields.MonitorTime(),
						fields.MonitorInterfaceReceive(),
						fields.MonitorInterfaceSend(),
					},
				},
				Tags: &schema.FieldTags{
					MapConv: "[]Interface",
				},
			},
		},
		NakedType: meta.Static(naked.MonitorValues{}),
	}
}

func (m *monitorsDef) routerModel() *schema.Model {
	return &schema.Model{
		Name: "RouterActivity",
		Fields: []*schema.FieldDesc{
			{
				Name: "Values",
				Type: &schema.Model{
					Name:      "MonitorRouterValue",
					NakedType: meta.Static(naked.MonitorRouterValue{}),
					IsArray:   true,
					Fields: []*schema.FieldDesc{
						fields.MonitorTime(),
						fields.MonitorRouterIn(),
						fields.MonitorRouterOut(),
					},
				},
				Tags: &schema.FieldTags{
					MapConv: "[]Router",
				},
			},
		},
		NakedType: meta.Static(naked.MonitorValues{}),
	}
}

func (m *monitorsDef) databaseModel() *schema.Model {
	return &schema.Model{
		Name: "DatabaseActivity",
		Fields: []*schema.FieldDesc{
			{
				Name: "Values",
				//Type: meta.Static([]naked.MonitorDatabaseValue{}),
				Type: &schema.Model{
					Name:      "MonitorDatabaseValue",
					NakedType: meta.Static(naked.MonitorDatabaseValue{}),
					IsArray:   true,
					Fields: []*schema.FieldDesc{
						fields.MonitorTime(),
						fields.MonitorDatabaseTotalMemorySize(),
						fields.MonitorDatabaseUsedMemorySize(),
						fields.MonitorDatabaseTotalDisk1Size(),
						fields.MonitorDatabaseUsedDisk1Size(),
						fields.MonitorDatabaseTotalDisk2Size(),
						fields.MonitorDatabaseUsedDisk2Size(),
						fields.MonitorDatabaseBinlogUsedSizeKiB(),
						fields.MonitorDatabaseDelayTimeSec(),
					},
				},
				Tags: &schema.FieldTags{
					MapConv: "[]Database",
				},
			},
		},
		NakedType: meta.Static(naked.MonitorValues{}),
	}
}

func (m *monitorsDef) freeDiskSizeModel() *schema.Model {
	return &schema.Model{
		Name: "FreeDiskSizeActivity",
		Fields: []*schema.FieldDesc{
			{
				Name: "Values",
				Type: &schema.Model{
					Name:      "MonitorFreeDiskSizeValue",
					NakedType: meta.Static(naked.MonitorFreeDiskSizeValue{}),
					IsArray:   true,
					Fields: []*schema.FieldDesc{
						fields.MonitorTime(),
						fields.MonitorFreeDiskSize(),
					},
				},
				Tags: &schema.FieldTags{
					MapConv: "[]FreeDiskSize",
				},
			},
		},
		NakedType: meta.Static(naked.MonitorValues{}),
	}
}

func (m *monitorsDef) responseTimeSecModel() *schema.Model {
	return &schema.Model{
		Name: "ResponseTimeSecActivity",
		Fields: []*schema.FieldDesc{
			{
				Name: "Values",
				Type: &schema.Model{
					Name:      "MonitorResponseTimeSecValue",
					NakedType: meta.Static(naked.MonitorResponseTimeSecValue{}),
					IsArray:   true,
					Fields: []*schema.FieldDesc{
						fields.MonitorTime(),
						fields.MonitorResponseTimeSec(),
					},
				},
				Tags: &schema.FieldTags{
					MapConv: "[]ResponseTimeSec",
				},
			},
		},
		NakedType: meta.Static(naked.MonitorValues{}),
	}
}

func (m *monitorsDef) linkModel() *schema.Model {
	return &schema.Model{
		Name: "LinkActivity",
		Fields: []*schema.FieldDesc{
			{
				Name: "Values",
				Type: &schema.Model{
					Name:      "MonitorLinkValue",
					NakedType: meta.Static(naked.MonitorLinkValue{}),
					IsArray:   true,
					Fields: []*schema.FieldDesc{
						fields.MonitorTime(),
						fields.MonitorUplinkBPS(),
						fields.MonitorDownlinkBPS(),
					},
				},
				Tags: &schema.FieldTags{
					MapConv: "[]Link",
				},
			},
		},
		NakedType: meta.Static(naked.MonitorValues{}),
	}
}

func (m *monitorsDef) connectionModel() *schema.Model {
	return &schema.Model{
		Name: "ConnectionActivity",
		Fields: []*schema.FieldDesc{
			{
				Name: "Values",
				//Type: meta.Static([]naked.MonitorConnectionValue{}),
				Type: &schema.Model{
					Name:      "MonitorConnectionValue",
					NakedType: meta.Static(naked.MonitorConnectionValue{}),
					IsArray:   true,
					Fields: []*schema.FieldDesc{
						fields.MonitorTime(),
						fields.MonitorActiveConnections(),
						fields.MonitorConnectionsPerSec(),
					},
				},
				Tags: &schema.FieldTags{
					MapConv: "[]Connection",
				},
			},
		},
		NakedType: meta.Static(naked.MonitorValues{}),
	}
}
