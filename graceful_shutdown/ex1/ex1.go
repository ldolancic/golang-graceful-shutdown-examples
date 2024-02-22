package main

/*
 * Without handling signals
 */

import (
	"fmt"
	"time"
)

func main() {
	defer fmt.Println("done")

	go func() {
		defer fmt.Println("ticker done")
		for {
			time.Sleep(time.Second)
			fmt.Println("tick")
		}
	}()

	time.Sleep(time.Second * 5)
	fmt.Println("exiting")
}
