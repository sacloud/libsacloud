// Copyright 2016-2019 The Libsacloud Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package define

import (
	"github.com/sacloud/libsacloud/v2/internal/dsl"
	"github.com/sacloud/libsacloud/v2/internal/dsl/meta"
	"github.com/sacloud/libsacloud/v2/sacloud/naked"
)

type monitorsDef struct{}

var monitors = &monitorsDef{}

func (m *monitorsDef) cpuTimeModel() *dsl.Model {
	return &dsl.Model{
		Name: "CPUTimeActivity",
		Fields: []*dsl.FieldDesc{
			{
				Name: "Values",
				Type: &dsl.Model{
					Name:      "MonitorCPUTimeValue",
					NakedType: meta.Static(naked.MonitorCPUTimeValue{}),
					IsArray:   true,
					Fields: []*dsl.FieldDesc{
						fields.MonitorTime(),
						fields.MonitorCPUTime(),
					},
				},
				Tags: &dsl.FieldTags{
					MapConv: "[]CPU",
				},
			},
		},
		NakedType: meta.Static(naked.MonitorValues{}),
	}
}

func (m *monitorsDef) diskModel() *dsl.Model {
	return &dsl.Model{
		Name: "DiskActivity",
		Fields: []*dsl.FieldDesc{
			{
				Name: "Values",
				Type: &dsl.Model{
					Name:      "MonitorDiskValue",
					NakedType: meta.Static(naked.MonitorDiskValue{}),
					IsArray:   true,
					Fields: []*dsl.FieldDesc{
						fields.MonitorTime(),
						fields.MonitorDiskRead(),
						fields.MonitorDiskWrite(),
					},
				},
				Tags: &dsl.FieldTags{
					MapConv: "[]Disk",
				},
			},
		},
		NakedType: meta.Static(naked.MonitorValues{}),
	}
}

func (m *monitorsDef) interfaceModel() *dsl.Model {
	return &dsl.Model{
		Name: "InterfaceActivity",
		Fields: []*dsl.FieldDesc{
			{
				Name: "Values",
				Type: &dsl.Model{
					Name:      "MonitorInterfaceValue",
					NakedType: meta.Static(naked.MonitorInterfaceValue{}),
					IsArray:   true,
					Fields: []*dsl.FieldDesc{
						fields.MonitorTime(),
						fields.MonitorInterfaceReceive(),
						fields.MonitorInterfaceSend(),
					},
				},
				Tags: &dsl.FieldTags{
					MapConv: "[]Interface",
				},
			},
		},
		NakedType: meta.Static(naked.MonitorValues{}),
	}
}

func (m *monitorsDef) routerModel() *dsl.Model {
	return &dsl.Model{
		Name: "RouterActivity",
		Fields: []*dsl.FieldDesc{
			{
				Name: "Values",
				Type: &dsl.Model{
					Name:      "MonitorRouterValue",
					NakedType: meta.Static(naked.MonitorRouterValue{}),
					IsArray:   true,
					Fields: []*dsl.FieldDesc{
						fields.MonitorTime(),
						fields.MonitorRouterIn(),
						fields.MonitorRouterOut(),
					},
				},
				Tags: &dsl.FieldTags{
					MapConv: "[]Router",
				},
			},
		},
		NakedType: meta.Static(naked.MonitorValues{}),
	}
}

func (m *monitorsDef) databaseModel() *dsl.Model {
	return &dsl.Model{
		Name: "DatabaseActivity",
		Fields: []*dsl.FieldDesc{
			{
				Name: "Values",
				//Type: meta.Static([]naked.MonitorDatabaseValue{}),
				Type: &dsl.Model{
					Name:      "MonitorDatabaseValue",
					NakedType: meta.Static(naked.MonitorDatabaseValue{}),
					IsArray:   true,
					Fields: []*dsl.FieldDesc{
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
				Tags: &dsl.FieldTags{
					MapConv: "[]Database",
				},
			},
		},
		NakedType: meta.Static(naked.MonitorValues{}),
	}
}

func (m *monitorsDef) freeDiskSizeModel() *dsl.Model {
	return &dsl.Model{
		Name: "FreeDiskSizeActivity",
		Fields: []*dsl.FieldDesc{
			{
				Name: "Values",
				Type: &dsl.Model{
					Name:      "MonitorFreeDiskSizeValue",
					NakedType: meta.Static(naked.MonitorFreeDiskSizeValue{}),
					IsArray:   true,
					Fields: []*dsl.FieldDesc{
						fields.MonitorTime(),
						fields.MonitorFreeDiskSize(),
					},
				},
				Tags: &dsl.FieldTags{
					MapConv: "[]FreeDiskSize",
				},
			},
		},
		NakedType: meta.Static(naked.MonitorValues{}),
	}
}

func (m *monitorsDef) responseTimeSecModel() *dsl.Model {
	return &dsl.Model{
		Name: "ResponseTimeSecActivity",
		Fields: []*dsl.FieldDesc{
			{
				Name: "Values",
				Type: &dsl.Model{
					Name:      "MonitorResponseTimeSecValue",
					NakedType: meta.Static(naked.MonitorResponseTimeSecValue{}),
					IsArray:   true,
					Fields: []*dsl.FieldDesc{
						fields.MonitorTime(),
						fields.MonitorResponseTimeSec(),
					},
				},
				Tags: &dsl.FieldTags{
					MapConv: "[]ResponseTimeSec",
				},
			},
		},
		NakedType: meta.Static(naked.MonitorValues{}),
	}
}

func (m *monitorsDef) linkModel() *dsl.Model {
	return &dsl.Model{
		Name: "LinkActivity",
		Fields: []*dsl.FieldDesc{
			{
				Name: "Values",
				Type: &dsl.Model{
					Name:      "MonitorLinkValue",
					NakedType: meta.Static(naked.MonitorLinkValue{}),
					IsArray:   true,
					Fields: []*dsl.FieldDesc{
						fields.MonitorTime(),
						fields.MonitorUplinkBPS(),
						fields.MonitorDownlinkBPS(),
					},
				},
				Tags: &dsl.FieldTags{
					MapConv: "[]Link",
				},
			},
		},
		NakedType: meta.Static(naked.MonitorValues{}),
	}
}

func (m *monitorsDef) connectionModel() *dsl.Model {
	return &dsl.Model{
		Name: "ConnectionActivity",
		Fields: []*dsl.FieldDesc{
			{
				Name: "Values",
				//Type: meta.Static([]naked.MonitorConnectionValue{}),
				Type: &dsl.Model{
					Name:      "MonitorConnectionValue",
					NakedType: meta.Static(naked.MonitorConnectionValue{}),
					IsArray:   true,
					Fields: []*dsl.FieldDesc{
						fields.MonitorTime(),
						fields.MonitorActiveConnections(),
						fields.MonitorConnectionsPerSec(),
					},
				},
				Tags: &dsl.FieldTags{
					MapConv: "[]Connection",
				},
			},
		},
		NakedType: meta.Static(naked.MonitorValues{}),
	}
}
