package sacloud

// SSHKey 公開鍵
type SSHKey struct {
	*Resource       // ID
	propName        // 名称
	propDescription // 説明
	propCreatedAt   // 作成日時

	PublicKey   string `json:",omitempty"` // 公開鍵
	Fingerprint string `json:",omitempty"` // フィンガープリント
}
