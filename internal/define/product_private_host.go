package define

import (
	"github.com/sacloud/libsacloud/v2/internal/define/ops"
	"github.com/sacloud/libsacloud/v2/internal/dsl"
	"github.com/sacloud/libsacloud/v2/internal/dsl/meta"
	"github.com/sacloud/libsacloud/v2/sacloud/naked"
)

const (
	productPrivateHostAPIName     = "PrivateHostPlan"
	productPrivateHostAPIPathName = "product/privatehost"
)

var productPrivateHostAPI = &dsl.Resource{
	Name:       productPrivateHostAPIName,
	PathName:   productPrivateHostAPIPathName,
	PathSuffix: dsl.CloudAPISuffix,
	IsGlobal:   true,
	Operations: dsl.Operations{
		ops.Find(productPrivateHostAPIName, productPrivateHostNakedType, findParameter, productPrivateHostView),
		ops.Read(productPrivateHostAPIName, productPrivateHostNakedType, productPrivateHostView),
	},
}

var (
	productPrivateHostNakedType = meta.Static(naked.ProductPrivateHost{})
	productPrivateHostView      = &dsl.Model{
		Name:      productPrivateHostAPIName,
		NakedType: productPrivateHostNakedType,
		Fields: []*dsl.FieldDesc{
			fields.ID(),
			fields.Name(),
			fields.Class(),
			fields.CPU(),
			fields.MemoryMB(),
		},
	}
)
