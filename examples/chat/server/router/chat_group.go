package router

import (
	"xins/core"
	"xins/examples/chat/object"
)

func ChatGroup(request core.Context) {
	session := request.Session()
	message := request.Message()

	data := message.(*core.Message)

	var gm = new(object.GroupMessage)

	session.Debugf("[recv] [json] %s", data.Data())

	if err := codc.Unmarshal(data.Data(), gm); nil != err {
		session.Debugf("[send] [error] %s", err)
		return
	}

	session.Debugf("[recv] [gid:%s] [uid:%s] [content:%s]", gm.GID, gm.UID, gm.Content)

	if err := object.DefaultGroupManager.SendMessage(gm); nil != err {
		session.Debugf("[send] [error] %s", err)
	}

}
