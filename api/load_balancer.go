package api

import (
	"encoding/json"
	"fmt"
	"github.com/yamamoto-febc/libsacloud/sacloud"
	"time"
)

//HACK: さくらのAPI側仕様: Applianceの内容によってJSONフォーマットが異なるため
//      ロードバランサ/VPCルータそれぞれでリクエスト/レスポンスデータ型を定義する。

type SearchLoadBalancerResponse struct {
	Total         int                    `json:",omitempty"`
	From          int                    `json:",omitempty"`
	Count         int                    `json:",omitempty"`
	LoadBalancers []sacloud.LoadBalancer `json:"Appliances,omitempty"`
}

type loadBalancerRequest struct {
	LoadBalancer *sacloud.LoadBalancer  `json:"Appliance,omitempty"`
	From         int                    `json:",omitempty"`
	Count        int                    `json:",omitempty"`
	Sort         []string               `json:",omitempty"`
	Filter       map[string]interface{} `json:",omitempty"`
	Exclude      []string               `json:",omitempty"`
	Include      []string               `json:",omitempty"`
}

type loadBalancerResponse struct {
	*sacloud.ResultFlagValue
	*sacloud.LoadBalancer `json:"Appliance,omitempty"`
	Success               interface{} `json:",omitempty"` //HACK: さくらのAPI側仕様: 戻り値:Successがbool値へ変換できないためinterface{}
}

type LoadBalancerAPI struct {
	*baseAPI
}

func NewLoadBalancerAPI(client *Client) *LoadBalancerAPI {
	return &LoadBalancerAPI{
		&baseAPI{
			client: client,
			FuncGetResourceURL: func() string {
				return "appliance"
			},
			FuncBaseSearchCondition: func() *sacloud.Request {
				res := &sacloud.Request{}
				res.AddFilter("Class", "loadbalancer")
				return res
			},
		},
	}
}

func (api *LoadBalancerAPI) Find() (*SearchLoadBalancerResponse, error) {
	data, err := api.client.newRequest("GET", api.getResourceURL(), api.getSearchState())
	if err != nil {
		return nil, err
	}
	var res SearchLoadBalancerResponse
	if err := json.Unmarshal(data, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (api *LoadBalancerAPI) request(f func(*loadBalancerResponse) error) (*sacloud.LoadBalancer, error) {
	res := &loadBalancerResponse{}
	err := f(res)
	if err != nil {
		return nil, err
	}
	return res.LoadBalancer, nil
}

func (api *LoadBalancerAPI) createRequest(value *sacloud.LoadBalancer) *loadBalancerResponse {
	return &loadBalancerResponse{LoadBalancer: value}
}

//func (api *LoadBalancerAPI) New() *sacloud.LoadBalancer {
//	return sacloud.CreateNewLoadBalancer()
//}

func (api *LoadBalancerAPI) Create(value *sacloud.LoadBalancer) (*sacloud.LoadBalancer, error) {
	return api.request(func(res *loadBalancerResponse) error {
		return api.create(api.createRequest(value), res)
	})
}

func (api *LoadBalancerAPI) Read(id int64) (*sacloud.LoadBalancer, error) {
	return api.request(func(res *loadBalancerResponse) error {
		return api.read(id, nil, res)
	})
}

func (api *LoadBalancerAPI) Update(id int64, value *sacloud.LoadBalancer) (*sacloud.LoadBalancer, error) {
	return api.request(func(res *loadBalancerResponse) error {
		return api.update(id, api.createRequest(value), res)
	})
}

func (api *LoadBalancerAPI) Delete(id int64) (*sacloud.LoadBalancer, error) {
	return api.request(func(res *loadBalancerResponse) error {
		return api.delete(id, nil, res)
	})
}

func (api *LoadBalancerAPI) Config(id int64) (bool, error) {
	var (
		method = "PUT"
		uri    = fmt.Sprintf("%s/%d/config", api.getResourceURL(), id)
	)
	return api.modify(method, uri, nil)
}

func (api *LoadBalancerAPI) IsUp(id int64) (bool, error) {
	lb, err := api.Read(id)
	if err != nil {
		return false, err
	}
	return lb.Instance.IsUp(), nil
}

func (api *LoadBalancerAPI) IsDown(id int64) (bool, error) {
	lb, err := api.Read(id)
	if err != nil {
		return false, err
	}
	return lb.Instance.IsDown(), nil
}

// Boot power on
func (api *LoadBalancerAPI) Boot(id int64) (bool, error) {
	var (
		method = "PUT"
		uri    = fmt.Sprintf("%s/%d/power", api.getResourceURL(), id)
	)
	return api.modify(method, uri, nil)
}

// Shutdown power off
func (api *LoadBalancerAPI) Shutdown(id int64) (bool, error) {
	var (
		method = "DELETE"
		uri    = fmt.Sprintf("%s/%d/power", api.getResourceURL(), id)
	)

	return api.modify(method, uri, nil)
}

// Stop force shutdown
func (api *LoadBalancerAPI) Stop(id int64) (bool, error) {
	var (
		method = "DELETE"
		uri    = fmt.Sprintf("%s/%d/power", api.getResourceURL(), id)
	)

	return api.modify(method, uri, map[string]bool{"Force": true})
}

func (api *LoadBalancerAPI) RebootForce(id int64) (bool, error) {
	var (
		method = "PUT"
		uri    = fmt.Sprintf("%s/%d/reset", api.getResourceURL(), id)
	)

	return api.modify(method, uri, nil)
}

func (api *LoadBalancerAPI) ResetForce(id int64, recycleProcess bool) (bool, error) {
	var (
		method = "PUT"
		uri    = fmt.Sprintf("%s/%d/reset", api.getResourceURL(), id)
	)

	return api.modify(method, uri, map[string]bool{"RecycleProcess": recycleProcess})
}

func (api *LoadBalancerAPI) SleepUntilUp(id int64, timeout time.Duration) error {
	current := 0 * time.Second
	interval := 5 * time.Second
	for {

		up, err := api.IsUp(id)
		if err != nil {
			return err
		}

		if up {
			return nil
		}
		time.Sleep(interval)
		current += interval

		if timeout > 0 && current > timeout {
			return fmt.Errorf("Timeout: WaitforAvailable")
		}
	}
}

func (api *LoadBalancerAPI) SleepUntilDown(id int64, timeout time.Duration) error {
	current := 0 * time.Second
	interval := 5 * time.Second
	for {

		down, err := api.IsDown(id)
		if err != nil {
			return err
		}

		if down {
			return nil
		}
		time.Sleep(interval)
		current += interval

		if timeout > 0 && current > timeout {
			return fmt.Errorf("Timeout: WaitforAvailable")
		}
	}
}

// SleepWhileCopying wait until became to available
func (api *LoadBalancerAPI) SleepWhileCopying(id int64, timeout time.Duration, maxRetryCount int) error {
	current := 0 * time.Second
	interval := 5 * time.Second
	errCount := 0

	for {
		loadBalancer, err := api.Read(id)
		if err != nil {
			errCount++
			if errCount > maxRetryCount {
				return err
			}
		}

		if loadBalancer != nil && loadBalancer.IsAvailable() {
			return nil
		}
		time.Sleep(interval)
		current += interval

		if timeout > 0 && current > timeout {
			return fmt.Errorf("Timeout: SleepWhileCopying")
		}
	}
}
