package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	defer fmt.Println("done")

	// signal channel
	sigCh := make(chan os.Signal)
	signal.Notify(sigCh, syscall.SIGINT)

	tick := time.NewTicker(time.Second)

	exit := make(chan struct{})
	workDone := make(chan struct{})

	go func() {
		for {
			select {
			case <-tick.C:
				fmt.Println("tick")
			case <-exit:
				time.Sleep(time.Second)
				fmt.Println("done with ticker")
				close(workDone)
				return
			}
		}
	}()

	<-sigCh
	close(exit)
	<-workDone
	fmt.Println("exiting")
}
