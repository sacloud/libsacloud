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
		ops.Find(serviceClassAPIName, serviceClassNakedType, findParameter, serviceClassView),
	},
}

var (
	serviceClassNakedType = meta.Static(naked.ServiceClass{})
	serviceClassView      = &dsl.Model{
		Name:      serviceClassAPIName,
		NakedType: serviceClassNakedType,
		Fields: []*dsl.FieldDesc{
			fields.ID(),
			fields.New("ServiceClassName", meta.TypeString),
			fields.New("ServiceClassPath", meta.TypeString),
			fields.New("DisplayName", meta.TypeString),
			fields.New("IsPublic", meta.TypeFlag),
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
			fields.New("Base", meta.TypeInt),
			fields.New("Daily", meta.TypeInt),
			fields.New("Hourly", meta.TypeInt),
			fields.New("Monthly", meta.TypeInt),
			fields.New("Zone", meta.TypeString),
		},
	}
)
