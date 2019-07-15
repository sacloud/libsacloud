package define

import (
	"github.com/sacloud/libsacloud/v2/internal/define/ops"
	"github.com/sacloud/libsacloud/v2/internal/dsl"
	"github.com/sacloud/libsacloud/v2/internal/dsl/meta"
	"github.com/sacloud/libsacloud/v2/sacloud/naked"
)

const (
	diskPlanAPIName     = "DiskPlan"
	diskPlanAPIPathName = "product/disk"
)

var diskPlanAPI = &dsl.Resource{
	Name:       diskPlanAPIName,
	PathName:   diskPlanAPIPathName,
	PathSuffix: dsl.CloudAPISuffix,
	Operations: dsl.Operations{
		ops.Find(diskPlanAPIName, diskPlanNakedType, findParameter, diskPlanView),
		ops.Read(diskPlanAPIName, diskPlanNakedType, diskPlanView),
	},
}

var (
	diskPlanNakedType = meta.Static(naked.DiskPlan{})
	diskPlanView      = &dsl.Model{
		Name:      diskPlanAPIName,
		NakedType: diskPlanNakedType,
		Fields: []*dsl.FieldDesc{
			fields.ID(),
			fields.Name(),
			fields.New("StorageClass", meta.TypeString),
			fields.Availability(),
			{
				Name: "Size",
				Type: &dsl.Model{
					Name:      "DiskPlanSizeInfo",
					NakedType: meta.Static(naked.DiskPlanSizeInfo{}),
					IsArray:   true,
					Fields: []*dsl.FieldDesc{
						fields.Availability(),
						fields.New("DisplaySize", meta.TypeInt),
						fields.New("DisplaySuffix", meta.TypeString),
						fields.SizeMB(),
					},
				},
				Tags: &dsl.FieldTags{
					MapConv: "[]Size,recursive",
				},
			},
		},
	}
)
