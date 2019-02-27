package main

import "fmt"

type(
	PrettyBadger interface{
		printBadge()
	}

	worker struct {
		id int
		name string
	}
	supervisor struct {
		worker
		level int
	}
)

func (w *worker ) printBadge () (string){
	return fmt.Sprintf("%d %s", w.id, w.name)

}

func (w *supervisor ) printBadge () {
	 p := w.worker.printBadge()
	 fmt.Printf("%d %d", p, w.level)
}

func (w *worker ) badge (id int, name string) {
	w.id = id
	w.name = name
}


func (w *supervisor ) badge (id int, name string, level int) {
	w.worker.badge(id, name)
	w.level = level

}

func main() {

	w := worker {}
	w.badge(100, "Fred")
	w.printBadge()

	s := supervisor{ }
	s.badge(200, "France", 10)

	s.printBadge()
}

