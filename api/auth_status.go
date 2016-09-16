package api

import (
	"encoding/json"
	"github.com/yamamoto-febc/libsacloud/sacloud"
)

type AuthStatusAPI struct {
	*baseAPI
}

func NewAuthStatusAPI(client *Client) *AuthStatusAPI {
	return &AuthStatusAPI{
		&baseAPI{
			client: client,
			FuncGetResourceURL: func() string {
				return "auth-status"
			},
		},
	}
}

func (b *AuthStatusAPI) Read() (*sacloud.AuthStatus, error) {

	data, err := b.client.newRequest("GET", b.getResourceURL(), nil)
	if err != nil {
		return nil, err
	}
	var res sacloud.AuthStatus
	if err := json.Unmarshal(data, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
