package sacloud

// Internet ルーター
type Internet struct {
	*Resource        // ID
	propName         // 名称
	propDescription  // 説明
	propScope        // スコープ
	propServiceClass // サービスクラス
	propSwitch       // 接続先スイッチ
	propIcon         // アイコン
	propTags         // タグ
	propCreatedAt    // 作成日時

	BandWidthMbps  int `json:",omitempty"` // 帯域
	NetworkMaskLen int `json:",omitempty"` // ネットワークマスク長

	//TODO Zone(API側起因のデータ型不一致のため)
	// ZoneType
}
