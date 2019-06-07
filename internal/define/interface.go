package define

import (
	"net/http"

	"github.com/sacloud/libsacloud-v2/internal/schema"
	"github.com/sacloud/libsacloud-v2/internal/schema/meta"
	"github.com/sacloud/libsacloud-v2/sacloud/naked"
)

func init() {
	nakedType := meta.Static(naked.Interface{})

	iface := &schema.Model{
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

	createParam := &schema.Model{
		Fields: []*schema.FieldDesc{
			fields.ServerID(),
		},
	}

	updateParam := &schema.Model{
		Fields: []*schema.FieldDesc{
			fields.UserIPAddress(),
		},
	}

	ifAPI := &schema.Resource{
		Name:       "Interface",
		PathName:   "interface",
		PathSuffix: schema.CloudAPISuffix,
	}
	ifAPI.Operations = []*schema.Operation{
		// find
		ifAPI.DefineOperationFind(nakedType, findParameter, iface),

		// create
		ifAPI.DefineOperationCreate(nakedType, createParam, iface),

		// read
		ifAPI.DefineOperationRead(nakedType, iface),

		// update
		ifAPI.DefineOperationUpdate(nakedType, updateParam, iface),

		// delete
		ifAPI.DefineOperationDelete(),

		// monitor
		ifAPI.DefineOperationMonitor(monitorParameter, monitors.interfaceModel()),

		ifAPI.DefineSimpleOperation("ConnectToSharedSegment", http.MethodPut, "to/switch/shared"),

		ifAPI.DefineSimpleOperation("ConnectToSwitch", http.MethodPut, "to/switch/{{.switchID}}",
			&schema.Argument{
				Name: "switchID",
				Type: meta.TypeID,
			},
		),

		ifAPI.DefineSimpleOperation("DisconnectFromSwitch", http.MethodDelete, "to/switch"),

		ifAPI.DefineSimpleOperation("ConnectToPacketFilter", http.MethodPut, "to/packetfilter/{{.packetFilterID}}",
			&schema.Argument{
				Name: "packetFilterID",
				Type: meta.TypeID,
			},
		),

		ifAPI.DefineSimpleOperation("DisconnectFromPacketFilter", http.MethodDelete, "to/packetfilter"),
	}
	Resources.Def(ifAPI)
}
