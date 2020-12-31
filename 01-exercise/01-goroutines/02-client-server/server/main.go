package main

import (
	"io"
	"net"
	"time"
  "log"
  "fmt"
)

type ClientInstance struct {
  cnn *net.Conn
  id int64
}

var cnnCounter int64 = 0

func main() {
	// TODO: write server program to handle concurrent client connections.
  listener, err := net.Listen("tcp", "localhost:8080")
  if err != nil {
    log.Fatal(err)
  }

  for {
    cnn, err := listener.Accept()
    cnnCounter++
    log.Printf("new connection [id:%v] arrives", cnnCounter)
    if err != nil {
      log.Printf("ERROR - can't accept connection[id:%v]: %v", cnnCounter, err)
      continue
    }

    go handleConn(&ClientInstance{
      cnn: &cnn,
      id: cnnCounter,
    })
  }
}

// handleConn - utility function
func handleConn(c *ClientInstance) {
  cnn := *c.cnn
	defer cnn.Close()
	for i := (c.id % 10)*5; i > 0; i-- {
		_, err := io.WriteString(
      cnn, 
      fmt.Sprintf("response from server to connection %v\n", c.id),
    )
		if err != nil {
			return
		}
		time.Sleep(time.Second/5)
	}
}
