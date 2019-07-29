package naked

import (
	"encoding/json"
	"time"

	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

// DNS DNSゾーン
type DNS struct {
	ID           types.ID            `json:",omitempty" yaml:"id,omitempty" structs:",omitempty"`
	Name         string              `json:",omitempty" yaml:"name,omitempty" structs:",omitempty"`
	Description  string              `yaml:"description"`
	Tags         types.Tags          `yaml:"tags"`
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
	ResourceRecordSets []*DNSRecord `yaml:"resource_record_sets"`
}

// MarshalJSON nullの場合に空配列を出力するための実装
func (ds DNSSetting) MarshalJSON() ([]byte, error) {
	if ds.ResourceRecordSets == nil {
		ds.ResourceRecordSets = make([]*DNSRecord, 0)
	}
	type alias DNSSetting
	tmp := alias(ds)
	return json.Marshal(&tmp)
}

// DNSRecord DNSレコード
type DNSRecord struct {
	Name  string               `json:",omitempty" yaml:"name,omitempty" structs:",omitempty"`  // ホスト名
	Type  types.EDNSRecordType `json:",omitempty" yaml:"type,omitempty" structs:",omitempty"`  // レコードタイプ
	RData string               `json:",omitempty" yaml:"rdata,omitempty" structs:",omitempty"` // レコードデータ
	TTL   int                  `json:",omitempty" yaml:"ttl,omitempty" structs:",omitempty"`   // TTL
}
