package model

import "github.com/soramitsu/iroha-gateway/action"

type Transaction struct {
	Command    action.Commander
	Creator    string
	Hash       string
	Signatures []Signatures
	Timestamp  int64
}
