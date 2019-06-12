package main

import (
"sync"
"fmt"
)

var (
	mutex sync.Mutex
	wg sync.WaitGroup
)

func main() {
	wg.Add(4)
	var counter = 0
	for i := 0 ; i < 4; i++ {
		go incCounter(counter)
	}
	wg.Wait()
	fmt.Print("counter", counter)
}


func incCounter(counter int) int {
	defer wg.Done()
	mutex.Lock()
	for i := 0 ; i < 2; i++ {
		counter = counter + 1
	}
	mutex.Unlock()
	return counter
}

