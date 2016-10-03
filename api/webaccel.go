package api

import (
	"encoding/json"
	"fmt"
	"github.com/yamamoto-febc/libsacloud/sacloud"
)

type WebAccelAPI struct {
	*baseAPI
}

func NewWebAccelAPI(client *Client) *WebAccelAPI {
	return &WebAccelAPI{
		&baseAPI{
			client:        client,
			apiRootSuffix: sakuraWebAccelAPIRootSuffix,
			FuncGetResourceURL: func() string {
				return ""
			},
		},
	}
}

type WebAccelDeleteCacheResponse struct {
	*sacloud.ResultFlagValue
	Results []*sacloud.DeleteCacheResult
}

func (api *WebAccelAPI) DeleteCache(urls ...string) (*WebAccelDeleteCacheResponse, error) {

	type request struct {
		URL []string
	}

	uri := fmt.Sprintf("%s/deletecache", api.getResourceURL())

	data, err := api.client.newRequest("POST", uri, &request{URL: urls})
	if err != nil {
		return nil, err
	}

	var res WebAccelDeleteCacheResponse
	if err := json.Unmarshal(data, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
