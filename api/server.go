package api

import (
	"fmt"
	"github.com/yamamoto-febc/libsacloud/sacloud"
	"time"
)

type ServerAPI struct {
	*baseAPI
}

func NewServerAPI(client *Client) *ServerAPI {
	return &ServerAPI{
		&baseAPI{
			client: client,
			FuncGetResourceURL: func() string {
				return "server"
			},
		},
	}
}

func (api *ServerAPI) WithPlan(planID string) *ServerAPI {
	return api.FilterBy("ServerPlan.ID", planID)
}

func (api *ServerAPI) WithStatus(status string) *ServerAPI {
	return api.FilterBy("Instance.Status", status)
}
func (api *ServerAPI) WithStatusUp() *ServerAPI {
	return api.WithStatus("up")
}
func (api *ServerAPI) WithStatusDown() *ServerAPI {
	return api.WithStatus("down")
}

func (api *ServerAPI) WithISOImage(imageID int64) *ServerAPI {
	return api.FilterBy("Instance.CDROM.ID", imageID)
}

func (api *ServerAPI) SortByCPU(reverse bool) *ServerAPI {
	api.sortBy("ServerPlan.CPU", reverse)
	return api
}

func (api *ServerAPI) SortByMemory(reverse bool) *ServerAPI {
	api.sortBy("ServerPlan.MemoryMB", reverse)
	return api
}

func (api *ServerAPI) DeleteWithDisk(id int64, disks []int64) (*sacloud.Server, error) {
	return api.request(func(res *sacloud.Response) error {
		return api.delete(id, map[string]interface{}{"WithDisk": disks}, res)
	})
}

// State get server state
func (api *ServerAPI) State(id int64) (string, error) {
	server, err := api.Read(id)
	if err != nil {
		return "", err
	}
	return server.Availability, nil
}

func (api *ServerAPI) IsUp(id int64) (bool, error) {
	server, err := api.Read(id)
	if err != nil {
		return false, err
	}
	return server.Instance.IsUp(), nil
}

func (api *ServerAPI) IsDown(id int64) (bool, error) {
	server, err := api.Read(id)
	if err != nil {
		return false, err
	}
	return server.Instance.IsDown(), nil
}

// Boot power on
func (api *ServerAPI) Boot(id int64) (bool, error) {
	var (
		method = "PUT"
		uri    = fmt.Sprintf("%s/%d/power", api.getResourceURL(), id)
	)
	return api.modify(method, uri, nil)
}

// Shutdown power off
func (api *ServerAPI) Shutdown(id int64) (bool, error) {
	var (
		method = "DELETE"
		uri    = fmt.Sprintf("%s/%d/power", api.getResourceURL(), id)
	)

	return api.modify(method, uri, nil)
}

// Stop force shutdown
func (api *ServerAPI) Stop(id int64) (bool, error) {
	var (
		method = "DELETE"
		uri    = fmt.Sprintf("%s/%d/power", api.getResourceURL(), id)
	)

	return api.modify(method, uri, map[string]bool{"Force": true})
}

func (api *ServerAPI) RebootForce(id int64) (bool, error) {
	var (
		method = "PUT"
		uri    = fmt.Sprintf("%s/%d/reset", api.getResourceURL(), id)
	)

	return api.modify(method, uri, nil)
}

func (api *ServerAPI) SleepUntilUp(id int64, timeout time.Duration) error {
	current := 0 * time.Second
	interval := 5 * time.Second
	for {

		up, err := api.IsUp(id)
		if err != nil {
			return err
		}

		if up {
			return nil
		}
		time.Sleep(interval)
		current += interval

		if timeout > 0 && current > timeout {
			return fmt.Errorf("Timeout: WaitforAvailable")
		}
	}
}

func (api *ServerAPI) SleepUntilDown(id int64, timeout time.Duration) error {
	current := 0 * time.Second
	interval := 5 * time.Second
	for {

		down, err := api.IsDown(id)
		if err != nil {
			return err
		}

		if down {
			return nil
		}
		time.Sleep(interval)
		current += interval

		if timeout > 0 && current > timeout {
			return fmt.Errorf("Timeout: WaitforAvailable")
		}
	}
}

