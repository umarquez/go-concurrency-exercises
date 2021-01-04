package main

import (
	"fmt"
	"sync"
	//"time"
)

var sharedRsc = make(map[string]interface{})

func main() {
	var wg sync.WaitGroup
  var mx = sync.Mutex{}
  var signal = sync.NewCond(&mx)

	wg.Add(2)
	go func() {
		defer wg.Done()

		//TODO: suspend goroutine until sharedRsc is populated.
    signal.L.Lock()
		for len(sharedRsc) == 0 {
			signal.Wait()
		}

		fmt.Println(sharedRsc["rsc1"])
    signal.L.Unlock()
	}()

  go func(){
	// writes changes to sharedRsc
    defer wg.Done()
    signal.L.Lock()
	  sharedRsc["rsc1"] = "lorem ipsum"
    signal.Signal() // This will only unlock one goroutine
    signal.L.Unlock()
  }()

	wg.Wait()
}
