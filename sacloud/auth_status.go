package sacloud

// AuthStatus 現在の認証状態
type AuthStatus struct {
	Account            *Account
	AuthClass          EAuthClass          `json:",omitempty"`
	AuthMethod         EAuthMethod         `json:",omitempty"`
	ExternalPermission EExternalPermission `json:",omitempty"`
	IsAPIKey           bool                `json:",omitempty"`
	Member             *Member
	OperationPenalty   EOperationPenalty `json:",omitempty"`
	Permission         EPermission       `json:",omitempty"`
	IsOk               bool              `json:"is_ok,omitempty"`

	// RESTFilter [unknown type] `json:",omitempty"`
	// User [unknown type] `json:",omitempty"`

}

// --------------------------------------------------------

// EAuthClass 認証種別
type EAuthClass string

// EAuthClassAccount アカウント認証
var EAuthClassAccount = EAuthClass("account")

// --------------------------------------------------------

// EAuthMethod 認証方法
type EAuthMethod string

// EAuthMethodAPIKey APIキー認証
var EAuthMethodAPIKey = EAuthMethod("apikey")

// --------------------------------------------------------

// EExternalPermission 他サービスへのアクセス権
type EExternalPermission string

// EExternalPermissionBill 請求情報
var EExternalPermissionBill = EExternalPermission("bill")

// EExternalPermissionCDN ウェブアクセラレータ
var EExternalPermissionCDN = EExternalPermission("cdn")

// --------------------------------------------------------

// EOperationPenalty ペナルティ
type EOperationPenalty string

// EOperationPenaltyNone ペナルティなし
var EOperationPenaltyNone = EOperationPenalty("none")

// --------------------------------------------------------

// EPermission アクセスレベル
type EPermission string

// EPermissionCreate 作成・削除権限
var EPermissionCreate = EPermission("create")

// EPermissionArrange 設定変更権限
var EPermissionArrange = EPermission("arrange")

// EPermissionPower 電源操作権限
var EPermissionPower = EPermission("power")

// EPermissionView リソース閲覧権限
var EPermissionView = EPermission("view")
