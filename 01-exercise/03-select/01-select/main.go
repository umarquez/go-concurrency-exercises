package main

import (
	"time"
  "fmt"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "one"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "two"
	}()

	// TODO: multiplex recv on channel - ch1, ch2
  completed := 0
  for completed < 2 {
    select {
      case v := <- ch1:
        fmt.Println(v)
      case v := <- ch2:
        fmt.Println(v)
    }

    completed++
  }
}
