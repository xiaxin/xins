package main

import (
	"fmt"
	"xins"
	"xins/examples/chat/object"
	"xins/examples/chat/server/middleware"
	"xins/examples/chat/server/router"

	protocol "xins/protocol/xins"

	"os"
	"os/signal"
	"syscall"
)

func main() {
	protocol := protocol.NewProtocol()

	protocol.AddRoute(11, router.ChatUser, middleware.AuthMiddleware)
	protocol.AddRoute(12, router.ChatGroup, middleware.AuthMiddleware)

	protocol.AddRoute(1000, router.Login)

	s := xins.NewServer(
		xins.ServerProtocol(protocol),
		xins.SessionOnStart(OnSessionStart),
		xins.SessionOnStop(OnSessionStop),
		// TODO 增加最大连接数
	)

	go s.Run(":9900")

	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		s.Stop()
		done <- true
	}()

	<-done

}

func OnSessionStart(session *xins.Session) {
	// TODO
	user := object.NewUser(session, session.ID(), session.ID())
	object.DefaultUserManager.Add(user)
	group, err := object.DefaultGroupManager.Get("1")

	if nil != err {
		session.Debugf("[group add error] %s", err)
		// TODO
		return
	}
	group.AddUser(user)
}

func OnSessionStop(session *xins.Session) {
	session.Debug("stop")
	user, err := object.DefaultUserManager.Get(session.ID())
	if nil != err {
		return
	}
	object.DefaultUserManager.Del(user)

	group, err := object.DefaultGroupManager.Get("1")

	if nil != err {
		session.Debugf("[group add error] %s", err)
		// TODO
		return
	}
	group.DelUser(user)
}
