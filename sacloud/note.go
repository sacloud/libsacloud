package sacloud

import (
	"context"

	"github.com/sacloud/libsacloud-v2/mapconv"
	"github.com/sacloud/libsacloud-v2/sacloud/naked"
)

// NoteAPI hhh
type NoteAPI interface {
	Read(ctx context.Context, id int64) (*NoteCommonResponse, error)
	Create(ctx context.Context, request *NoteCreateRequest) (*NoteCommonResponse, error)
	Update(ctx context.Context, request *NoteUpdateRequest) (*NoteCommonResponse, error)
	Delete(ctx context.Context, id int64) (*NoteCommonResponse, error)
}

// NoteCreateRequest ccc
type NoteCreateRequest struct {
	Name    string
	Tags    []string
	IconID  int64 `mapconv:"Icon.ID"`
	Class   string
	Content string
}

// NoteUpdateRequest bbb
type NoteUpdateRequest struct {
	ID      int64
	Name    string
	Tags    []string
	IconID  int64 `mapconv:"Icon.ID"`
	Class   string
	Content string
}

// ToNaked returns naked Note *コード生成対象*
func (n *NoteUpdateRequest) ToNaked() (*naked.Note, error) {
	dest := &naked.Note{}
	err := mapconv.ToNaked(n, dest)
	return dest, err
}

// ParseNaked parse values from naked Note *コード生成対象*
func (n *NoteUpdateRequest) ParseNaked(naked *naked.Note) error {
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
