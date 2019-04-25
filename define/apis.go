// Package define .
//go:generate go run ../tools/internal/gen-api-interfaces/main.go
//go:generate go run ../tools/internal/gen-api-structures/main.go
//go:generate go run ../tools/internal/gen-api-op/main.go
package define

import "github.com/sacloud/libsacloud-v2/schema"

// Resources APIでの操作対象リソースの定義
var Resources schema.Resources
