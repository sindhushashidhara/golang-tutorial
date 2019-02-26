package main

import "fmt"

type node struct {
	value string
	next *node
}

func main() {
	a := 10
	fmt.Printf("a: %-10T %p:%v\n", a, &a, a)

	b := &a
	fmt.Printf("b: %-10T %p:%p:%v\n", b, &b, b, *b)

	var c *int
	c = &a
	fmt.Printf("c: %-10T %p:%p:%v\n", c, &c, c, *c)

	var d = new(int)
	*d = a
	fmt.Printf("d: %-10T %p:%p:%v\n", d, &d, d, *d)


	na, nb, nc, nd, ne := node{value: "a"}, node{value: "b"}, node{value: "c"}, node{value: "d"}, node{value: "e"}
	na.next, nb.next, nc.next, nd.next = &nb, &nc, &nd, &ne
	h := &na
	printList(h, "before")

	//Delete  c
	nc.value = nc.next.value
	nc.next = nc.next.next

	printList(h, "after")

}


func printList(n *node, m string) {

	for c := n; c != nil; c = c.next {
		fmt.Printf("%s ->", c.value)
	}
	fmt.Printf("nil\n")
}
