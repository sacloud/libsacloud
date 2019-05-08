// Package define .
//go:generate go run ../tools/gen-api-models/main.go
//go:generate go run ../tools/gen-api-interfaces/main.go
//go:generate go run ../tools/gen-api-envelope/main.go
//go:generate go run ../tools/gen-api-op/main.go
package define

import "github.com/sacloud/libsacloud-v2/internal/schema"

// Resources APIでの操作対象リソースの定義
var Resources schema.Resources
