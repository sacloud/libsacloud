package naked

import (
	"time"

	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

// Bridge ブリッジ
type Bridge struct {
	ID           types.ID          `json:",omitempty" yaml:"id,omitempty" structs:",omitempty"`
	Name         string            `json:",omitempty" yaml:"name,omitempty" structs:",omitempty"`
	Description  string            `json:",omitempty" yaml:"description,omitempty" structs:",omitempty"`
	ServiceClass string            `json:",omitempty" yaml:"service_class,omitempty" structs:",omitempty"`
	CreatedAt    *time.Time        `json:",omitempty" yaml:"created_at,omitempty" structs:",omitempty"`
	Region       *Region           `json:",omitempty" yaml:"region,omitempty" structs:",omitempty"`
	SwitchInZone *BridgeSwitchInfo `json:",omitempty" yaml:"switch_in_zone,omitempty" structs:",omitempty"`
}

// BridgeInfo ブリッジに接続されているスイッチの情報
type BridgeInfo struct {
	Switches []*Switch `json:",omitempty" yaml:"switches,omitempty" structs:",omitempty"`
}

// BridgeSwitchInfo ゾーン内での接続スイッチ情報
type BridgeSwitchInfo struct {
	ID             types.ID     `json:",omitempty" yaml:"id,omitempty" structs:",omitempty"`
	Scope          types.EScope `json:",omitempty" yaml:"scope,omitempty" structs:",omitempty"`
	Name           string       `json:",omitempty" yaml:"name,omitempty" structs:",omitempty"`
	ServerCount    int          `json:",omitempty" yaml:"server_count,omitempty" structs:",omitempty"`
	ApplianceCount int          `json:",omitempty" yaml:"appliance_count,omitempty" structs:",omitempty"`
}
