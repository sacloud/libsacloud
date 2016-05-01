package api

import (
	"fmt"
	sakura "github.com/yamamoto-febc/libsacloud/resources"
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

func (api *ServerAPI) request(f func(*sakura.Response) error) (*sakura.Server, error) {
	res := &sakura.Response{}
	err := f(res)
	if err != nil {
		return nil, err
	}
	return res.Server, nil
}

func (api *ServerAPI) createRequest(value *sakura.Server) *sakura.Request {
	return &sakura.Request{Server: value}
}

func (api *ServerAPI) Create(value *sakura.Server) (*sakura.Server, error) {
	return api.request(func(res *sakura.Response) error {
		return api.create(api.createRequest(value), res)
	})
}

func (api *ServerAPI) Read(id string) (*sakura.Server, error) {
	return api.request(func(res *sakura.Response) error {
		return api.read(id, nil, res)
	})
}

func (api *ServerAPI) Update(id string, value *sakura.Server) (*sakura.Server, error) {
	return api.request(func(res *sakura.Response) error {
		return api.update(id, api.createRequest(value), res)
	})
}

func (api *ServerAPI) Delete(id string, disks []string) (*sakura.Server, error) {
	return api.request(func(res *sakura.Response) error {
		body := &struct{ WithDisk []string }{disks}
		if disks == nil {
			body = nil
		}
		return api.delete(id, body, res)
	})
}

// CreateWithAdditionalIP create server
func (api *ServerAPI) CreateWithAdditionalIP(spec *sakura.Server, addIPAddress string) (*sakura.Server, error) {

	server, err := api.Create(spec)
	if err != nil {
		return nil, err
	}

	if addIPAddress != "" && len(server.Interfaces) > 1 {
		if err := api.updateIPAddress(server, addIPAddress); err != nil {
			return nil, err
		}
	}

	return server, nil
}

func (api *ServerAPI) updateIPAddress(server *sakura.Server, ip string) error {
	//TODO 高レベルAPIへ移動
	var (
		method = "PUT"
		uri    = fmt.Sprintf("interface/%s", server.Interfaces[1].ID)
		body   = sakura.Request{
			Interface: &sakura.Interface{UserIPAddress: ip},
		}
	)

	_, err := api.client.newRequest(method, uri, body)
	if err != nil {
		return err
	}

	return nil

}

// State get server state
func (api *ServerAPI) State(id string) (string, error) {
	server, err := api.Read(id)
	if err != nil {
		return "", err
	}
	return server.Availability, nil
}

// PowerOn power on
func (api *ServerAPI) PowerOn(id string) error {
	var (
		method = "PUT"
		uri    = fmt.Sprintf("%s/%s/power", api.getResourceURL(), id)
	)

	res := &sakura.ResultFlagValue{}
	err := api.baseAPI.request(method, uri, nil, res)
	if err != nil {
		return err
	}
	return nil
}

// PowerOff power off
func (api *ServerAPI) PowerOff(id string) error {
	var (
		method = "DELETE"
		uri    = fmt.Sprintf("%s/%s/power", api.getResourceURL(), id)
	)

	res := &sakura.ResultFlagValue{}
	err := api.baseAPI.request(method, uri, nil, res)
	if err != nil {
		return err
	}
	return nil
}

//TODO 高レベルAPIヘ移動
//// GetIP get public ip address
//func (c *Client) GetIP(id string, privateIPOnly bool) (string, error) {
//
//
//
//	var (
//		method = "GET"
//		uri    = fmt.Sprintf("%s/%s", "server", id)
//	)
//
//	data, err := c.newRequest(method, uri, nil)
//	if err != nil {
//		return "", err
//	}
//	var s sakura.Response
//	if err := json.Unmarshal(data, &s); err != nil {
//		return "", err
//	}
//
//	if privateIPOnly && len(s.Server.Interfaces) > 1 {
//		return s.Server.Interfaces[1].UserIPAddress, nil
//	}
//
//	return s.Server.Interfaces[0].IPAddress, nil
//}

// SearchServerByName Search server by name
func (api *ServerAPI) SearchServerByName(name string) (*sakura.Server, error) {

	req := &sakura.Request{}
	req.AddFilter("Name", name)

	res, err := api.Find(req)
	if err != nil {
		return nil, err
	}
	if res.Count > 0 {
		return &res.Servers[0], nil
	}
	return nil, fmt.Errorf("server [%s] is not found", name)
}
