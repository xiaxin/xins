package xins

import (
	"fmt"

	"github.com/golang/protobuf/proto"
)

var _ Codec = &ProtobufCodec{}

type ProtobufCodec struct{}

func (p *ProtobufCodec) Encode(v interface{}) ([]byte, error) {
	m, ok := v.(proto.Message)
	if !ok {
		return nil, fmt.Errorf("v should be proto.Message but %T", v)
	}
	return proto.Marshal(m)
}

func (p *ProtobufCodec) Decode(data []byte, v interface{}) error {
	m, ok := v.(proto.Message)
	if !ok {
		return fmt.Errorf("v should be proto.Message but %T", v)
	}
	return proto.Unmarshal(data, m)
}
