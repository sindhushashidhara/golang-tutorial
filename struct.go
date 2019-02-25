package main

import "fmt"

const (
	rotten, filmNoir = iota, 100 + iota
	green, comedy = iota + 1, filmNoir + iota
	fresh, documentary = iota + 2, comedy + iota
)

type movie struct {
	actor Actor
	genre , title string
	year int
	rating float64
}

type Actor struct  {
	fName, lName string
}

func main() {

	actor := Actor {
		fName: "John",
		lName: "Cleese",
	}

	movie := movie {
		actor:actor,
		genre: "Classic",
		year: 2018,
		title: "Rango II",
		rating: 9.4,
	}

	fmt.Printf("COMING SOON IN %v!!\n", movie.year)
	fmt.Printf("\t Don't miss `%s starring %s %s!\n\n", movie.title, movie.actor.fName, movie.actor.lName)


	fmt.Printf("It has rating %v! Movie is of type %v\n", green, comedy)

}