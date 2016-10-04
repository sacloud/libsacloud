package sacloud

import (
	"strconv"
	"time"
)

// Resource type of sakuracloud resource(have ID:string)
type Resource struct {
	ID int64 `json:",omitempty"`
}

type ResourceIDHolder interface {
	SetID(int64)
	GetID() int64
}

var EmptyID int64

func NewResource(id int64) *Resource {
	return &Resource{ID: id}
}
func NewResourceByStringID(id string) *Resource {
	intID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		panic(err)
	}
	return &Resource{ID: intID}
}

func (n *Resource) SetID(id int64) {
	n.ID = id
}
func (n *Resource) GetID() int64 {
	if n == nil {
		return -1
	}
	return n.ID
}

// EAvailability Enum of sakuracloud
type EAvailability struct {
	Availability string `json:",omitempty"`
}

// IsAvailable Return availability is "available"
func (a *EAvailability) IsAvailable() bool {
	return a.Availability == "available"
}

func (a *EAvailability) IsFailed() bool {
	return a.Availability == "failed"
}

//EServerInstanceStatus Enum [up / cleaning / down]
type EServerInstanceStatus struct {
	Status       string `json:",omitempty"`
	BeforeStatus string `json:",omitempty"`
}

func (e *EServerInstanceStatus) IsUp() bool {
	return e.Status == "up"
}

func (e *EServerInstanceStatus) IsDown() bool {
	return e.Status == "down"
}

// EScope Enum [shared / user]
type EScope string

var ESCopeShared = EScope("shared")
var ESCopeUser = EScope("user")

// EDiskConnection Enum [virtio / ide]
type EDiskConnection string

// SakuraCloudResources type of resources
type SakuraCloudResources struct {
	Server       *Server       `json:",omitempty"`
	Disk         *Disk         `json:",omitempty"`
	Note         *Note         `json:",omitempty"`
	Archive      *Archive      `json:",omitempty"`
	PacketFilter *PacketFilter `json:",omitempty"`
	Bridge       *Bridge       `json:",omitempty"`
	Icon         *Icon         `json:",omitempty"`
	Image        *Image        `json:",omitempty"`
	Interface    *Interface    `json:",omitempty"`
	Internet     *Internet     `json:",omitempty"`
	IPAddress    *IPAddress    `json:",omitempty"`
	License      *License      `json:",omitempty"`
	Switch       *Switch       `json:",omitempty"`
	CDROM        *CDROM        `json:",omitempty"`
	SSHKey       *SSHKey       `json:",omitempty"`
	Subnet       *Subnet       `json:",omitempty"`

	DiskPlan     *ProductDisk     `json:",omitempty"`
	InternetPlan *ProductInternet `json:",omitempty"`
	LicenseInfo  *ProductLicense  `json:",omitempty"`
	ServerPlan   *ProductServer   `json:",omitempty"`

	Region    *Region    `json:",omitempty"`
	Zone      *Zone      `json:",omitempty"`
	FTPServer *FTPServer `json:",omitempty"`
	//CommonServiceItemとApplianceはapiパッケージにて別途定義
}

// SakuraCloudResourceList type of resources
type SakuraCloudResourceList struct {
	Servers       []Server       `json:",omitempty"`
	Disks         []Disk         `json:",omitempty"`
	Notes         []Note         `json:",omitempty"`
	Archives      []Archive      `json:",omitempty"`
	PacketFilters []PacketFilter `json:",omitempty"`
	Bridges       []Bridge       `json:",omitempty"`
	Icons         []Icon         `json:",omitempty"`
	Interfaces    []Interface    `json:",omitempty"`
	Internet      []Internet     `json:",omitempty"`
	IPAddress     []IPAddress    `json:",omitempty"`
	Licenses      []License      `json:",omitempty"`
	Switches      []Switch       `json:",omitempty"`
	CDROMs        []CDROM        `json:",omitempty"`
	SSHKeys       []SSHKey       `json:",omitempty"`
	Subnets       []Subnet       `json:",omitempty"`

	DiskPlans     []ProductDisk     `json:",omitempty"`
	InternetPlans []ProductInternet `json:",omitempty"`
	LicenseInfo   []ProductLicense  `json:",omitempty"`
	ServerPlans   []ProductServer   `json:",omitempty"`

	Regions []Region `json:",omitempty"`
	Zones   []Zone   `json:",omitempty"`

	ServiceClasses []PublicPrice `json:",omitempty"`

	//CommonServiceItemとApplianceはapiパッケージにて別途定義
}

