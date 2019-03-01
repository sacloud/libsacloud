package sacloud

import (
	"encoding/json"
	"time"
)

// ProxyLB ProxyLB(CommonServiceItem)
type ProxyLB struct {
	*Resource        // ID
	propName         // 名称
	propDescription  // 説明
	propServiceClass // サービスクラス
	propIcon         // アイコン
	propTags         // タグ
	propCreatedAt    // 作成日時
	propModifiedAt   // 変更日時

	Status   ProxyLBStatus   `json:",omitempty"` // ステータス
	Provider ProxyLBProvider `json:",omitempty"` // プロバイダ
	Settings ProxyLBSettings `json:",omitempty"` // ProxyLB設定

}

// ProxyLBSettings ProxyLB設定
type ProxyLBSettings struct {
	ProxyLB ProxyLBSetting `json:",omitempty"` // ProxyLB ProxyLBエントリー
}

// ProxyLBStatus ProxyLBステータス
type ProxyLBStatus struct {
	VirtualIPAddress string   `json:",omitempty"` // 割り当てられたVIP
	ProxyNetworks    []string `json:",omitempty"` // プロキシ元ネットワークアドレス(CIDR)
}

// ProxyLBProvider プロバイダ
type ProxyLBProvider struct {
	Class string `json:",omitempty"` // クラス
}

// CreateNewProxyLB ProxyLB作成
func CreateNewProxyLB(name string) *ProxyLB {
	return &ProxyLB{
		Resource: &Resource{},
		propName: propName{Name: name},
		Provider: ProxyLBProvider{
			Class: "proxylb",
		},
		Settings: ProxyLBSettings{
			ProxyLB: ProxyLBSetting{
				HealthCheck: defaultProxyLBHealthCheck,
				SorryServer: &ProxyLBSorryServer{},
				Servers:     []ProxyLBServer{},
			},
		},
	}

}

// SetHTTPHealthCheck HTTPヘルスチェック 設定
func (p *ProxyLB) SetHTTPHealthCheck(hostHeader, path string, delayLoop int) {
	if delayLoop <= 0 {
		delayLoop = 10
	}

	p.Settings.ProxyLB.HealthCheck.Protocol = "http"
	p.Settings.ProxyLB.HealthCheck.Host = hostHeader
	p.Settings.ProxyLB.HealthCheck.Path = path
	p.Settings.ProxyLB.HealthCheck.DelayLoop = delayLoop
	p.Settings.ProxyLB.HealthCheck.Port = 0
}

// SetTCPHealthCheck TCPヘルスチェック 設定
func (p *ProxyLB) SetTCPHealthCheck(port, delayLoop int) {
	if delayLoop <= 0 {
		delayLoop = 10
	}

	p.Settings.ProxyLB.HealthCheck.Protocol = "tcp"
	p.Settings.ProxyLB.HealthCheck.Host = ""
	p.Settings.ProxyLB.HealthCheck.Path = ""
	p.Settings.ProxyLB.HealthCheck.DelayLoop = delayLoop
	p.Settings.ProxyLB.HealthCheck.Port = port
}

// SetSorryServer ソーリーサーバ 設定
func (p *ProxyLB) SetSorryServer(ipaddress string, port int) {
	p.Settings.ProxyLB.SorryServer = &ProxyLBSorryServer{
		IPAddress: ipaddress,
		Port:      port,
	}
}

// HasProxyLBServer ProxyLB配下にサーバーを保持しているか判定
func (p *ProxyLB) HasProxyLBServer() bool {
	return len(p.Settings.ProxyLB.Servers) > 0
}

// ClearProxyLBServer ProxyLB配下のサーバーをクリア
func (p *ProxyLB) ClearProxyLBServer() {
	p.Settings.ProxyLB.Servers = []ProxyLBServer{}
}

// AddBindPort バインドポート追加
func (p *ProxyLB) AddBindPort(mode string, port int) {
	p.Settings.ProxyLB.AddBindPort(mode, port)
}

// DeleteBindPort バインドポート削除
func (p *ProxyLB) DeleteBindPort(mode string, port int) {
	p.Settings.ProxyLB.DeleteBindPort(mode, port)
}

// AddServer ProxyLB配下のサーバーを追加
func (p *ProxyLB) AddServer(ip string, port int) {
	p.Settings.ProxyLB.AddServer(ip, port)
}

// DeleteServer ProxyLB配下のサーバーを削除
func (p *ProxyLB) DeleteServer(ip string, port int) {
	p.Settings.ProxyLB.DeleteServer(ip, port)
}

// ProxyLBSetting ProxyLBセッティング
type ProxyLBSetting struct {
	HealthCheck ProxyLBHealthCheck  `json:",omitempty"` // ヘルスチェック
	SorryServer *ProxyLBSorryServer `json:",omitempty"` // ソーリーサーバー
	BindPorts   []*ProxyLBBindPorts // プロキシ方式(プロトコル&ポート)
	Servers     []ProxyLBServer     // サーバー
}

// ProxyLBSorryServer ソーリーサーバ
type ProxyLBSorryServer struct {
	IPAddress string `json:",omitempty"` // IPアドレス
	Port      int    `json:",omitempty"` // ポート
}

