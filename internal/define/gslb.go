package define

import (
	"github.com/sacloud/libsacloud-v2/internal/schema"
	"github.com/sacloud/libsacloud-v2/internal/schema/meta"
	"github.com/sacloud/libsacloud-v2/sacloud/naked"
)

func init() {
	nakedType := meta.Static(naked.GSLB{})

	gslb := &schema.Model{
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

	createParam := &schema.Model{
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

	updateParam := &schema.Model{
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
	gslbAPI := &schema.Resource{
		Name:       "GSLB",
		PathName:   "commonserviceitem",
		PathSuffix: schema.CloudAPISuffix,
		IsGlobal:   true,
	}
	gslbAPI.Operations = []*schema.Operation{
		// find
		gslbAPI.DefineOperationCommonServiceItemFind(nakedType, findParameter, gslb),

		// create
		gslbAPI.DefineOperationCommonServiceItemCreate(nakedType, createParam, gslb),

		// read
		gslbAPI.DefineOperationCommonServiceItemRead(nakedType, gslb),

		// update
		gslbAPI.DefineOperationCommonServiceItemUpdate(nakedType, updateParam, gslb),

		// delete
		gslbAPI.DefineOperationDelete(),
	}
	Resources.Def(gslbAPI)
}
