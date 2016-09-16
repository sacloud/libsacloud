package api

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"github.com/yamamoto-febc/libsacloud/sacloud"
	"io"
	"strings"
	"time"
)

type BillAPI struct {
	*baseAPI
}

func NewBillAPI(client *Client) *BillAPI {
	return &BillAPI{
		&baseAPI{
			client:        client,
			apiRootSuffix: sakuraBillingAPIRootSuffix,
			FuncGetResourceURL: func() string {
				return "bill"
			},
		},
	}
}

type BillResponse struct {
	*sacloud.ResultFlagValue
	Count       int        `json:",omitempty"`
	ResponsedAt *time.Time `json:",omitempty"`
	Bills       []*sacloud.Bill
}

type BillDetailResponse struct {
	*sacloud.ResultFlagValue
	Count       int        `json:",omitempty"`
	ResponsedAt *time.Time `json:",omitempty"`
	BillDetails []*sacloud.BillDetail
}

type BillDetailCSVResponse struct {
	*sacloud.ResultFlagValue
	Count       int        `json:",omitempty"`
	ResponsedAt *time.Time `json:",omitempty"`
	Filename    string     `json:",omitempty"`
	RawBody     string     `json:"Body,omitempty"`
	HeaderRow   []string
	BodyRows    [][]string
}

func (res *BillDetailCSVResponse) buildCSVBody() {

	if res == nil || res.RawBody == "" {
		return
	}

	//CSV分割(先頭行/それ以降)、
	reader := csv.NewReader(strings.NewReader(res.RawBody))
	reader.LazyQuotes = true

	isFirst := true
	res.BodyRows = [][]string{}
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}

		if isFirst {
			res.HeaderRow = record
			isFirst = false
		} else {
			res.BodyRows = append(res.BodyRows, record)
		}
	}
}

func (api *BillAPI) ByContract(accountID int64) (*BillResponse, error) {

	uri := fmt.Sprintf("%s/by-contract/%d", api.getResourceURL(), accountID)
	return api.getContract(uri)
}

func (api *BillAPI) ByContractYear(accountID int64, year int) (*BillResponse, error) {
	uri := fmt.Sprintf("%s/by-contract/%d/%d", api.getResourceURL(), accountID, year)
	return api.getContract(uri)
}

func (api *BillAPI) ByContractYearMonth(accountID int64, year int, month int) (*BillResponse, error) {
	uri := fmt.Sprintf("%s/by-contract/%d/%d/%d", api.getResourceURL(), accountID, year, month)
	return api.getContract(uri)
}

func (api *BillAPI) Read(billNo int64) (*BillResponse, error) {
	uri := fmt.Sprintf("%s/id/%d/", api.getResourceURL(), billNo)
	return api.getContract(uri)

}

func (api *BillAPI) getContract(uri string) (*BillResponse, error) {

	data, err := api.client.newRequest("GET", uri, nil)
	if err != nil {
		return nil, err
	}
	var res BillResponse
	if err := json.Unmarshal(data, &res); err != nil {
		return nil, err
	}
	return &res, nil

}

func (api *BillAPI) GetDetail(memberCD string, billNo int64) (*BillDetailResponse, error) {

	//TODO マルチスレッド非対応
	oldFunc := api.FuncGetResourceURL
	defer func() { api.FuncGetResourceURL = oldFunc }()
	api.FuncGetResourceURL = func() string {
		return "billdetail"
	}

	uri := fmt.Sprintf("%s/%s/%d", api.getResourceURL(), memberCD, billNo)
	data, err := api.client.newRequest("GET", uri, nil)
	if err != nil {
		return nil, err
	}
	var res BillDetailResponse
	if err := json.Unmarshal(data, &res); err != nil {
		return nil, err
	}
	return &res, nil

}

func (api *BillAPI) GetDetailCSV(memberCD string, billNo int64) (*BillDetailCSVResponse, error) {

	//TODO マルチスレッド非対応
	oldFunc := api.FuncGetResourceURL
	defer func() { api.FuncGetResourceURL = oldFunc }()
	api.FuncGetResourceURL = func() string {
		return "billdetail"
	}

	uri := fmt.Sprintf("%s/%s/%d/csv", api.getResourceURL(), memberCD, billNo)
	data, err := api.client.newRequest("GET", uri, nil)
	if err != nil {
		return nil, err
	}
	var res BillDetailCSVResponse
	if err := json.Unmarshal(data, &res); err != nil {
		return nil, err
	}

	// build HeaderRow and BodyRows from RawBody
	res.buildCSVBody()

	return &res, nil

}
