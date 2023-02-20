package xins

import (
	"fmt"
	"io"
	"runtime/debug"

	// TODO
	protocol "xins/protocol/default"
)

type Protocol interface {
	Pack(interface{}) ([]byte, error)
	Unpack(io.Reader) (interface{}, error)
	// TODO
	Handle(rrouter *Router, request Request)
}

type defaultProtocol struct {
	packer protocol.Packer
}

func NewDefaultProtocol() Protocol {
	return &defaultProtocol{
		packer: protocol.NewDefaultPacker(),
	}
}

func (dp *defaultProtocol) Pack(message interface{}) ([]byte, error) {
	return dp.packer.Pack(message.(*protocol.Message))
}

func (dp *defaultProtocol) Unpack(reader io.Reader) (interface{}, error) {
	return dp.packer.Unpack(reader)
}

func (dp *defaultProtocol) Handle(router *Router, request Request) {
	// todo
	defer func() {
		if r := recover(); r != nil {
			logger.Errorf("%s", fmt.Sprintf("PANIC | %s | %+v \n%s", r, request.SessionID(), debug.Stack()))
		}
	}()

	// todo
	message := request.Message()
	id := message.(*protocol.Message).ID()

	route, err := router.Get(id)

	if nil != err {
		fmt.Printf("[error] [router handle] error:%s", err)
		return
	}
	route.Handle(request)
}
