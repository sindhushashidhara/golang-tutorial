package main

import "fmt"

var para string = `Thank you for coming in today. Please come again!
Your total is printed below...`

func main(){
	a, b := 10.50, 8.75

	c := a + b

	fmt.Println(para)
	fmt.Print("$ ", c)
}
