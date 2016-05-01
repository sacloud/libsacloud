package api

import (
	"encoding/json"
	//	"strings"
	sakura "github.com/yamamoto-febc/libsacloud/resources"
)

//HACK: さくらのAPI側仕様: CommonServiceItemsの内容によってJSONフォーマットが異なるため
//      DNS/GSLB/シンプル監視それぞれでリクエスト/レスポンスデータ型を定義する。

type SearchGSLBResponse struct {
	Total                  int                            `json:",omitempty"`
	From                   int                            `json:",omitempty"`
	Count                  int                            `json:",omitempty"`
	CommonServiceGSLBItems []sakura.CommonServiceGSLBItem `json:"CommonServiceItems,omitempty"`
}

type gslbRequest struct {
	CommonServiceGSLBItem *sakura.CommonServiceGSLBItem `json:"CommonServiceItem,omitempty"`
	From                  int                           `json:",omitempty"`
	Count                 int                           `json:",omitempty"`
	Sort                  []string                      `json:",omitempty"`
	Filter                map[string]interface{}        `json:",omitempty"`
	Exclude               []string                      `json:",omitempty"`
	Include               []string                      `json:",omitempty"`
}

type gslbResponse struct {
	*sakura.ResultFlagValue
	*sakura.CommonServiceGSLBItem `json:"CommonServiceItem,omitempty"`
}

// GSLBAPI API Client for SAKURA CLOUD GSLB
type GSLBAPI struct {
	*baseAPI
}

func NewGSLBAPI(client *Client) *GSLBAPI {
	return &GSLBAPI{
		&baseAPI{
			client: client,
			FuncGetResourceURL: func() string {
				return "commonserviceitem"
			},
		},
	}
}

func (api *GSLBAPI) Find(condition *sakura.Request) (*SearchGSLBResponse, error) {

	//DNS固定
	condition.AddFilter("Provider.Class", "gslb")
	data, err := api.client.newRequest("GET", api.getResourceURL(), condition)
	if err != nil {
		return nil, err
	}
	var res SearchGSLBResponse
	if err := json.Unmarshal(data, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (api *GSLBAPI) request(f func(*gslbResponse) error) (*sakura.CommonServiceGSLBItem, error) {
	res := &gslbResponse{}
	err := f(res)
	if err != nil {
		return nil, err
	}
	return res.CommonServiceGSLBItem, nil
}

func (api *GSLBAPI) createRequest(value *sakura.CommonServiceGSLBItem) *gslbResponse {
	return &gslbResponse{CommonServiceGSLBItem: value}
}

func (api *GSLBAPI) Create(value *sakura.CommonServiceGSLBItem) (*sakura.CommonServiceGSLBItem, error) {
	return api.request(func(res *gslbResponse) error {
		return api.create(api.createRequest(value), res)
	})
}

func (api *GSLBAPI) Read(id string) (*sakura.CommonServiceGSLBItem, error) {
	return api.request(func(res *gslbResponse) error {
		return api.read(id, nil, res)
	})
}

func (api *GSLBAPI) Update(id string, value *sakura.CommonServiceGSLBItem) (*sakura.CommonServiceGSLBItem, error) {
	return api.request(func(res *gslbResponse) error {
		return api.update(id, api.createRequest(value), res)
	})
}

func (api *GSLBAPI) Delete(id string) (*sakura.CommonServiceGSLBItem, error) {
	return api.request(func(res *gslbResponse) error {
		return api.delete(id, nil, res)
	})
}

// SetupGSLBRecord create or update Gslb
func (api *GSLBAPI) SetupGSLBRecord(gslbName string, ip string) ([]string, error) {

	gslbItem, err := api.findOrCreateBy(gslbName)

	if err != nil {
		return nil, err
	}
	gslbItem.Settings.GSLB.AddServer(ip)
	res, err := api.updateGSLBServers(gslbItem)
	if err != nil {
		return nil, err
	}

	if gslbItem.ID == "" {
		return []string{res.Status.FQDN}, nil
	}
	return nil, nil

}

// DeleteGSLBServer delete gslb server
func (api *GSLBAPI) DeleteGSLBServer(gslbName string, ip string) error {
	gslbItem, err := api.findOrCreateBy(gslbName)
	if err != nil {
		return err
	}
	gslbItem.Settings.GSLB.DeleteServer(ip)

	if gslbItem.HasGSLBServer() {
		_, err = api.updateGSLBServers(gslbItem)
		if err != nil {
			return err
		}

	} else {
		_, err = api.Delete(gslbItem.ID)
		if err != nil {
			return err
		}

	}
	return nil
}

func (api *GSLBAPI) findOrCreateBy(gslbName string) (*sakura.CommonServiceGSLBItem, error) {

	req := &sakura.Request{}
	req.AddFilter("Name", gslbName)
	res, err := api.Find(req)
	if err != nil {
		return nil, err
	}

	//すでに登録されている場合
	var gslbItem *sakura.CommonServiceGSLBItem
	if res.Count > 0 {
		gslbItem = &res.CommonServiceGSLBItems[0]
	} else {
		gslbItem = sakura.CreateNewGSLBCommonServiceItem(gslbName)
	}

	return gslbItem, nil
}

func (api *GSLBAPI) updateGSLBServers(gslbItem *sakura.CommonServiceGSLBItem) (*sakura.CommonServiceGSLBItem, error) {

	var item *sakura.CommonServiceGSLBItem
	var err error

	if gslbItem.ID == "" {
		item, err = api.Create(gslbItem)
	} else {
		item, err = api.Update(gslbItem.ID, gslbItem)
	}

	if err != nil {
		return nil, err
	}

	return item, nil
}
