package xins

import (
	"io"
)

type Protocol interface {
	Codec
	Pack(interface{}) ([]byte, error)
	Unpack(io.Reader) (interface{}, error)

	Handle(session *Session) error
}
