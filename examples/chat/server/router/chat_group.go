package router

import (
	"xins"
	"xins/examples/chat/server/object"
	protocol "xins/protocol/default"
)

type ChatGroup struct {
}

func (cg *ChatGroup) Handle(request xins.Request) {
	session := request.Session()
	message := request.Message()

	data := message.(*protocol.Message)

	var gm = new(object.GroupMessage)

	session.Debugf("[recv] [json] %s", data.Data())

	if err := codc.Decode(data.Data(), gm); nil != err {
		session.Debugf("[send] [error] %s", err)
		return
	}

	session.Debugf("[recv] [gid:%s] [uid:%s] [content:%s]", gm.GID, gm.UID, gm.Content)

	if err := object.DefaultGroupManager.SendMessage(gm); nil != err {
		session.Debugf("[send] [error] %s", err)
	}

}
