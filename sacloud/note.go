package sacloud

import (
	"github.com/sacloud/libsacloud-v2/mapconv"
)

// NoteCreateRequest ccc
type NoteCreateRequest struct {
	Name    string
	Tags    []string
	IconID  int64 `mapconv:"Icon.ID"`
	Class   string
	Content string
}

// ToNaked returns naked Note *コード生成対象*
func (n *NoteCreateRequest) ToNaked() (*NakedNote, error) {
	dest := &NakedNote{}
	err := mapconv.ToNaked(n, dest)
	return dest, err
}

// ParseNaked parse values from naked Note *コード生成対象*
func (n *NoteCreateRequest) ParseNaked(naked *NakedNote) error {
	return mapconv.FromNaked(naked, n)
}

// NoteUpdateRequest bbb
type NoteUpdateRequest struct {
	Name    string
	Tags    []string
	IconID  int64 `mapconv:"Icon.ID"`
	Class   string
	Content string
}

// ToNaked returns naked Note *コード生成対象*
func (n *NoteUpdateRequest) ToNaked() (*NakedNote, error) {
	dest := &NakedNote{}
	err := mapconv.ToNaked(n, dest)
	return dest, err
}

// ParseNaked parse values from naked Note *コード生成対象*
func (n *NoteUpdateRequest) ParseNaked(naked *NakedNote) error {
	return mapconv.FromNaked(naked, n)
}

// NoteCommonResponse aaa
type NoteCommonResponse struct {
	ID           int64
	Name         string
	Tags         []string
	IconID       int64 `mapconv:"Icon.ID"`
	Class        string
	Content      string
	Availability EAvailability
}

// ToNaked returns naked Note *コード生成対象*
func (n *NoteCommonResponse) ToNaked() (*NakedNote, error) {
	dest := &NakedNote{}
	err := mapconv.ToNaked(n, dest)
	return dest, err
}

// ParseNaked parse values from naked Note *コード生成対象*
func (n *NoteCommonResponse) ParseNaked(naked *NakedNote) error {
	return mapconv.FromNaked(naked, n)
}

// IntID idを返す
func (n *NoteCommonResponse) IntID() int64 {
	return n.ID
}
