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
	switchAPIName     = "Switch"
	switchAPIPathName = "switch"
)

var switchAPI = &schema.Resource{
	Name:       switchAPIName,
	PathName:   switchAPIPathName,
	PathSuffix: schema.CloudAPISuffix,
	Operations: schema.Operations{
		// find
		ops.Find(switchAPIName, switchNakedType, findParameter, switchView),

		// create
		ops.Create(switchAPIName, switchNakedType, switchCreateParam, switchView),

		// read
		ops.Read(switchAPIName, switchNakedType, switchView),

		// update
		ops.Update(switchAPIName, switchNakedType, switchUpdateParam, switchView),

		// delete
		ops.Delete(switchAPIName),

		// connect from bridge
		ops.WithIDAction(switchAPIName, "ConnectToBridge", http.MethodPut, "to/bridge/{{.bridgeID}}",
			&schema.Argument{
				Name: "bridgeID",
				Type: meta.TypeID,
			},
		),

		// disconnect from bridge
		ops.WithIDAction(switchAPIName, "DisconnectFromBridge", http.MethodDelete, "to/bridge/"),
	},
}

var (
	switchNakedType = meta.Static(naked.Switch{})

	switchView = &schema.Model{
		Name:      switchAPIName,
		NakedType: switchNakedType,
		Fields: []*schema.FieldDesc{
			fields.ID(),
			fields.Name(),
			fields.Description(),
			fields.Tags(),
			fields.IconID(),
			fields.CreatedAt(),
			fields.ModifiedAt(),
			fields.Scope(),
			fields.UserSubnetNetworkMaskLen(),
			fields.UserSubnetDefaultRoute(),
			{
				Name: "Subnets",
				Type: models.switchSubnet(),
				Tags: &schema.FieldTags{
					MapConv: "[]Subnets,omitempty,recursive",
					JSON:    ",omitempty",
				},
			},
			fields.BridgeID(),
		},
	}

	switchCreateParam = &schema.Model{
		Name:      names.CreateParameterName(switchAPIName),
		NakedType: switchNakedType,
		Fields: []*schema.FieldDesc{
			fields.Name(),
			fields.UserSubnetNetworkMaskLen(),
			fields.UserSubnetDefaultRoute(),
			fields.Description(),
			fields.Tags(),
			fields.IconID(),
		},
	}

	switchUpdateParam = &schema.Model{
		Name:      names.UpdateParameterName(switchAPIName),
		NakedType: switchNakedType,
		Fields: []*schema.FieldDesc{
			fields.Name(),
			fields.UserSubnetNetworkMaskLen(),
			fields.UserSubnetDefaultRoute(),
			fields.Description(),
			fields.Tags(),
			fields.IconID(),
		},
	}
)
