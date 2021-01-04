package main

import (
  "io"
  "log"
  "net"
  "os"
  "time"
)

func main() {
  // TODO: connect to server on localhost port 8000

  for i := 20; i > 0; i-- {
    go func(){
      cnn, err := net.Dial("tcp", "localhost:8080")
      if err != nil {
        log.Fatal(err)
      }
      defer cnn.Close()

      mustCopy(os.Stdout, cnn)
    }()
  }

  time.Sleep(30*time.Second)
}

// mustCopy - utility function
func mustCopy(dst io.Writer, src io.Reader) {
  if _, err := io.Copy(dst, src); err != nil {
    log.Fatal(err)
  }
}