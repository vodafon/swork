package swork

import (
	"bytes"
	"fmt"
	"io"
	"math/rand"
	"strconv"
	"strings"
	"testing"
	"time"
)

func TestSwork(t *testing.T) {
	wr := bytes.Buffer{}
	w := NewWorkerGroup(3, TestProbe{"TEST: ", &wr})

	for i := 0; i < 8; i++ {
		w.StringC <- strconv.Itoa(i)
	}
	close(w.StringC)
	w.Wait()
	els := []int{}
	sum := 0
	for _, el := range strings.Split(wr.String(), "|") {
		if el == "" {
			continue
		}
		i, err := strconv.Atoi(el)
		if err != nil {
			t.Fatal(err)
		}
		sum += i
		els = append(els, i)
	}
	if len(els) != 8 {
		t.Errorf("Incorrect result. Expected %d, got %d\n", 8, len(els))
	}
	if sum != 28 {
		t.Errorf("Incorrect result. Expected %d, got %d\n", 28, sum)
	}
}

type TestProbe struct {
	prefix string
	writer io.Writer
}

func (obj TestProbe) Process(el string) {
	time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)
	fmt.Fprintf(obj.writer, "|%s", el)
}
