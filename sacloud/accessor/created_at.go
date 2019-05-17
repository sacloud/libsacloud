package accessor

import "time"

// CreatedAt 作成日時
type CreatedAt interface {
	GetCreatedAt() *time.Time
	SetCreatedAt(t *time.Time)
}
