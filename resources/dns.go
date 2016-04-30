package resources

import "time"

// CommonServiceDNSItem type of CommonServiceDNSItem
type CommonServiceDNSItem struct {
	*Resource
	Name         string
	Description  string                   `json:",omitempty"`
	Status       CommonServiceDNSStatus   `json:",omitempty"`
	Provider     CommonServiceDNSProvider `json:",omitempty"`
	Settings     CommonServiceDNSSettings `json:",omitempty"`
	ServiceClass string                   `json:",omitempty"`
	CreatedAt    time.Time                `json:",omitempty"`
	ModifiedAt   time.Time                `json:",omitempty"`
	Icon         *Icon                    `json:",omitempty"`
	Tags         []string                 `json:",omitempty"`
}

// CommonServiceDNSSettings type of CommonServiceDnsSettings
type CommonServiceDNSSettings struct {
	DNS DNSRecordSets `json:",omitempty"`
}

// CommonServiceDNSStatus type of CommonServiceDNSStatus
type CommonServiceDNSStatus struct {
	Zone string   `json:",omitempty"`
	NS   []string `json:",omitempty"`
}

// CommonServiceDNSProvider type of CommonServiceDNSProvider
type CommonServiceDNSProvider struct {
	Class string `json:",omitempty"`
}

// CreateNewDNSCommonServiceItem Create new CommonServiceDNSItem
func CreateNewDNSCommonServiceItem(zoneName string) *CommonServiceDNSItem {
	return &CommonServiceDNSItem{
		Resource: &Resource{ID: ""},
		Name:     zoneName,
		Status: CommonServiceDNSStatus{
			Zone: zoneName,
		},
		Provider: CommonServiceDNSProvider{
			Class: "dns",
		},
		Settings: CommonServiceDNSSettings{
			DNS: DNSRecordSets{},
		},
	}

}

// HasDNSRecord return has record
func (d *CommonServiceDNSItem) HasDNSRecord() bool {
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