// Request type of SakuraCloud API Request
type Request struct {
	SakuraCloudResources
	From    int                    `json:",omitempty"`
	Count   int                    `json:",omitempty"`
	Sort    []string               `json:",omitempty"`
	Filter  map[string]interface{} `json:",omitempty"`
	Exclude []string               `json:",omitempty"`
	Include []string               `json:",omitempty"`
}

func (r *Request) AddFilter(key string, value interface{}) *Request {
	if r.Filter == nil {
		r.Filter = map[string]interface{}{}
	}
	r.Filter[key] = value
	return r
}

func (r *Request) AddSort(keyName string) *Request {
	if r.Sort == nil {
		r.Sort = []string{}
	}
	r.Sort = append(r.Sort, keyName)
	return r
}

func (r *Request) AddExclude(keyName string) *Request {
	if r.Exclude == nil {
		r.Exclude = []string{}
	}
	r.Exclude = append(r.Exclude, keyName)
	return r
}

func (r *Request) AddInclude(keyName string) *Request {
	if r.Include == nil {
		r.Include = []string{}
	}
	r.Include = append(r.Include, keyName)
	return r
}

// ResultFlagValue type of api result
type ResultFlagValue struct {
	IsOk    bool `json:"is_ok,omitempty"`
	Success bool `json:",omitempty"`
}

// SearchResponse  type of search/find response
type SearchResponse struct {
	Total int `json:",omitempty"`
	From  int `json:",omitempty"`
	Count int `json:",omitempty"`
	*SakuraCloudResourceList
	ResponsedAt *time.Time `json:",omitempty"`
}

// Response type of GET response
type Response struct {
	*ResultFlagValue
	*SakuraCloudResources
}

type ResultErrorValue struct {
	IsFatal      bool   `json:"is_fatal,omitempty"`
	Serial       string `json:"serial,omitempty"`
	Status       string `json:"status,omitempty"`
	ErrorCode    string `json:"error_code,omitempty"`
	ErrorMessage string `json:"error_msg,omitempty"`
}

type MigrationJobStatus struct {
	Status string `json:",omitempty"`
	Delays *struct {
		Start *struct {
			Max int `json:",omitempty"`
			Min int `json:",omitempty"`
		} `json:",omitempty"`
		Finish *struct {
			Max int `json:",omitempty"`
			Min int `json:",omitempty"`
		} `json:",omitempty"`
	}
}

type TagsType struct {
	Tags []string `json:",omitempty"`
}

var (
	// TagGroupA サーバをグループ化し起動ホストを分離します(グループA)
	TagGroupA = "@group=a"
	// TagGroupB サーバをグループ化し起動ホストを分離します(グループB)
	TagGroupB = "@group=b"
	// TagGroupC サーバをグループ化し起動ホストを分離します(グループC)
	TagGroupC = "@group=b"
	// TagGroupD サーバをグループ化し起動ホストを分離します(グループD)
	TagGroupD = "@group=b"

	// TagAutoReboot サーバ停止時に自動起動します
	TagAutoReboot = "@aut-reboot"

	// TagKeyboardUS リモートスクリーン画面でUSキーボード入力します
	TagKeyboardUS = "@keyboard-us"

	// TagBootCDROM 優先ブートデバイスをCD-ROMに設定します
	TagBootCDROM = "@boot-cdrom"
	// TagBootNetwork 優先ブートデバイスをPXE bootに設定します
	TagBootNetwork = "@boot-network"

	// TagVirtIONetPCI サーバの仮想NICをvirtio-netに変更します
	TagVirtIONetPCI = "@virtio-net-pci"
)

func (t *TagsType) HasTag(target string) bool {

	for _, tag := range t.Tags {
		if target == tag {
			return true
		}
	}

	return false
}

func (t *TagsType) AppendTag(target string) {
	if t.HasTag(target) {
		return
	}

	t.Tags = append(t.Tags, target)
}

func (t *TagsType) RemoveTag(target string) {
	if !t.HasTag(target) {
		return
	}
	res := []string{}
	for _, tag := range t.Tags {
		if tag != target {
			res = append(res, tag)
		}
	}

	t.Tags = res
}
