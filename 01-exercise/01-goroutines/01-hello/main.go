package main

import (
	"fmt"
	"time"
)

const sleepTime = 1000

func fun(s string) {
	for i := 0; i < 3; i++ {
		fmt.Println(s)
		time.Sleep(1 * time.Millisecond)
	}
}

func main() {
	// Direct call
	fun("direct call")

	// TODO: write goroutine with different variants for function call.

	// goroutine function call
  go fun("goroutine function call")

	// goroutine with anonymous function
  go func () {
    fun("anonymous")
  }()

	// goroutine with function value call
  fn := fun
  go fn("value call")

	// wait for goroutines to end
  fmt.Printf("waiting %vms...\n", sleepTime)
  time.Sleep(time.Millisecond*sleepTime)
	fmt.Println("done..")
}
