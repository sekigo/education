package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {

	var atom atomic.Uint32

	var wg sync.WaitGroup

	for range 500 {
		wg.Go(func() {
			for range 1000 {
				atom.Add(1)
			}
		})
	}

	wg.Wait()
	fmt.Println("atomic =", atom.Load())

}
