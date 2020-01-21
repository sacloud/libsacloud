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
	"github.com/sacloud/libsacloud/v2/internal/define/names"
	"github.com/sacloud/libsacloud/v2/internal/define/ops"
	"github.com/sacloud/libsacloud/v2/internal/dsl"
	"github.com/sacloud/libsacloud/v2/internal/dsl/meta"
	"github.com/sacloud/libsacloud/v2/sacloud/naked"
)

const (
	localRouterAPIName     = "LocalRouter"
	localRouterAPIPathName = "commonserviceitem"
)

var localRouterAPI = &dsl.Resource{
	Name:       localRouterAPIName,
	PathName:   localRouterAPIPathName,
	PathSuffix: dsl.CloudAPISuffix,
	IsGlobal:   true,
	Operations: dsl.Operations{
		// find
		ops.FindCommonServiceItem(localRouterAPIName, localRouterNakedType, findParameter, localRouterView),

		// create
		ops.CreateCommonServiceItem(localRouterAPIName, localRouterNakedType, localRouterCreateParam, localRouterView),

		// read
		ops.ReadCommonServiceItem(localRouterAPIName, localRouterNakedType, localRouterView),

		// update
		ops.UpdateCommonServiceItem(localRouterAPIName, localRouterNakedType, localRouterUpdateParam, localRouterView),
		// updateSettings
		ops.UpdateCommonServiceItemSettings(localRouterAPIName, localRouterUpdateSettingsNakedType, localRouterUpdateSettingsParam, localRouterView),

		// delete
		ops.Delete(localRouterAPIName),

		// Health
		ops.HealthStatus(localRouterAPIName, meta.Static(naked.LocalRouterHealth{}), localRouterHealth),

		// Monitor
		ops.MonitorChild(localRouterAPIName, "LocalRouter", "activity/localrouter",
			monitorParameter, monitors.localRouterModel()),
	},
}

var (
	localRouterNakedType               = meta.Static(naked.LocalRouter{})
	localRouterUpdateSettingsNakedType = meta.Static(naked.LocalRouterSettingsUpdate{})

	localRouterView = &dsl.Model{
		Name:      localRouterAPIName,
		NakedType: localRouterNakedType,
		Fields: []*dsl.FieldDesc{
			fields.ID(),
			fields.Name(),
			fields.Description(),
			fields.Tags(),
			fields.Availability(),
			fields.IconID(),
			fields.CreatedAt(),
			fields.ModifiedAt(),

			// settings
			fields.LocalRouterSwitch(),
			fields.LocalRouterInterface(),
			fields.LocalRouterPeers(),
			fields.LocalRouterStaticRoutes(),
			fields.SettingsHash(),

			// status
			fields.LocalRouterSecretKeys(),
		},
	}

	localRouterCreateParam = &dsl.Model{
		Name:      names.CreateParameterName(localRouterAPIName),
		NakedType: localRouterNakedType,
		ConstFields: []*dsl.ConstFieldDesc{
			{
				Name: "Class",
				Type: meta.TypeString,
				Tags: &dsl.FieldTags{
					MapConv: "Provider.Class",
				},
				Value: `"localrouter"`,
			},
		},
		Fields: []*dsl.FieldDesc{
			// common fields
			fields.Name(),
			fields.Description(),
			fields.Tags(),
			fields.IconID(),
		},
	}

	localRouterUpdateParam = &dsl.Model{
		Name:      names.UpdateParameterName(localRouterAPIName),
		NakedType: localRouterNakedType,
		Fields: []*dsl.FieldDesc{
			// settings
			fields.LocalRouterSwitch(),
			fields.LocalRouterInterface(),
			fields.LocalRouterPeers(),
			fields.LocalRouterStaticRoutes(),
			// settings hash
			fields.SettingsHash(),

			// common fields
			fields.Name(),
			fields.Description(),
			fields.Tags(),
			fields.IconID(),
		},
	}

	localRouterUpdateSettingsParam = &dsl.Model{
		Name:      names.UpdateSettingsParameterName(localRouterAPIName),
		NakedType: localRouterNakedType,
		Fields: []*dsl.FieldDesc{
			// settings
			fields.LocalRouterSwitch(),
			fields.LocalRouterInterface(),
			fields.LocalRouterPeers(),
			fields.LocalRouterStaticRoutes(),
			// settings hash
			fields.SettingsHash(),
		},
	}

	localRouterHealth = &dsl.Model{
		Name: "LocalRouterHealth",
		Fields: []*dsl.FieldDesc{
			{
				Name: "Peers",
				Type: &dsl.Model{
					Name:    "LocalRouterHealthPeer",
					IsArray: true,
					Fields: []*dsl.FieldDesc{
						fields.Def("ID", meta.TypeID),
						fields.Def("Status", meta.TypeInstanceStatus),
						fields.Def("Routes", meta.TypeStringSlice),
					},
				},
				Tags: &dsl.FieldTags{
					MapConv: "LocalRouter.[]Peers,recursive",
				},
			},
		},
		NakedType: meta.Static(naked.LocalRouterHealth{}),
	}
)
