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

			r.DefineOperation("ConnectToSharedSegment").
				Method(http.MethodPut).
				PathFormat(schema.IDAndSuffixPathFormat("to/switch/shared")).
				Argument(schema.ArgumentZone).
				Argument(schema.ArgumentID),

			r.DefineOperation("ConnectToSwitch").
				Method(http.MethodPut).
				PathFormat(schema.IDAndSuffixPathFormat("to/switch/{{.switchID}}")).
				Argument(schema.ArgumentZone).
				Argument(schema.ArgumentID).
				Argument(&schema.SimpleArgument{
					Name: "switchID",
					Type: meta.TypeID,
				}),

			r.DefineOperation("DisconnectFromSwitch").
				Method(http.MethodDelete).
				PathFormat(schema.IDAndSuffixPathFormat("to/switch")).
				Argument(schema.ArgumentZone).
				Argument(schema.ArgumentID),

			r.DefineOperation("ConnectToPacketFilter").
				Method(http.MethodPut).
				PathFormat(schema.IDAndSuffixPathFormat("to/packetfilter/{{.packetFilterID}}")).
				Argument(schema.ArgumentZone).
				Argument(schema.ArgumentID).
				Argument(&schema.SimpleArgument{
					Name: "packetFilterID",
					Type: meta.TypeID,
				}),

			r.DefineOperation("DisconnectFromPacketFilter").
				Method(http.MethodDelete).
				PathFormat(schema.IDAndSuffixPathFormat("to/packetfilter")).
				Argument(schema.ArgumentZone).
				Argument(schema.ArgumentID),
		)
	})
}
