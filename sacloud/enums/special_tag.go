package enums

// SpecialTag 特殊タグ
type SpecialTag string

// SpecialTags 特殊タグ一覧
var SpecialTags = struct {
	// GroupA サーバをグループ化し起動ホストを分離します(グループA)
	GroupA SpecialTag
	// GroupB サーバをグループ化し起動ホストを分離します(グループB)
	GroupB SpecialTag
	// GroupC サーバをグループ化し起動ホストを分離します(グループC)
	GroupC SpecialTag
	// GroupD サーバをグループ化し起動ホストを分離します(グループD)
	GroupD SpecialTag
	// AutoReboot サーバ停止時に自動起動します
	AutoReboot SpecialTag
	// KeyboardUS リモートスクリーン画面でUSキーボード入力します
	KeyboardUS SpecialTag
	// BootCDROM 優先ブートデバイスをCD-ROMに設定します
	BootCDROM SpecialTag
	// BootNetwork 優先ブートデバイスをPXE bootに設定します
	BootNetwork SpecialTag
}{
	GroupA:      SpecialTag("@group=a"),
	GroupB:      SpecialTag("@group=b"),
	GroupC:      SpecialTag("@group=c"),
	GroupD:      SpecialTag("@group=d"),
	AutoReboot:  SpecialTag("@auto-reboot"),
	KeyboardUS:  SpecialTag("@keyboard-us"),
	BootCDROM:   SpecialTag("@boot-cdrom"),
	BootNetwork: SpecialTag("@boot-network"),
}
