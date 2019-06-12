package main

import (
	"sync"
	"fmt"
)

//var wg sync.WaitGroup

func main() {
	wg.Add(2)
	go count(10, 100)
	go count(100, 200)

	wg.Wait()

	fmt.Print("All done")
}

func count(min, max int) {
	defer func() {
		fmt.Print(" done counter")
		wg.Done()
	}()

	for i := min; i < max; i++ {
		fmt.Printf(" %d", i)
	}
}