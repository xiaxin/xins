package router

import (
	"xins/core"
	"xins/examples/chat/object"
)

func Login(request core.Context) {

	session := request.Session()
	message := request.Message()

	data := message.(*core.Message)

	var login object.Login

	session.Debugf("[login] [recv] json: %s", data.Data())

	if err := codc.Unmarshal(data.Data(), &login); nil != err {
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
