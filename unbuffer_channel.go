package main

import (
	"fmt"
)

func main() {
	/* Example 1
	var c = make(chan struct{})

	go func(c chan <- struct{} ) {
		defer func() {
			fmt.Print("DB backup to S3 is complete")
		}()
		c <- struct{}{}
		close(c)
	}(c)

	_ <- c // Receiver is blocked until DB backup is complete

	fmt.Print("Done") */

	//  Example 1
	c := make(chan int)

	go func(c chan <- int ) {
		defer func() {
			fmt.Print("DB backup to S3 is complete")
		}()
		for v := range c {
			fmt.Printf("IMG-%-2d ", v)
		}
	}(c)

	for i := 0 ; i < 10 ; i++ {
		c <- i
	}
	close(c)
}
