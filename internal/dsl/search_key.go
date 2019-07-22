package dsl

// SearchKeyDesc 検索条件となるキーの情報
type SearchKeyDesc interface {
	// KeyName キー名 コード生成時にフィールド名となる
	KeyName() string
	// SourceFieldName APIリクエスト時に指定するキーとなるフィールド名
	SourceFieldName() string
}

// SearchKey 検索条件となるキーの情報
type SearchKey struct {
	// Name キー名
	Name string
	// SourceField APIリクエスト時のフィールド名
	SourceField string
}

// KeyName キー名 コード生成時にフィールド名となる
func (k *SearchKey) KeyName() string {
	return k.Name
}

// SourceFieldName フィールド名 SourceFieldが空の場合はNameを返す
func (k *SearchKey) SourceFieldName() string {
	if k.SourceField != "" {
		return k.SourceField
	}
	return k.Name
}

// SearchKeyDef SearchKeyを作成
func SearchKeyDef(name string, sourceField string) *SearchKey {
	return &SearchKey{Name: name, SourceField: sourceField}
}
