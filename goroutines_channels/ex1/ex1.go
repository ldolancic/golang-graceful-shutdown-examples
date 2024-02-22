package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
 Unbuffered channels
*/

func main() {
	ch := make(chan int)

	writerDone := make(chan struct{})
	// writer goroutine
	go func() {
		defer fmt.Println("writer done")
		for {
			select {
			case <-writerDone:
				close(ch)
				return
			default:
				ch <- rand.Intn(10)
				time.Sleep(time.Millisecond * 500)
			}
		}
	}()

	readerDone := make(chan struct{})
	// reader goroutine
	go func() {
		defer close(readerDone)
		defer fmt.Println("reader done")
		for {
			val, ok := <-ch
			if !ok {
				return
			}
			fmt.Println(val)
			time.Sleep(time.Second)
		}
	}()

	time.Sleep(time.Second * 2)
	close(writerDone)
	<-readerDone
	fmt.Println("program done")
}
