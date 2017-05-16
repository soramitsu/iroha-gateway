package model

import (
	"github.com/soramitsu/iroha-gateway/command"
)

type Transaction struct {
	Command    command.Commander
	Creator    string
	Hash       string
	Signatures []Signatures
	Timestamp  int64
}
