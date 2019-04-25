package schema

import "strings"

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

// Resources []*Resourceのエイリアス
type Resources []*Resource

// ImportStatements コード生成時に利用するimport文を生成する
func (r Resources) ImportStatements() []string {
	var ss []string

	for _, re := range r {
		ss = append(ss, re.ImportStatements()...)
	}

	return uniqStrings(ss)
}

// Resource APIで操作する対象のリソース
type Resource struct {
	Name       string       // リソース名 e.g.: Server
	PathName   string       // リソースのパス名 APIのURLで利用される e.g.: server 省略した場合はNameを小文字にしたものとなる
	PathSuffix string       // APIのURLで利用されるプレフィックス e.g.: api/cloud/1.1
	Operations []*Operation // このリソースに対する操作
}

// LowerName 小文字にしたResourceの名前、生成したコードの保存先ファイル名に利用される
func (r *Resource) LowerName() string {
	return strings.ReplaceAll(strings.ToLower(r.Name), "-", "_")
}

// ImportStatements コード生成時に利用するimport文を生成する
func (r *Resource) ImportStatements() []string {
	var ss []string
	for _, o := range r.Operations {
		ss = append(ss, o.ImportStatements()...)
	}

	return uniqStrings(ss)
}
