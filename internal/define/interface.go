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

	Resources.DefineWith("Interface", func(r *schema.Resource) {
		r.Operations(
			// find
			r.DefineOperationFind(nakedType, findParameter, iface),

			// create
			r.DefineOperationCreate(nakedType, createParam, iface),

			// read
			r.DefineOperationRead(nakedType, iface),

			// update
			r.DefineOperationUpdate(nakedType, updateParam, iface),

			// delete
			r.DefineOperationDelete(),

			// monitor
			r.DefineOperationMonitor(monitorParameter, monitors.interfaceModel()),

			r.DefineSimpleOperation("ConnectToSharedSegment", http.MethodPut, "to/switch/shared"),

			r.DefineSimpleOperation("ConnectToSwitch", http.MethodPut, "to/switch/{{.switchID}}",
				&schema.SimpleArgument{
					Name: "switchID",
					Type: meta.TypeID,
				},
			),

			r.DefineSimpleOperation("DisconnectFromSwitch", http.MethodDelete, "to/switch"),

			r.DefineSimpleOperation("ConnectToPacketFilter", http.MethodPut, "to/packetfilter/{{.packetFilterID}}",
				&schema.SimpleArgument{
					Name: "packetFilterID",
					Type: meta.TypeID,
				},
			),

			r.DefineSimpleOperation("DisconnectFromPacketFilter", http.MethodDelete, "to/packetfilter"),
		)
	})
}
