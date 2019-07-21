package define

import (
	"github.com/sacloud/libsacloud/v2/internal/define/ops"
	"github.com/sacloud/libsacloud/v2/internal/dsl"
	"github.com/sacloud/libsacloud/v2/internal/dsl/meta"
	"github.com/sacloud/libsacloud/v2/sacloud/naked"
)

const (
	serverPlanAPIName     = "ServerPlan"
	serverPlanAPIPathName = "product/server"
)

var serverPlanAPI = &dsl.Resource{
	Name:       serverPlanAPIName,
	PathName:   serverPlanAPIPathName,
	PathSuffix: dsl.CloudAPISuffix,
	Operations: dsl.Operations{
		ops.Find(serverPlanAPIName, serverPlanNakedType, findParameter, serverPlanView, nil),
		ops.Read(serverPlanAPIName, serverPlanNakedType, serverPlanView),
	},
}

var (
	serverPlanNakedType = meta.Static(naked.ServerPlan{})
	serverPlanView      = &dsl.Model{
		Name:      serverPlanAPIName,
		NakedType: serverPlanNakedType,
		Fields: []*dsl.FieldDesc{
			fields.ID(),
			fields.Name(),
			fields.CPU(),
			fields.MemoryMB(),
			fields.Def("Commitment", meta.TypeCommitment),
			fields.Def("Generation", meta.TypeInt),
			fields.Availability(),
		},
	}
)
