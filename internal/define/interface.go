package define

import (
	"net/http"

	"github.com/sacloud/libsacloud/internal/schema"
	"github.com/sacloud/libsacloud/internal/schema/meta"
	"github.com/sacloud/libsacloud/sacloud/naked"
)

var interfaceAPI = &schema.Resource{
	Name:       "Interface",
	PathName:   "interface",
	PathSuffix: schema.CloudAPISuffix,
	OperationsDefineFunc: func(r *schema.Resource) []*schema.Operation {
		return []*schema.Operation{
			// find
			r.DefineOperationFind(interfaceNakedType, findParameter, interfaceView),

			// create
			r.DefineOperationCreate(interfaceNakedType, interfaceCreateParam, interfaceView),

			// read
			r.DefineOperationRead(interfaceNakedType, interfaceView),

			// update
			r.DefineOperationUpdate(interfaceNakedType, interfaceUpdateParam, interfaceView),

			// delete
			r.DefineOperationDelete(),

			// monitor
			r.DefineOperationMonitor(monitorParameter, monitors.interfaceModel()),

			r.DefineSimpleOperation("ConnectToSharedSegment", http.MethodPut, "to/switch/shared"),

			r.DefineSimpleOperation("ConnectToSwitch", http.MethodPut, "to/switch/{{.switchID}}",
				&schema.Argument{
					Name: "switchID",
					Type: meta.TypeID,
				},
			),

			r.DefineSimpleOperation("DisconnectFromSwitch", http.MethodDelete, "to/switch"),

			r.DefineSimpleOperation("ConnectToPacketFilter", http.MethodPut, "to/packetfilter/{{.packetFilterID}}",
				&schema.Argument{
					Name: "packetFilterID",
					Type: meta.TypeID,
				},
			),

			r.DefineSimpleOperation("DisconnectFromPacketFilter", http.MethodDelete, "to/packetfilter"),
		}
	},
}
var (
	interfaceNakedType = meta.Static(naked.Interface{})

	interfaceView = &schema.Model{
		Fields: []*schema.FieldDesc{
			fields.ID(),
			fields.MACAddress(),
			fields.IPAddress(),
			fields.UserIPAddress(),
			fields.HostName(),
			fields.SwitchID(), // TODO どこまで実装する? デフォルトゲートウェイなどはSubnetまでみないとわからないかも
			fields.PacketFilterID(),
			fields.ServerID(),
			fields.CreatedAt(),
			fields.ModifiedAt(), // TODO ModifiedAtがないかも?
		},
	}

	interfaceCreateParam = &schema.Model{
		Fields: []*schema.FieldDesc{
			fields.ServerID(),
		},
	}

	interfaceUpdateParam = &schema.Model{
		Fields: []*schema.FieldDesc{
			fields.UserIPAddress(),
		},
	}
)
