package sacloud

// MobileGateway モバイルゲートウェイ
type MobileGateway struct {
	*Appliance // アプライアンス共通属性

	Remark   *MobileGatewayRemark   `json:",omitempty"` // リマーク
	Settings *MobileGatewaySettings `json:",omitempty"` // モバイルゲートウェイ設定
}

// MobileGatewayRemark リマーク
type MobileGatewayRemark struct {
	*ApplianceRemarkBase
	// TODO Zone
	//Zone *Resource
}

// MobileGatewaySettings モバイルゲートウェイ設定
type MobileGatewaySettings struct {
	MobileGateway *MobileGatewaySetting `json:",omitempty"` // モバイルゲートウェイ設定リスト
}

// MobileGatewaySetting モバイルゲートウェイ設定
type MobileGatewaySetting struct {
	InternetConnection *MGWInternetConnection `json:",omitempty"` // インターネット接続
	Interfaces         []*MGWInterface        `json:",omitempty"` // インターフェース
	StaticRoutes       []*MGWStaticRoute      `json:",omitempty"` // スタティックルート
}

// MGWInternetConnection インターネット接続
type MGWInternetConnection struct {
	Enabled string `json:",omitempty"`
}

// MGWInterface インターフェース
type MGWInterface struct {
	IPAddress      []string `json:",omitempty"`
	NetworkMaskLen int      `json:",omitempty"`
}

// MGWStaticRoute スタティックルート
type MGWStaticRoute struct {
	Prefix  string `json:",omitempty"`
	NextHop string `json:",omitempty"`
}

// MobileGatewayPlan モバイルゲートウェイプラン
type MobileGatewayPlan int

var (
	// MobileGatewayPlanStandard スタンダードプラン // TODO 正式名称不明なため暫定の名前
	MobileGatewayPlanStandard = MobileGatewayPlan(1)
)

// CreateMobileGatewayValue モバイルゲートウェイ作成用パラメーター
type CreateMobileGatewayValue struct {
	Name        string   // 名称
	Description string   // 説明
	Tags        []string // タグ
	IconID      int64    // アイコン
}

// CreateNewMobileGateway モバイルゲートウェイ作成
func CreateNewMobileGateway(values *CreateMobileGatewayValue, setting *MobileGatewaySetting) (*MobileGateway, error) {

	lb := &MobileGateway{
		Appliance: &Appliance{
			Class:           "mobilegateway",
			propName:        propName{Name: values.Name},
			propDescription: propDescription{Description: values.Description},
			propTags:        propTags{Tags: values.Tags},
			propPlanID:      propPlanID{Plan: &Resource{ID: int64(MobileGatewayPlanStandard)}},
			propIcon: propIcon{
				&Icon{
					Resource: NewResource(values.IconID),
				},
			},
		},
		Remark: &MobileGatewayRemark{
			ApplianceRemarkBase: &ApplianceRemarkBase{
				Switch: &ApplianceRemarkSwitch{
					propScope: propScope{
						Scope: "shared",
					},
				},
				Servers: []interface{}{
					nil,
				},
			},
		},
		Settings: &MobileGatewaySettings{
			MobileGateway: setting,
		},
	}

	return lb, nil
}

// MobileGatewayResolver DNS登録用パラメータ
type MobileGatewayResolver struct {
	SimGroup *MobileGatewaySIMGroup `json:"sim_group,omitempty"`
}

// MobileGatewaySIMGroup DNS登録用SIMグループ値
type MobileGatewaySIMGroup struct {
	DNS1 string `json:"dns_1,omitempty"`
	DNS2 string `json:"dns_2,omitempty"`
}
