package main

import "fmt"

func owner () chan int {
  messages := make(chan int)

  go func(ch chan<- int){
    for i := 0; i < 10; i++ {
      ch <- i
    }
    close(ch)
  }(messages)

  return messages
}

func main() {
	//TODO: create channel owner goroutine which return channel and
	// writes data into channel and
	// closes the channel when done.

	consumer := func(ch <-chan int) {
		// read values from channel
		for v := range ch {
			fmt.Printf("Received: %d\n", v)
		}
		fmt.Println("Done receiving!")
	}

	ch := owner()
	consumer(ch)
}
