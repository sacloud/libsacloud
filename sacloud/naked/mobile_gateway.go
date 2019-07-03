package naked

import (
	"encoding/json"
	"strings"
	"time"

	"github.com/sacloud/libsacloud/sacloud/types"
)

// MobileGateway モバイルゲートウェイ
type MobileGateway struct {
	ID           int64                  `json:",omitempty" yaml:"id,omitempty" structs:",omitempty"`
	Class        string                 `json:",omitempty" yaml:"class,omitempty" structs:",omitempty"`
	Name         string                 `json:",omitempty" yaml:"name,omitempty" structs:",omitempty"`
	Tags         []string               `yaml:"tags"`
	Description  string                 `json:",omitempty" yaml:"description,omitempty" structs:",omitempty"`
	Plan         *AppliancePlan         `json:",omitempty" yaml:"plan,omitempty" structs:",omitempty"`
	Settings     *MobileGatewaySettings `json:",omitempty" yaml:"settings,omitempty" structs:",omitempty"`
	SettingsHash string                 `json:",omitempty" yaml:"settings_hash,omitempty" structs:",omitempty"`
	Remark       *ApplianceRemark       `json:",omitempty" yaml:"remark,omitempty" structs:",omitempty"`
	Availability types.EAvailability    `json:",omitempty" yaml:"availability,omitempty" structs:",omitempty"`
	Instance     *Instance              `json:",omitempty" yaml:"instance,omitempty" structs:",omitempty"`
	ServiceClass string                 `json:",omitempty" yaml:"service_class,omitempty" structs:",omitempty"`
	CreatedAt    time.Time              `json:",omitempty" yaml:"created_at,omitempty" structs:",omitempty"`
	Icon         *Icon                  `json:",omitempty" yaml:"icon,omitempty" structs:",omitempty"`
	Switch       *Switch                `json:",omitempty" yaml:"switch,omitempty" structs:",omitempty"`
	Interfaces   []*Interface           `json:",omitempty" yaml:"interfaces,omitempty" structs:",omitempty"`
}

// MobileGatewaySettings モバイルゲートウェイ セッティング
type MobileGatewaySettings struct {
	MobileGateway *MobileGatewaySetting `json:",omitempty" yaml:"mobile_gateway,omitempty" structs:",omitempty"`
}

// MobileGatewaySetting モバイルゲートウェイ セッティング
type MobileGatewaySetting struct {
	Interfaces               []*MobileGatewayInterface              `json:",omitempty" yaml:"interfaces,omitempty" structs:",omitempty"`
	InternetConnection       *MobileGatewayInternetConnection       `json:",omitempty" yaml:"internet_connection,omitempty" structs:",omitempty"`
	StaticRoutes             []*MobileGatewayStaticRoute            `json:",omitempty" yaml:"static_routes,omitempty" structs:",omitempty"`
	InterDeviceCommunication *MobileGatewayInterDeviceCommunication `json:",omitempty" yaml:"inter_device_communication,omitempty" structs:",omitempty"`
}

// MobileGatewayInterDeviceCommunication デバイス間通信
type MobileGatewayInterDeviceCommunication struct {
	Enabled string `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
}

// MobileGatewayInternetConnection インターネット接続
type MobileGatewayInternetConnection struct {
	Enabled string `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
}

// MobileGatewayInterface インターフェース
type MobileGatewayInterface struct {
	IPAddress      []string `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	NetworkMaskLen int      `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
}

// MobileGatewayStaticRoute スタティックルート
type MobileGatewayStaticRoute struct {
	Prefix  string `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	NextHop string `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
}

// MobileGatewayResolver DNS登録用パラメータ
type MobileGatewayResolver struct {
	SimGroup *MobileGatewaySIMGroup `json:"sim_group,omitempty" yaml:"sim_group,omitempty" structs:",omitempty"`
}

