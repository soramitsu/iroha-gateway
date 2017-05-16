package model

type Account struct {
	PublicKey   string
	Alias       string
	Sigtatories []string
	Quorum      int
	Permissions Permissions
}
