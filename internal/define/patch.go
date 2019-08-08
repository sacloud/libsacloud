package define

import (
	"strings"

	"github.com/sacloud/libsacloud/v2/internal/dsl"
	"github.com/sacloud/libsacloud/v2/internal/dsl/meta"
)

func patchModel(updateModel *dsl.Model) *dsl.Model {
	pm := &dsl.Model{
		Name:        strings.Replace(updateModel.Name, "Update", "Patch", -1),
		ConstFields: updateModel.ConstFields,
		Methods:     updateModel.Methods,
		NakedType:   updateModel.NakedType,
		IsArray:     updateModel.IsArray,
	}

	var fields []*dsl.FieldDesc
	for _, f := range updateModel.Fields {

		fields = append(fields, f)
		if f.IsNeedPatchEmpty() {
			fields = append(fields, &dsl.FieldDesc{
				Name: "PatchEmptyTo" + f.Name,
				Type: meta.TypeFlag,
			})
		}
	}
	pm.Fields = fields
	return pm
}
