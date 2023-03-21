package router

import (
	"xins/core"
)

func Ping(request core.Context) {

	session := request.Session()

	session.Debugf("ping")

	if val, exists := request.Get("test"); exists {
		session.Debugf("middleware getval %s", val)
	}

	ctx := request.Session().NewContext()
	ctx.SetMessage(core.NewMessage(1, []byte("pong")))

	request.Session().Send(ctx)
}
