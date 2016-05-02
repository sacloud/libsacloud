package sacloud

import "time"

type appliance struct {
	*Resource
	Class       string `json:",omitempty"`
	Name        string `json:",omitempty"`
	Description string `json:",omitempty"`
	Plan        *NumberResource
	//Settings
	SettingHash string `json:",omitempty"`
	//Remark      *ApplianceRemark `json:",omitempty"`
	*EAvailability
	Instance struct {
		Status          string    `json:",omitempty"`
		StatusChangedAt time.Time `json:",omitempty"`
	} `json:",omitempty"`
	ServiceClass string      `json:",omitempty"`
	CreatedAt    time.Time   `json:",omitempty"`
	Icon         *Icon       `json:",omitempty"`
	Switch       *Switch     `json:",omitempty"`
	Interfaces   []Interface `json:",omitempty"`
	Tags         []string    `json:",omitempty"`
}

//HACK Appliance:Zone.IDがRoute/LoadBalancerの場合でデータ型が異なるため
//それぞれのstruct定義でZoneだけ上書きした構造体を定義して使う

type applianceRemarkBase struct {
	Servers []interface{}
	Switch  *struct {
		ID    string `json:",omitempty"`
		Scope string `json:",omitempty"`
	} `json:",omitempty"`
	//Zone *NumberResource `json:",omitempty"`
	VRRP *struct {
		VRID int `json:",omitempty"`
	} `json:",omitempty"`
	Network *struct {
		NetworkMaskLen int    `json:",omitempty"`
		DefaultRoute   string `json:",omitempty"`
	} `json:",omitempty"`
	Plan *NumberResource
}
