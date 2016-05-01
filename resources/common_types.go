package resources

// Resource type of sakuracloud resource(have ID:string)
type Resource struct {
	ID string `json:",omitempty"`
}

// NumberResource type of sakuracloud resource(int64)
type NumberResource struct {
	ID int64 `json:",omitempty"`
}

// EAvailability Enum of sakuracloud
type EAvailability struct {
	Availability string `json:",omitempty"`
}

// IsAvailable Return availability
func (a *EAvailability) IsAvailable() bool {
	return a.Availability == "available"
}

// SakuraCloudResources type of resources
type SakuraCloudResources struct {
	Server        *Server        `json:",omitempty"`
	Disk          *Disk          `json:",omitempty"`
	Note          *Note          `json:",omitempty"`
	PacketFilter  *PacketFilter  `json:",omitempty"`
	ProductServer *ProductServer `json:"ServerPlan,omitempty"`
	Archive       *Archive       `json:",omitempty"`
	//CommonServiceDnsItem  *CommonServiceDnsItem  `json:"CommonServiceItem,omitempty"`
	//CommonServiceGslbItem *CommonServiceGslbItem `json:"CommonServiceItem,omitempty"`
}

// SakuraCloudResourceList type of resources
type SakuraCloudResourceList struct {
	Servers        []Server        `json:",omitempty"`
	Disks          []Disk          `json:",omitempty"`
	Notes          []Note          `json:",omitempty"`
	Archives       []Archive       `json:",omitempty"`
	PacketFilters  []PacketFilter  `json:",omitempty"`
	ProductServers []ProductServer `json:"ServerPlans,omitempty"`
	//CommonServiceDnsItems  []CommonServiceDnsItem  `json:"CommonServiceItems,omitempty"`
	//CommonServiceGslbItems []CommonServiceGslbItem `json:"CommonServiceItems,omitempty"`
}

// Request type of SakuraCloud API Request
type Request struct {
	// *SakuraCloudResources
	Server       *Server        `json:",omitempty"`
	Disk         *Disk          `json:",omitempty"`
	Note         *Note          `json:",omitempty"`
	Archive      *Archive       `json:",omitempty"`
	PacketFilter *PacketFilter  `json:",omitempty"`
	ServerPlan   *ProductServer `json:",omitempty"`
	//CommonServiceDnsItem  *CommonServiceDnsItem  `json:"CommonServiceItem,omitempty"`
	//CommonServiceGslbItem *CommonServiceGslbItem `json:"CommonServiceItem,omitempty"`
	Interface *Interface             `json:",omitempty"`
	From      int                    `json:",omitempty"`
	Count     int                    `json:",omitempty"`
	Sort      []string               `json:",omitempty"`
	Filter    map[string]interface{} `json:",omitempty"`
	Exclude   []string               `json:",omitempty"`
	Include   []string               `json:",omitempty"`
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
}

// Response type of GET response
type Response struct {
	*ResultFlagValue
	*SakuraCloudResources
}
