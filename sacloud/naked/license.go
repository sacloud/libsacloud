package naked

import (
	"time"

	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

// License ライセンス
type License struct {
	ID          types.ID     `json:",omitempty" yaml:"id,omitempty" structs:",omitempty"`
	Name        string       `json:",omitempty" yaml:"name,omitempty" structs:",omitempty"`
	Description string       `yaml:"description"`
	CreatedAt   *time.Time   `json:",omitempty" yaml:"created_at,omitempty" structs:",omitempty"`
	ModifiedAt  *time.Time   `json:",omitempty" yaml:"modified_at,omitempty" structs:",omitempty"`
	LicenseInfo *LicenseInfo `json:",omitempty" yaml:"license_info,omitempty" structs:",omitempty"` // ライセンス情報
}

// LicenseInfo ライセンスプラン
type LicenseInfo struct {
	ID           types.ID   `json:",omitempty" yaml:"id,omitempty" structs:",omitempty"`
	Name         string     `json:",omitempty" yaml:"name,omitempty" structs:",omitempty"`
	ServiceClass string     `json:",omitempty" yaml:"service_class,omitempty" structs:",omitempty"`
	CreatedAt    *time.Time `json:",omitempty" yaml:"created_at,omitempty" structs:",omitempty"`
	ModifiedAt   *time.Time `json:",omitempty" yaml:"modified_at,omitempty" structs:",omitempty"`
	TermsOfUse   string     `json:",omitempty" yaml:"terms_of_use,omitempty" structs:",omitempty"` // 利用規約
}
