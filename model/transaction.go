package model

type Transaction struct {
	Command    string
	Creator    string
	Hash       string
	Signatures []Signatures
	Timestamp  int64
}
