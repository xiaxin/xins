package core

import "io"

type Package interface {
	// 打包相关
	Pack(interface{}) ([]byte, error)
	Unpack(io.Reader) (interface{}, error)
}
