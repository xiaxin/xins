package main

import (
	"fmt"
	"xins"
	"xins/core"
	"xins/examples/chat/server/middleware"
	"xins/examples/chat/server/router"

	"os"
	"os/signal"
	"syscall"
)

func main() {
	protocol := core.NewProtocol()

	protocol.AddMiddleware(middleware.AuthMiddleware())

	protocol.AddRoute(1, router.Ping, middleware.AuthMiddleware())

	protocol.AddRoute(11, router.ChatUser, middleware.AuthMiddleware())
	protocol.AddRoute(12, router.ChatGroup, middleware.AuthMiddleware())

	protocol.AddRoute(1000, router.Login)

	s := xins.NewServer(
		xins.ServerProtocol(protocol),
		// xins.SessionOnStart(OnSessionStart),
		// xins.SessionOnStop(OnSessionStop),
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
