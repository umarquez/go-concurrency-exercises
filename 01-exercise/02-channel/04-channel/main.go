package main

import "fmt"

// TODO: Implement relaying of message with Channel Direction
var chanA, chanB chan string

func genMsg(ch1 chan<- string) {
	// send message on ch1
  for i := 0; i < 10; i++ {
    fmt.Printf("sending msg %v\n", i)
    ch1 <- fmt.Sprintf("This is the message #%v", i)
  }
  close(ch1)
}

func relayMsg(ch1 <-chan string, ch2 chan<- string) {
	// recv message on ch1
	// send it on ch2

  for msg := range ch1 {
    ch2 <- msg+" [chan1]"
  }
  close(ch2)
}

func main() {
	// create ch1 and ch2
  chanA = make(chan string)
  chanB = make(chan string)

	// spine goroutine genMsg and relayMsg
  go genMsg(chanA)
	// recv message on ch2
  go relayMsg(chanA, chanB)

  for msg := range chanB {
    fmt.Println(msg)
  }
}
