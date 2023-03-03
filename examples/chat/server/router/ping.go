package router

import (
	"xins"
	protocol "xins/protocol/xins"
)

func Ping(request xins.Request) {

	writeRequest := xins.NewRequest(request.Session(), protocol.NewMessage(1, []byte("pong")))

	request.Session().SendRequest(writeRequest)
}
