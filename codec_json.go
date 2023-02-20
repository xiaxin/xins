package xins

import "encoding/json"

var _ Codec = &JsonCodec{}

type JsonCodec struct{}

func (c *JsonCodec) Encode(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

// Decode implements the Codec Decode method.
func (c *JsonCodec) Decode(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}
