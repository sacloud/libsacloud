// Copyright 2016-2022 The Libsacloud Authors
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
	"github.com/sacloud/libsacloud/v2/internal/define/names"
	"github.com/sacloud/libsacloud/v2/internal/define/ops"
	"github.com/sacloud/libsacloud/v2/internal/dsl"
	"github.com/sacloud/libsacloud/v2/internal/dsl/meta"
	"github.com/sacloud/libsacloud/v2/sacloud/naked"
)

const (
	nfsAPIName     = "NFS"
	nfsAPIPathName = "appliance"
)

var nfsAPI = &dsl.Resource{
	Name:       nfsAPIName,
	PathName:   nfsAPIPathName,
	PathSuffix: dsl.CloudAPISuffix,
	Operations: dsl.Operations{
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
		ops.MonitorChild(nfsAPIName, "CPU", "cpu",
			monitorParameter, monitors.cpuTimeModel()),
		ops.MonitorChild(nfsAPIName, "FreeDiskSize", "database",
			monitorParameter, monitors.freeDiskSizeModel()),
		ops.MonitorChild(nfsAPIName, "Interface", "interface",
			monitorParameter, monitors.interfaceModel()),
	},
}

var (
	nfsNakedType = meta.Static(naked.NFS{})

	nfsView = &dsl.Model{
		Name:      nfsAPIName,
		NakedType: nfsNakedType,
		Fields: []*dsl.FieldDesc{
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
			fields.Def("SwitchName", meta.TypeString, mapConvTag("Switch.Name")),
		},
	}

	nfsCreateParam = &dsl.Model{
		Name:      names.CreateParameterName(nfsAPIName),
		NakedType: nfsNakedType,
		ConstFields: []*dsl.ConstFieldDesc{
			{
				Name:  "Class",
				Type:  meta.TypeString,
				Value: `"nfs"`,
			},
		},
		Fields: []*dsl.FieldDesc{
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

	nfsUpdateParam = &dsl.Model{
		Name:      names.UpdateParameterName(nfsAPIName),
		NakedType: nfsNakedType,
		Fields: []*dsl.FieldDesc{
			fields.Name(),
			fields.Description(),
			fields.Tags(),
			fields.IconID(),
		},
	}
)
