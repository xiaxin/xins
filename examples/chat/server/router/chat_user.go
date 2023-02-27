package router

import (
	"xins"
	"xins/examples/chat/server/object"
	protocol "xins/protocol/default"
)

var (
	codc = &xins.JsonCodec{}
)

type ChatUser struct {
}

func (cu *ChatUser) Handle(request xins.Request) {
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