func (api *ServerAPI) ChangePlan(serverID int64, planID string) (*sacloud.Server, error) {
	var (
		method = "PUT"
		uri    = fmt.Sprintf("%s/%d/to/plan/%s", api.getResourceURL(), serverID, planID)
	)

	return api.request(func(res *sacloud.Response) error {
		return api.baseAPI.request(method, uri, nil, res)
	})
}

func (api *ServerAPI) FindDisk(serverID int64) ([]sacloud.Disk, error) {
	server, err := api.Read(serverID)
	if err != nil {
		return nil, err
	}
	return server.Disks, nil
}

func (api *ServerAPI) InsertCDROM(serverID int64, cdromID int64) (bool, error) {
	var (
		method = "PUT"
		uri    = fmt.Sprintf("%s/%d/cdrom", api.getResourceURL(), serverID)
	)

	req := &sacloud.Request{
		SakuraCloudResources: sacloud.SakuraCloudResources{
			CDROM: &sacloud.CDROM{Resource: &sacloud.Resource{ID: cdromID}},
		},
	}

	return api.modify(method, uri, req)
}

func (api *ServerAPI) EjectCDROM(serverID int64, cdromID int64) (bool, error) {
	var (
		method = "DELETE"
		uri    = fmt.Sprintf("%s/%d/cdrom", api.getResourceURL(), serverID)
	)

	req := &sacloud.Request{
		SakuraCloudResources: sacloud.SakuraCloudResources{
			CDROM: &sacloud.CDROM{Resource: &sacloud.Resource{ID: cdromID}},
		},
	}

	return api.modify(method, uri, req)
}

func (api *ServerAPI) NewKeyboardRequest() *sacloud.KeyboardRequest {
	return &sacloud.KeyboardRequest{}
}

func (api *ServerAPI) SendKey(serverID int64, body *sacloud.KeyboardRequest) (bool, error) {
	var (
		method = "PUT"
		uri    = fmt.Sprintf("%s/%d/keyboard", api.getResourceURL(), serverID)
	)

	return api.modify(method, uri, body)
}

func (api *ServerAPI) NewMouseRequest() *sacloud.MouseRequest {
	return &sacloud.MouseRequest{
		Buttons: &sacloud.MouseRequestButtons{},
	}
}

func (api *ServerAPI) SendMouse(serverID int64, mouseIndex string, body *sacloud.MouseRequest) (bool, error) {
	var (
		method = "PUT"
		uri    = fmt.Sprintf("%s/%d/mouse/%s", api.getResourceURL(), serverID, mouseIndex)
	)

	return api.modify(method, uri, body)
}

func (api *ServerAPI) NewVNCSnapshotRequest() *sacloud.VNCSnapshotRequest {
	return &sacloud.VNCSnapshotRequest{}
}

func (api *ServerAPI) GetVNCProxy(serverID int64) (*sacloud.VNCProxyResponse, error) {
	var (
		method = "GET"
		uri    = fmt.Sprintf("%s/%d/vnc/proxy", api.getResourceURL(), serverID)
		res    = &sacloud.VNCProxyResponse{}
	)
	err := api.baseAPI.request(method, uri, nil, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (api *ServerAPI) GetVNCSize(serverID int64) (*sacloud.VNCSizeResponse, error) {
	var (
		method = "GET"
		uri    = fmt.Sprintf("%s/%d/vnc/size", api.getResourceURL(), serverID)
		res    = &sacloud.VNCSizeResponse{}
	)
	err := api.baseAPI.request(method, uri, nil, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (api *ServerAPI) GetVNCSnapshot(serverID int64, body *sacloud.VNCSnapshotRequest) (*sacloud.VNCSnapshotResponse, error) {
	var (
		method = "GET"
		uri    = fmt.Sprintf("%s/%d/vnc/snapshot", api.getResourceURL(), serverID)
		res    = &sacloud.VNCSnapshotResponse{}
	)
	err := api.baseAPI.request(method, uri, body, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (api *ServerAPI) Monitor(id int64, body *sacloud.ResourceMonitorRequest) (*sacloud.MonitorValues, error) {
	return api.baseAPI.monitor(id, body)
}
