package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	defer fmt.Println("done")

	// signal channel
	sigCh := make(chan os.Signal)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGHUP, syscall.SIGTERM)

	fmt.Println("waiting for interrupt")
	<-sigCh
	fmt.Println("exiting")
}
