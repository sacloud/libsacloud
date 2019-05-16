package naked

import "github.com/sacloud/libsacloud-v2/sacloud/types"

// PacketFilter パケットフィルタ
//
// TODO 後で
type PacketFilter struct {
	ID                  types.ID `json:",omitempty" yaml:"id,omitempty" structs:",omitempty"`
	Name                string   `json:",omitempty" yaml:"name,omitempty" structs:",omitempty"`
	RequiredHostVersion int      `json:",omitempty" yaml:"require_host_version,omitempty" structs:",omitempty"`
}
