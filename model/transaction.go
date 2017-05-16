package model

type Transaction struct {
	Command    interface{}
	Creator    string
	Hash       string
	Signatures []Signatures
	Timestamp  int64
}
