package sacloud

// Account type of SakuraCloud Account
type Account struct {
	// *Resource //HACK 現状ではAPI戻り値が文字列なためパースエラーになる
	ID    string `json:",omitempty"`
	Class string `json:",omitempty"`
	Code  string `json:",omitempty"`
	Name  string `json:",omitempty"`
}
