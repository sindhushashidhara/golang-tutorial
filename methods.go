package main

import "fmt"

type (
	worker struct {
		id int
		name string
	}
	supervisor struct {
		worker
		level int
	}
)

func main() {
  w := worker {}

  w.badge(100, "Fred")
  w.info()
  w.setName("Freddy")

  s := supervisor{ }

  s.badge(200, "France", 10)
  s.info()
}


func (w *worker ) badge (id int, name string) {
	w.id = id
	w.name = name
}

func (w *worker ) setName (name string) {
	w.name = name
}

func (w *worker ) info () {
	fmt.Print(w)
}

func (w *supervisor ) badge (id int, name string, level int) {
	w.worker.badge(id, name)
	w.level = level

}

func (w *supervisor ) setName (name string) {
	w.name = name
}


