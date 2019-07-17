package sacloud

import (
	"encoding/json"

	"github.com/sacloud/libsacloud/v2/sacloud/naked"
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
