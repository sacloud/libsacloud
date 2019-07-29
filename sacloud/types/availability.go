package types

// EAvailability 有効状態
type EAvailability string

// Availabilities 有効状態
var Availabilities = struct {
	// Unknown 不明
	Unknown EAvailability
	// Available 有効
	Available EAvailability // 有効
	// Uploading アップロード中
	Uploading EAvailability // アップロード中
	// Failed 失敗
	Failed EAvailability // 失敗
	// Migrating マイグレーション中
	Migrating EAvailability
	// Transfering 転送中(他ゾーンからのアーカイブコピー作成など)
	Transfering EAvailability
	// Discontinued 切断(ソースアーカイブ/ソースディスクなどで利用される)
	Discontinued EAvailability
}{
	Unknown:      EAvailability(""),
	Available:    EAvailability("available"),
	Uploading:    EAvailability("uploading"),
	Failed:       EAvailability("failed"),
	Migrating:    EAvailability("migrating"),
	Transfering:  EAvailability("transfering"),
	Discontinued: EAvailability("discontinued"),
}

// IsAvailable 有効状態が"有効"か判定
func (e EAvailability) IsAvailable() bool {
	return e == Availabilities.Available
}

// IsUploading 有効状態が"アップロード中"か判定
func (e EAvailability) IsUploading() bool {
	return e == Availabilities.Uploading
}

// IsFailed 有効状態が"失敗"か判定
func (e EAvailability) IsFailed() bool {
	return e == Availabilities.Failed
}

// IsMigrating 有効状態が"マイグレーション中"か判定
func (e EAvailability) IsMigrating() bool {
	return e == Availabilities.Migrating
}

// IsTransfering 有効状態が"転送中"か判定
func (e EAvailability) IsTransfering() bool {
	return e == Availabilities.Transfering
}

// IsDiscontinued 有効状態が"切断"か判定
func (e EAvailability) IsDiscontinued() bool {
	return e == Availabilities.Discontinued
}
