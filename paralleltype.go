package main

import "fmt"

func main() {
	type multiplier float64  // NOTE!! Parallel type ie multiplier type != float64 type
	// multiplier is alias type
	// You can declare your own type using the type keyword
	var m multiplier = 10

	s := 4.5 * m  //  m is type multiplier , not float64

 	fmt.Println(s)


 	const(
 		x = 42
	)

 	type pie float64
 	var y pie = 3.14
	z := 2.5 * x + 54.75 + 10*y

	fmt.Print("Computing....", z)
}