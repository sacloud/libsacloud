package types

// ECommitment サーバプランCPUコミットメント
//
// 通常 or コア専有
type ECommitment string

// Commitments サーバプランCPUコミットメント
var Commitments = struct {
	// Unknown 不明
	Unknown ECommitment
	// Standard 通常
	Standard ECommitment
	// DedicatedCPU コア専有
	DedicatedCPU ECommitment
}{
	Unknown:      ECommitment(""),
	Standard:     ECommitment("standard"),
	DedicatedCPU: ECommitment("dedicatedcpu"),
}

// IsStandard Standardであるか判定
func (c ECommitment) IsStandard() bool {
	return c == Commitments.Standard
}

// IsDedicatedCPU DedicatedCPUであるか判定
func (c ECommitment) IsDedicatedCPU() bool {
	return c == Commitments.DedicatedCPU
}

// String ECommitmentの文字列表現
func (c ECommitment) String() string {
	return string(c)
}
