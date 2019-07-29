package accessor

import (
	"time"
)

// Instance インスタンス
type Instance interface {
	InstanceStatus
	GetInstanceHostName() string
	SetInstanceHostName(v string)
	GetInstanceHostInfoURL() string
	SetInstanceHostInfoURL(v string)
	GetInstanceStatusChangedAt() time.Time
	SetInstanceStatusChangedAt(v time.Time)
}
