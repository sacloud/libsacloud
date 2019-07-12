package naked

import (
	"time"

	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

// Coupon クーポン情報
type Coupon struct {
	ID             types.ID   `json:"CouponID,omitempty" yaml:",omitempty" structs:",omitempty"`         // クーポンID
	MemberID       string     `json:",omitempty" yaml:"member_id,omitempty" structs:",omitempty"`        // メンバーID
	ContractID     types.ID   `json:",omitempty" yaml:"contract_id,omitempty" structs:",omitempty"`      // 契約ID
	ServiceClassID types.ID   `json:",omitempty" yaml:"service_class_id,omitempty" structs:",omitempty"` // サービスクラスID
	Discount       int64      `json:",omitempty" yaml:"discount,omitempty" structs:",omitempty"`         // クーポン残高
	AppliedAt      *time.Time `json:",omitempty" yaml:"applied_at,omitempty" structs:",omitempty"`       // 適用開始日
	UntilAt        *time.Time `json:",omitempty" yaml:"until_at,omitempty" structs:",omitempty"`         // 有効期限
}
