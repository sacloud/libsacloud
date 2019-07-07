package schema

import (
	"fmt"

	"github.com/sacloud/libsacloud/v2/internal/schema/meta"
)

// ResultType Operationからの戻り値の型情報
type ResultType struct {
	resourceName string
	operation    *Operation
	results      Results
}

// Type モデルの型情報
func (r *ResultType) Type() meta.Type {
	return r
}

// GoType 型名
func (r *ResultType) GoType() string {
	return fmt.Sprintf("%s%sResult", r.resourceName, r.operation.Name)
}

// GoPkg パッケージ名
func (r *ResultType) GoPkg() string {
	if IsOutOfSacloudPackage {
		return "sacloud"
	}
	return ""
}

// GoImportPath インポートパス
func (r *ResultType) GoImportPath() string {
	if IsOutOfSacloudPackage {
		return "github.com/sacloud/libsacloud/v2/sacloud"
	}
	return ""
}

// GoTypeSourceCode ソースコードでの型表現
func (r *ResultType) GoTypeSourceCode() string {
	name := r.GoType()
	prefix := ""
	if IsOutOfSacloudPackage {
		prefix = "sacloud."
	}
	return fmt.Sprintf("*%s%s", prefix, name)
}

// ZeroInitializeSourceCode 型に応じたzero値での初期化コード
func (r *ResultType) ZeroInitializeSourceCode() string {
	name := r.GoType()
	if IsOutOfSacloudPackage {
		name = "sacloud." + name
	}
	return fmt.Sprintf("&%s{}", name)
}

// ZeroValueSourceCode 型に応じたzero値コード
func (r *ResultType) ZeroValueSourceCode() string {
	return "nil"
}
