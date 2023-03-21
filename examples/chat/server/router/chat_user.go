package router

import (
	"xins/core"
	"xins/examples/chat/object"
)

func ChatUser(request core.Context) {
	session := request.Session()
	message := request.Message()

	data := message.(*core.Message)

	var um object.UserMessage

	session.Debugf("[recv] [json] %s", data.Data())

	if err := codc.Unmarshal(data.Data(), &um); nil != err {
		return
	}

	session.Debugf("[recv] [uid:%s] [content:%s]", um.UID, um.Content)
	// 获取用户

}
