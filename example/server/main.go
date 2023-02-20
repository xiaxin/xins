package main

import (
	"fmt"
	"xins"
	"xins/example/server/router"

	"os"
	"os/signal"
	"syscall"
)

func main() {
	s := xins.NewServer()

	s.AddRoute(1, &router.Ping{})
	s.AddRoute(2, &router.Panic{})
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
