package main

import (
	"fmt"
	"time"
	"sync"
)
func main() {
	var wg sync.WaitGroup
	size := 5
	choc := make(chan int, size)
	wg.Add(2)
	go wrapChoc("Ethel", choc )
	go wrapChoc("Lucy", choc )
	makeChoc(choc)
	wg.Wait()
}

func wrapChoc(person string, choc <- chan int) {
	defer wg.Done()
	for {
		select {
		case v , ok := <- choc:
			if !ok {
				break
			}
			fmt.Printf("[%s] ->  wraps (%d) ", person, v)
		case <- time.After(time.Duration(100) * time.Millisecond): {
			fmt.Printf("[%s] ->  misses", person)
		}
	}
}

func makeChoc(choc chan<- int) {
	for i := 0 ;i < 10 ; i ++ {
		fmt.Printf("[Roller] ->  choc (%d) ", i)
		choc <- i
	}
	close(choc)
}
