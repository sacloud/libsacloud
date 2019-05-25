package naked

import (
	"encoding/json"
	"time"

	"github.com/sacloud/libsacloud-v2/sacloud/types"
)

// VPCRouter VPCルータ
type VPCRouter struct {
	ID           types.ID            `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	Class        string              `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	Name         string              `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	Description  string              `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	Plan         *AppliancePlan      `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	Settings     *VPCRouterSettings  `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	SettingsHash string              `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	Remark       *ApplianceRemark    `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	Availability types.EAvailability `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	Instance     *Instance           `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	ServiceClass string              `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	CreatedAt    *time.Time          `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	Icon         *Icon               `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	Switch       *Switch             `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	Interfaces   Interfaces          `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	Tags         []string            `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
}

// VPCRouterSettings VPCルータ 設定
type VPCRouterSettings struct {
	Router *VPCRouterSetting `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
}

// VPCRouterSetting VPCルータ 設定
type VPCRouterSetting struct {
	InternetConnection *VPCRouterInternetConnection `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	Interfaces         []*VPCRouterInterface        `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	VRID               int                          `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	StaticNAT          *VPCRouterStaticNAT          `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	Firewall           *VPCRouterFirewall           `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	DHCPServer         *VPCRouterDHCPServer         `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	DHCPStaticMapping  *VPCRouterDHCPStaticMappings `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	PPTPServer         *VPCRouterPPTPServer         `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	L2TPIPsecServer    *VPCRouterL2TPIPsecServer    `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	RemoteAccessUsers  *VPCRouterRemoteAccessUsers  `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	SiteToSiteIPsecVPN *VPCRouterSiteToSiteIPsecVPN `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	StaticRoutes       *VPCRouterStaticRoutes       `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
}

// VPCRouterInternetConnection インターフェース
type VPCRouterInternetConnection struct {
	Enabled types.StringFlag `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
}

// VPCRouterInterface インターフェース
type VPCRouterInterface struct {
	IPAddress        []string `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	VirtualIPAddress string   `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	IPAliases        []string `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	NetworkMaskLen   int      `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	// Index 仮想フィールド、VPCルータなどでInterfaces(実体は[]*Interface)を扱う場合にUnmarshalJSONの中で設定される
	//
	// Findした際のAPIからの応答にも同名のフィールドが含まれるが無関係。
	Index int `json:"-" yaml:"-" structs:"-"`
}

// VPCRouterInterfaces Interface配列
//
// 配列中にnullが返ってくる(VPCルータなど)への対応のためのtype
type VPCRouterInterfaces []*VPCRouterInterface

// UnmarshalJSON 配列中にnullが返ってくる(VPCルータなど)への対応
func (i *VPCRouterInterfaces) UnmarshalJSON(b []byte) error {
	type alias VPCRouterInterfaces
	var a alias
	if err := json.Unmarshal(b, &a); err != nil {
		return err
	}

	var dest []*VPCRouterInterface
	for i, v := range a {
		if v != nil {
			v.Index = i
			dest = append(dest, v)
		}
	}

	*i = VPCRouterInterfaces(dest)
	return nil
}

// MarshalJSON 配列中にnullが入る場合(VPCルータなど)への対応
func (i *VPCRouterInterfaces) MarshalJSON() ([]byte, error) {
	max := 0
	for _, iface := range *i {
		if max < iface.Index {
			max = iface.Index
		}
	}

	var dest = make([]*VPCRouterInterface, max)
	for _, iface := range *i {
		dest[iface.Index] = iface
	}

	return json.Marshal(dest)
}

// UnmarshalJSON JSON
func (i *VPCRouterInterface) UnmarshalJSON(b []byte) error {
	type alias VPCRouterInterface
	var a alias
	if err := json.Unmarshal(b, &a); err != nil {
		return err
	}
	*i = VPCRouterInterface(a)
	return nil
}

// VPCRouterStaticNAT スタティックNAT
type VPCRouterStaticNAT struct {
	Config  []*VPCRouterStaticNATConfig `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	Enabled types.StringFlag            `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
}

// MarshalJSON Configが一つ以上ある場合にEnabledをtrueに設定する
func (f *VPCRouterStaticNAT) MarshalJSON() ([]byte, error) {
	if f == nil || f.Config == nil {
		return nil, nil
	}
	if len(f.Config) > 0 {
		f.Enabled = types.StringTrue
	}
	type alias VPCRouterStaticNAT
	a := alias(*f)
	return json.Marshal(&a)
}

