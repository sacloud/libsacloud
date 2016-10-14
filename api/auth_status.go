package api

import (
	"encoding/json"
	"github.com/yamamoto-febc/libsacloud/sacloud"
)

// AuthStatusAPI 認証状態API
type AuthStatusAPI struct {
	*baseAPI
}

// NewAuthStatusAPI 認証状態API作成
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

// Read 読み取り
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

// Find 検索
func (b *AuthStatusAPI) Find() (*sacloud.AuthStatus, error) {
	return b.Read()
}
