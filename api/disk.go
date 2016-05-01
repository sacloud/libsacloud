package api

import (
	"fmt"
	sakura "github.com/yamamoto-febc/libsacloud/resources"
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

func (api *DiskAPI) request(f func(*sakura.Response) error) (*sakura.Disk, error) {
	res := &sakura.Response{}
	err := f(res)
	if err != nil {
		return nil, err
	}
	return res.Disk, nil
}

func (api *DiskAPI) createRequest(value *sakura.Disk) *sakura.Request {
	return &sakura.Request{Disk: value}
}

func (api *DiskAPI) Create(value *sakura.Disk) (*sakura.Disk, error) {
	//HACK: さくらのAPI側仕様: 戻り値:Successがbool値へ変換できないため文字列で受ける
	type diskResponse struct {
		*sakura.Response
		Success string `json:",omitempty"`
	}
	res := &diskResponse{}
	err := api.create(api.createRequest(value), res)
	if err != nil {
		return nil, err
	}
	return res.Disk, nil
}

func (api *DiskAPI) Read(id string) (*sakura.Disk, error) {
	return api.request(func(res *sakura.Response) error {
		return api.read(id, nil, res)
	})
}

func (api *DiskAPI) Update(id string, value *sakura.Disk) (*sakura.Disk, error) {
	return api.request(func(res *sakura.Response) error {
		return api.update(id, api.createRequest(value), res)
	})
}

func (api *DiskAPI) Delete(id string) (*sakura.Disk, error) {
	return api.request(func(res *sakura.Response) error {
		return api.delete(id, nil, res)
	})
}

func (api *DiskAPI) Config(id string, disk *sakura.DiskEditValue) (bool, error) {
	var (
		method = "PUT"
		uri    = fmt.Sprintf("%s/%s/config", api.getResourceURL(), id)
	)

	res := &sakura.ResultFlagValue{}
	err := api.baseAPI.request(method, uri, nil, res)
	if err != nil {
		return false, err
	}
	return res.IsOk, nil
}

func (api *DiskAPI) install(id string, body *sakura.Disk) (bool, error) {
	var (
		method = "PUT"
		uri    = fmt.Sprintf("%s/%s/install", api.getResourceURL(), id)
	)

	res := &sakura.ResultFlagValue{}
	err := api.baseAPI.request(method, uri, body, res)
	if err != nil {
		return false, err
	}
	return res.IsOk, nil

}

func (api *DiskAPI) ReinstallFromBlank(id string, sizeMB int) (bool, error) {
	var body = &sakura.Disk{
		SizeMB: sizeMB,
	}
	return api.install(id, body)
}

func (api *DiskAPI) ReinstallFromArchive(id string, archiveID string) (bool, error) {
	var body = &sakura.Disk{
		SourceArchive: &sakura.Archive{
			Resource: &sakura.Resource{ID: archiveID},
		},
	}
	return api.install(id, body)
}

func (api *DiskAPI) ReinstallFromDisk(id string, diskID string) (bool, error) {
	var body = &sakura.Disk{
		SourceDisk: &sakura.Disk{
			Resource: &sakura.Resource{ID: diskID},
		},
	}
	return api.install(id, body)
}

func (api *DiskAPI) ToBlank(diskID string) (bool, error) {
	var (
		method = "PUT"
		uri    = fmt.Sprintf("%s/%s/to/blank", api.getResourceURL(), diskID)
	)
	res := &sakura.ResultFlagValue{}
	err := api.baseAPI.request(method, uri, nil, res)
	if err != nil {
		return false, err
	}
	return res.IsOk, nil
}

func (api *DiskAPI) DisconnectFromServer(diskID string) (bool, error) {
	var (
		method = "DELETE"
		uri    = fmt.Sprintf("%s/%s/to/server", api.getResourceURL(), diskID)
	)
	res := &sakura.ResultFlagValue{}
	err := api.baseAPI.request(method, uri, nil, res)
	if err != nil {
		return false, err
	}
	return res.IsOk, nil
}

func (api *DiskAPI) ConnectToServer(diskID string, serverID string) (bool, error) {
	var (
		method = "PUT"
		uri    = fmt.Sprintf("%s/%s/to/server/%s", api.getResourceURL(), diskID, serverID)
	)
	res := &sakura.ResultFlagValue{}
	err := api.baseAPI.request(method, uri, nil, res)
	if err != nil {
		return false, err
	}
	return res.IsOk, nil
}

// State get disk state
func (api *DiskAPI) State(diskID string) (bool, error) {
	disk, err := api.Read(diskID)
	if err != nil {
		return false, err
	}
	return disk.IsAvailable(), nil
}

// WaitForAvailable wait until became to available
func (api *DiskAPI) WaitForAvailable(diskID string, timeout time.Duration) error {
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

		if current > timeout {
			return fmt.Errorf("Timeout: WaitforAvailable")
		}
	}
}
