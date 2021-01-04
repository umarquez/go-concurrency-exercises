package main

import "fmt"

func main() {
  chVals := make(chan int)
	go func() {
		for i := 0; i < 6; i++ {
			chVals <- i
		}

    close(chVals)
	}()

	// TODO: range over channel to recv values
  for v := range chVals {
    fmt.Println(v)
  }
}
