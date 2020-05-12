// Copyright 2016-2020 The Libsacloud Authors
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
	"net/http"

	"github.com/sacloud/libsacloud/v2/internal/define/names"
	"github.com/sacloud/libsacloud/v2/internal/define/ops"
	"github.com/sacloud/libsacloud/v2/internal/dsl"
	"github.com/sacloud/libsacloud/v2/internal/dsl/meta"
	"github.com/sacloud/libsacloud/v2/sacloud/naked"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

const (
	serverAPIName     = "Server"
	serverAPIPathName = "server"
)

var serverAPI = &dsl.Resource{
	Name:       serverAPIName,
	PathName:   serverAPIPathName,
	PathSuffix: dsl.CloudAPISuffix,
	Operations: dsl.Operations{
		// find
		ops.Find(serverAPIName, serverNakedType, findParameter, serverView),

		// create
		ops.Create(serverAPIName, serverNakedType, serverCreateParam, serverView),

		// read
		ops.Read(serverAPIName, serverNakedType, serverView),

		// update
		ops.Update(serverAPIName, serverNakedType, serverUpdateParam, serverView),

		// delete
		ops.Delete(serverAPIName),

		// delete with disks
		{
			ResourceName: serverAPIName,
			Name:         "DeleteWithDisks",
			PathFormat:   dsl.DefaultPathFormatWithID,
			Method:       http.MethodDelete,
			RequestEnvelope: dsl.RequestEnvelope(
				&dsl.EnvelopePayloadDesc{
					Type: meta.Static([]types.ID{}),
					Name: "WithDisk",
				},
			),
			Arguments: dsl.Arguments{
				dsl.ArgumentID,
				dsl.PassthroughModelArgument("disks", serverDeleteParam),
			},
		},

		// change plan
		{
			ResourceName:    serverAPIName,
			Name:            "ChangePlan",
			PathFormat:      dsl.IDAndSuffixPathFormat("plan"),
			Method:          http.MethodPut,
			RequestEnvelope: dsl.RequestEnvelopeFromModel(serverChangePlanParam),
			Arguments: dsl.Arguments{
				dsl.ArgumentID,
				dsl.PassthroughModelArgument("plan", serverChangePlanParam),
			},
			ResponseEnvelope: dsl.ResponseEnvelope(&dsl.EnvelopePayloadDesc{
				Name: serverAPIName,
				Type: meta.Static(naked.Server{}),
			}),
			Results: dsl.Results{
				{
					SourceField: names.ResourceFieldName(serverAPIName, dsl.PayloadForms.Singular),
					DestField:   names.ResourceFieldName(serverAPIName, dsl.PayloadForms.Singular),
					IsPlural:    false,
					Model:       serverView,
				},
			},
		},

		// insert cdrom
		{
			ResourceName: serverAPIName,
			Name:         "InsertCDROM",
			PathFormat:   dsl.IDAndSuffixPathFormat("cdrom"),
			Method:       http.MethodPut,
			RequestEnvelope: dsl.RequestEnvelope(
				&dsl.EnvelopePayloadDesc{
					Type: meta.Static(naked.CDROM{}),
					Name: "CDROM",
				},
			),
			Arguments: dsl.Arguments{
				dsl.ArgumentID,
				{
					Name: "insertParam",
					Type: &dsl.Model{
						Name: "InsertCDROMRequest",
						Fields: []*dsl.FieldDesc{
							fields.ID(),
						},
						NakedType: meta.Static(naked.CDROM{}),
					},
					MapConvTag: "CDROM",
				},
			},
		},

		// eject cdrom
		{
			ResourceName: serverAPIName,
			Name:         "EjectCDROM",
			PathFormat:   dsl.IDAndSuffixPathFormat("cdrom"),
			Method:       http.MethodDelete,
			RequestEnvelope: dsl.RequestEnvelope(
				&dsl.EnvelopePayloadDesc{
					Type: meta.Static(naked.CDROM{}),
					Name: "CDROM",
				},
			),
			Arguments: dsl.Arguments{
				dsl.ArgumentID,
				{
					Name: "ejectParam",
					Type: &dsl.Model{
						Name: "EjectCDROMRequest",
						Fields: []*dsl.FieldDesc{
							fields.ID(),
						},
						NakedType: meta.Static(naked.CDROM{}),
					},
					MapConvTag: "CDROM",
				},
			},
		},

		// power management(boot/shutdown/reset)
		ops.Boot(serverAPIName),
		ops.Shutdown(serverAPIName),
		ops.Reset(serverAPIName),

		// send key
		{
			ResourceName: serverAPIName,
			Name:         "SendKey",
			PathFormat:   dsl.IDAndSuffixPathFormat("keyboard"),
			Method:       http.MethodPut,
			Arguments: dsl.Arguments{
				dsl.ArgumentID,
				dsl.PassthroughModelArgument("keyboardParam", serverSendKeyParam),
			},
			RequestEnvelope: dsl.RequestEnvelopeFromModel(serverSendKeyParam),
		},

		// get vnc proxy
		{
			ResourceName: serverAPIName,
			Name:         "GetVNCProxy",
			PathFormat:   dsl.IDAndSuffixPathFormat("vnc/proxy"),
			Method:       http.MethodGet,
			Arguments: dsl.Arguments{
				dsl.ArgumentID,
			},
			ResponseEnvelope: dsl.ResponseEnvelope(
				&dsl.EnvelopePayloadDesc{
					Name: "VNCProxyInfo",
					Type: meta.Static(naked.VNCProxyInfo{}),
				},
			),
			Results: dsl.Results{
				{
					SourceField: "VNCProxyInfo",
					DestField:   serverVNCProxyView.Name,
					IsPlural:    false,
					Model:       serverVNCProxyView,
				},
			},
		},

		// monitor
		ops.Monitor(serverAPIName, monitorParameter, monitors.cpuTimeModel()),
		ops.MonitorChild(serverAPIName, "CPU", "", monitorParameter, monitors.cpuTimeModel()),
	},
}

