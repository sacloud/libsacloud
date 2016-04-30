package api

import (
	"encoding/json"
	"fmt"
	sakura "github.com/yamamoto-febc/libsacloud/resources"
	"strings"
)

type searchDNSResponse struct {
	Total                 int                           `json:",omitempty"`
	From                  int                           `json:",omitempty"`
	Count                 int                           `json:",omitempty"`
	CommonServiceDNSItems []sakura.CommonServiceDNSItem `json:"CommonServiceItems,omitempty"`
}

type dnsRequest struct {
	CommonServiceDNSItem *sakura.CommonServiceDNSItem `json:"CommonServiceItem,omitempty"`
	From                 int                          `json:",omitempty"`
	Count                int                          `json:",omitempty"`
	Sort                 []string                     `json:",omitempty"`
	Filter               map[string]interface{}       `json:",omitempty"`
	Exclude              []string                     `json:",omitempty"`
	Include              []string                     `json:",omitempty"`
}
type dnsResponse struct {
	*sakura.ResultFlagValue
	*sakura.CommonServiceDNSItem `json:"CommonServiceItem,omitempty"`
}

// SetupDNSRecord get dns zone commonserviceitem id
func (c *Client) SetupDNSRecord(zoneName string, hostName string, ip string) ([]string, error) {

	dnsItem, err := c.getDNSCommonServiceItem(zoneName)
	if err != nil {
		return nil, err
	}

	if strings.HasSuffix(hostName, zoneName) {
		hostName = strings.Replace(hostName, zoneName, "", -1)
	}

	dnsItem.Settings.DNS.AddDNSRecordSet(hostName, ip)

	res, err := c.updateDNSRecord(dnsItem)
	if err != nil {
		return nil, err
	}

	if dnsItem.ID == "" {
		return res.Status.NS, nil
	}

	return nil, nil

}

// DeleteDNSRecord delete dns record
func (c *Client) DeleteDNSRecord(zoneName string, hostName string, ip string) error {
	dnsItem, err := c.getDNSCommonServiceItem(zoneName)
	if err != nil {
		return err
	}
	dnsItem.Settings.DNS.DeleteDNSRecordSet(hostName, ip)

	if dnsItem.HasDNSRecord() {
		_, err = c.updateDNSRecord(dnsItem)
		if err != nil {
			return err
		}

	} else {
		err = c.deleteCommonServiceDNSItem(dnsItem)
		if err != nil {
			return err
		}

	}
	return nil
}

func (c *Client) getDNSCommonServiceItem(zoneName string) (*sakura.CommonServiceDNSItem, error) {

	var (
		method = "GET"
		uri    = "commonserviceitem"
		body   = sakura.Request{
			Filter: map[string]interface{}{
				"Name":           zoneName,
				"Provider.Class": "dns",
			},
		}
	)

	bodyJSON, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	uri = fmt.Sprintf("%s?%s", uri, bodyJSON)
	data, err := c.newRequest(method, uri, nil)
	if err != nil {
		return nil, err
	}
	var dnsZone searchDNSResponse
	if err := json.Unmarshal(data, &dnsZone); err != nil {
		return nil, err
	}

	//すでに登録されている場合
	var dnsItem *sakura.CommonServiceDNSItem
	if dnsZone.Count > 0 {
		dnsItem = &dnsZone.CommonServiceDNSItems[0]
	} else {
		dnsItem = sakura.CreateNewDNSCommonServiceItem(zoneName)
	}

	return dnsItem, nil
}

func (c *Client) updateDNSRecord(dnsItem *sakura.CommonServiceDNSItem) (*sakura.CommonServiceDNSItem, error) {

	var (
		method string
		uri    string
	)
	if dnsItem.ID == "" {
		method = "POST"
		uri = "/commonserviceitem"

	} else {
		method = "PUT"
		uri = fmt.Sprintf("/commonserviceitem/%s", dnsItem.ID)
	}
	n := dnsRequest{
		CommonServiceDNSItem: dnsItem,
	}

	data, err := c.newRequest(method, uri, n)
	if err != nil {
		return nil, err
	}
	var res dnsResponse
	if err := json.Unmarshal(data, &res); err != nil {
		return nil, err
	}

	return res.CommonServiceDNSItem, nil
}

func (c *Client) deleteCommonServiceDNSItem(item *sakura.CommonServiceDNSItem) error {
	var (
		method string
		uri    string
	)
	method = "DELETE"
	uri = fmt.Sprintf("/commonserviceitem/%s", item.ID)

	_, err := c.newRequest(method, uri, item)
	if err != nil {
		return err
	}

	return nil

}
