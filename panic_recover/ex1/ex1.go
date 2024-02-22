package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("before panic")
	go panicFunc()
	time.Sleep(time.Second)
	fmt.Println("main after panic")
}

func panicFunc() {
	defer func() {
		fmt.Println("panicked")
		if r := recover(); r != nil {
			fmt.Println("recovered")
		}
	}()
	panic("panicking")
}
