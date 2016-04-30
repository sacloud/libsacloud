package resources

import "time"

// Note type of startup script
type Note struct {
	*Resource
	Name        string
	Class       string `json:",omitempty"`
	Scope       string `json:",omitempty"`
	Content     string `json:",omitempty"`
	Description string `json:",omitempty"`
	*EAvailability
	CreatedAt  time.Time `json:",omitempty"`
	ModifiedAt time.Time `json:",omitempty"`
	Icon       *Icon     `json:",omitempty"`
	Tags       []string  `json:",omitempty"`
	//TODO Remarkオブジェクトのパース
}
