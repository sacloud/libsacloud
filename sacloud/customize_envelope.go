package sacloud

import (
	"encoding/json"

	"github.com/sacloud/libsacloud/v2/sacloud/naked"
	"github.com/sacloud/libsacloud/v2/sacloud/search"
)

// UnmarshalJSON APIからの戻り値でレスポンスボディ直下にデータを持つことへの対応
func (a *authStatusReadResponseEnvelope) UnmarshalJSON(data []byte) error {
	type alias authStatusReadResponseEnvelope

	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	var nakedAuthStatus naked.AuthStatus
	if err := json.Unmarshal(data, &nakedAuthStatus); err != nil {
		return err
	}
	tmp.AuthStatus = &nakedAuthStatus

	*a = authStatusReadResponseEnvelope(tmp)
	return nil
}

func (b *billDetailsCSVResponseEnvelope) UnmarshalJSON(data []byte) error {
	type alias billDetailsCSVResponseEnvelope

	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	var nakedBillDetailCSV naked.BillDetailCSV
	if err := json.Unmarshal(data, &nakedBillDetailCSV); err != nil {
		return err
	}
	tmp.CSV = &nakedBillDetailCSV

	*b = billDetailsCSVResponseEnvelope(tmp)
	return nil
}

func (m *mobileGatewaySetSIMRoutesRequestEnvelope) MarshalJSON() ([]byte, error) {
	type alias struct {
		SIMRoutes []*naked.MobileGatewaySIMRoute `json:"sim_routes"`
	}
	tmp := &alias{
		SIMRoutes: m.SIMRoutes,
	}
	if len(tmp.SIMRoutes) == 0 {
		tmp.SIMRoutes = make([]*naked.MobileGatewaySIMRoute, 0)
	}
	return json.Marshal(tmp)
}

// UnmarshalJSON APIからの戻り値でレスポンスボディ直下にデータを持つことへの対応
func (s *serverGetVNCProxyResponseEnvelope) UnmarshalJSON(data []byte) error {
	type alias serverGetVNCProxyResponseEnvelope

	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	var nakedVNCProxy naked.VNCProxyInfo
	if err := json.Unmarshal(data, &nakedVNCProxy); err != nil {
		return err
	}
	tmp.VNCProxyInfo = &nakedVNCProxy

	*s = serverGetVNCProxyResponseEnvelope(tmp)
	return nil
}

/*
 * 検索時に固定パラメータを設定するための実装
 */

func (s autoBackupFindRequestEnvelope) MarshalJSON() ([]byte, error) {
	type alias autoBackupFindRequestEnvelope
	tmp := alias(s)
	if tmp.Filter == nil {
		tmp.Filter = search.Filter{}
	}
	tmp.Filter[search.Key("Provider.Class")] = "autobackup"
	return json.Marshal(tmp)
}

func (s dNSFindRequestEnvelope) MarshalJSON() ([]byte, error) {
	type alias dNSFindRequestEnvelope
	tmp := alias(s)
	if tmp.Filter == nil {
		tmp.Filter = search.Filter{}
	}
	tmp.Filter[search.Key("Provider.Class")] = "dns"
	return json.Marshal(tmp)
}

func (s simpleMonitorFindRequestEnvelope) MarshalJSON() ([]byte, error) {
	type alias simpleMonitorFindRequestEnvelope
	tmp := alias(s)
	if tmp.Filter == nil {
		tmp.Filter = search.Filter{}
	}
	tmp.Filter[search.Key("Provider.Class")] = "simplemon"
	return json.Marshal(tmp)
}

func (s gSLBFindRequestEnvelope) MarshalJSON() ([]byte, error) {
	type alias gSLBFindRequestEnvelope
	tmp := alias(s)
	if tmp.Filter == nil {
		tmp.Filter = search.Filter{}
	}
	tmp.Filter[search.Key("Provider.Class")] = "gslb"
	return json.Marshal(tmp)
}

func (s proxyLBFindRequestEnvelope) MarshalJSON() ([]byte, error) {
	type alias proxyLBFindRequestEnvelope
	tmp := alias(s)
	if tmp.Filter == nil {
		tmp.Filter = search.Filter{}
	}
	tmp.Filter[search.Key("Provider.Class")] = "proxylb"
	return json.Marshal(tmp)
}

func (s sIMFindRequestEnvelope) MarshalJSON() ([]byte, error) {
	type alias sIMFindRequestEnvelope
	tmp := alias(s)
	if tmp.Filter == nil {
		tmp.Filter = search.Filter{}
	}
	tmp.Filter[search.Key("Provider.Class")] = "sim"
	return json.Marshal(tmp)
}

func (s databaseFindRequestEnvelope) MarshalJSON() ([]byte, error) {
	type alias databaseFindRequestEnvelope
	tmp := alias(s)
	if tmp.Filter == nil {
		tmp.Filter = search.Filter{}
	}
	tmp.Filter[search.Key("Class")] = "database"
	return json.Marshal(tmp)
}

func (s loadBalancerFindRequestEnvelope) MarshalJSON() ([]byte, error) {
	type alias loadBalancerFindRequestEnvelope
	tmp := alias(s)
	if tmp.Filter == nil {
		tmp.Filter = search.Filter{}
	}
	tmp.Filter[search.Key("Class")] = "loadbalancer"
	return json.Marshal(tmp)
}

func (s vPCRouterFindRequestEnvelope) MarshalJSON() ([]byte, error) {
	type alias vPCRouterFindRequestEnvelope
	tmp := alias(s)
	if tmp.Filter == nil {
		tmp.Filter = search.Filter{}
	}
	tmp.Filter[search.Key("Class")] = "vpcrouter"
	return json.Marshal(tmp)
}

func (s nFSFindRequestEnvelope) MarshalJSON() ([]byte, error) {
	type alias nFSFindRequestEnvelope
	tmp := alias(s)
	if tmp.Filter == nil {
		tmp.Filter = search.Filter{}
	}
	tmp.Filter[search.Key("Class")] = "nfs"
	return json.Marshal(tmp)
}

func (s mobileGatewayFindRequestEnvelope) MarshalJSON() ([]byte, error) {
	type alias mobileGatewayFindRequestEnvelope
	tmp := alias(s)
	if tmp.Filter == nil {
		tmp.Filter = search.Filter{}
	}
	tmp.Filter[search.Key("Class")] = "mobilegateway"
	return json.Marshal(tmp)
}
