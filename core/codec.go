package core

import "encoding/json"

// 数据相关
type Codec interface {
	Marshal(v any) ([]byte, error)
	Unmarshal(data []byte, v interface{}) error
}

var _ Codec = &JsonCodec{}

type JsonCodec struct{}

func NewJsonCodec() Codec {
	return &JsonCodec{}
}

func (c *JsonCodec) Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func (c *JsonCodec) Unmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}
