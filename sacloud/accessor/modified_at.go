package accessor

import "time"

// ModifiedAt 更新日時
type ModifiedAt interface {
	GetModifiedAt() time.Time
	SetModifiedAt(t time.Time)
}
