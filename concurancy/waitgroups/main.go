package main

import (
	"fmt"
	"sync"
	"time"
)

func tmp(i int) {
	fmt.Println("Fisrt operation number:", i)
	time.Sleep(500 * time.Microsecond)
	fmt.Println("Finished with ", i, "th operation")

}

func main() {
	var wg sync.WaitGroup
	slice := []int{1, 2, 3, 4, 5}

	for i := range slice {
		wg.Go(func() {
			tmp(i)
		})
	}

	wg.Wait()
}
