package define

import (
	"github.com/sacloud/libsacloud/v2/internal/define/ops"
	"github.com/sacloud/libsacloud/v2/internal/dsl"
	"github.com/sacloud/libsacloud/v2/internal/dsl/meta"
	"github.com/sacloud/libsacloud/v2/sacloud/naked"
)

const (
	regionAPIName     = "Region"
	regionAPIPathName = "region"
)

var regionAPI = &dsl.Resource{
	Name:       regionAPIName,
	PathName:   regionAPIPathName,
	PathSuffix: dsl.CloudAPISuffix,
	IsGlobal:   true,
	Operations: dsl.Operations{
		ops.Find(regionAPIName, regionNakedType, findParameter, regionView),
		ops.Read(regionAPIName, regionNakedType, regionView),
	},
}

var (
	regionNakedType = meta.Static(naked.Region{})
	regionView      = models.region()
)
