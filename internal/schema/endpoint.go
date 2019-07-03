package schema

import "fmt"

const (
	// DefaultPathFormat デフォルトのパスフォーマット
	DefaultPathFormat = "{{.rootURL}}/{{.zone}}/{{.pathSuffix}}/{{.pathName}}"
	// CloudAPISuffix IaaSリソースでのAPIサフィックス
	CloudAPISuffix = "api/cloud/1.1"
	// BillingAPISuffix 課金関連でのAPIサフィックス
	BillingAPISuffix = "api/system/1.0"
	// WebAccelAPISuffix ウェブアクセラレータ関連でのAPIサフィックス
	WebAccelAPISuffix = "api/webaccel/1.0"
)

var (
	// DefaultPathFormatWithID デフォルトのパス+IDのパスフォーマット
	DefaultPathFormatWithID = fmt.Sprintf("%s/{{.%s}}", DefaultPathFormat, ArgumentID.ArgName())
)

// IDAndSuffixPathFormat デフォルトのパス+ID+指定のサフィックスのパスフォーマット
func IDAndSuffixPathFormat(suffix string) string {
	return fmt.Sprintf(DefaultPathFormatWithID+"/%s", suffix)
}
