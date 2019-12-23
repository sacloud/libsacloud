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

package names

import (
	"fmt"
	"strings"

	"github.com/sacloud/libsacloud/v2/internal/dsl"
)

// ResourceFieldName リソース名がペイロードなどで利用される場合のフィールド名、コード生成時に利用される
func ResourceFieldName(resourceName string, form dsl.PayloadForm) string {
	switch {
	case form.IsSingular():
		return resourceName
	case form.IsPlural():
		switch {
		case
			resourceName == "NFS",
			resourceName == "DNS",
			resourceName == "Internet",
			resourceName == "IPAddress",
			strings.HasSuffix(resourceName, "Info"):
			return resourceName
		case resourceName == "ContainerRegistry":
			return "ContainerRegistries"
		case
			strings.HasSuffix(resourceName, "ch"),
			strings.HasSuffix(resourceName, "ss"):
			return resourceName + "es"
		default:
			return resourceName + "s"
		}
	default:
		return ""
	}
}

// CreateParameterName Create操作に渡すパラメータの名称
func CreateParameterName(resourceName string) string {
	return RequestParameterName(resourceName, "Create")
}

// UpdateParameterName Update操作に渡すパラメータの名称
func UpdateParameterName(resourceName string) string {
	return RequestParameterName(resourceName, "Update")
}

// UpdateSettingsParameterName UpdateSettings操作に渡すパラメータの名称
func UpdateSettingsParameterName(resourceName string) string {
	return RequestParameterName(resourceName, "UpdateSettings")
}

// RequestParameterName 任意の操作に渡すパラメータの名称
func RequestParameterName(resourceName, funcName string) string {
	return fmt.Sprintf("%s%sRequest", resourceName, funcName)
}
