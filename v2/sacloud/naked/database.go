// Copyright 2016-2022 The Libsacloud Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package naked

import (
	"encoding/json"
	"strconv"
	"strings"
	"time"

	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

// Database データベース
type Database struct {
	ID           types.ID            `json:",omitempty" yaml:"id,omitempty" structs:",omitempty"`
	Name         string              `json:",omitempty" yaml:"name,omitempty" structs:",omitempty"`
	Description  string              `yaml:"description"`
	Tags         types.Tags          `yaml:"tags"`
	Icon         *Icon               `json:",omitempty" yaml:"icon,omitempty" structs:",omitempty"`
	CreatedAt    *time.Time          `json:",omitempty" yaml:"created_at,omitempty" structs:",omitempty"`
	ModifiedAt   *time.Time          `json:",omitempty" yaml:"modified_at,omitempty" structs:",omitempty"`
	Availability types.EAvailability `json:",omitempty" yaml:"availability,omitempty" structs:",omitempty"`
	Class        string              `json:",omitempty" yaml:"class,omitempty" structs:",omitempty"`
	ServiceClass string              `json:",omitempty" yaml:"service_class,omitempty" structs:",omitempty"`
	Plan         *AppliancePlan      `json:",omitempty" yaml:"plan,omitempty" structs:",omitempty"`
	Instance     *Instance           `json:",omitempty" yaml:"instance,omitempty" structs:",omitempty"`
	Interfaces   []*Interface        `json:",omitempty" yaml:"interfaces,omitempty" structs:",omitempty"`
	Switch       *Switch             `json:",omitempty" yaml:"switch,omitempty" structs:",omitempty"`
	Settings     *DatabaseSettings   `json:",omitempty" yaml:"settings,omitempty" structs:",omitempty"`
	SettingsHash string              `json:",omitempty" yaml:"settings_hash,omitempty" structs:",omitempty"`
	Remark       *ApplianceRemark    `json:",omitempty" yaml:"remark,omitempty" structs:",omitempty"`

	Generation interface{}
}

// DatabaseSettingsUpdate データベース
type DatabaseSettingsUpdate struct {
	Settings     *DatabaseSettings `json:",omitempty" yaml:"settings,omitempty" structs:",omitempty"`
	SettingsHash string            `json:",omitempty" yaml:"settings_hash,omitempty" structs:",omitempty"`
}

// DatabaseSettings データベース設定
type DatabaseSettings struct {
	DBConf *DatabaseSetting `json:",omitempty" yaml:"db_conf,omitempty" structs:",omitempty"`
}

// DatabaseSetting データベース設定
type DatabaseSetting struct {
	Common      *DatabaseSettingCommon      `json:",omitempty" yaml:"common,omitempty" structs:",omitempty"`
	Backup      *DatabaseSettingBackup      `json:",omitempty" yaml:"backup,omitempty" structs:",omitempty"`
	Replication *DatabaseSettingReplication `json:",omitempty" yaml:"replication,omitempty" structs:",omitempty"`
}

// DatabaseSettingCommon データベース設定 汎用項目設定
type DatabaseSettingCommon struct {
	// WebUI WebUIの有効/無効、またはアクセスするためのアドレス
	//
	// [HACK] Create時はbool型、Read/Update時は文字列(FQDN or IP)となる。
	// また、無効にするにはJSONで要素自体を指定しないことで行う。
	WebUI           interface{}                   `yaml:"web_ui"`
	ServicePort     int                           `json:",omitempty" yaml:"service_port,omitempty" structs:",omitempty"`
	SourceNetwork   DatabaseSettingSourceNetworks `yaml:"source_network"`
	DefaultUser     string                        `json:",omitempty" yaml:"default_user,omitempty" structs:",omitempty"`
	UserPassword    string                        `json:",omitempty" yaml:"user_password,omitempty" structs:",omitempty"`
	ReplicaUser     string                        `json:",omitempty" yaml:"replica_user,omitempty" structs:",omitempty"`
	ReplicaPassword string                        `json:",omitempty" yaml:"replica_password,omitempty" structs:",omitempty"`
}

// DatabaseSettingSourceNetworks データベースへのアクセスを許可するCIDRリスト
//
// Note: すべての接続先を許可する場合は"0.0.0.0/0"を指定する。
// この処理はMarshalJSON時にDatabaseSettingSourceNetwork側で行われるため、
// APIクライアント側は許可したいCIDRブロックのリストを指定する。
// libsacloudではすべての接続を拒否する設定はサポートしない。
type DatabaseSettingSourceNetworks []string

// MarshalJSON すべての接続先を許可する場合は"0.0.0.0/0"を指定するための対応
func (d DatabaseSettingSourceNetworks) MarshalJSON() ([]byte, error) {
	type alias DatabaseSettingSourceNetworks
	dest := alias(d)

	if dest == nil || len(dest) == 0 {
		dest = append(dest, "0.0.0.0/0")
	}

	return json.Marshal(dest)
}

func (d *DatabaseSettingSourceNetworks) UnmarshalJSON(b []byte) error {
	if string(b) == `""` || string(b) == "" {
		return nil
	}
	type alias DatabaseSettingSourceNetworks

	var a alias
	if err := json.Unmarshal(b, &a); err != nil {
		return err
	}
	if len(a) == 1 && a[0] == "0.0.0.0/0" {
		return nil
	}
	*d = DatabaseSettingSourceNetworks(a)
	return nil
}

// DatabaseSettingBackup データベース設定 バックアップ設定
type DatabaseSettingBackup struct {
	Rotate    int                        `json:",omitempty" yaml:"rotate,omitempty" structs:",omitempty"`
	Time      string                     `json:",omitempty" yaml:"time,omitempty" structs:",omitempty"`
	DayOfWeek []types.EBackupSpanWeekday `json:",omitempty" yaml:"day_of_week,omitempty" structs:",omitempty"`
}

// UnmarshalJSON 配列/オブジェクトが混在することへの対応
func (d *DatabaseSettingBackup) UnmarshalJSON(b []byte) error {
	if string(b) == "[]" {
		return nil
	}
	type alias DatabaseSettingBackup

	var a alias
	if err := json.Unmarshal(b, &a); err != nil {
		return err
	}
	*d = DatabaseSettingBackup(a)
	return nil
}

// DatabaseSettingReplication レプリケーション設定
type DatabaseSettingReplication struct {
	Model     types.EDatabaseReplicationModel `json:",omitempty" yaml:"model,omitempty" structs:",omitempty"`
	Appliance *struct {
		ID types.ID
	} `json:",omitempty" yaml:"appliance,omitempty" structs:",omitempty"`
	IPAddress string `json:",omitempty" yaml:"ip_address,omitempty" structs:",omitempty"`
	Port      int    `json:",omitempty" yaml:"port,omitempty" structs:",omitempty"`
	User      string `json:",omitempty" yaml:"user,omitempty" structs:",omitempty"`
	Password  string `json:",omitempty" yaml:"password,omitempty" structs:",omitempty"`
}

// DatabaseStatusResponse Status APIの戻り値
type DatabaseStatusResponse struct {
	SettingsResponse *DatabaseStatus `json:",omitempty" yaml:"settings_response,omitempty" structs:",omitempty"`
}

// DatabaseStatus データベースのステータス
type DatabaseStatus struct {
	Status  types.EServerInstanceStatus `json:",omitempty" yaml:"status,omitempty" structs:",omitempty"`
	IsFatal bool                        `json:"is_fatal"`
	DBConf  *DatabaseStatusDBConf       `json:",omitempty" yaml:"db_conf,omitempty" structs:",omitempty"`
}

// DatabaseStatusDBConf データベース設定
type DatabaseStatusDBConf struct {
	Version  *DatabaseStatusVersion    `json:"version,omitempty" yaml:"version,omitempty" structs:",omitempty"`
	Log      []*DatabaseLog            `json:"log,omitempty" yaml:"log,omitempty" structs:",omitempty"`
	Backup   *DatabaseBackupInfo       `json:"backup,omitempty" yaml:"backup,omitempty" structs:",omitempty"`
	MariaDB  *DatabaseStatusMariaDB    `json:",omitempty" yaml:"maria_db,omitempty" structs:",omitempty"`
	Postgres *DatabaseStatusPostgreSQL `json:"postgres,omitempty" yaml:"postgres,omitempty" structs:",omitempty"`

	// 以下フィールドはサポートしない
	// Replication
}

type DatabaseStatusMariaDB struct {
	Status string `json:"status,omitempty"`
}
type DatabaseStatusPostgreSQL struct {
	Status string `json:"status,omitempty"`
}

// DatabaseStatusVersion データベース設定バージョン情報
type DatabaseStatusVersion struct {
	LastModified string `json:"lastmodified,omitempty" yaml:"last_modified,omitempty" structs:",omitempty"`
	CommitHash   string `json:"commithash,omitempty" yaml:"commit_hash,omitempty" structs:",omitempty"`
	Status       string `json:"status,omitempty" yaml:"status,omitempty" structs:",omitempty"`
	Tag          string `json:"tag,omitempty" yaml:"tag,omitempty" structs:",omitempty"`
	Expire       string `json:"expire,omitempty" yaml:"expire,omitempty" structs:",omitempty"`
}

// DatabaseLog データベースログ
type DatabaseLog struct {
	Name string             `json:"name,omitempty" yaml:"name,omitempty" structs:",omitempty"`
	Data string             `json:"data,omitempty" yaml:"data,omitempty" structs:",omitempty"`
	Size types.StringNumber `json:"size,omitempty" yaml:"size,omitempty" structs:",omitempty"`
}

// IsSystemdLog systemcltのログか判定
func (l *DatabaseLog) IsSystemdLog() bool {
	return l.Name == "systemctl"
}

// Logs ログボディ取得
func (l *DatabaseLog) Logs() []string {
	return strings.Split(l.Data, "\n")
}

// ID ログのID取得
func (l *DatabaseLog) ID() string {
	return l.Name
}

// DatabaseBackupInfo データベースバックアップ情報
type DatabaseBackupInfo struct {
	History []*DatabaseBackupHistory `json:"history,omitempty" yaml:"history,omitempty" structs:",omitempty"`
}

// DatabaseBackupHistory データベースバックアップ履歴情報
type DatabaseBackupHistory struct {
	CreatedAt    time.Time  `json:"createdat,omitempty" yaml:"created_at,omitempty" structs:",omitempty"`
	Availability string     `json:"availability,omitempty" yaml:"availability,omitempty" structs:",omitempty"`
	RecoveredAt  *time.Time `json:"recoveredat,omitempty" yaml:"recovered_at,omitempty" structs:",omitempty"`
	Size         int64      `json:"size,omitempty" yaml:"size,omitempty" structs:",omitempty"`
}

// ID バックアップ履歴のID取得
func (h *DatabaseBackupHistory) ID() string {
	return h.CreatedAt.Format(time.RFC3339)
}

// FormatCreatedAt 指定のレイアウトで作成日時を文字列化
func (h *DatabaseBackupHistory) FormatCreatedAt(layout string) string {
	return h.CreatedAt.Format(layout)
}

// FormatRecoveredAt 指定のレイアウトで復元日時を文字列化
//
// 復元日時がnilの場合は空の文字列を返す
func (h *DatabaseBackupHistory) FormatRecoveredAt(layout string) string {
	if h.RecoveredAt == nil {
		return ""
	}
	return h.RecoveredAt.Format(layout)
}

// UnmarshalJSON JSON復号処理
func (h *DatabaseBackupHistory) UnmarshalJSON(data []byte) error {
	var tmpMap = map[string]interface{}{}
	if err := json.Unmarshal(data, &tmpMap); err != nil {
		return err
	}

	if recoveredAt, ok := tmpMap["recoveredat"]; ok {
		if strRecoveredAt, ok := recoveredAt.(string); ok {
			if _, err := time.Parse(time.RFC3339, strRecoveredAt); err != nil {
				tmpMap["recoveredat"] = nil
			}
		}
	}

	data, err := json.Marshal(tmpMap)
	if err != nil {
		return err
	}

	tmp := &struct {
		CreatedAt    time.Time  `json:"createdat,omitempty"`
		Availability string     `json:"availability,omitempty"`
		RecoveredAt  *time.Time `json:"recoveredat,omitempty"`
		Size         string     `json:"size,omitempty"`
	}{}
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	h.CreatedAt = tmp.CreatedAt
	h.Availability = tmp.Availability
	h.RecoveredAt = tmp.RecoveredAt
	s, err := strconv.ParseInt(tmp.Size, 10, 64)
	if err == nil {
		h.Size = s
	} else {
		return err
	}

	return nil
}

// DatabaseParameter RDBMSごとに固有のパラメータ設定
type DatabaseParameter struct {
	Parameter *DatabaseParameterSetting `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	Remark    *DatabaseParameterRemark  `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
}

type DatabaseParameterSetting struct {
	NoteID types.ID                     `json:",omitempty" yaml:"note_id,omitempty" structs:",omitempty"`
	Attr   DatabaseParameterSettingAttr `json:",omitempty" yaml:"attr,omitempty" structs:",omitempty"`
}

type DatabaseParameterSettingAttr map[string]interface{}

// UnmarshalJSON 配列/オブジェクトが混在することへの対応
func (d *DatabaseParameterSettingAttr) UnmarshalJSON(b []byte) error {
	if string(b) == "[]" {
		return nil
	}
	type alias map[string]interface{}

	var a alias
	if err := json.Unmarshal(b, &a); err != nil {
		return err
	}
	*d = DatabaseParameterSettingAttr(a)
	return nil
}

type DatabaseParameterRemark struct {
	Settings []interface{}                // どのような値が入るのか不明
	Form     []*DatabaseParameterFormMeta `json:",omitempty" yaml:"form,omitempty" structs:",omitempty"`
}

type DatabaseParameterFormMeta struct {
	Type    string                            `json:"type" yaml:"yaml"`
	Name    string                            `json:"name" yaml:"name"`
	Label   string                            `json:"label" yaml:"label"`
	Options *DatabaseParameterFormMetaOptions `json:"options" yaml:"options"`
	Items   [][]interface{}                   `json:"items,omitempty" yaml:"items,omitempty" structs:",omitempty"` // 例: [["value1", "text1"],[ "value2", "text2"]] ※ valueは数値となる可能性がある
}

type DatabaseParameterFormMetaOptions struct {
	Validator string  `json:"validator" yaml:"validator"`
	Example   string  `json:"ex" yaml:"ex"`
	Min       float64 `json:"min" yaml:"min"`
	Max       float64 `json:"max" yaml:"max"`
	MaxLen    int     `json:"maxlen" yaml:"maxlen"`
	Text      string  `json:"text" yaml:"text"`
	Reboot    string  `json:"reboot" yaml:"reboot"`
	Type      string  `json:"type" yaml:"type"`
	Integer   bool    `json:"integer" yaml:"integer"` // postgres用のパラメータにだけ存在する模様
}
