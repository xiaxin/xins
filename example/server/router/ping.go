package router

import (
	"xins"
	protocol "xins/protocol/default"
)

type Ping struct {
}

func (bm *Ping) Handle(request xins.Request) {

	writeRequest := xins.NewRequest(request.Session(), protocol.NewMessage(1, []byte("pong")))

	request.Session().Send(writeRequest)
}
