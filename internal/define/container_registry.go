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
	containerRegistryAPIName     = "ContainerRegistry"
	containerRegistryAPIPathName = "commonserviceitem"
)

var containerRegistryAPI = &dsl.Resource{
	Name:       containerRegistryAPIName,
	PathName:   containerRegistryAPIPathName,
	PathSuffix: dsl.CloudAPISuffix,
	IsGlobal:   true,
	Operations: dsl.Operations{
		// find
		ops.FindCommonServiceItem(containerRegistryAPIName, containerRegistryNakedType, findParameter, containerRegistryView),

		// create
		ops.CreateCommonServiceItem(containerRegistryAPIName, containerRegistryNakedType, containerRegistryCreateParam, containerRegistryView),

		// read
		ops.ReadCommonServiceItem(containerRegistryAPIName, containerRegistryNakedType, containerRegistryView),

		// update
		ops.UpdateCommonServiceItem(containerRegistryAPIName, containerRegistryNakedType, containerRegistryUpdateParam, containerRegistryView),
		// updateSettings
		ops.UpdateCommonServiceItemSettings(containerRegistryAPIName, containerRegistryUpdateSettingsNakedType, containerRegistryUpdateSettingsParam, containerRegistryView),

		// delete
		ops.Delete(containerRegistryAPIName),

		// list users
		{
			ResourceName: containerRegistryAPIName,
			Name:         "ListUsers",
			PathFormat:   dsl.IDAndSuffixPathFormat("containerregistry/users"),
			Method:       http.MethodGet,
			Arguments: dsl.Arguments{
				dsl.ArgumentID,
			},
			ResponseEnvelope: dsl.ResponseEnvelope(&dsl.EnvelopePayloadDesc{
				Type: meta.Static(naked.ContainerRegistryUsers{}),
				Name: containerRegistryAPIName,
			}),
			Results: dsl.Results{
				{
					SourceField: containerRegistryAPIName,
					DestField:   containerRegistryUserListView.Name,
					Model:       containerRegistryUserListView,
				},
			},
		},
		// Add User
		{
			ResourceName: containerRegistryAPIName,
			Name:         "AddUser",
			PathFormat:   dsl.IDAndSuffixPathFormat("containerregistry/users"),
			Method:       http.MethodPost,
			RequestEnvelope: dsl.RequestEnvelope(&dsl.EnvelopePayloadDesc{
				Type: containerRegistryUserNakedType,
				Name: containerRegistryAPIName,
			}),
			Arguments: dsl.Arguments{
				dsl.ArgumentID,
				dsl.MappableArgument("param", containerRegistryUserCreateParam, containerRegistryAPIName),
			},
		},
		// Update User
		{
			ResourceName: containerRegistryAPIName,
			Name:         "UpdateUser",
			PathFormat:   dsl.IDAndSuffixPathFormat("containerregistry/users/{{.username}}"),
			Method:       http.MethodPut,
			RequestEnvelope: dsl.RequestEnvelope(&dsl.EnvelopePayloadDesc{
				Type: containerRegistryUserNakedType,
				Name: containerRegistryAPIName,
			}),
			Arguments: dsl.Arguments{
				dsl.ArgumentID,
				&dsl.Argument{
					Name: "username",
					Type: meta.TypeString,
				},
				dsl.MappableArgument("param", containerRegistryUserUpdateParam, containerRegistryAPIName),
			},
		},
		// delete certificates
		{
			ResourceName: containerRegistryAPIName,
			Name:         "DeleteUser",
			PathFormat:   dsl.IDAndSuffixPathFormat("containerregistry/users/{{.username}}"),
			Method:       http.MethodDelete,
			Arguments: dsl.Arguments{
				dsl.ArgumentID,
				&dsl.Argument{
					Name: "username",
					Type: meta.TypeString,
				},
			},
		},
	},
}

