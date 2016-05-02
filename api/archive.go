package api

import (
	"errors"
	"github.com/yamamoto-febc/libsacloud/sacloud"
)

const (
	sakuraCloudPublicImageSearchWords = "Ubuntu%20Server%2014%2064bit"
)

type ArchiveAPI struct {
	*baseAPI
}

func NewArchiveAPI(client *Client) *ArchiveAPI {
	return &ArchiveAPI{
		&baseAPI{
			client: client,
			FuncGetResourceURL: func() string {
				return "archive"
			},
		},
	}
}

func (api *ArchiveAPI) request(f func(*sacloud.Response) error) (*sacloud.Archive, error) {
	res := &sacloud.Response{}
	err := f(res)
	if err != nil {
		return nil, err
	}
	return res.Archive, nil
}

func (api *ArchiveAPI) createRequest(value *sacloud.Archive) *sacloud.Request {
	return &sacloud.Request{Archive: value}
}

func (api *ArchiveAPI) Create(value *sacloud.Archive) (*sacloud.Archive, error) {
	return api.request(func(res *sacloud.Response) error {
		return api.create(api.createRequest(value), res)
	})
}

func (api *ArchiveAPI) Read(id string) (*sacloud.Archive, error) {
	return api.request(func(res *sacloud.Response) error {
		return api.read(id, nil, res)
	})
}

func (api *ArchiveAPI) Update(id string, value *sacloud.Archive) (*sacloud.Archive, error) {
	return api.request(func(res *sacloud.Response) error {
		return api.update(id, api.createRequest(value), res)
	})
}

func (api *ArchiveAPI) Delete(id string) (*sacloud.Archive, error) {
	return api.request(func(res *sacloud.Response) error {
		return api.delete(id, nil, res)
	})
}

// GetUbuntuArchiveID get ubuntu archive id
func (api *ArchiveAPI) GetUbuntuArchiveID() (string, error) {

	var req = &sacloud.Request{}
	req.AddFilter("Name", sakuraCloudPublicImageSearchWords).
		AddFilter("Scope", "shared").
		AddInclude("ID").
		AddInclude("Name")

	res, err := api.Find(req)
	if err != nil {
		return "", err
	}

	//すでに登録されている場合
	if res.Count > 0 {
		return res.Archives[0].ID, nil
	}

	return "", errors.New("Archive'Ubuntu Server 14.04 LTS 64bit' not found.")
}
