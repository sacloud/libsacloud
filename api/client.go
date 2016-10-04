package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/yamamoto-febc/libsacloud/sacloud"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	sakuraCloudAPIRoot = "https://secure.sakura.ad.jp/cloud/zone"
)

var (
	client *Client
)

// Client type of sakuracloud api client config values
type Client struct {
	AccessToken       string
	AccessTokenSecret string
	Zone              string
	*api
	TraceMode bool
}

// NewClient Create new API client
func NewClient(token, tokenSecret, zone string) *Client {
	c := &Client{AccessToken: token, AccessTokenSecret: tokenSecret, Zone: zone, TraceMode: false}
	c.api = newAPI(c)
	return c
}

func (c *Client) Clone() *Client {
	n := &Client{AccessToken: c.AccessToken, AccessTokenSecret: c.AccessTokenSecret, Zone: c.Zone, TraceMode: c.TraceMode}
	n.api = newAPI(n)
	return n
}

type api struct {
	AuthStatus    *AuthStatusAPI
	AutoBackup    *AutoBackupAPI
	Archive       *ArchiveAPI
	Bill          *BillAPI
	Bridge        *BridgeAPI
	CDROM         *CDROMAPI
	Database      *DatabaseAPI
	Disk          *DiskAPI
	DNS           *DNSAPI
	Facility      *facilityAPI
	GSLB          *GSLBAPI
	Icon          *IconAPI
	Interface     *InterfaceAPI
	Internet      *InternetAPI
	IPAddress     *IPAddressAPI
	IPv6Addr      *IPv6AddrAPI
	IPv6Net       *IPv6NetAPI
	License       *LicenseAPI
	LoadBalancer  *LoadBalancerAPI
	Note          *NoteAPI
	PacketFilter  *PacketFilterAPI
	Product       *productAPI
	Server        *ServerAPI
	SimpleMonitor *SimpleMonitorAPI
	SSHKey        *SSHKeyAPI
	Subnet        *SubnetAPI
	Switch        *SwitchAPI
	VPCRouter     *VPCRouterAPI
	WebAccel      *WebAccelAPI
}
type productAPI struct {
	Server   *ProductServerAPI
	License  *ProductLicenseAPI
	Disk     *ProductDiskAPI
	Internet *ProductInternetAPI
	Price    *PublicPriceAPI
}

type facilityAPI struct {
	Region *RegionAPI
	Zone   *ZoneAPI
}

func newAPI(client *Client) *api {
	return &api{
		AuthStatus: NewAuthStatusAPI(client),
		AutoBackup: NewAutoBackupAPI(client),
		Archive:    NewArchiveAPI(client),
		Bill:       NewBillAPI(client),
		Bridge:     NewBridgeAPI(client),
		CDROM:      NewCDROMAPI(client),
		Database:   NewDatabaseAPI(client),
		Disk:       NewDiskAPI(client),
		DNS:        NewDNSAPI(client),
		Facility: &facilityAPI{
			Region: NewRegionAPI(client),
			Zone:   NewZoneAPI(client),
		},
		GSLB:         NewGSLBAPI(client),
		Icon:         NewIconAPI(client),
		Interface:    NewInterfaceAPI(client),
		Internet:     NewInternetAPI(client),
		IPAddress:    NewIPAddressAPI(client),
		IPv6Addr:     NewIPv6AddrAPI(client),
		IPv6Net:      NewIPv6NetAPI(client),
		License:      NewLicenseAPI(client),
		LoadBalancer: NewLoadBalancerAPI(client),
		Note:         NewNoteAPI(client),
		PacketFilter: NewPacketFilterAPI(client),
		Product: &productAPI{
			Server:   NewProductServerAPI(client),
			License:  NewProductLicenseAPI(client),
			Disk:     NewProductDiskAPI(client),
			Internet: NewProductInternetAPI(client),
			Price:    NewPublicPriceAPI(client),
		},
		Server:        NewServerAPI(client),
		SimpleMonitor: NewSimpleMonitorAPI(client),
		SSHKey:        NewSSHKeyAPI(client),
		Subnet:        NewSubnetAPI(client),
		Switch:        NewSwitchAPI(client),
		VPCRouter:     NewVPCRouterAPI(client),
		WebAccel:      NewWebAccelAPI(client),
	}
}

func (c *Client) getEndpoint() string {
	return fmt.Sprintf("%s/%s", sakuraCloudAPIRoot, c.Zone)
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
		if c.TraceMode {
			b, _ := json.MarshalIndent(body, "", "\t")
			log.Printf("[libsacloud:Client#request] method : %#v , url : %s , \nbody : %s", method, url, b)
		}

	} else {
		req, err = http.NewRequest(method, url, nil)
		if c.TraceMode {
			log.Printf("[libsacloud:Client#request] method : %#v , url : %s ", method, url)
		}
	}

	if err != nil {
		return nil, fmt.Errorf("Error with request: %v - %q", url, err)
	}

	req.SetBasicAuth(c.AccessToken, c.AccessTokenSecret)
	req.Header.Add("X-Sakura-Bigint-As-Int", "1") //Use BigInt on resource ids.
	//if c.TraceMode {
	//	req.Header.Add("X-Sakura-API-Beautify", "1") // format response-JSON
	//}
	req.Method = method

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if c.TraceMode {
		v := &map[string]interface{}{}
		json.Unmarshal(data, v)
		b, _ := json.MarshalIndent(v, "", "\t")
		log.Printf("[libsacloud:Client#response] : %s", b)
	}
	if !c.isOkStatus(resp.StatusCode) {

		errResponse := &sacloud.ResultErrorValue{}
		err := json.Unmarshal(data, errResponse)

		if err != nil {
			return nil, fmt.Errorf("Error in response: %s", string(data))
		}
		return nil, fmt.Errorf("Error in response: %#v", errResponse)

	}
	if err != nil {
		return nil, err
	}

	return data, nil
}
