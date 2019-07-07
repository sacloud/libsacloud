package names

import (
	"strings"

	"github.com/sacloud/libsacloud/v2/internal/schema"
)

// ResourceFieldName リソース名がペイロードなどで利用される場合のフィールド名、コード生成時に利用される
func ResourceFieldName(resourceName string, form schema.PayloadForm) string {
	switch {
	case form.IsSingular():
		return resourceName
	case form.IsPlural():
		// TODO とりあえずワードで例外指定
		switch {
		case resourceName == "NFS":
			return resourceName
		case strings.HasSuffix(resourceName, "ch"):
			return resourceName + "es"
		default:
			return resourceName + "s"
		}
	default:
		return ""
	}
}
