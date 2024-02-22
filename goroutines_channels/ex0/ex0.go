package main

import "fmt"

/*
 All goroutines are asleep - deadlock!
*/

func main() {
	ch := make(chan int)

	ch <- 5

	val := <-ch
	fmt.Println(val)
}
