package resources

// SSHKey type of sshkey
type SSHKey struct {
	ID        string `json:",omitempty"`
	PublicKey string `json:",omitempty"`
}
