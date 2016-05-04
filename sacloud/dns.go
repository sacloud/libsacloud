package sacloud

import "time"

// DNS type of DNS(CommonServiceItem)
type DNS struct {
	*Resource
	Name         string
	Description  string      `json:",omitempty"`
	Status       DNSStatus   `json:",omitempty"`
	Provider     DNSProvider `json:",omitempty"`
	Settings     DNSSettings `json:",omitempty"`
	ServiceClass string      `json:",omitempty"`
	CreatedAt    *time.Time  `json:",omitempty"`
	ModifiedAt   *time.Time  `json:",omitempty"`
	Icon         *Icon       `json:",omitempty"`
	Tags         []string    `json:",omitempty"`
}

// DNSSettings type of DNSSettings
type DNSSettings struct {
	DNS DNSRecordSets `json:",omitempty"`
}

// DNSStatus type of DNSStatus
type DNSStatus struct {
	Zone string   `json:",omitempty"`
	NS   []string `json:",omitempty"`
}

// DNSProvider type of CommonServiceDNSProvider
type DNSProvider struct {
	Class string `json:",omitempty"`
}

// CreateNewDNS Create new CommonServiceDNSItem
func CreateNewDNS(zoneName string) *DNS {
	return &DNS{
		Resource: &Resource{ID: ""},
		Name:     zoneName,
		Status: DNSStatus{
			Zone: zoneName,
		},
		Provider: DNSProvider{
			Class: "dns",
		},
		Settings: DNSSettings{
			DNS: DNSRecordSets{},
		},
	}

}

// HasDNSRecord return has record
func (d *DNS) HasDNSRecord() bool {
	return len(d.Settings.DNS.ResourceRecordSets) > 0
}

// DNSRecordSets type of dns records
type DNSRecordSets struct {
	ResourceRecordSets []DNSRecordSet
}

// AddDNSRecordSet Add dns record
func (d *DNSRecordSets) AddDNSRecordSet(name string, ip string) {
	var record DNSRecordSet
	var isExist = false
	for i := range d.ResourceRecordSets {
		if d.ResourceRecordSets[i].Name == name && d.ResourceRecordSets[i].Type == "A" {
			d.ResourceRecordSets[i].RData = ip
			isExist = true
		}
	}

	if !isExist {
		record = DNSRecordSet{
			Name:  name,
			Type:  "A",
			RData: ip,
		}
		d.ResourceRecordSets = append(d.ResourceRecordSets, record)
	}
}

// DeleteDNSRecordSet Delete dns record
func (d *DNSRecordSets) DeleteDNSRecordSet(name string, ip string) {
	res := []DNSRecordSet{}
	for i := range d.ResourceRecordSets {
		if d.ResourceRecordSets[i].Name != name || d.ResourceRecordSets[i].Type != "A" || d.ResourceRecordSets[i].RData != ip {
			res = append(res, d.ResourceRecordSets[i])
		}
	}

	d.ResourceRecordSets = res
}

// DNSRecordSet type of dns records
type DNSRecordSet struct {
	Name  string `json:",omitempty"`
	Type  string `json:",omitempty"`
	RData string `json:",omitempty"`
	TTL   int    `json:",omitempty"`
}
