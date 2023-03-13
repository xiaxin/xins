package core

type Protocol interface {
	Codec
	Package

	// 生命周期
	Handle(session Session) error

	// 协议信息
	Info() string

	// TODO
	Run()
}
