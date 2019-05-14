package naked

import (
	"time"

	"github.com/sacloud/libsacloud-v2/sacloud/types"
)

// LoadBalancer ロードバランサ
type LoadBalancer struct {
	ID           types.ID              `json:",omitempty" yaml:"id,omitempty" structs:",omitempty"`
	Name         string                `json:",omitempty" yaml:"name,omitempty" structs:",omitempty"`
	Description  string                `json:",omitempty" yaml:"description,omitempty" structs:",omitempty"`
	Tags         []string              `json:"" yaml:"tags"`
	Icon         *Icon                 `json:",omitempty" yaml:"icon,omitempty" structs:",omitempty"`
	CreatedAt    *time.Time            `json:",omitempty" yaml:"created_at,omitempty" structs:",omitempty"`
	ModifiedAt   *time.Time            `json:",omitempty" yaml:"modified_at,omitempty" structs:",omitempty"`
	Availability types.EAvailability   `json:",omitempty" yaml:"availability,omitempty" structs:",omitempty"`
	Class        string                `json:",omitempty" yaml:"class,omitempty" structs:",omitempty"`
	ServiceClass string                `json:",omitempty" yaml:"service_class,omitempty" structs:",omitempty"`
	Plan         *AppliancePlan        `json:",omitempty" yaml:"plan,omitempty" structs:",omitempty"`
	Instance     *Instance             `json:",omitempty" yaml:"instance,omitempty" structs:",omitempty"`
	Interfaces   []*Interface          `json:",omitempty" yaml:"interfaces,omitempty" structs:",omitempty"`
	Switch       *Switch               `json:",omitempty" yaml:"switch,omitempty" structs:",omitempty"`
	Settings     *LoadBalancerSettings `json:",omitempty" yaml:"settings,omitempty" structs:",omitempty"`
	SettingsHash string                `json:",omitempty" yaml:"settings_hash,omitempty" structs:",omitempty"`
	Remark       *ApplianceRemark      `json:",omitempty" yaml:"remark,omitempty" structs:",omitempty"`
}

// LoadBalancerSettings ロードバランサの設定
type LoadBalancerSettings struct {
	LoadBalancer []*LoadBalancerSetting `json:",omitempty" yaml:"load_balancer,omitempty" structs:",omitempty"`
}

// LoadBalancerSetting ロードバランサの設定
type LoadBalancerSetting struct {
	VirtualIPAddress string                           `json:",omitempty" yaml:"virtual_ip_address,omitempty" structs:",omitempty"`
	Port             types.StringNumber               `json:",omitempty" yaml:"port,omitempty" structs:",omitempty"`
	DelayLoop        types.StringNumber               `json:",omitempty" yaml:"delay_loop,omitempty" structs:",omitempty"`
	SorryServer      string                           `json:",omitempty" yaml:"sorry_server,omitempty" structs:",omitempty"`
	Description      string                           `json:",omitempty" yaml:"description,omitempty" structs:",omitempty"`
	Servers          []*LoadBalancerDestinationServer `json:",omitempty" yaml:"servers,omitempty" structs:",omitempty"`
}

// LoadBalancerDestinationServer ロードバランサ配下の実サーバ
type LoadBalancerDestinationServer struct {
	IPAddress   string             `json:",omitempty" yaml:"ip_address,omitempty" structs:",omitempty"`
	Port        types.StringNumber `json:",omitempty" yaml:"port,omitempty" structs:",omitempty"`
	Enabled     types.StringFlag   `json:",omitempty" yaml:"enabled,omitempty" structs:",omitempty"`
	HealthCheck *HealthCheck       `json:",omitempty" yaml:"health_check,omitempty" structs:",omitempty"`
}

// LoadBalancerStatus ロードバランサのステータス
type LoadBalancerStatus struct {
	VirtualIPAddress string                      `json:",omitempty" yaml:"virtual_ip_address,omitempty" structs:",omitempty"`
	Port             types.StringNumber          `json:",omitempty" yaml:"port,omitempty" structs:",omitempty"`
	CPS              types.StringNumber          `json:",omitempty" yaml:"cps,omitempty" structs:",omitempty"`
	Servers          []*LoadBalancerServerStatus `json:",omitempty" yaml:"servers,omitempty" structs:",omitempty"`
}

// LoadBalancerServerStatus ロードバランサの実サーバのステータス
type LoadBalancerServerStatus struct {
	ActiveConn types.StringNumber          `json:",omitempty" yaml:"active_conn,omitempty" structs:",omitempty"`
	Status     types.EServerInstanceStatus `json:",omitempty" yaml:"status,omitempty" structs:",omitempty"`
	IPAddress  string                      `json:",omitempty" yaml:"ip_address,omitempty" structs:",omitempty"`
	Port       types.StringNumber          `json:",omitempty" yaml:"port,omitempty" structs:",omitempty"`
	CPS        types.StringNumber          `json:",omitempty" yaml:"cps,omitempty" structs:",omitempty"`
}
