package archive

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/ostype"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
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

// SourceInfoReader アーカイブソースを取得するためのインターフェース
type SourceInfoReader struct {
	ArchiveReader SourceArchiveReader
	DiskReader    SourceDiskReader
}

// SourceArchiveReader アーカイブ参照インターフェース
type SourceArchiveReader interface {
	Read(ctx context.Context, zone string, id types.ID) (*sacloud.Archive, error)
}

// SourceDiskReader ディスク参照インターフェース
type SourceDiskReader interface {
	Read(ctx context.Context, zone string, id types.ID) (*sacloud.Disk, error)
}

var (
	// allowDiskEditTags ディスクの編集可否判定に用いるタグ
	allowDiskEditTags = []string{
		"os-unix",
		"os-linux",
	}

	// bundleInfoWindowsHostClass ディスクの編集可否判定に用いる、BundleInfoでのWindows判定文字列
	bundleInfoWindowsHostClass = "ms_windows"
)

func isSophosUTM(archive *sacloud.Archive) bool {
	// SophosUTMであれば編集不可
	if archive.BundleInfo != nil && strings.Contains(strings.ToLower(archive.BundleInfo.ServiceClass), "sophosutm") {
		return true
	}
	return false
}

// CanEditDisk ディスクの修正が可能か判定
func CanEditDisk(ctx context.Context, zone string, reader *SourceInfoReader, id types.ID) (bool, error) {

	disk, err := reader.DiskReader.Read(ctx, zone, id)
	if err != nil {
		if !sacloud.IsNotFoundError(err) {
			return false, err
		}
	}
	if disk != nil {
		// 無限ループ予防
		if disk.ID == disk.SourceDiskID || disk.ID == disk.SourceArchiveID {
			return false, errors.New("invalid state: disk has invalid ID or SourceDiskID or SourceArchiveID")
		}

		if disk.SourceDiskID.IsEmpty() && disk.SourceArchiveID.IsEmpty() {
			return false, nil
		}
		if !disk.SourceDiskID.IsEmpty() {
			return CanEditDisk(ctx, zone, reader, disk.SourceDiskID)
		}
		if !disk.SourceArchiveID.IsEmpty() {
			id = disk.SourceArchiveID
		}
	}

	archive, err := reader.ArchiveReader.Read(ctx, zone, id)
	if err != nil {
		return false, err
	}

	// 無限ループ予防
	if archive.ID == archive.SourceDiskID || archive.ID == archive.SourceArchiveID {
		return false, errors.New("invalid state: archive has invalid ID or SourceDiskID or SourceArchiveID")
	}

	// BundleInfoがあれば編集不可
	if archive.BundleInfo != nil && archive.BundleInfo.HostClass == bundleInfoWindowsHostClass {
		// Windows
		return false, nil
	}

	// SophosUTMであれば編集不可
	if archive.HasTag("pkg-sophosutm") || isSophosUTM(archive) {
		return false, nil
	}
	// OPNsenseであれば編集不可
	if archive.HasTag("distro-opnsense") {
		return false, nil
	}
	// Netwiser VEであれば編集不可
	if archive.HasTag("pkg-netwiserve") {
		return false, nil
	}

	for _, t := range allowDiskEditTags {
		if archive.HasTag(t) {
			// 対応OSインストール済みディスク
			return true, nil
		}
	}

	// ここまできても判定できないならソースに投げる
	if !archive.SourceDiskID.IsEmpty() && archive.SourceDiskAvailability != types.Availabilities.Discontinued {
		return CanEditDisk(ctx, zone, reader, archive.SourceDiskID)
	}
	if !archive.SourceArchiveID.IsEmpty() && archive.SourceArchiveAvailability != types.Availabilities.Discontinued {
		return CanEditDisk(ctx, zone, reader, archive.SourceArchiveID)
	}
	return false, nil

}