// VPCRouterStaticNATConfig スタティックNAT
type VPCRouterStaticNATConfig struct {
	GlobalAddress  string `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	PrivateAddress string `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	Description    string `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
}

// VPCRouterFirewall ファイアウォール
type VPCRouterFirewall struct {
	Config  []*VPCRouterFirewallConfig `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	Enabled types.StringFlag           `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
}

// MarshalJSON Configが一つ以上ある場合にEnabledをtrueに設定する
func (f *VPCRouterFirewall) MarshalJSON() ([]byte, error) {
	if f == nil {
		return nil, nil
	}
	if len(f.Config) > 0 {
		f.Enabled = types.StringTrue
	}
	type alias VPCRouterFirewall
	a := alias(*f)
	return json.Marshal(&a)
}

// VPCRouterFirewallConfig ファイアウォール
type VPCRouterFirewallConfig struct {
	Receive []*VPCRouterFirewallRule `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	Send    []*VPCRouterFirewallRule `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
}

// VPCRouterFirewallRule ファイアウォール ルール
type VPCRouterFirewallRule struct {
	Protocol           types.Protocol           `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	SourceNetwork      types.VPCFirewallNetwork `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	SourcePort         types.VPCFirewallPort    `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	DestinationNetwork types.VPCFirewallNetwork `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	DestinationPort    types.VPCFirewallPort    `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	Action             types.Action             `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	Logging            types.StringFlag         `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	Description        string                   `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
}

// VPCRouterDHCPServer DHCPサーバ
type VPCRouterDHCPServer struct {
	Config  []*VPCRouterDHCPServerConfig `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	Enabled types.StringFlag             `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
}

// MarshalJSON Configが一つ以上ある場合にEnabledをtrueに設定する
func (f *VPCRouterDHCPServer) MarshalJSON() ([]byte, error) {
	if f == nil {
		return nil, nil
	}
	if len(f.Config) > 0 {
		f.Enabled = types.StringTrue
	}
	type alias VPCRouterDHCPServer
	a := alias(*f)
	return json.Marshal(&a)
}

// VPCRouterDHCPServerConfig DHCPサーバ
type VPCRouterDHCPServerConfig struct {
	Interface  string   `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	RangeStop  string   `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	RangeStart string   `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	DNSServers []string `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
}

// VPCRouterDHCPStaticMappings DHCPスタティックマッピング
type VPCRouterDHCPStaticMappings struct {
	Config  []*VPCRouterDHCPStaticMappingConfig `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	Enabled types.StringFlag                    `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
}

// MarshalJSON Configが一つ以上ある場合にEnabledをtrueに設定する
func (f *VPCRouterDHCPStaticMappings) MarshalJSON() ([]byte, error) {
	if f == nil {
		return nil, nil
	}
	if len(f.Config) > 0 {
		f.Enabled = types.StringTrue
	}
	type alias VPCRouterDHCPStaticMappings
	a := alias(*f)
	return json.Marshal(&a)
}

// VPCRouterDHCPStaticMappingConfig DHCPスタティックマッピング
type VPCRouterDHCPStaticMappingConfig struct {
	MACAddress string `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	IPAddress  string `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
}

// VPCRouterPPTPServer PPTP
type VPCRouterPPTPServer struct {
	Config  *VPCRouterPPTPServerConfig `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	Enabled types.StringFlag           `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
}

// VPCRouterPPTPServerConfig PPTP
type VPCRouterPPTPServerConfig struct {
	RangeStart string `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	RangeStop  string `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
}

// VPCRouterL2TPIPsecServer L2TP
type VPCRouterL2TPIPsecServer struct {
	Config  *VPCRouterL2TPIPsecServerConfig `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	Enabled types.StringFlag                `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
}

// VPCRouterL2TPIPsecServerConfig L2TP
type VPCRouterL2TPIPsecServerConfig struct {
	RangeStart      string `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	RangeStop       string `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	PreSharedSecret string `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
}

// VPCRouterRemoteAccessUsers リモートアクセスユーザー
type VPCRouterRemoteAccessUsers struct {
	Config  []*VPCRouterRemoteAccessUserConfig `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	Enabled types.StringFlag                   `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
}

// MarshalJSON Configが一つ以上ある場合にEnabledをtrueに設定する
func (f *VPCRouterRemoteAccessUsers) MarshalJSON() ([]byte, error) {
	if f == nil {
		return nil, nil
	}
	if len(f.Config) > 0 {
		f.Enabled = types.StringTrue
	}
	type alias VPCRouterRemoteAccessUsers
	a := alias(*f)
	return json.Marshal(&a)
}

// VPCRouterRemoteAccessUserConfig リモートアクセスユーザー
type VPCRouterRemoteAccessUserConfig struct {
	UserName string `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	Password string `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
}

// VPCRouterSiteToSiteIPsecVPN サイト間VPN
type VPCRouterSiteToSiteIPsecVPN struct {
	Config  []*VPCRouterSiteToSiteIPsecVPNConfig `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	Enabled types.StringFlag                     `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
}

// MarshalJSON Configが一つ以上ある場合にEnabledをtrueに設定する
func (f *VPCRouterSiteToSiteIPsecVPN) MarshalJSON() ([]byte, error) {
	if f == nil {
		return nil, nil
	}
	if len(f.Config) > 0 {
		f.Enabled = types.StringTrue
	}
	type alias VPCRouterSiteToSiteIPsecVPN
	a := alias(*f)
	return json.Marshal(&a)
}

// VPCRouterSiteToSiteIPsecVPNConfig サイト間VPN
type VPCRouterSiteToSiteIPsecVPNConfig struct {
	Peer            string   `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	PreSharedSecret string   `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	RemoteID        string   `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	Routes          []string `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	LocalPrefix     []string `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
}

// VPCRouterStaticRoutes スタティックルート
type VPCRouterStaticRoutes struct {
	Config  []*VPCRouterStaticRouteConfig `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	Enabled types.StringFlag              `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
}

// MarshalJSON Configが一つ以上ある場合にEnabledをtrueに設定する
func (f *VPCRouterStaticRoutes) MarshalJSON() ([]byte, error) {
	if f == nil {
		return nil, nil
	}
	if len(f.Config) > 0 {
		f.Enabled = types.StringTrue
	}
	type alias VPCRouterStaticRoutes
	a := alias(*f)
	return json.Marshal(&a)
}

// VPCRouterStaticRouteConfig スタティックルート
type VPCRouterStaticRouteConfig struct {
	Prefix  string `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	NextHop string `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
}
