package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 5)

	// writer
	go func() {
		defer fmt.Println("writer done")
		for i := 1; i < 6; i++ {
			ch <- i
		}
		close(ch)
	}()

	readerDone := make(chan struct{})
	// slow reader
	go func() {
		for v := range ch {
			fmt.Println(v)
			time.Sleep(time.Millisecond * 500)
		}
		fmt.Println("reader done")
		close(readerDone)
	}()

	<-readerDone
	fmt.Println("program done")
}
