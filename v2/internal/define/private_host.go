// Copyright 2016-2022 The Libsacloud Authors
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
	privateHostAPIName     = "PrivateHost"
	privateHostAPIPathName = "privatehost"
)

var privateHostAPI = &dsl.Resource{
	Name:       privateHostAPIName,
	PathName:   privateHostAPIPathName,
	PathSuffix: dsl.CloudAPISuffix,
	Operations: dsl.Operations{
		// find
		ops.Find(privateHostAPIName, privateHostNakedType, findParameter, privateHostView),

		// create
		ops.Create(privateHostAPIName, privateHostNakedType, privateHostCreateParam, privateHostView),

		// read
		ops.Read(privateHostAPIName, privateHostNakedType, privateHostView),

		// update
		ops.Update(privateHostAPIName, privateHostNakedType, privateHostUpdateParam, privateHostView),

		// delete
		ops.Delete(privateHostAPIName),
	},
}

var (
	privateHostNakedType = meta.Static(naked.PrivateHost{})

	privateHostView = &dsl.Model{
		Name:      privateHostAPIName,
		NakedType: privateHostNakedType,
		Fields: []*dsl.FieldDesc{
			fields.ID(),
			fields.Name(),
			fields.Description(),
			fields.Tags(),
			fields.IconID(),
			fields.CreatedAt(),
			fields.PrivateHostPlanID(),
			{
				Name: "PlanName",
				Type: meta.TypeString,
				Tags: &dsl.FieldTags{
					MapConv: "Plan.Name",
				},
			},
			{
				Name: "PlanClass",
				Type: meta.TypeString,
				Tags: &dsl.FieldTags{
					MapConv: "Plan.Class",
				},
			},
			{
				Name: "CPU",
				Type: meta.TypeInt,
				Tags: &dsl.FieldTags{
					MapConv: "Plan.CPU",
				},
			},
			{
				Name: "MemoryMB",
				Type: meta.TypeInt,
				Tags: &dsl.FieldTags{
					MapConv: "Plan.MemoryMB",
				},
				Methods: []*dsl.MethodDesc{
					{
						Name:        "GetMemoryGB",
						ResultTypes: []meta.Type{meta.TypeInt},
					},
					{
						Name: "SetMemoryGB",
						Arguments: dsl.Arguments{
							{
								Name: "memory",
								Type: meta.TypeInt,
							},
						},
					},
				},
			},
			fields.Def("AssignedCPU", meta.TypeInt),
			{
				Name: "AssignedMemoryMB",
				Type: meta.TypeInt,
				Methods: []*dsl.MethodDesc{
					{
						Name:        "GetAssignedMemoryGB",
						ResultTypes: []meta.Type{meta.TypeInt},
					},
					{
						Name: "SetAssignedMemoryGB",
						Arguments: dsl.Arguments{
							{
								Name: "memory",
								Type: meta.TypeInt,
							},
						},
					},
				},
			},
			fields.PrivateHostHostName(),
		},
	}

	privateHostCreateParam = &dsl.Model{
		Name:      names.CreateParameterName(privateHostAPIName),
		NakedType: privateHostNakedType,
		Fields: []*dsl.FieldDesc{
			fields.Name(),
			fields.Description(),
			fields.Tags(),
			fields.IconID(),
			fields.PrivateHostPlanID(),
		},
	}

	privateHostUpdateParam = &dsl.Model{
		Name:      names.UpdateParameterName(privateHostAPIName),
		NakedType: privateHostNakedType,
		Fields: []*dsl.FieldDesc{
			fields.Name(),
			fields.Description(),
			fields.Tags(),
			fields.IconID(),
		},
	}
)