// AddBindPort バインドポート追加
func (s *ProxyLBSetting) AddBindPort(mode string, port int) {
	var isExist bool
	for i := range s.BindPorts {
		if s.BindPorts[i].ProxyMode == mode && s.BindPorts[i].Port == port {
			isExist = true
		}
	}

	if !isExist {
		s.BindPorts = append(s.BindPorts, &ProxyLBBindPorts{
			ProxyMode: mode,
			Port:      port,
		})
	}
}

// DeleteBindPort バインドポート削除
func (s *ProxyLBSetting) DeleteBindPort(mode string, port int) {
	var res []*ProxyLBBindPorts
	for i := range s.BindPorts {
		if s.BindPorts[i].ProxyMode != mode || s.BindPorts[i].Port != port {
			res = append(res, s.BindPorts[i])
		}
	}
	s.BindPorts = res
}

// AddServer ProxyLB配下のサーバーを追加
func (s *ProxyLBSetting) AddServer(ip string, port int) {
	var record ProxyLBServer
	var isExist = false
	for i := range s.Servers {
		if s.Servers[i].IPAddress == ip && s.Servers[i].Port == port {
			isExist = true
		}
	}

	if !isExist {
		record = ProxyLBServer{
			IPAddress: ip,
			Port:      port,
			Enabled:   true,
		}
		s.Servers = append(s.Servers, record)
	}
}

// DeleteServer ProxyLB配下のサーバーを削除
func (s *ProxyLBSetting) DeleteServer(ip string, port int) {
	var res []ProxyLBServer
	for i := range s.Servers {
		if s.Servers[i].IPAddress != ip || s.Servers[i].Port != port {
			res = append(res, s.Servers[i])
		}
	}

	s.Servers = res
}

// AllowProxyLBBindModes プロキシ方式
var AllowProxyLBBindModes = []string{"http", "https"}

// ProxyLBBindPorts プロキシ方式
type ProxyLBBindPorts struct {
	ProxyMode string `json:",omitempty"` // モード(プロトコル)
	Port      int    `json:",omitempty"` // ポート
}

// ProxyLBServer ProxyLB配下のサーバー
type ProxyLBServer struct {
	IPAddress string `json:",omitempty"` // IPアドレス
	Port      int    `json:",omitempty"` // ポート
	Enabled   bool   `json:",omitempty"` // 有効/無効
}

// NewProxyLBServer ProxyLB配下のサーバ作成
func NewProxyLBServer(ipaddress string, port int) *ProxyLBServer {
	return &ProxyLBServer{
		IPAddress: ipaddress,
		Port:      port,
		Enabled:   true,
	}
}

// AllowProxyLBHealthCheckProtocols プロキシLBで利用できるヘルスチェックプロトコル
var AllowProxyLBHealthCheckProtocols = []string{"http", "tcp"}

// ProxyLBHealthCheck ヘルスチェック
type ProxyLBHealthCheck struct {
	Protocol  string `json:",omitempty"` // プロトコル
	Host      string `json:",omitempty"` // 対象ホスト
	Path      string `json:",omitempty"` // HTTPの場合のリクエストパス
	Port      int    `json:",omitempty"` // ポート番号
	DelayLoop int    `json:",omitempty"` // 監視間隔

}

var defaultProxyLBHealthCheck = ProxyLBHealthCheck{
	Protocol:  "http",
	Host:      "",
	Path:      "/",
	DelayLoop: 10,
}

// ProxyLBCertificates ProxyLBのSSL証明書
type ProxyLBCertificates struct {
	ServerCertificate       string    // サーバ証明書
	IntermediateCertificate string    // 中間証明書
	PrivateKey              string    // 秘密鍵
	CertificateEndDate      time.Time `json:",omitempty"` // 有効期限
}

// UnmarshalJSON UnmarshalJSON(CertificateEndDateのtime.TimeへのUnmarshal対応)
func (p *ProxyLBCertificates) UnmarshalJSON(data []byte) error {
	tmp := &struct {
		ServerCertificate       string // サーバ証明書
		IntermediateCertificate string // 中間証明書
		PrivateKey              string // 秘密鍵
		CertificateEndDate      string `json:",omitempty"` // 有効期限
	}{}
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	p.ServerCertificate = tmp.ServerCertificate
	p.IntermediateCertificate = tmp.IntermediateCertificate
	p.PrivateKey = tmp.PrivateKey
	if tmp.CertificateEndDate != "" {
		date, err := time.Parse("Jan _2 15:04:05 2006 MST", tmp.CertificateEndDate)
		if err != nil {
			return err
		}
		p.CertificateEndDate = date
	}

	return nil
}

// ProxyLBHealth ProxyLBのヘルスチェック戻り値
type ProxyLBHealth struct {
	ActiveConn int                    // アクティブなコネクション数
	CPS        int                    // 秒あたりコネクション数
	Servers    []*ProxyLBHealthServer // 実サーバのステータス
}

// ProxyLBHealthServer ProxyLBの実サーバのステータス
type ProxyLBHealthServer struct {
	ActiveConn int    // アクティブなコネクション数
	Status     string // ステータス(UP or DOWN)
	IPAddress  string // IPアドレス
	Port       string // ポート
	CPS        int    // 秒あたりコネクション数
}
