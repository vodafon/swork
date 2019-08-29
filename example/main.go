package main

import (
	"fmt"

	"github.com/vodafon/swork"
)

type Probe struct {
	prefix string
}

func (obj Probe) Process(el string) {
	fmt.Printf("%s%s\n", obj.prefix, el)
}

func main() {
	w := swork.NewWorkerGroup(3, Probe{"TEST: "})

	for _, v := range []string{"A", "B", "C", "D", "E"} {
		w.StringC <- v
	}
	close(w.StringC)
	w.Wait()
}
