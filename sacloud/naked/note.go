package naked

import "time"

// Note スタートアップスクリプト
type Note struct {
	ID           int64      `json:"ID,omitempty" yaml:"id,omitempty"`
	Name         string     `json:"Name,omitempty" yaml:"name,omitempty"`
	Description  string     `json:"Description,omitempty" yaml:"description,omitempty"`
	Tags         []string   `json:"Tags" yaml:"tags"`
	Availability string     `json:"Availability,omitempty" yaml:"availability,omitempty"`
	Scope        string     `json:"Scope,omitempty" yaml:"scope,omitempty"`
	Class        string     `json:"Class,omitempty" yaml:"class,omitempty"`
	Content      string     `json:"Content,omitempty" yaml:"content,omitempty"`
	Icon         *Icon      `json:"Icon,omitempty" yaml:"icon,omitempty"`
	CreatedAt    *time.Time `json:"CreatedAt,omitempty" yaml:"created_at,omitempty"`
	ModifiedAt   *time.Time `json:"ModifiedAt,omitempty" yaml:"modified_at,omitempty"`
}

// Icon アイコン
type Icon struct {
	ID           int64      `json:"ID,omitempty" yaml:"id,omitempty"`
	Name         string     `json:"Name,omitempty" yaml:"name,omitempty"`
	Tags         []string   `json:"Tags" yaml:"tags"`
	Availability string     `json:"Availability,omitempty" yaml:"availability,omitempty"`
	Scope        string     `json:"Scope,omitempty" yaml:"scope,omitempty"`
	URL          string     `json:"URL,omitempty" yaml:"url,omitempty"`
	CreatedAt    *time.Time `json:"CreatedAt,omitempty" yaml:"created_at,omitempty"`
	ModifiedAt   *time.Time `json:"ModifiedAt,omitempty" yaml:"modified_at,omitempty"`
}
