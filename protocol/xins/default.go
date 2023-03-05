package protocol

import (
	"fmt"
	"io"
	"runtime/debug"
	"xins"
)

type defaultProtocol struct {
	packer Packer
	codec  xins.Codec
	router *Router
}

func NewDefaultProtocol() *defaultProtocol {
	return &defaultProtocol{
		packer: NewDefaultPacker(),
		codec:  &xins.JsonCodec{},
		router: NewRouter(),
	}
}

func (dp *defaultProtocol) NewMessage(id uint32, data interface{}) (*Message, error) {

	bytes, err := dp.codec.Encode(data)

	if nil != err {
		return nil, err
	}

	return NewMessage(id, bytes), nil
}

func (dp *defaultProtocol) Encode(v interface{}) ([]byte, error) {
	return dp.codec.Encode(v)
}

func (dp *defaultProtocol) Decode(data []byte, v interface{}) error {
	return dp.codec.Decode(data, v)
}

func (dp *defaultProtocol) Pack(message interface{}) ([]byte, error) {
	return dp.packer.Pack(message.(*Message))
}

func (dp *defaultProtocol) Unpack(reader io.Reader) (interface{}, error) {
	return dp.packer.Unpack(reader)
}

func (dp *defaultProtocol) AddRoute(id uint32, route RouteFunc, middlewares ...MiddlewareFunc) {
	dp.router.Add(id, route, middlewares...)
}

func (dp *defaultProtocol) AddMiddleware(middlewares ...MiddlewareFunc) {
	dp.router.AddMiddleware(middlewares...)
}

func (dp *defaultProtocol) Handle(session *xins.Session) error {

	message, err := dp.Unpack(session.Conn().GetTCPConn())

	if nil != err {
		return err
	}

	request := xins.NewRequest(session, message)

	go dp.handle(request)

	return nil
}

func (dp *defaultProtocol) handle(request xins.Request) {
	// todo
	defer func() {
		if r := recover(); r != nil {
			// TODO
			fmt.Printf("%s", fmt.Sprintf("PANIC | %s | %+v \n%s", r, "", debug.Stack()))
		}
	}()

	// todo
	message := request.Message()
	id := message.(*Message).ID()

	router := dp.router

	route, err := dp.router.Get(id)

	if nil != err {
		fmt.Printf("[error] [router handle] error:%s", err)
		return
	}

	var mws = router.middlewares
	if v, has := router.routeMiddlewares[id]; has {
		mws = append(mws, v...) // append to global ones
	}

	wrapped := route
	for i := len(mws) - 1; i >= 0; i-- {
		m := mws[i]
		wrapped = m(wrapped)
	}

	wrapped(request)
}
