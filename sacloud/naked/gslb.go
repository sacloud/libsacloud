package naked

import (
	"time"

	"github.com/sacloud/libsacloud-v2/sacloud/types"
)

// GSLB GSLB
type GSLB struct {
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
	Settings     *GSLBSettings       `json:",omitempty" yaml:"settings,omitempty" structs:",omitempty"`
	SettingsHash string              `json:",omitempty" yaml:"settings_hash,omitempty" structs:",omitempty"`
	Status       *GSLBStatus         `json:",omitempty" yaml:"status,omitempty" structs:",omitempty"`
}

// GSLBSettings GSLBの設定
type GSLBSettings struct {
	GSLB *GSLBSetting `json:",omitempty" yaml:"gslb,omitempty" structs:",omitempty"`
}

// GSLBSetting GSLBの設定
type GSLBSetting struct {
	DelayLoop   int              `json:",omitempty" yaml:"delay_loop,omitempty" structs:",omitempty"`
	HealthCheck *HealthCheck     `json:",omitempty" yaml:"health_check,omitempty" structs:",omitempty"`
	Weighted    types.StringFlag `yaml:"weighted"`
	Servers     []*GSLBServer    `yaml:"servers"`
	SorryServer string           `json:",omitempty" yaml:",omitempty" structs:",omitempty"` // ソーリーサーバー
}

// GSLBServer GSLB配下のサーバー
type GSLBServer struct {
	IPAddress string             `json:",omitempty" yaml:"ip_address,omitempty" structs:",omitempty"` // IPアドレス
	Enabled   types.StringFlag   `yaml:"enabled" `                                                    // 有効/無効
	Weight    types.StringNumber `json:",omitempty" yaml:"weight,omitempty" structs:",omitempty"`     // ウェイト
}

// GSLBStatus GSLBステータス
type GSLBStatus struct {
	FQDN string `json:",omitempty" yaml:"fqdn,omitempty" structs:",omitempty"`
}
