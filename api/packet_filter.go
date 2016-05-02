package api

import (
	"fmt"
	"github.com/yamamoto-febc/libsacloud/sacloud"
	"regexp"
)

type PacketFilterAPI struct {
	*baseAPI
}

func NewPacketFilterAPI(client *Client) *PacketFilterAPI {
	return &PacketFilterAPI{
		&baseAPI{
			client: client,
			FuncGetResourceURL: func() string {
				return "packetfilter"
			},
		},
	}
}

func (api *PacketFilterAPI) request(f func(*sacloud.Response) error) (*sacloud.PacketFilter, error) {
	res := &sacloud.Response{}
	err := f(res)
	if err != nil {
		return nil, err
	}
	return res.PacketFilter, nil
}

func (api *PacketFilterAPI) createRequest(value *sacloud.PacketFilter) *sacloud.Request {
	return &sacloud.Request{PacketFilter: value}
}

func (api *PacketFilterAPI) Create(value *sacloud.PacketFilter) (*sacloud.PacketFilter, error) {
	return api.request(func(res *sacloud.Response) error {
		return api.create(api.createRequest(value), res)
	})
}

func (api *PacketFilterAPI) Read(id string) (*sacloud.PacketFilter, error) {
	return api.request(func(res *sacloud.Response) error {
		return api.read(id, nil, res)
	})
}

func (api *PacketFilterAPI) Update(id string, value *sacloud.PacketFilter) (*sacloud.PacketFilter, error) {
	return api.request(func(res *sacloud.Response) error {
		return api.update(id, api.createRequest(value), res)
	})
}

func (api *PacketFilterAPI) Delete(id string) (*sacloud.PacketFilter, error) {
	return api.request(func(res *sacloud.Response) error {
		return api.delete(id, nil, res)
	})
}

// ConnectPacketFilterToSharedNIC connect packet filter to eth0(shared)
func (api *PacketFilterAPI) ConnectPacketFilterToSharedNIC(server *sacloud.Server, idOrNameFilter string) error {
	if server.Interfaces != nil && len(server.Interfaces) > 0 {
		return api.connectPacketFilter(&server.Interfaces[0], idOrNameFilter)
	}
	return nil
}

// ConnectPacketFilterToPrivateNIC connect packet filter to eth1(private)
func (api *PacketFilterAPI) ConnectPacketFilterToPrivateNIC(server *sacloud.Server, idOrNameFilter string) error {
	if server.Interfaces != nil && len(server.Interfaces) > 1 {
		return api.connectPacketFilter(&server.Interfaces[1], idOrNameFilter)
	}
	return nil
}

// ConnectPacketFilter connect filter to nic
func (api *PacketFilterAPI) connectPacketFilter(nic *sacloud.Interface, idOrNameFilter string) error {
	if idOrNameFilter == "" {
		return nil
	}

	var id string
	//id or name ?
	if match, _ := regexp.MatchString(`^[0-9]+$`, idOrNameFilter); match {
		res, err := api.Read(idOrNameFilter)

		if err == nil {
			id = res.ID
		}
	}

	//search
	if id == "" {
		//名前での検索
		req := &sacloud.Request{}
		req.AddFilter("Name", idOrNameFilter)
		res, err := api.Find(req)
		if err != nil {
			return err
		}
		if res.Count > 0 {
			id = res.PacketFilters[0].ID
		} else {
			return fmt.Errorf("PacketFilter [%s](name):Not Found", idOrNameFilter)
		}
	}

	// not found
	if id == "" {
		return nil
	}

	_, err := api.ConnectToInterface(nic.ID, id)
	return err
}

func (api *PacketFilterAPI) ConnectToInterface(nicID string, packetFilterID string) (bool, error) {
	var (
		method = "PUT"
		uri    = fmt.Sprintf("/%s/%s/to/packetfilter/%s", api.getResourceURL(), nicID, packetFilterID)
	)
	res := &sacloud.ResultFlagValue{}
	err := api.baseAPI.request(method, uri, nil, res)
	if err != nil {
		return false, err
	}
	return res.IsOk, nil
}
