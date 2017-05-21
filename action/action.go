package action

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type Commander interface {
	Serialize(builder *flatbuffers.Builder) flatbuffers.UOffsetT
}
