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
	"net/http"

	"github.com/sacloud/libsacloud/v2/internal/define/names"
	"github.com/sacloud/libsacloud/v2/internal/define/ops"
	"github.com/sacloud/libsacloud/v2/internal/dsl"
	"github.com/sacloud/libsacloud/v2/internal/dsl/meta"
	"github.com/sacloud/libsacloud/v2/sacloud/naked"
)

const (
	interfaceAPIName     = "Interface"
	interfaceAPIPathName = "interface"
)

var interfaceAPI = &dsl.Resource{
	Name:       interfaceAPIName,
	PathName:   interfaceAPIPathName,
	PathSuffix: dsl.CloudAPISuffix,
	Operations: dsl.Operations{
		// find
		ops.Find(interfaceAPIName, interfaceNakedType, findParameter, interfaceView),

		// create
		ops.Create(interfaceAPIName, interfaceNakedType, interfaceCreateParam, interfaceView),

		// read
		ops.Read(interfaceAPIName, interfaceNakedType, interfaceView),

		// update
		ops.Update(interfaceAPIName, interfaceNakedType, interfaceUpdateParam, interfaceView),

		// patch
		ops.Patch(interfaceAPIName, interfaceNakedType, patchModel(interfaceUpdateParam), interfaceView),

		// delete
		ops.Delete(interfaceAPIName),

		// monitor
		ops.Monitor(interfaceAPIName, monitorParameter, monitors.interfaceModel()),

		ops.WithIDAction(interfaceAPIName, "ConnectToSharedSegment", http.MethodPut, "to/switch/shared"),

		ops.WithIDAction(interfaceAPIName, "ConnectToSwitch", http.MethodPut, "to/switch/{{.switchID}}",
			&dsl.Argument{
				Name: "switchID",
				Type: meta.TypeID,
			},
		),

		ops.WithIDAction(interfaceAPIName, "DisconnectFromSwitch", http.MethodDelete, "to/switch"),

		ops.WithIDAction(interfaceAPIName, "ConnectToPacketFilter", http.MethodPut, "to/packetfilter/{{.packetFilterID}}",
			&dsl.Argument{
				Name: "packetFilterID",
				Type: meta.TypeID,
			},
		),

		ops.WithIDAction(interfaceAPIName, "DisconnectFromPacketFilter", http.MethodDelete, "to/packetfilter"),
	},
}
var (
	interfaceNakedType = meta.Static(naked.Interface{})

	interfaceView = &dsl.Model{
		Name:      interfaceAPIName,
		NakedType: interfaceNakedType,
		Fields: []*dsl.FieldDesc{
			fields.ID(),
			fields.MACAddress(),
			fields.IPAddress(),
			fields.UserIPAddress(),
			fields.HostName(),
			fields.SwitchID(),
			fields.PacketFilterID(),
			fields.ServerID(),
			fields.CreatedAt(),
		},
	}

	interfaceCreateParam = &dsl.Model{
		Name:      names.CreateParameterName(interfaceAPIName),
		NakedType: interfaceNakedType,
		Fields: []*dsl.FieldDesc{
			fields.ServerID(),
		},
	}

	interfaceUpdateParam = &dsl.Model{
		Name:      names.UpdateParameterName(interfaceAPIName),
		NakedType: interfaceNakedType,
		Fields: []*dsl.FieldDesc{
			fields.UserIPAddress(),
		},
	}
)
