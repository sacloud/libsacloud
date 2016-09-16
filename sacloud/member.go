package sacloud

// Member type of SakuraCloud Account
type Member struct {
	Class string `json:",omitempty"`
	Code  string `json:",omitempty"`
	// Errors [unknown type] `json:",omitempty"`
}
