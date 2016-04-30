package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	//"log"
	"net/http"
)

const (
	sakuraCloudAPIRoot       = "https://secure.sakura.ad.jp/cloud/zone"
	sakuraCloudAPIRootSuffix = "api/cloud/1.1"
)

var (
	client *Client
)

// Client type of sakuracloud api client config values
type Client struct {
	AccessToken       string
	AccessTokenSecret string
	Region            string
	*api
}

// NewClient Create new API client
func NewClient(token, tokenSecret, region string) *Client {
	c := &Client{AccessToken: token, AccessTokenSecret: tokenSecret, Region: region}
	c.api = newAPI(c)
	return c
}

type api struct {
	Archive      *ArchiveAPI
	Note         *NoteAPI
	Disk         *DiskAPI
	DNS          *DNSAPI
	GSLB         *GSLBAPI
	PacketFilter *PacketFilterAPI
	Product      *productAPI
	Server       *ServerAPI
}
type productAPI struct {
	Server *ProductServerAPI
}

func newAPI(client *Client) *api {
	return &api{
		Archive:      NewArchiveAPI(client),
		Note:         NewNoteAPI(client),
		Disk:         NewDiskAPI(client),
		DNS:          NewDNSAPI(client),
		GSLB:         NewGSLBAPI(client),
		PacketFilter: NewPacketFilterAPI(client),
		Product: &productAPI{
			Server: NewProductServerAPI(client),
		},
		Server: NewServerAPI(client),
	}
}

func (c *Client) getEndpoint() string {
	return fmt.Sprintf("%s/%s/%s", sakuraCloudAPIRoot, c.Region, sakuraCloudAPIRootSuffix)
}

func (c *Client) isOkStatus(code int) bool {
	codes := map[int]bool{
		200: true,
		201: true,
		202: true,
		204: true,
		305: false,
		400: false,
		401: false,
		403: false,
		404: false,
		405: false,
		406: false,
		408: false,
		409: false,
		411: false,
		413: false,
		415: false,
		500: false,
		503: false,
	}
	return codes[code]
}

func (c *Client) newRequest(method, uri string, body interface{}) ([]byte, error) {
	var (
		client = &http.Client{}
		url    = fmt.Sprintf("%s/%s", c.getEndpoint(), uri)
		err    error
		req    *http.Request
	)

	if body != nil {
		bodyJSON, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		if method == "GET" {

			url = fmt.Sprintf("%s/%s?%s", c.getEndpoint(), uri, bytes.NewBuffer(bodyJSON))
			req, err = http.NewRequest(method, url, nil)
		} else {
			req, err = http.NewRequest(method, url, bytes.NewBuffer(bodyJSON))
		}
		//log.Printf("********* method : %#v , url : %s , body : %#v", method, url, string(bodyJSON))

	} else {
		req, err = http.NewRequest(method, url, nil)
		//log.Printf("********* method : %#v , url : %s ", method, url)
	}

	if err != nil {
		return nil, fmt.Errorf("Error with request: %v - %q", url, err)
	}

	req.SetBasicAuth(c.AccessToken, c.AccessTokenSecret)
	req.Method = method

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	//log.Printf("******** response: %s", string(data))
	if !c.isOkStatus(resp.StatusCode) {
		return nil, fmt.Errorf("Error in response: %s", data)
	}
	if err != nil {
		return nil, err
	}

	return data, nil
}
