package main

import "fmt"

func main() {
  chVal := make(chan int)
	go func(a, b int) {
		c := a + b
    chVal <- c
	}(1, 2)
	// TODO: get the value computed from goroutine
	fmt.Printf("computed value %v\n", <-chVal)
}
