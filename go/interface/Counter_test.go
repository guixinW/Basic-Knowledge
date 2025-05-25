package _interface

import (
	"fmt"
	"testing"
)

func TestWordCounter(t *testing.T) {
	var count WordCount
	fmt.Fprintf(&count, "hello, this is a test")
	fmt.Println(count)
}

func TestLineCounter(t *testing.T) {
	var count LineCount
	fmt.Fprintf(&count, "hello, this is a test\ntest for line\ncount should be 3")
	fmt.Println(count)
}