// MobileGatewaySIMGroup DNS登録用SIMグループ値
type MobileGatewaySIMGroup struct {
	DNS1 string `json:"dns_1,omitempty" yaml:"dns_1,omitempty" structs:",omitempty"`
	DNS2 string `json:"dns_2,omitempty" yaml:"dns_2,omitempty" structs:",omitempty"`
}

// UnmarshalJSON JSONアンマーシャル(配列、オブジェクトが混在するためここで対応)
func (m *MobileGatewaySIMGroup) UnmarshalJSON(data []byte) error {
	targetData := strings.Replace(strings.Replace(string(data), " ", "", -1), "\n", "", -1)
	if targetData == `[]` {
		return nil
	}

	type alias MobileGatewaySIMGroup
	tmp := alias{}
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*m = MobileGatewaySIMGroup(tmp)
	return nil
}

// MobileGatewaySIMRoute SIルート
type MobileGatewaySIMRoute struct {
	ICCID      string `json:"iccid,omitempty" yaml:"iccid,omitempty" structs:",omitempty"`
	Prefix     string `json:"prefix,omitempty" yaml:"prefix,omitempty" structs:",omitempty"`
	ResourceID string `json:"resource_id,omitempty" yaml:"resource_id,omitempty" structs:",omitempty"`
}

// MobileGatewaySIMRoutes SIMルート一覧
type MobileGatewaySIMRoutes struct {
	SIMRoutes []*MobileGatewaySIMRoute `json:"sim_routes" yaml:"sim_routes,omitempty" structs:",omitempty"`
}

// TrafficStatus トラフィックコントロール 当月通信量
type TrafficStatus struct {
	UplinkBytes    types.StringNumber `json:"uplink_bytes,omitempty" yaml:"uplink_bytes,omitempty" structs:",omitempty"`
	DownlinkBytes  types.StringNumber `json:"downlink_bytes,omitempty" yaml:"downlink_bytes,omitempty" structs:",omitempty"`
	TrafficShaping bool               `json:"traffic_shaping" yaml:"traffic_shaping,omitempty" structs:",omitempty"` // 帯域制限
}

// UnmarshalJSON JSONアンマーシャル(uint64文字列対応)
func (s *TrafficStatus) UnmarshalJSON(data []byte) error {
	type alias TrafficStatus
	tmp := alias{}
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*s = TrafficStatus(tmp)
	return nil
}

// TrafficMonitoringConfig トラフィックコントロール 設定
type TrafficMonitoringConfig struct {
	TrafficQuotaInMB     int                           `json:"traffic_quota_in_mb,omitempty" yaml:"traffic_quota_in_mb" structs:",omitempty"`
	BandWidthLimitInKbps int                           `json:"bandwidth_limit_in_kbps,omitempty" yaml:"bandwidth_limit_in_kbps" structs:",omitempty"`
	EMailConfig          *TrafficMonitoringNotifyEmail `json:"email_config,omitempty" yaml:"email_config,omitempty" structs:",omitempty"`
	SlackConfig          *TrafficMonitoringNotifySlack `json:"slack_config,omitempty" yaml:"slack_config,omitempty" structs:",omitempty"`
	AutoTrafficShaping   bool                          `json:"auto_traffic_shaping,omitempty" yaml:"auto_traffic_shaping,omitempty" structs:",omitempty"`
}

// TrafficMonitoringNotifyEmail トラフィックコントロール通知設定
type TrafficMonitoringNotifyEmail struct {
	Enabled bool `json:"enabled" yaml:"enabled"` // 有効/無効
}

// TrafficMonitoringNotifySlack トラフィックコントロール通知設定
type TrafficMonitoringNotifySlack struct {
	Enabled             bool   `json:"enabled" yaml:"enabled"`                         // 有効/無効
	IncomingWebhooksURL string `json:"slack_url,omitempty" yaml:"slack_url,omitempty"` // Slack通知の場合のWebhook URL
}
