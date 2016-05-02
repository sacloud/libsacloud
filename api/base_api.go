package api

import (
	"encoding/json"
	"fmt"
	"github.com/yamamoto-febc/libsacloud/sacloud"
)

type baseAPI struct {
	client             *Client
	FuncGetResourceURL func() string
}

func (b *baseAPI) getResourceURL() string {
	if b.FuncGetResourceURL != nil {
		return b.FuncGetResourceURL()
	}
	return ""
}

func (b *baseAPI) Find(condition *sacloud.Request) (*sacloud.SearchResponse, error) {

	data, err := b.client.newRequest("GET", b.getResourceURL(), condition)
	if err != nil {
		return nil, err
	}
	var res sacloud.SearchResponse
	if err := json.Unmarshal(data, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (b *baseAPI) request(method string, uri string, body interface{}, res interface{}) error {
	data, err := b.client.newRequest(method, uri, body)
	if err != nil {
		return err
	}

	if res != nil {
		if err := json.Unmarshal(data, &res); err != nil {
			return err
		}
	}
	return nil
}

func (b *baseAPI) create(body interface{}, res interface{}) error {
	var (
		method = "POST"
		uri    = b.getResourceURL()
	)

	return b.request(method, uri, body, res)
}

func (b *baseAPI) read(id string, body interface{}, res interface{}) error {
	var (
		method = "GET"
		uri    = fmt.Sprintf("%s/%s", b.getResourceURL(), id)
	)

	return b.request(method, uri, body, res)
}

func (b *baseAPI) update(id string, body interface{}, res interface{}) error {
	var (
		method = "PUT"
		uri    = fmt.Sprintf("%s/%s", b.getResourceURL(), id)
	)
	return b.request(method, uri, body, res)
}

func (b *baseAPI) delete(id string, body interface{}, res interface{}) error {
	var (
		method = "DELETE"
		uri    = fmt.Sprintf("%s/%s", b.getResourceURL(), id)
	)
	return b.request(method, uri, body, res)
}
