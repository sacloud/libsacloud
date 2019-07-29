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
			resourceName == "IPAddress",
			strings.HasSuffix(resourceName, "Info"):
			return resourceName
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

// RequestParameterName 任意の操作に渡すパラメータの名称
func RequestParameterName(resourceName, funcName string) string {
	return fmt.Sprintf("%s%sRequest", resourceName, funcName)
}
