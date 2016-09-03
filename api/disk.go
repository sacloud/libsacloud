package api

import (
	"fmt"
	"github.com/yamamoto-febc/libsacloud/sacloud"
	"time"
)

type DiskAPI struct {
	*baseAPI
}

func NewDiskAPI(client *Client) *DiskAPI {
	return &DiskAPI{
		&baseAPI{
			client: client,
			FuncGetResourceURL: func() string {
				return "disk"
			},
		},
	}
}

func (api *DiskAPI) SortByConnectionOrder(reverse bool) *DiskAPI {
	api.sortBy("ConnectionOrder", reverse)
	return api
}

func (api *DiskAPI) WithServerID(id int64) *DiskAPI {
	api.FilterBy("Server.ID", id)
	return api
}

func (api *DiskAPI) Create(value *sacloud.Disk) (*sacloud.Disk, error) {
	//HACK: さくらのAPI側仕様: 戻り値:Successがbool値へ変換できないため文字列で受ける
	type diskResponse struct {
		*sacloud.Response
		Success string `json:",omitempty"`
	}
	res := &diskResponse{}
	err := api.create(api.createRequest(value), res)
	if err != nil {
		return nil, err
	}
	return res.Disk, nil
}

func (api *DiskAPI) NewCondig() *sacloud.DiskEditValue {
	return &sacloud.DiskEditValue{}
}

func (api *DiskAPI) Config(id int64, disk *sacloud.DiskEditValue) (bool, error) {
	var (
		method = "PUT"
		uri    = fmt.Sprintf("%s/%d/config", api.getResourceURL(), id)
	)

	return api.modify(method, uri, disk)
}

func (api *DiskAPI) install(id int64, body *sacloud.Disk) (bool, error) {
	var (
		method = "PUT"
		uri    = fmt.Sprintf("%s/%d/install", api.getResourceURL(), id)
	)

	return api.modify(method, uri, body)
}

func (api *DiskAPI) ReinstallFromBlank(id int64, sizeMB int) (bool, error) {
	var body = &sacloud.Disk{
		SizeMB: sizeMB,
	}
	return api.install(id, body)
}

func (api *DiskAPI) ReinstallFromArchive(id int64, archiveID int64) (bool, error) {
	var body = &sacloud.Disk{
		SourceArchive: &sacloud.Archive{},
	}
	body.SourceArchive.ID = id
	return api.install(id, body)
}

func (api *DiskAPI) ReinstallFromDisk(id int64, diskID int64) (bool, error) {
	var body = &sacloud.Disk{
		SourceDisk: &sacloud.Disk{
			Resource: &sacloud.Resource{ID: diskID},
		},
	}
	return api.install(id, body)
}

func (api *DiskAPI) ToBlank(diskID int64) (bool, error) {
	var (
		method = "PUT"
		uri    = fmt.Sprintf("%s/%d/to/blank", api.getResourceURL(), diskID)
	)
	return api.modify(method, uri, nil)
}

func (api *DiskAPI) DisconnectFromServer(diskID int64) (bool, error) {
	var (
		method = "DELETE"
		uri    = fmt.Sprintf("%s/%d/to/server", api.getResourceURL(), diskID)
	)
	return api.modify(method, uri, nil)
}

func (api *DiskAPI) ConnectToServer(diskID int64, serverID int64) (bool, error) {
	var (
		method = "PUT"
		uri    = fmt.Sprintf("%s/%d/to/server/%d", api.getResourceURL(), diskID, serverID)
	)
	return api.modify(method, uri, nil)
}

// State get disk state
func (api *DiskAPI) State(diskID int64) (bool, error) {
	disk, err := api.Read(diskID)
	if err != nil {
		return false, err
	}
	return disk.IsAvailable(), nil
}

// SleepWhileCopying wait until became to available
func (api *DiskAPI) SleepWhileCopying(diskID int64, timeout time.Duration) error {
	current := 0 * time.Second
	interval := 5 * time.Second
	for {
		available, err := api.State(diskID)
		if err != nil {
			return err
		}

		if available {
			return nil
		}
		time.Sleep(interval)
		current += interval

		if timeout > 0 && current > timeout {
			return fmt.Errorf("Timeout: WaitforAvailable")
		}
	}
}

func (api *DiskAPI) Monitor(id int64, body *sacloud.ResourceMonitorRequest) (*sacloud.MonitorValues, error) {
	return api.baseAPI.monitor(id, body)
}
