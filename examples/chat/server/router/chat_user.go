package router

import (
	"xins"
	"xins/examples/chat/object"
	protocol "xins/protocol/xins"
)

var (
	codc = &xins.JsonCodec{}
)

func ChatUser(request *xins.Request) {
	session := request.Session()
	message := request.Message()

	data := message.(*protocol.Message)

	var um object.UserMessage

	session.Debugf("[recv] [json] %s", data.Data())

	if err := codc.Decode(data.Data(), &um); nil != err {
		return
	}

	session.Debugf("[recv] [uid:%s] [content:%s]", um.UID, um.Content)
	// 获取用户

}
