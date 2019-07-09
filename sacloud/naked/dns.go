package naked

import (
	"time"

	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

// DNS DNSゾーン
type DNS struct {
	ID           types.ID            `json:",omitempty" yaml:"id,omitempty" structs:",omitempty"`
	Name         string              `json:",omitempty" yaml:"name,omitempty" structs:",omitempty"`
	Description  string              `json:",omitempty" yaml:"description,omitempty" structs:",omitempty"`
	Tags         []string            `json:"" yaml:"tags"`
	Icon         *Icon               `json:",omitempty" yaml:"icon,omitempty" structs:",omitempty"`
	CreatedAt    *time.Time          `json:",omitempty" yaml:"created_at,omitempty" structs:",omitempty"`
	ModifiedAt   *time.Time          `json:",omitempty" yaml:"modified_at,omitempty" structs:",omitempty"`
	Availability types.EAvailability `json:",omitempty" yaml:"availability,omitempty" structs:",omitempty"`
	ServiceClass string              `json:",omitempty" yaml:"service_class,omitempty" structs:",omitempty"`
	Provider     *Provider           `json:",omitempty" yaml:"provider,omitempty" structs:",omitempty"`
	Settings     *DNSSettings        `json:",omitempty" yaml:"settings,omitempty" structs:",omitempty"`
	SettingsHash string              `json:",omitempty" yaml:"settings_hash,omitempty" structs:",omitempty"`
	Status       *DNSStatus          `json:",omitempty" yaml:"status,omitempty" structs:",omitempty"`
}

// DNSStatus DNSステータス
type DNSStatus struct {
	Zone string   `json:",omitempty" yaml:"zone,omitempty" structs:",omitempty"`
	NS   []string `json:",omitempty" yaml:"ns,omitempty" structs:",omitempty"`
}

// DNSSettings DNSセッティング
type DNSSettings struct {
	DNS *DNSSetting `json:",omitempty" yaml:"dns,omitempty" structs:",omitempty"`
}

// DNSSetting DNSセッティング
type DNSSetting struct {
	ResourceRecordSets []*DNSRecord `json:",omitempty" yaml:"resource_record_sets,omitempty" structs:",omitempty"`
}

// DNSRecord DNSレコード
type DNSRecord struct {
	Name  string               `json:",omitempty" yaml:"name,omitempty" structs:",omitempty"`  // ホスト名
	Type  types.EDNSRecordType `json:",omitempty" yaml:"type,omitempty" structs:",omitempty"`  // レコードタイプ
	RData string               `json:",omitempty" yaml:"rdata,omitempty" structs:",omitempty"` // レコードデータ
	TTL   int                  `json:",omitempty" yaml:"ttl,omitempty" structs:",omitempty"`   // TTL
}
