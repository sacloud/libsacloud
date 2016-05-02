package sacloud

import "time"

// SimpleMonitor type of SimpleMonitor(CommonServiceItem)
type SimpleMonitor struct {
	*Resource
	Name         string
	Description  string                 `json:",omitempty"`
	Settings     *SimpleMonitorSettings `json:",omitempty"`
	Status       *SimpleMonitorStatus   `json:",omitempty"`
	ServiceClass string                 `json:",omitempty"`
	CreatedAt    time.Time              `json:",omitempty"`
	ModifiedAt   time.Time              `json:",omitempty"`
	Provider     *SimpleMonitorProvider `json:",omitempty"`
	Icon         *Icon                  `json:",omitempty"`
	Tags         []string               `json:",omitempty"`
}

// SimpleMonitorSettings type of SimpleMonitorSettings
type SimpleMonitorSettings struct {
	SimpleMonitor *SimpleMonitorSetting `json:",omitempty"`
}

// SimpleMonitorSetting type of SimpleMonitorSetting
type SimpleMonitorSetting struct {
	DelayLoop   int `json:",omitempty"`
	HealthCheck struct {
		Protocol     string `json:",omitempty"`
		Port         string `json:",omitempty"`
		Path         string `json:",omitempty"`
		Status       string `json:",omitempty"`
		QName        string `json:",omitempty"`
		ExpectedData string `json:",omitempty"`
	}
	Enabled     string `json:",omitempty"`
	NotifyEmail struct {
		Enabled string `json:",omitempty"`
	}
	NotifySlack struct {
		Enable string `json:",omitempty"`
	}
}

// SimpleMonitorStatus type of CommonServiceDNSStatus
type SimpleMonitorStatus struct {
	Target string `json:",omitempty"`
}

// SimpleMonitorProvider type of CommonServiceDNSProvider
type SimpleMonitorProvider struct {
	*NumberResource
	Class        string `json:",omitempty"`
	Name         string `json:",omitempty"`
	ServiceClass string `json:",omitempty"`
}

// CreateNewSimpleMonitor Create new CommonServiceSimpleMonitorItem
func CreateNewSimpleMonitor(target string) *SimpleMonitor {
	return &SimpleMonitor{
		Resource: &Resource{ID: ""},
		Name:     target,
		Provider: &SimpleMonitorProvider{
			Class: "simplemon",
		},
		Status: &SimpleMonitorStatus{
			Target: target,
		},
		Settings: &SimpleMonitorSettings{
			SimpleMonitor: &SimpleMonitorSetting{},
		},
	}

}
