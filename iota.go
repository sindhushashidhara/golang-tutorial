package main

import "fmt"

func main() {

	const (
		Win = iota
		Loose = iota
		Draw = iota
	)

	fmt.Println("Player status:", Win, Loose, Draw)


	type Weather int

	const (
		Cloudy Weather = iota + 1 // 0 + 1
		Sunny
		Snowy
		_
		_
		Rainy
	)


	fmt.Println("weather status:", Cloudy, Sunny, Snowy, Rainy)

	const (
		a = 1        // 0
		b = 2        // 1
		c = iota + 1 // 2 + 1
		d            // 3 + 1
	)
	fmt.Println("Block:", a, b, c, d)


	const (
		storeName = "ACME"
		taxRate = 7.5
		item1 = 10.50
		item2 = 8.75
	)

	item1AfterTax := (taxRate * item1 ) / 100 + item1
	item2AfterTax := (taxRate * item2 ) /100 + item1
	totalItemsAfterTax := item1AfterTax + item2AfterTax

	fmt.Printf("Thanks for stopping by the %s Store. Please come again!\n", storeName)
	fmt.Println("You owe us: $ ", totalItemsAfterTax)

}