var (
	containerRegistryNakedType               = meta.Static(naked.ContainerRegistry{})
	containerRegistryUpdateSettingsNakedType = meta.Static(naked.ContainerRegistrySettingsUpdate{})
	containerRegistryUserNakedType           = meta.Static(naked.ContainerRegistryUser{})

	containerRegistryView = &dsl.Model{
		Name:      containerRegistryAPIName,
		NakedType: containerRegistryNakedType,
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
			fields.ContainerRegistryAccessLevel(),
			fields.ContainerRegistryVirtualDomain(),
			fields.SettingsHash(),

			// status
			fields.ContainerRegistrySubDomainLabel(),
			fields.ContainerRegistryFQDN(),
		},
	}

	containerRegistryCreateParam = &dsl.Model{
		Name:      names.CreateParameterName(containerRegistryAPIName),
		NakedType: containerRegistryNakedType,
		ConstFields: []*dsl.ConstFieldDesc{
			{
				Name: "Class",
				Type: meta.TypeString,
				Tags: &dsl.FieldTags{
					MapConv: "Provider.Class",
				},
				Value: `"containerregistry"`,
			},
		},
		Fields: []*dsl.FieldDesc{
			// common fields
			fields.Name(),
			fields.Description(),
			fields.Tags(),
			fields.IconID(),

			// settings
			fields.ContainerRegistryAccessLevel(),
			fields.ContainerRegistryVirtualDomain(),
			// status
			fields.ContainerRegistrySubDomainLabel(),
		},
	}

	containerRegistryUpdateParam = &dsl.Model{
		Name:      names.UpdateParameterName(containerRegistryAPIName),
		NakedType: containerRegistryNakedType,
		Fields: []*dsl.FieldDesc{
			// common fields
			fields.Name(),
			fields.Description(),
			fields.Tags(),
			fields.IconID(),

			// settings
			fields.ContainerRegistryAccessLevel(),
			fields.ContainerRegistryVirtualDomain(),
			// settings hash
			fields.SettingsHash(),
		},
	}

	containerRegistryUpdateSettingsParam = &dsl.Model{
		Name:      names.UpdateSettingsParameterName(containerRegistryAPIName),
		NakedType: containerRegistryNakedType,
		Fields: []*dsl.FieldDesc{
			// settings
			fields.ContainerRegistryAccessLevel(),
			fields.ContainerRegistryVirtualDomain(),
			// settings hash
			fields.SettingsHash(),
		},
	}

	containerRegistryUserListView = &dsl.Model{
		Name:      containerRegistryAPIName + "Users",
		NakedType: meta.Static(naked.ContainerRegistryUsers{}),
		Fields: []*dsl.FieldDesc{
			{
				Name: "Users",
				Type: &dsl.Model{
					Name:      containerRegistryAPIName + "User",
					NakedType: meta.Static(naked.ContainerRegistryUser{}),
					IsArray:   true,
					Fields: []*dsl.FieldDesc{
						fields.Def("UserName", meta.TypeString),
						fields.Def("Permission", meta.Static(types.EContainerRegistryPermission(""))),
					},
				},
			},
		},
	}
	containerRegistryUserCreateParam = &dsl.Model{
		Name:      names.CreateParameterName(containerRegistryAPIName + "User"),
		NakedType: meta.Static(naked.ContainerRegistryUser{}),
		Fields: []*dsl.FieldDesc{
			fields.Def("UserName", meta.TypeString),
			fields.Def("Password", meta.TypeString),
			fields.Def("Permission", meta.Static(types.EContainerRegistryPermission(""))),
		},
	}
	containerRegistryUserUpdateParam = &dsl.Model{
		Name:      names.UpdateParameterName(containerRegistryAPIName + "User"),
		NakedType: meta.Static(naked.ContainerRegistryUser{}),
		Fields: []*dsl.FieldDesc{
			fields.Def("Password", meta.TypeString),
			fields.Def("Permission", meta.Static(types.EContainerRegistryPermission(""))),
		},
	}
)
