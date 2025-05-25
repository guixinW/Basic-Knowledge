package _interface

import (
	"bufio"
	"bytes"
)

//func Fprintf(w io.Writer, format string, args ...interface{}) (int, error)

type WordCount int

func (c *WordCount) WordCountAdd(addCount WordCount) {
	*c += addCount
}

func (c *WordCount) Write(p []byte) (n int, err error) {
	reader := bytes.NewReader(p)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		c.WordCountAdd(1)
	}
	return len(p), nil
}

type LineCount int

func (c *LineCount) LineCountAdd(addCount LineCount) {
	*c += addCount
}

func (c *LineCount) Write(p []byte) (n int, err error) {
	reader := bytes.NewReader(p)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		c.LineCountAdd(1)
	}
	return len(p), nil
}
