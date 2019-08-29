# Swork

Helper for processing a slice of strings like Stdin concurrently.

## Examlple

```go
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
```

```bash
$ go run main.go
TEST: A
TEST: D
TEST: E
TEST: B
TEST: C
```
