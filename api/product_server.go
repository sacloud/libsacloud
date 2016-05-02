package api

import (
	"fmt"
	"github.com/yamamoto-febc/libsacloud/sacloud"
	"strconv"
)

type ProductServerAPI struct {
	*baseAPI
}

func NewProductServerAPI(client *Client) *ProductServerAPI {
	return &ProductServerAPI{
		&baseAPI{
			client: client,
			FuncGetResourceURL: func() string {
				return "product/server"
			},
		},
	}
}

func (api *ProductServerAPI) request(f func(*sacloud.Response) error) (*sacloud.ProductServer, error) {
	res := &sacloud.Response{}
	err := f(res)
	if err != nil {
		return nil, err
	}
	return res.ProductServer, nil
}

func (api *ProductServerAPI) Read(id int64) (*sacloud.ProductServer, error) {
	return api.request(func(res *sacloud.Response) error {
		return api.read(fmt.Sprintf("%d", id), nil, res)
	})
}

// IsValidPlan return valid plan
func (api *ProductServerAPI) IsValidPlan(core int, memGB int) (bool, error) {
	//assert args
	if core <= 0 {
		return false, fmt.Errorf("Invalid Parameter: CPU Core")
	}
	if memGB <= 0 {
		return false, fmt.Errorf("Invalid Parameter: Memory Size(GB)")
	}

	planID, _ := strconv.ParseInt(fmt.Sprintf("%d%03d", memGB, core), 10, 64)

	productServer, err := api.Read(planID)

	if err != nil {
		return false, err
	}

	if productServer != nil {
		return true, nil
	}

	return false, fmt.Errorf("Server Plan[%d] Not Found", planID)

}
