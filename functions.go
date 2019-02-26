package main

import "fmt"

func main() {
	fred("c1", 1)
	fred("c2", 1, 2)
	fred("c3", 1, 2, 3)
	// NOTE! Deconstruction
	fred("c4", []int{4, 5, 6}...)

	type multiplier func(int) int

	curry := func (m int) multiplier {
		return func(x int) int {return m * x}
	}

	double, triple := curry(2), curry(3)

	fmt.Print(double(10))
	fmt.Print(triple(30))
}

// NOTE! Must be the last arg!
func fred(s string, args ...int) {
	fmt.Println(s, args)
}





