// Copyright 2016-2019 The Libsacloud Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
