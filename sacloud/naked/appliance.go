package naked

import "github.com/sacloud/libsacloud-v2/sacloud/types"

// ApplianceRemark アプライアンスの設定/ステータスなど
//
// Appliance.Remarkを表現する
type ApplianceRemark struct {
	Zone    *ApplianceRemarkZone     `json:",omitempty" yaml:"zone,omitempty" structs:",omitempty"`
	Switch  *ApplianceRemarkSwitch   `json:",omitempty" yaml:"switch,omitempty" structs:",omitempty"`
	VRRP    *ApplianceVRRP           `json:",omitempty" yaml:"vrrp,omitempty" structs:",omitempty"`
	Network *ApplianceRemarkNetwork  `json:",omitempty" yaml:"network,omitempty" structs:",omitempty"`
	Servers []*ApplianceRemarkServer `json:",omitempty" yaml:"servers,omitempty" structs:",omitempty"`
	Plan    *AppliancePlan           `json:",omitempty" yaml:"plan,omitempty" structs:",omitempty"`
}

// AppliancePlan アプライアンスプラン
type AppliancePlan struct {
	ID types.ID `json:",omitempty" yaml:"id,omitempty" structs:",omitempty"`
}

// ApplianceVRRP アプライアンスのVRRPの設定
type ApplianceVRRP struct {
	VRID int `json:",omitempty" yaml:"vrid,omitempty" structs:",omitempty"`
}

// ApplianceRemarkNetwork Appliance ネットワーク設定
type ApplianceRemarkNetwork struct {
	DefaultRoute   string `json:",omitempty" yaml:"default_route,omitempty" structs:",omitempty"`
	NetworkMaskLen int    `json:",omitempty" yaml:"network_mask_len,omitempty" structs:",omitempty"`
}

// ApplianceRemarkServer Applianceの稼働している仮想サーバのIPアドレス
type ApplianceRemarkServer struct {
	IPAddress string `json:",omitempty" yaml:"ipaddress,omitempty" structs:",omitempty"`
}

// ApplianceRemarkSwitch Applianceに接続されているスイッチのID
type ApplianceRemarkSwitch struct {
	ID types.ID `json:",omitempty" yaml:"id,omitempty" structs:",omitempty"`
}

// ApplianceRemarkZone Applianceの属するゾーンのID
type ApplianceRemarkZone struct {
	ID types.ID `json:",omitempty" yaml:"id,omitempty" structs:",omitempty"`
}
