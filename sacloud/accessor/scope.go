package accessor

import "github.com/sacloud/libsacloud/v2/sacloud/types"

// Scope スコープ
type Scope interface {
	GetScope() types.EScope
	SetScope(scope types.EScope)
}
