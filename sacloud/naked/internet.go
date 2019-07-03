package naked

import (
	"time"

	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

// Internet ルータ+スイッチのルータ部分
type Internet struct {
	ID             types.ID     `json:",omitempty" yaml:"id,omitempty" structs:",omitempty"`
	Name           string       `json:",omitempty" yaml:"name,omitempty" structs:",omitempty"`
	Description    string       `json:",omitempty" yaml:"description,omitempty" structs:",omitempty"`
	Tags           []string     `json:"" yaml:"tags"`
	Icon           *Icon        `json:",omitempty" yaml:"icon,omitempty" structs:",omitempty"`
	CreatedAt      *time.Time   `json:",omitempty" yaml:"created_at,omitempty" structs:",omitempty"`
	Scope          types.EScope `json:",omitempty" yaml:"scope,omitempty" structs:",omitempty"`
	ServiceClass   string       `json:",omitempty" yaml:"service_class,omitempty" structs:",omitempty"`
	Switch         *Switch      `json:",omitempty" yaml:"switch,omitempty" structs:",omitempty"`
	BandWidthMbps  int          `json:",omitempty" yaml:"band_width_mbps,omitempty" structs:",omitempty"`
	NetworkMaskLen int          `json:",omitempty" yaml:"network_mask_len,omitempty" structs:",omitempty"`
}

// SubnetOperationRequest サブネット追加時のリクエストパラメータ
type SubnetOperationRequest struct {
	NetworkMaskLen int    `json:",omitempty" yaml:"network_mask_len,omitempty" structs:",omitempty"`
	NextHop        string `json:",omitempty" yaml:"next_hop,omitempty" structs:",omitempty"`
}
