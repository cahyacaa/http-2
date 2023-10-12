package main

import (
	"log"
	"os"
	"os/signal"
	"server/server"
	"syscall"
)

func main() {
	logger := log.New(os.Stderr, "", 0)
	logger.Println("[WARNING] DON'T USE THE EMBED CERTS FROM THIS EXAMPLE IN PRODUCTION ENVIRONMENT, GENERATE YOUR OWN!")

	go func() {
		server.GinInitHttp2()
	}()

	sc := make(chan os.Signal, 1)
	signal.Notify(
		sc,
		os.Interrupt,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)
	<-sc

	signal.Stop(sc)
}
