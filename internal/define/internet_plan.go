package define

import (
	"github.com/sacloud/libsacloud/v2/internal/define/ops"
	"github.com/sacloud/libsacloud/v2/internal/dsl"
	"github.com/sacloud/libsacloud/v2/internal/dsl/meta"
	"github.com/sacloud/libsacloud/v2/sacloud/naked"
)

const (
	internetPlanAPIName     = "InternetPlan"
	internetPlanAPIPathName = "product/internet"
)

var internetPlanAPI = &dsl.Resource{
	Name:       internetPlanAPIName,
	PathName:   internetPlanAPIPathName,
	PathSuffix: dsl.CloudAPISuffix,
	Operations: dsl.Operations{
		ops.Find(internetPlanAPIName, internetPlanNakedType, findParameter, internetPlanView, nil),
		ops.Read(internetPlanAPIName, internetPlanNakedType, internetPlanView),
	},
}

var (
	internetPlanNakedType = meta.Static(naked.InternetPlan{})
	internetPlanView      = &dsl.Model{
		Name:      internetPlanAPIName,
		NakedType: internetPlanNakedType,
		Fields: []*dsl.FieldDesc{
			fields.ID(),
			fields.Name(),
			fields.Def("BandWidthMbps", meta.TypeInt),
			fields.Availability(),
		},
	}
)
