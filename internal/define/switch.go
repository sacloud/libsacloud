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
)

const (
	switchAPIName     = "Switch"
	switchAPIPathName = "switch"
)

var switchAPI = &dsl.Resource{
	Name:       switchAPIName,
	PathName:   switchAPIPathName,
	PathSuffix: dsl.CloudAPISuffix,
	Operations: dsl.Operations{
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
			&dsl.Argument{
				Name: "bridgeID",
				Type: meta.TypeID,
			},
		),

		// disconnect from bridge
		ops.WithIDAction(switchAPIName, "DisconnectFromBridge", http.MethodDelete, "to/bridge/"),

		// find connected servers
		{
			ResourceName:     switchAPIName,
			Name:             "GetServers",
			PathFormat:       dsl.IDAndSuffixPathFormat("server"),
			Method:           http.MethodGet,
			UseWrappedResult: true,
			Arguments: dsl.Arguments{
				dsl.ArgumentID,
			},
			ResponseEnvelope: dsl.ResponseEnvelopePlural(&dsl.EnvelopePayloadDesc{
				Type: serverNakedType,
				Name: names.ResourceFieldName(serverAPIName, dsl.PayloadForms.Plural),
			}),
			Results: dsl.Results{
				{
					SourceField: names.ResourceFieldName(serverAPIName, dsl.PayloadForms.Plural),
					DestField:   names.ResourceFieldName(serverAPIName, dsl.PayloadForms.Plural),
					IsPlural:    true,
					Model:       serverView,
				},
			},
		},
	},
}

var (
	switchNakedType = meta.Static(naked.Switch{})

	switchView = &dsl.Model{
		Name:      switchAPIName,
		NakedType: switchNakedType,
		Fields: []*dsl.FieldDesc{
			fields.ID(),
			fields.Name(),
			fields.Description(),
			fields.Tags(),
			fields.IconID(),
			fields.CreatedAt(),
			fields.ModifiedAt(),
			fields.Scope(),
			fields.Def("ServerCount", meta.TypeInt),
			fields.UserSubnetNetworkMaskLen(),
			fields.UserSubnetDefaultRoute(),
			{
				Name: "Subnets",
				Type: models.switchSubnet(),
				Tags: &dsl.FieldTags{
					MapConv: "[]Subnets,omitempty,recursive",
					JSON:    ",omitempty",
				},
			},
			fields.BridgeID(),
		},
	}

	switchCreateParam = &dsl.Model{
		Name:      names.CreateParameterName(switchAPIName),
		NakedType: switchNakedType,
		Fields: []*dsl.FieldDesc{
			fields.Name(),
			fields.UserSubnetNetworkMaskLen(),
			fields.UserSubnetDefaultRoute(),
			fields.Description(),
			fields.Tags(),
			fields.IconID(),
		},
	}

	switchUpdateParam = &dsl.Model{
		Name:      names.UpdateParameterName(switchAPIName),
		NakedType: switchNakedType,
		Fields: []*dsl.FieldDesc{
			fields.Name(),
			fields.UserSubnetNetworkMaskLen(),
			fields.UserSubnetDefaultRoute(),
			fields.Description(),
			fields.Tags(),
			fields.IconID(),
		},
	}
)
