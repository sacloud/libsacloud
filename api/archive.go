package api

import (
	"fmt"
	"github.com/yamamoto-febc/libsacloud/sacloud"
	"strings"
	"time"
)

type ArchiveAPI struct {
	*baseAPI
	FindFuncMapPerOSType map[sacloud.ArchiveOSTypes]func() (*sacloud.Archive, error)
}

var (
	ArchiveLatestStableCentOSTags   = []string{"current-stable", "distro-centos"}
	ArchiveLatestStableUbuntuTags   = []string{"current-stable", "distro-ubuntu"}
	ArchiveLatestStableDebianTags   = []string{"current-stable", "distro-debian"}
	ArchiveLatestStableVyOSTags     = []string{"current-stable", "distro-vyos"}
	ArchiveLatestStableCoreOSTags   = []string{"current-stable", "distro-coreos"}
	ArchiveLatestStableKusanagiTags = []string{"current-stable", "pkg-kusanagi"}
	//ArchiveLatestStableSiteGuardTags = []string{"current-stable", "pkg-siteguard"} //tk1aではcurrent-stableタグが付いていないため絞り込めない
)

func NewArchiveAPI(client *Client) *ArchiveAPI {
	api := &ArchiveAPI{
		baseAPI: &baseAPI{
			client: client,
			FuncGetResourceURL: func() string {
				return "archive"
			},
		},
	}

	api.FindFuncMapPerOSType = map[sacloud.ArchiveOSTypes]func() (*sacloud.Archive, error){
		sacloud.CentOS:   api.FindLatestStableCentOS,
		sacloud.Ubuntu:   api.FindLatestStableUbuntu,
		sacloud.Debian:   api.FindLatestStableDebian,
		sacloud.VyOS:     api.FindLatestStableVyOS,
		sacloud.CoreOS:   api.FindLatestStableCoreOS,
		sacloud.Kusanagi: api.FindLatestStableKusanagi,
	}

	return api
}

func (api *ArchiveAPI) OpenFTP(id int64) (*sacloud.FTPServer, error) {
	var (
		method = "PUT"
		uri    = fmt.Sprintf("%s/%d/ftp", api.getResourceURL(), id)
		//body   = map[string]bool{"ChangePassword": reset}
		res = &sacloud.Response{}
	)

	result, err := api.action(method, uri, nil, res)
	if !result || err != nil {
		return nil, err
	}

	return res.FTPServer, nil
}

func (api *ArchiveAPI) CloseFTP(id int64) (bool, error) {
	var (
		method = "DELETE"
		uri    = fmt.Sprintf("%s/%d/ftp", api.getResourceURL(), id)
	)
	return api.modify(method, uri, nil)

}

func (api *ArchiveAPI) SleepWhileCopying(id int64, timeout time.Duration) error {

	current := 0 * time.Second
	interval := 5 * time.Second
	for {
		archive, err := api.Read(id)
		if err != nil {
			return err
		}

		if archive.IsAvailable() {
			return nil
		}
		time.Sleep(interval)
		current += interval

		if timeout > 0 && current > timeout {
			return fmt.Errorf("Timeout: SleepWhileCopying[disk:%d]", id)
		}
	}
}

func (api *ArchiveAPI) CanEditDisk(id int64) (bool, error) {

	archive, err := api.Read(id)
	if err != nil {
		return false, err
	}

	if archive == nil {
		return false, nil
	}

	// BundleInfoがあれば編集不可
	if archive.BundleInfo != nil {
		// Windows
		return false, nil
	}

	// ソースアーカイブ/ソースディスクともに持っていない場合
	if archive.SourceArchive == nil && archive.SourceDisk == nil {
		//ブランクディスクがソース
		return false, nil
	}

	for _, t := range allowDiskEditTags {
		if archive.HasTag(t) {
			// 対応OSインストール済みディスク
			return true, nil
		}
	}

	// ここまできても判定できないならソースに投げる
	if archive.SourceDisk != nil {
		return api.CanEditDisk(archive.SourceDisk.ID)
	}
	return api.client.Archive.CanEditDisk(archive.SourceArchive.ID)

}

func (api *ArchiveAPI) FindLatestStableCentOS() (*sacloud.Archive, error) {
	return api.findByOSTags(ArchiveLatestStableCentOSTags)
}
func (api *ArchiveAPI) FindLatestStableDebian() (*sacloud.Archive, error) {
	return api.findByOSTags(ArchiveLatestStableDebianTags)
}
func (api *ArchiveAPI) FindLatestStableUbuntu() (*sacloud.Archive, error) {
	return api.findByOSTags(ArchiveLatestStableUbuntuTags)
}
func (api *ArchiveAPI) FindLatestStableVyOS() (*sacloud.Archive, error) {
	return api.findByOSTags(ArchiveLatestStableVyOSTags)
}
func (api *ArchiveAPI) FindLatestStableCoreOS() (*sacloud.Archive, error) {
	return api.findByOSTags(ArchiveLatestStableCoreOSTags)
}
func (api *ArchiveAPI) FindLatestStableKusanagi() (*sacloud.Archive, error) {
	return api.findByOSTags(ArchiveLatestStableKusanagiTags)
}
func (api *ArchiveAPI) FindByOSType(os sacloud.ArchiveOSTypes) (*sacloud.Archive, error) {
	if f, ok := api.FindFuncMapPerOSType[os]; ok {
		return f()
	}

	return nil, fmt.Errorf("OSType [%s] is invalid", os)
}

func (api *ArchiveAPI) findByOSTags(tags []string) (*sacloud.Archive, error) {
	res, err := api.Reset().WithTags(tags).Find()
	if err != nil {
		return nil, fmt.Errorf("Archive [%s] error : %s", strings.Join(tags, ","), err)
	}

	if len(res.Archives) == 0 {
		return nil, fmt.Errorf("Archive [%s] Not Found", strings.Join(tags, ","))
	}

	return &res.Archives[0], nil

}
