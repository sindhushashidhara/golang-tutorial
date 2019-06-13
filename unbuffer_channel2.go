package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {

	in := make(chan int)
	wg.Add(2)
	go player(1, in)
	go player(2, in)

	in <- 0
	wg.Wait()
}

func player (id int , in chan int) {

	defer wg.Done()
	for val := range in {
		fmt.Printf(" %d Got unicorn %d\n", id, val)
		if val % 5 == 0  {
			fmt.Print("Bored")
			close(in)
			return
		}
		val++
		fmt.Printf(" %d Send unicorn %d\n", id, val)
		in <- val
		time.Sleep(100 * time.Millisecond)
	}
}