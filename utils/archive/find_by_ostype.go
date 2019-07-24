package archive

import (
	"context"
	"fmt"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/ostype"
)

// Finder アーカイブ検索インターフェース
type Finder interface {
	Find(ctx context.Context, zone string, conditions *sacloud.FindCondition) (*sacloud.ArchiveFindResult, error)
}

// FindByOSType OS種別ごとの最新安定板のアーカイブを取得
func FindByOSType(ctx context.Context, api Finder, zone string, os ostype.ArchiveOSType) (*sacloud.Archive, error) {

	filter, ok := ostype.ArchiveCriteria[os]
	if !ok {
		return nil, fmt.Errorf("unsupported ostype.ArchiveOSType: %v", os)
	}

	searched, err := api.Find(ctx, zone, &sacloud.FindCondition{Filter: filter})
	if err != nil {
		return nil, err
	}
	if searched.Count == 0 {
		return nil, fmt.Errorf("archive not found with ostype.ArchiveOSType: %v", os)
	}
	return searched.Archives[0], nil
}
