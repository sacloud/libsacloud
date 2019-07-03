package naked

import "github.com/sacloud/libsacloud/v2/sacloud/types"

// BundleInfo バンドル情報
type BundleInfo struct {
	ID           types.ID `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	HostClass    string   `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	ServiceClass string   `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
}
