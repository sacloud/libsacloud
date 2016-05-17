package api

import (
	"fmt"
	"github.com/yamamoto-febc/libsacloud/sacloud"
)

type InternetAPI struct {
	*baseAPI
}

func NewInternetAPI(client *Client) *InternetAPI {
	return &InternetAPI{
		&baseAPI{
			client: client,
			FuncGetResourceURL: func() string {
				return "internet"
			},
		},
	}
}

func (api *InternetAPI) UpdateBandWidth(id string, bandWidth int) (*sacloud.Internet, error) {
	var (
		method = "PUT"
		uri    = fmt.Sprintf("%s/%s/bandwidth", api.getResourceURL(), id)
		body   = &sacloud.Request{}
	)
	body.Internet = &sacloud.Internet{BandWidthMbps: bandWidth}

	return api.request(func(res *sacloud.Response) error {
		return api.baseAPI.request(method, uri, body, res)
	})
}
