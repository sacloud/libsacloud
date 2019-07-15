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
	IsGlobal:   true,
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
			fields.New("PlanName", meta.TypeString),
			fields.New("PlanClass", meta.TypeString),
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
				ExtendAccessors: []*dsl.ExtendAccessor{
					{
						Name: "MemoryGB",
						Type: meta.TypeInt,
					},
				},
			},
			fields.New("AssignedCPU", meta.TypeInt),
			fields.New("AssignedMemoryMB", meta.TypeInt),
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
