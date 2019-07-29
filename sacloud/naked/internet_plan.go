package naked

import "github.com/sacloud/libsacloud/v2/sacloud/types"

// InternetPlan サーバープラン
type InternetPlan struct {
	ID            types.ID            `json:",omitempty" yaml:"id,omitempty" structs:",omitempty"`
	Name          string              `json:",omitempty" yaml:"name,omitempty" structs:",omitempty"`
	BandWidthMbps int                 `json:",omitempty" yaml:"band_width_mbps,omitempty" structs:",omitempty"`
	ServiceClass  string              `json:",omitempty" yaml:"service_class,omitempty" structs:",omitempty"`
	Availability  types.EAvailability `json:",omitempty" yaml:"availability,omitempty" structs:",omitempty"`
}
