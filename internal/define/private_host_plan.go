package define

import (
	"github.com/sacloud/libsacloud/v2/internal/define/ops"
	"github.com/sacloud/libsacloud/v2/internal/dsl"
	"github.com/sacloud/libsacloud/v2/internal/dsl/meta"
	"github.com/sacloud/libsacloud/v2/sacloud/naked"
)

const (
	privateHostPlanAPIName     = "PrivateHostPlan"
	privateHostPlanAPIPathName = "product/privatehost"
)

var privateHostPlanAPI = &dsl.Resource{
	Name:       privateHostPlanAPIName,
	PathName:   privateHostPlanAPIPathName,
	PathSuffix: dsl.CloudAPISuffix,
	Operations: dsl.Operations{
		ops.Find(privateHostPlanAPIName, privateHostPlanNakedType, findParameter, privateHostPlanView, nil),
		ops.Read(privateHostPlanAPIName, privateHostPlanNakedType, privateHostPlanView),
	},
}

var (
	privateHostPlanNakedType = meta.Static(naked.PrivateHostPlan{})
	privateHostPlanView      = &dsl.Model{
		Name:      privateHostPlanAPIName,
		NakedType: privateHostPlanNakedType,
		Fields: []*dsl.FieldDesc{
			fields.ID(),
			fields.Name(),
			fields.Class(),
			fields.CPU(),
			fields.MemoryMB(),
			fields.Availability(),
		},
	}
)
