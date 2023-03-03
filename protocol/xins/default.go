package protocol

import (
	"fmt"
	"io"
	"runtime/debug"
	"xins"
)

type defaultProtocol struct {
	packer Packer
	router *Router
}

func NewDefaultProtocol() *defaultProtocol {
	return &defaultProtocol{
		packer: NewDefaultPacker(),
		router: NewRouter(),
	}
}

func (dp *defaultProtocol) Pack(message interface{}) ([]byte, error) {
	return dp.packer.Pack(message.(*Message))
}

func (dp *defaultProtocol) Unpack(reader io.Reader) (interface{}, error) {
	return dp.packer.Unpack(reader)
}

func (dp *defaultProtocol) AddRoute(id uint32, route RouteFunc) {
	dp.router.Add(id, route)
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
