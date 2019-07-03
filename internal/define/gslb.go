package define

import (
	"github.com/sacloud/libsacloud/v2/internal/schema"
	"github.com/sacloud/libsacloud/v2/internal/schema/meta"
	"github.com/sacloud/libsacloud/v2/sacloud/naked"
)

var gslbAPI = &schema.Resource{
	Name:       "GSLB",
	PathName:   "commonserviceitem",
	PathSuffix: schema.CloudAPISuffix,
	IsGlobal:   true,
	OperationsDefineFunc: func(r *schema.Resource) []*schema.Operation {
		return []*schema.Operation{
			// find
			r.DefineOperationCommonServiceItemFind(gslbNakedType, findParameter, gslbView),

			// create
			r.DefineOperationCommonServiceItemCreate(gslbNakedType, gslbCreateParam, gslbView),

			// read
			r.DefineOperationCommonServiceItemRead(gslbNakedType, gslbView),

			// update
			r.DefineOperationCommonServiceItemUpdate(gslbNakedType, gslbUpdateParam, gslbView),

			// delete
			r.DefineOperationDelete(),
		}
	},
}

var (
	gslbNakedType = meta.Static(naked.GSLB{})

	gslbView = &schema.Model{
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
