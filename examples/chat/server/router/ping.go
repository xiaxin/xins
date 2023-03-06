package router

import (
	"xins"
	protocol "xins/protocol/xins"
)

func Ping(request *xins.Request) {

	session := request.Session()

	session.Debugf("ping")

	if val, exists := request.Get("test"); exists {
		session.Debugf("middleware getval %s", val)
	}

	writeRequest := xins.NewRequest(request.Session(), protocol.NewMessage(1, []byte("pong")))

	request.Session().SendRequest(writeRequest)
}
