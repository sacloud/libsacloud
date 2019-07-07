package define

import (
	"github.com/sacloud/libsacloud/v2/internal/define/names"
	"github.com/sacloud/libsacloud/v2/internal/define/ops"
	"github.com/sacloud/libsacloud/v2/internal/schema"
	"github.com/sacloud/libsacloud/v2/internal/schema/meta"
	"github.com/sacloud/libsacloud/v2/sacloud/naked"
)

const (
	gslbAPIName     = "GSLB"
	gslbAPIPathName = "commonserviceitem"
)

var gslbAPI = &schema.Resource{
	Name:       gslbAPIName,
	PathName:   gslbAPIPathName,
	PathSuffix: schema.CloudAPISuffix,
	IsGlobal:   true,
	Operations: schema.Operations{
		// find
		ops.FindCommonServiceItem(gslbAPIName, gslbNakedType, findParameter, gslbView),

		// create
		ops.CreateCommonServiceItem(gslbAPIName, gslbNakedType, gslbCreateParam, gslbView),

		// read
		ops.ReadCommonServiceItem(gslbAPIName, gslbNakedType, gslbView),

		// update
		ops.UpdateCommonServiceItem(gslbAPIName, gslbNakedType, gslbUpdateParam, gslbView),

		// delete
		ops.Delete(gslbAPIName),
	},
}

var (
	gslbNakedType = meta.Static(naked.GSLB{})

	gslbView = &schema.Model{
		Name:      gslbAPIName,
		NakedType: gslbNakedType,
		Fields: []*schema.FieldDesc{
			fields.ID(),
			fields.Name(),
			fields.Description(),
			fields.Tags(),
			fields.Availability(),
			fields.IconID(),
			fields.CreatedAt(),
			fields.ModifiedAt(),
			fields.GSLBProviderClass(),
			fields.SettingsHash(),
			fields.GSLBFQDN(),
			// settings
			fields.GSLBDelayLoop(),
			fields.GSLBWeighted(),
			fields.GSLBHealthCheckProtocol(),
			fields.GSLBHealthCheckHostHeader(),
			fields.GSLBHealthCheckPath(),
			fields.GSLBHealthCheckResponseCode(),
			fields.GSLBHealthCheckPort(),
			fields.GSLBSorryServer(),
			fields.GSLBDestinationServers(),
		},
	}

	gslbCreateParam = &schema.Model{
		Name:      names.CreateParameterName(gslbAPIName),
		NakedType: gslbNakedType,
		Fields: []*schema.FieldDesc{
			fields.GSLBProviderClass(),

			fields.GSLBHealthCheckProtocol(),
			fields.GSLBHealthCheckHostHeader(),
			fields.GSLBHealthCheckPath(),
			fields.GSLBHealthCheckResponseCode(),
			fields.GSLBHealthCheckPort(),
			fields.GSLBDelayLoop(),
			fields.GSLBWeighted(),
			fields.GSLBSorryServer(),
			fields.GSLBDestinationServers(),

			fields.Name(),
			fields.Description(),
			fields.Tags(),
			fields.IconID(),
		},
	}

	gslbUpdateParam = &schema.Model{
		Name:      names.UpdateParameterName(gslbAPIName),
		NakedType: gslbNakedType,
		Fields: []*schema.FieldDesc{
			fields.GSLBHealthCheckProtocol(),
			fields.GSLBHealthCheckHostHeader(),
			fields.GSLBHealthCheckPath(),
			fields.GSLBHealthCheckResponseCode(),
			fields.GSLBHealthCheckPort(),
			fields.GSLBDelayLoop(),
			fields.GSLBWeighted(),
			fields.GSLBSorryServer(),
			fields.GSLBDestinationServers(),

			fields.Name(),
			fields.Description(),
			fields.Tags(),
			fields.IconID(),
		},
	}
)
