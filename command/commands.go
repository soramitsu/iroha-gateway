package command

import flatbuffers "github.com/motxx/flatbuffers/go"

type Commander interface {
	Serialize() flatbuffers.UOffsetT
}
