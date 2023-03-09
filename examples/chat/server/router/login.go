package router

import (
	"xins"
	"xins/examples/chat/object"
	protocol "xins/protocol/xins"
)

func Login(request *xins.Request) {

	session := request.Session()
	message := request.Message()

	data := message.(*protocol.Message)

	var login object.Login

	session.Debugf("[login] [recv] json: %s", data.Data())

	if err := codc.Decode(data.Data(), &login); nil != err {
		return
	}

	id, ok := object.GetUIDByToken(login.Token)
	if !ok {
		session.Debugf("[error] token is not exists")
		session.Close()
		return
	}

	user, err := object.DefaultUserManager.Get(id)

	if nil != err {
		session.Debugf("[error] %s", err)
		session.Close()
		return
	}

	request.Set("login", true)
	request.Set("user", user)

}
