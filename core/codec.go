package core

type Codec interface {

	// 数据相关
	Marshal(v any) ([]byte, error)
	Unmarshal(data []byte, v interface{}) error
}