var (
	serverNakedType = meta.Static(naked.Server{})

	serverView = &dsl.Model{
		Name:      serverAPIName,
		NakedType: serverNakedType,
		Fields: []*dsl.FieldDesc{
			fields.ID(),
			fields.Name(),
			fields.Description(),
			fields.Tags(),
			fields.Availability(),
			fields.HostName(),
			fields.InterfaceDriver(),
			// server plan
			fields.ServerPlanID(),
			fields.ServerPlanName(),
			fields.ServerPlanCPU(),
			fields.ServerPlanMemoryMB(),
			fields.ServerPlanCommitment(),
			fields.ServerPlanGeneration(),
			// zone
			fields.Zone(),
			// instance
			fields.InstanceHostName(),
			fields.InstanceHostInfoURL(),
			fields.InstanceStatus(),
			fields.InstanceBeforeStatus(),
			fields.InstanceStatusChangedAt(),
			fields.InstanceWarnings(),
			fields.InstanceWarningsValue(),

			// disks
			{
				Name: "Disks",
				Type: serverConnectedDiskView,
				Tags: &dsl.FieldTags{
					JSON:    ",omitempty",
					MapConv: ",recursive",
				},
			},
			fields.Interfaces(),

			fields.CDROMID(),

			fields.PrivateHostID(),
			fields.PrivateHostName(),

			fields.BundleInfo(),

			fields.IconID(),
			fields.CreatedAt(),
			fields.ModifiedAt(),
		},
	}

	serverConnectedDiskView = &dsl.Model{
		Name:      "ServerConnectedDisk",
		NakedType: meta.Static(naked.Disk{}),
		IsArray:   true,
		Fields: []*dsl.FieldDesc{
			fields.ID(),
			fields.Name(),
			fields.Availability(),
			fields.DiskConnection(),
			fields.DiskConnectionOrder(),
			fields.DiskReinstallCount(),
			fields.SizeMB(),
			fields.DiskPlanID(),
			{
				Name: "Storage",
				Type: serverConnectedStorage,
				Tags: &dsl.FieldTags{
					MapConv: ",omitempty,recursive",
					JSON:    ",omitempty",
				},
			},
		},
	}

	serverConnectedStorage = &dsl.Model{
		Name:      "Storage",
		NakedType: meta.Static(naked.Storage{}),
		Fields: []*dsl.FieldDesc{
			fields.ID(),
			{
				Name: "Class",
				Type: meta.TypeString,
				Tags: &dsl.FieldTags{
					MapConv: ",omitempty",
					JSON:    ",omitempty",
				},
			},
			{
				Name: "Generation",
				Type: meta.TypeInt,
				Tags: &dsl.FieldTags{
					MapConv: ",omitempty",
					JSON:    ",omitempty",
				},
			},
		},
	}

	serverCreateParam = &dsl.Model{
		Name:      names.CreateParameterName(serverAPIName),
		NakedType: serverNakedType,
		Fields: []*dsl.FieldDesc{
			// server plan
			fields.ServerPlanCPU(),
			fields.ServerPlanMemoryMB(),
			fields.ServerPlanCommitment(),
			fields.ServerPlanGeneration(),
			fields.ServerConnectedSwitch(),
			fields.InterfaceDriver(),
			fields.Name(),
			fields.Description(),
			fields.Tags(),
			fields.IconID(),
			{
				Name: "WaitDiskMigration",
				Type: meta.TypeFlag,
				Tags: &dsl.FieldTags{
					MapConv: ",omitempty",
					JSON:    ",omitempty",
				},
			},
			fields.PrivateHostID(),
		},
	}

	serverUpdateParam = &dsl.Model{
		Name:      names.UpdateParameterName(serverAPIName),
		NakedType: serverNakedType,
		Fields: []*dsl.FieldDesc{
			fields.Name(),
			fields.Description(),
			fields.Tags(),
			fields.IconID(),
			fields.PrivateHostID(),
			fields.InterfaceDriver(),
		},
	}

	serverDeleteParam = &dsl.Model{
		Name: "ServerDeleteWithDisksRequest",
		Fields: []*dsl.FieldDesc{
			fields.Def("IDs", meta.Static([]types.ID{}), &dsl.FieldTags{
				MapConv: "WithDisk",
			}),
		},
		NakedType: meta.Static(naked.DeleteServerWithDiskParameter{}),
	}

	serverChangePlanParam = &dsl.Model{
		Name: "ServerChangePlanRequest",
		Fields: []*dsl.FieldDesc{
			fields.CPU(),
			fields.MemoryMB(),
			fields.Generation(),
			fields.ServerPlanCommitment(),
		},
		NakedType: meta.Static(naked.ServerPlan{}),
	}

	serverSendKeyParam = &dsl.Model{
		Name: "SendKeyRequest",
		Fields: []*dsl.FieldDesc{
			fields.Def("Key", meta.TypeString),
			fields.Def("Keys", meta.TypeStringSlice),
		},
	}

	serverVNCProxyView = &dsl.Model{
		Name:      "VNCProxyInfo",
		NakedType: meta.Static(naked.VNCProxyInfo{}),
		Fields: []*dsl.FieldDesc{
			fields.Def("Status", meta.TypeString),
			fields.Def("Host", meta.TypeString),
			fields.Def("IOServerHost", meta.TypeString),
			fields.Def("Port", meta.TypeStringNumber),
			fields.Def("Password", meta.TypeString),
			fields.Def("VNCFile", meta.TypeString),
		},
	}
)
