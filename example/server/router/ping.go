package router

import (
	"xins"
)

type Ping struct {
}

func (bm *Ping) Handle(request xins.Request) {

	writeRequest := xins.NewRequest(request.Session(), xins.NewMessage(1, []byte("pong")))

	request.Session().Send(writeRequest)
}
