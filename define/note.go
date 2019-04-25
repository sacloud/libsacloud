package define

import (
	"github.com/sacloud/libsacloud-v2/sacloud"
	"github.com/sacloud/libsacloud-v2/schema"
)

func init() {
	resourceName := "Note"
	pathName := "note"

	Resources = append(Resources, &schema.Resource{
		Name:       resourceName,
		PathName:   pathName,
		PathSuffix: schema.CloudAPISuffix,
		Operations: []*schema.Operation{
			schema.CreateOperation(&schema.CreateOperationParam{
				FieldName:       resourceName,
				ParameterStruct: &sacloud.NoteCreateRequest{},
				ResponseStruct:  &sacloud.NoteCommonResponse{},
			}),
			schema.ReadOperation(&schema.ReadOperationParam{
				FieldName:      resourceName,
				ResponseStruct: &sacloud.NoteCommonResponse{},
			}),
			schema.UpdateOperation(&schema.UpdateOperationParam{
				FieldName:       resourceName,
				ParameterStruct: &sacloud.NoteUpdateRequest{},
				ResponseStruct:  &sacloud.NoteCommonResponse{},
			}),
			schema.DeleteOperation(),
		},
	})
}
