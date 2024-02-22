package main

import (
	"fmt"
	"math/rand"
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

	data := make(chan int, 10)
	exitWriter := make(chan struct{})
	// writer goroutine
	go func() {
		defer fmt.Println("done with writer")
		for {
			select {
			case <-exitWriter:
				close(data)
				return
			default:
				data <- rand.Intn(10)
			}
		}
	}()

	// reader goroutine
	readerDone := make(chan struct{})
	go func() {
		for i := range data {
			fmt.Println(i)
			time.Sleep(time.Second)
		}

		fmt.Println("done with reader")
		close(readerDone)
	}()

	<-sigCh
	close(exitWriter)
	<-readerDone
	fmt.Println("exiting")
}
