package api

import (
	"fmt"
	"github.com/yamamoto-febc/libsacloud/sacloud"
)

type SwitchAPI struct {
	*baseAPI
}

func NewSwitchAPI(client *Client) *SwitchAPI {
	return &SwitchAPI{
		&baseAPI{
			client: client,
			FuncGetResourceURL: func() string {
				return "switch"
			},
		},
	}
}

func (api *SwitchAPI) DisconnectFromBridge(switchID int64) (bool, error) {
	var (
		method = "DELETE"
		uri    = fmt.Sprintf("%s/%d/to/bridge", api.getResourceURL(), switchID)
	)
	return api.modify(method, uri, nil)
}

func (api *SwitchAPI) ConnectToBridge(switchID int64, bridgeID int64) (bool, error) {
	var (
		method = "PUT"
		uri    = fmt.Sprintf("%s/%d/to/bridge/%d", api.getResourceURL(), switchID, bridgeID)
	)
	return api.modify(method, uri, nil)
}

func (api *SwitchAPI) GetServers(switchID int64) ([]sacloud.Server, error) {
	var (
		method = "GET"
		uri    = fmt.Sprintf("%s/%d/server", api.getResourceURL(), switchID)
		res    = &sacloud.SearchResponse{}
	)
	err := api.baseAPI.request(method, uri, nil, res)
	if err != nil {
		return nil, err
	}
	return res.Servers, nil
}
