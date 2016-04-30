package resources

import "time"

// CommonServiceGSLBItem type of CommonServiceGSLBItem
type CommonServiceGSLBItem struct {
	*Resource
	Name         string
	Description  string                    `json:",omitempty"`
	Status       CommonServiceGSLBStatus   `json:",omitempty"`
	Provider     CommonServiceGSLBProvider `json:",omitempty"`
	Settings     CommonServiceGSLBSettings `json:",omitempty"`
	ServiceClass string                    `json:",omitempty"`
	CreatedAt    time.Time                 `json:",omitempty"`
	ModifiedAt   time.Time                 `json:",omitempty"`
	Icon         *Icon                     `json:",omitempty"`
	Tags         []string                  `json:",omitempty"`
}

// CommonServiceGSLBSettings type of CommonServiceGSLBSettings
type CommonServiceGSLBSettings struct {
	GSLB GSLBRecordSets `json:",omitempty"`
}

// CommonServiceGSLBStatus type of CommonServiceGSLBStatus
type CommonServiceGSLBStatus struct {
	FQDN string `json:",omitempty"`
}

// CommonServiceGSLBProvider type of CommonServiceGSLBProvider
type CommonServiceGSLBProvider struct {
	Class string `json:",omitempty"`
}

// CreateNewGSLBCommonServiceItem Create new CommonServiceItem
func CreateNewGSLBCommonServiceItem(gslbName string) *CommonServiceGSLBItem {
	return &CommonServiceGSLBItem{
		Resource: &Resource{ID: ""},
		Name:     gslbName,
		Provider: CommonServiceGSLBProvider{
			Class: "gslb",
		},
		Settings: CommonServiceGSLBSettings{
			GSLB: GSLBRecordSets{
				DelayLoop:   10,
				HealthCheck: defaultGSLBHealthCheck,
				Weighted:    "True",
			},
		},
	}

}

// HasGSLBServer return has server
func (d *CommonServiceGSLBItem) HasGSLBServer() bool {
	return len(d.Settings.GSLB.Servers) > 0
}

// GSLBRecordSets type of GSLBRecordSets
type GSLBRecordSets struct {
	DelayLoop   int             `json:",omitempty"`
	HealthCheck GSLBHealthCheck `json:",omitempty"`
	Weighted    string          `json:",omitempty"`
	Servers     []GSLBServer    `json:",omitempty"`
}

// AddServer Add server to GSLB
func (g *GSLBRecordSets) AddServer(ip string) {
	var record GSLBServer
	var isExist = false
	for i := range g.Servers {
		if g.Servers[i].IPAddress == ip {
			isExist = true
		}
	}

	if !isExist {
		record = GSLBServer{
			IPAddress: ip,
			Enabled:   "True",
			Weight:    "1",
		}
		g.Servers = append(g.Servers, record)
	}
}

//DeleteServer Delete server from GSLB
func (g *GSLBRecordSets) DeleteServer(ip string) {
	res := []GSLBServer{}
	for i := range g.Servers {
		if g.Servers[i].IPAddress != ip {
			res = append(res, g.Servers[i])
		}
	}

	g.Servers = res
}

// GSLBServer type of GSLBServer
type GSLBServer struct {
	IPAddress string `json:",omitempty"`
	Enabled   string `json:",omitempty"`
	Weight    string `json:",omitempty"`
}

// GSLBHealthCheck type of GSLBHealthCheck
type GSLBHealthCheck struct {
	Protocol string `json:",omitempty"`
	Path     string `json:",omitempty"`
	Status   string `json:",omitempty"`
}

var defaultGSLBHealthCheck = GSLBHealthCheck{
	Protocol: "http",
	Path:     "/",
	Status:   "200",
}
