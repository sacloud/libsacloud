package sacloud

import "time"

// Bill type of SakuraCloud Bill
type Bill struct {
	Amount         int64      `json:",omitempty"`
	BillID         int64      `json:",omitempty"`
	Date           *time.Time `json:",omitempty"`
	MemberID       string     `json:",omitempty"`
	Paid           bool       `json:",omitempty"`
	PayLimit       *time.Time `json:",omitempty"`
	PaymentClassID int        `json:",omitempty"`
}

// BillDetail type of SakuraCloud BillDetail
type BillDetail struct {
	Amount         int64  `json:",omitempty"`
	ContractID     int64  `json:",omitempty"`
	Description    string `json:",omitempty"`
	Index          int    `json:",omitempty"`
	ServiceClassID int64  `json:",omitempty"`
	Usage          int64  `json:",omitempty"`
	Zone           string `json:",omitempty"`
}
