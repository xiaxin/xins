package main

import (
	"fmt"
	"xins"
	"xins/example/server/router"
	protocol "xins/protocol/default"

	"os"
	"os/signal"
	"syscall"
)

func main() {
	protocol := protocol.NewDefaultProtocol()
	protocol.AddRoute(1, &router.Ping{})
	protocol.AddRoute(2, &router.Panic{})

	s := xins.NewServer(xins.ServerProtocol(protocol))

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
