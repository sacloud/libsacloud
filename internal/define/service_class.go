package define

import (
	"github.com/sacloud/libsacloud/v2/internal/define/ops"
	"github.com/sacloud/libsacloud/v2/internal/dsl"
	"github.com/sacloud/libsacloud/v2/internal/dsl/meta"
	"github.com/sacloud/libsacloud/v2/sacloud/naked"
)

const (
	serviceClassAPIName     = "ServiceClass"
	serviceClassAPIPathName = "public/price"
)

var serviceClassAPI = &dsl.Resource{
	Name:       serviceClassAPIName,
	PathName:   serviceClassAPIPathName,
	PathSuffix: dsl.CloudAPISuffix,
	Operations: dsl.Operations{
		ops.Find(serviceClassAPIName, serviceClassNakedType, findParameter, serviceClassView, nil),
	},
}

var (
	serviceClassNakedType = meta.Static(naked.ServiceClass{})
	serviceClassView      = &dsl.Model{
		Name:      serviceClassAPIName,
		NakedType: serviceClassNakedType,
		Fields: []*dsl.FieldDesc{
			fields.ID(),
			fields.Def("ServiceClassName", meta.TypeString),
			fields.Def("ServiceClassPath", meta.TypeString),
			fields.Def("DisplayName", meta.TypeString),
			fields.Def("IsPublic", meta.TypeFlag),
			{
				Name: "Price",
				Type: priceModel,
				Tags: &dsl.FieldTags{
					MapConv: ",recursive",
				},
			},
		},
	}

	priceModel = &dsl.Model{
		Name:      "Price",
		NakedType: meta.Static(naked.Price{}),
		Fields: []*dsl.FieldDesc{
			fields.Def("Base", meta.TypeInt),
			fields.Def("Daily", meta.TypeInt),
			fields.Def("Hourly", meta.TypeInt),
			fields.Def("Monthly", meta.TypeInt),
			fields.Def("Zone", meta.TypeString),
		},
	}
)
