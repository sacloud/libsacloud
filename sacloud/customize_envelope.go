package sacloud

import (
	"encoding/json"

	"github.com/sacloud/libsacloud/v2/sacloud/naked"
)

// UnmarshalJSON APIからの戻り値でレスポンスボディ直下にデータを持つことへの対応
func (a *authstatusReadResponseEnvelope) UnmarshalJSON(data []byte) error {
	type alias authstatusReadResponseEnvelope

	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	var nakedAuthStatus naked.AuthStatus
	if err := json.Unmarshal(data, &nakedAuthStatus); err != nil {
		return err
	}
	tmp.AuthStatus = &nakedAuthStatus

	*a = authstatusReadResponseEnvelope(tmp)
	return nil
}
