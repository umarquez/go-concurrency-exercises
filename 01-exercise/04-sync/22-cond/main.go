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

	wg.Add(1)
	go func() {
		defer wg.Done()

		//TODO: suspend goroutine until sharedRsc is populated.
    signal.L.Lock()
		for len(sharedRsc) == 0 {
      signal.Wait()
			//time.Sleep(1 * time.Millisecond)
		}

		fmt.Println(sharedRsc["rsc1"])
    signal.L.Unlock()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		//TODO: suspend goroutine until sharedRsc is populated.
    signal.L.Lock()
		for len(sharedRsc) == 0 {
      signal.Wait()
			//time.Sleep(1 * time.Millisecond)
		}

		fmt.Println(sharedRsc["rsc2"])
    signal.L.Unlock()
	}()

  wg.Add(1)
  go func(){
    signal.L.Lock()
    // writes changes to sharedRsc
    sharedRsc["rsc1"] = "foo"
    sharedRsc["rsc2"] = "bar"
    signal.L.Unlock()
    signal.Broadcast() // This will unlock all goroutines, not only one
    wg.Done()
  }()

	wg.Wait()
}
