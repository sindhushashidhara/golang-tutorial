package main

import (
	"fmt"
	"testing"
	"github.com/stretchr/testify/assert"
)

/*
type user struct {
	firstName string
	lastName string
	age int
}

type userFuns interface {
	greet()
}
*/

func main() {
	var (
		firstName = "Ammy"
		lastName = "Ray"
		age = 100
	)
	//user := user()
	greet(firstName, lastName, age)
}

func greet(firstName string, lastName string, age int ) string{
	return fmt.Sprintf("Hello %s %s, Welcome to Wonderland. You are %d old today", firstName, lastName, age)
}

func TestGreeter1(t *testing.T) {
	assert.Equal(t, "Hello Ammy Ray, Welcome to Wonderland. You are 10 old today", greet("Ammy", "Ray", 100))
}

func BrenchmarkTest(b *testing.B) {
	for i := 0 ; i < b.N; i ++ {
		fmt.Sprint(" Hello")
	}
}