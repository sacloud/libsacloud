package sacloud

import "context"

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
	IconID  int64
	Class   string
	Content string
}

// NoteUpdateRequest bbb
type NoteUpdateRequest struct {
	ID      int64
	Name    string
	Tags    []string
	IconID  int64
	Class   string
	Content string
}

// NoteCommonResponse aaa
type NoteCommonResponse struct {
	ID           int64
	Name         string
	Tags         []string
	IconID       int64
	Class        string
	Content      string
	Availability EAvailability
}
