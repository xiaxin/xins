package core

import "io"

type Protocol interface {
	Codec
	Package

	Codec() Codec

	// 处理逻辑
	Handle(session Session) error

	// 协议信息
	Info(addr string) string
}

var _ Protocol = NewProtocol()

type protocol struct {
	codec Codec

	packer *packer
	router *Router
}

// 测试协议
func NewProtocol() *protocol {
	return &protocol{
		codec: NewJsonCodec(),

		packer: NewPacker(),
		router: NewRouter(),
	}
}

func (d *protocol) Codec() Codec {
	return d.codec
}

func (d *protocol) NewMessage(id uint32, data interface{}) (*Message, error) {

	bytes, err := d.codec.Marshal(data)

	if nil != err {
		return nil, err
	}

	return NewMessage(id, bytes), nil
}

func (d *protocol) Handle(session Session) error {

	message, err := d.packer.Unpack(session.Conn())

	if nil != err {
		return err
	}

	if nil != err {
		return err
	}

	ctx := session.NewContext()
	ctx.SetMessage(message)

	go d.router.HandleRequest(ctx)

	return nil
}

func (p *protocol) Info(addr string) string {
	return "default"
}

// package
func (p *protocol) Pack(message interface{}) ([]byte, error) {
	return p.packer.Pack(message.(*Message))
}

func (p *protocol) Unpack(reader io.Reader) (interface{}, error) {
	return p.packer.Unpack(reader)
}

// codec
func (p *protocol) Marshal(v interface{}) ([]byte, error) {
	return p.codec.Marshal(v)
}

func (p *protocol) Unmarshal(data []byte, v interface{}) error {
	return p.codec.Unmarshal(data, v)
}

//

func (p *protocol) AddRoute(id uint32, route RouteFunc, middlewares ...RouteFunc) {
	p.router.Add(id, route, middlewares...)
}

func (p *protocol) AddMiddleware(middlewares ...RouteFunc) {
	p.router.AddMiddleware(middlewares...)
}
