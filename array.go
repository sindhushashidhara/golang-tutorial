package main

import "fmt"

func main() {
	a := [3][3]int{ {1, 1, 1}, {2, 2, 2}, {3, 3, 3} }

	b := [3][3]int{ {3,3,3}, {4,4,4}, {5,5,5} }

	c := addTwoArrays(a, b)

	printArray(c)

	d := multiplyTwoArrays(a,b)

	printArray(d)

	cslice := addTwoSlices(a,b)

	printSlice(cslice)

	dSlice := multiplyTwoSlices(a,b)

	printSlice(dSlice)

}

func addTwoArrays(a , b [3][3]int) (c [3][3]int){
	for i := range a {
		for j := range b {
			c[i][j] = a[i][j] + b[i][j]
		}
	}
	return
}


func multiplyTwoArrays(a , b [3][3]int) (c [3][3]int) {

	for  i := range a {
		for j := range b {
			for k := range a {
				c[i][j] += a[i][k] * b[k][j]
			}
		}
	}
	return
}

func printArray(a [3][3]int) {

	for  i := range a  {
		for j := range a {
			fmt.Printf("%d ", a[i][j])
		}
		fmt.Printf("\n")
	}
}


func addTwoSlices(a , b [3][3]int) ([][]int){
	c := make([][]int, len(a))
	for i := range a {
		for j := range b {
			c[i]= append(c[i], a[i][j] + b[i][j])
		}
	}
	return c
}

func multiplyTwoSlices(a , b [3][3]int) ([][]int) {
	c := make([][]int, len(a))

	for  i := range a {
		for j := range b {
			c[i] = append(c[i], 0)
			for k := range a {
				c[i][j] += a[i][k] * b[k][j]
			}
		}
	}
	return c
}


func printSlice(a [][]int) {

	for  i := range a  {
		for j := range a {
			fmt.Printf("%d ", a[i][j])
		}
		fmt.Printf("\n")
	}
}