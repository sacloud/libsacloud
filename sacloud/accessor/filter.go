package accessor

import (
	"github.com/sacloud/libsacloud/v2/sacloud/search"
)

// Filter 検索フィルタ
type Filter interface {
	GetFilter() search.Filter
	SetFilter(v search.Filter)
}

// ClearFilter 設定されたフィルタをクリアします
func ClearFilter(f Filter) {
	f.SetFilter(search.Filter{})
}
