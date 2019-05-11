package naked

import (
	"time"

	"github.com/sacloud/libsacloud-v2/sacloud/types"
)

// NFSPlan NFSプラン
type NFSPlan struct {
	ID types.ID `json:",omitempty" yaml:"id,omitempty" structs:",omitempty"`
}

// NFSRemarkNetwork NFS ネットワーク設定
type NFSRemarkNetwork struct {
	DefaultRoute   string `json:",omitempty" yaml:"default_route,omitempty" structs:",omitempty"`
	NetworkMaskLen int    `json:",omitempty" yaml:"network_mask_len,omitempty" structs:",omitempty"`
}

// NFSRemarkPlan NFSプラン
type NFSRemarkPlan struct {
	ID types.ID `json:",omitempty" yaml:"id,omitempty" structs:",omitempty"`
}

// NFSRemarkServer NFSの稼働している仮想サーバのIPアドレス
type NFSRemarkServer struct {
	IPAddress string `json:",omitempty" yaml:"ipaddress,omitempty" structs:",omitempty"`
}

// NFSRemarkSwitch NFSに接続されているスイッチのID
type NFSRemarkSwitch struct {
	ID types.ID `json:",omitempty" yaml:"id,omitempty" structs:",omitempty"`
}

// NFSRemarkZone NFSの属するゾーンのID
type NFSRemarkZone struct {
	ID types.ID `json:",omitempty" yaml:"id,omitempty" structs:",omitempty"`
}

// NFSRemark NFSの設定/ステータス
type NFSRemark struct {
	Network *NFSRemarkNetwork  `json:",omitempty" yaml:"network,omitempty" structs:",omitempty"`
	Plan    *NFSRemarkPlan     `json:",omitempty" yaml:"plan,omitempty" structs:",omitempty"`
	Servers []*NFSRemarkServer `json:",omitempty" yaml:"servers,omitempty" structs:",omitempty"`
	Switch  *NFSRemarkSwitch   `json:",omitempty" yaml:"switch,omitempty" structs:",omitempty"`
	Zone    *NFSRemarkZone     `json:",omitempty" yaml:"zone,omitempty" structs:",omitempty"`
}

// NFS NFS
type NFS struct {
	ID           types.ID            `json:",omitempty" yaml:"id,omitempty" structs:",omitempty"`
	Name         string              `json:",omitempty" yaml:"name,omitempty" structs:",omitempty"`
	Description  string              `json:",omitempty" yaml:"description,omitempty" structs:",omitempty"`
	Tags         []string            `json:"" yaml:"tags"`
	Icon         *Icon               `json:",omitempty" yaml:"icon,omitempty" structs:",omitempty"`
	CreatedAt    *time.Time          `json:",omitempty" yaml:"created_at,omitempty" structs:",omitempty"`
	ModifiedAt   *time.Time          `json:",omitempty" yaml:"modified_at,omitempty" structs:",omitempty"`
	Availability types.EAvailability `json:",omitempty" yaml:"availability,omitempty" structs:",omitempty"`
	Class        string              `json:",omitempty" yaml:"class,omitempty" structs:",omitempty"`
	Instance     *Instance           `json:",omitempty" yaml:"instance,omitempty" structs:",omitempty"`
	Interfaces   []*Interface        `json:",omitempty" yaml:"interfaces,omitempty" structs:",omitempty"`
	Plan         *NFSPlan            `json:",omitempty" yaml:"plan,omitempty" structs:",omitempty"`
	Remark       *NFSRemark          `json:",omitempty" yaml:"remark,omitempty" structs:",omitempty"`
	ServiceClass string              `json:",omitempty" yaml:"service_class,omitempty" structs:",omitempty"`
	Switch       *Switch             `json:",omitempty" yaml:"switch,omitempty" structs:",omitempty"`
}
