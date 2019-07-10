package naked

import (
	"encoding/json"

	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

// ApplianceRemark アプライアンスの設定/ステータスなど
//
// Appliance.Remarkを表現する
type ApplianceRemark struct {
	Zone    *ApplianceRemarkZone    `json:",omitempty" yaml:"zone,omitempty" structs:",omitempty"`
	Switch  *ApplianceRemarkSwitch  `json:",omitempty" yaml:"switch,omitempty" structs:",omitempty"`
	VRRP    *ApplianceVRRP          `json:",omitempty" yaml:"vrrp,omitempty" structs:",omitempty"`
	Network *ApplianceRemarkNetwork `json:",omitempty" yaml:"network,omitempty" structs:",omitempty"`
	Servers ApplianceRemarkServers  `yaml:"servers"`
	Plan    *AppliancePlan          `json:",omitempty" yaml:"plan,omitempty" structs:",omitempty"`
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

// ApplianceRemarkServers Applianceの稼働している仮想サーバのIPアドレス
type ApplianceRemarkServers []*ApplianceRemarkServer

// ApplianceRemarkServer Applianceの稼働している仮想サーバのIPアドレス
type ApplianceRemarkServer struct {
	IPAddress string `json:",omitempty" yaml:"ip_address,omitempty" structs:",omitempty"`
}

// ApplianceRemarkSwitch Applianceに接続されているスイッチのID
type ApplianceRemarkSwitch struct {
	ID    types.ID     `json:",omitempty" yaml:"id,omitempty" structs:",omitempty"`
	Scope types.EScope `json:",omitempty" yaml:"scope,omitempty" structs:",omitempty"`
}

// ApplianceRemarkZone Applianceの属するゾーンのID
type ApplianceRemarkZone struct {
	ID types.ID `json:",omitempty" yaml:"id,omitempty" structs:",omitempty"`
}

// UnmarshalJSON 配列/オブジェクトが混在することへの対応
func (s *ApplianceRemarkNetwork) UnmarshalJSON(b []byte) error {
	if string(b) == "[]" {
		return nil
	}
	type alias ApplianceRemarkNetwork

	var a alias
	if err := json.Unmarshal(b, &a); err != nil {
		return err
	}
	*s = ApplianceRemarkNetwork(a)
	return nil
}

// UnmarshalJSON 配列/オブジェクトが混在することへの対応
func (s *ApplianceRemarkServer) UnmarshalJSON(b []byte) error {
	if string(b) == "[]" {
		return nil
	}
	type alias ApplianceRemarkServer

	var a alias
	if err := json.Unmarshal(b, &a); err != nil {
		return err
	}
	*s = ApplianceRemarkServer(a)
	return nil
}

// UnmarshalJSON 配列/オブジェクトが混在することへの対応
func (s *ApplianceRemarkServers) UnmarshalJSON(b []byte) error {
	if string(b) == "[[]]" {
		return nil
	}
	type alias ApplianceRemarkServers

	var a alias
	if err := json.Unmarshal(b, &a); err != nil {
		return err
	}
	*s = ApplianceRemarkServers(a)
	return nil
}

// MarshalJSON APIの要求するJSONフォーマットへの変換
//
// 値がからの場合に配列、かつ内部に空オブジェクトを指定する。(主にVPCルータへの対応)
func (s *ApplianceRemarkServers) MarshalJSON() ([]byte, error) {
	if s == nil || len(*s) == 0 {
		return []byte("[{}]"), nil
	}

	type alias ApplianceRemarkServers

	a := alias(*s)
	return json.Marshal(a)
}
