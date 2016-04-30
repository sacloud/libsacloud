package api

import (
	"errors"
	sakura "github.com/yamamoto-febc/libsacloud/resources"
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

func (api *ArchiveAPI) request(f func(*sakura.Response) error) (*sakura.Archive, error) {
	res := &sakura.Response{}
	err := f(res)
	if err != nil {
		return nil, err
	}
	return res.Archive, nil
}

func (api *ArchiveAPI) createRequest(value *sakura.Archive) *sakura.Request {
	return &sakura.Request{Archive: value}
}

func (api *ArchiveAPI) Create(value *sakura.Archive) (*sakura.Archive, error) {
	return api.request(func(res *sakura.Response) error {
		return api.create(api.createRequest(value), res)
	})
}

func (api *ArchiveAPI) Read(id string) (*sakura.Archive, error) {
	return api.request(func(res *sakura.Response) error {
		return api.read(id, nil, res)
	})
}

func (api *ArchiveAPI) Update(id string, value *sakura.Archive) (*sakura.Archive, error) {
	return api.request(func(res *sakura.Response) error {
		return api.update(id, api.createRequest(value), res)
	})
}

func (api *ArchiveAPI) Delete(id string) (*sakura.Archive, error) {
	return api.request(func(res *sakura.Response) error {
		return api.delete(id, nil, res)
	})
}

// GetUbuntuArchiveID get ubuntu archive id
func (api *ArchiveAPI) GetUbuntuArchiveID() (string, error) {

	var req = &sakura.Request{}
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
