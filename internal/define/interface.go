package define

import (
	"net/http"

	"github.com/sacloud/libsacloud/v2/internal/define/names"
	"github.com/sacloud/libsacloud/v2/internal/define/ops"
	"github.com/sacloud/libsacloud/v2/internal/schema"
	"github.com/sacloud/libsacloud/v2/internal/schema/meta"
	"github.com/sacloud/libsacloud/v2/sacloud/naked"
)

const (
	interfaceAPIName     = "Interface"
	interfaceAPIPathName = "interface"
)

var interfaceAPI = &schema.Resource{
	Name:       interfaceAPIName,
	PathName:   interfaceAPIPathName,
	PathSuffix: schema.CloudAPISuffix,
	Operations: schema.Operations{
		// find
		ops.Find(interfaceAPIName, interfaceNakedType, findParameter, interfaceView),

		// create
		ops.Create(interfaceAPIName, interfaceNakedType, interfaceCreateParam, interfaceView),

		// read
		ops.Read(interfaceAPIName, interfaceNakedType, interfaceView),

		// update
		ops.Update(interfaceAPIName, interfaceNakedType, interfaceUpdateParam, interfaceView),

		// delete
		ops.Delete(interfaceAPIName),

		// monitor
		ops.Monitor(interfaceAPIName, monitorParameter, monitors.interfaceModel()),

		ops.WithIDAction(interfaceAPIName, "ConnectToSharedSegment", http.MethodPut, "to/switch/shared"),

		ops.WithIDAction(interfaceAPIName, "ConnectToSwitch", http.MethodPut, "to/switch/{{.switchID}}",
			&schema.Argument{
				Name: "switchID",
				Type: meta.TypeID,
			},
		),

		ops.WithIDAction(interfaceAPIName, "DisconnectFromSwitch", http.MethodDelete, "to/switch"),

		ops.WithIDAction(interfaceAPIName, "ConnectToPacketFilter", http.MethodPut, "to/packetfilter/{{.packetFilterID}}",
			&schema.Argument{
				Name: "packetFilterID",
				Type: meta.TypeID,
			},
		),

		ops.WithIDAction(interfaceAPIName, "DisconnectFromPacketFilter", http.MethodDelete, "to/packetfilter"),
	},
}
var (
	interfaceNakedType = meta.Static(naked.Interface{})

	interfaceView = &schema.Model{
		Name:      interfaceAPIName,
		NakedType: interfaceNakedType,
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
		Name:      names.CreateParameterName(interfaceAPIName),
		NakedType: interfaceNakedType,
		Fields: []*schema.FieldDesc{
			fields.ServerID(),
		},
	}

	interfaceUpdateParam = &schema.Model{
		Name:      names.UpdateParameterName(interfaceAPIName),
		NakedType: interfaceNakedType,
		Fields: []*schema.FieldDesc{
			fields.UserIPAddress(),
		},
	}
)
