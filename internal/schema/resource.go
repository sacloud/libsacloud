package schema

// Resources []*Resourceのエイリアス
type Resources []*Resource

// ImportStatements コード生成時に利用するimport文を生成する
func (r Resources) ImportStatements(additionalImports ...string) []string {
	ss := wrapByDoubleQuote(additionalImports...)

	for _, re := range r {
		ss = append(ss, re.ImportStatements()...)
	}

	return uniqStrings(ss)
}

// ImportStatementsForModelDef Resources配下に含まれる全てのモデルのフィールドを含めたimport文を生成する
func (r Resources) ImportStatementsForModelDef(additionalImports ...string) []string {
	ss := wrapByDoubleQuote(additionalImports...)
	for _, m := range r.Models() {
		ss = append(ss, m.ImportStatementsForModelDef()...)
	}
	return uniqStrings(ss)
}

// Define リソースの定義
func (r *Resources) Define(rs *Resource) {
	if *r == nil {
		rr := Resources{}
		*r = rr
	}
	*r = append(*r, rs)
}

// Models モデル一覧を取得
func (r Resources) Models() Models {
	ms := Models{}
	for _, res := range r {
		for _, o := range res.Operations {
			ms = append(ms, o.Models()...)
		}
	}
	return ms.UniqByName()
}

// OperationsDefineFunc リソースに対するオペレーション定義用Func
type OperationsDefineFunc func(r *Resource) []*Operation

// Resource APIで操作する対象のリソース
type Resource struct {
	Name       string       // リソース名 e.g.: Server
	PathName   string       // リソースのパス名 APIのURLで利用される e.g.: server 省略した場合はNameを小文字にしたものとなる
	PathSuffix string       // APIのURLで利用されるプレフィックス e.g.: api/cloud/1.1
	IsGlobal   bool         // 全ゾーンで共通リソース(グローバルリソース)
	Operations []*Operation // このリソースに対する操作、OperationsDefineFuncが設定されている場合はそちらを呼び出して設定される
}

// GetPathName リソースのパス名 APIのエンドポイントURLの算出で利用される 例: server
//
// 省略した場合はNameをスネークケース(小文字+アンダーバー)に変換したものが利用される
func (r *Resource) GetPathName() string {
	if r.PathName != "" {
		return r.PathName
	}
	return toSnakeCaseName(r.Name)
}

// GetPathSuffix PathSuffixの取得
func (r *Resource) GetPathSuffix() string {
	if r.PathSuffix != "" {
		return r.PathSuffix
	}
	return CloudAPISuffix
}

// FileSafeName スネークケースにしたResourceの名前、コード生成時の保存先ファイル名に利用される
func (r *Resource) FileSafeName() string {
	return toSnakeCaseName(r.Name)
}

// TypeName 型名を返す、コード生成時の型定義などで利用される
func (r *Resource) TypeName() string {
	return r.Name
}

// ImportStatements コード生成時に利用するimport文を生成する
func (r *Resource) ImportStatements(additionalImports ...string) []string {
	ss := wrapByDoubleQuote(additionalImports...)
	for _, o := range r.Operations {
		ss = append(ss, o.ImportStatements()...)
	}

	return uniqStrings(ss)
}
