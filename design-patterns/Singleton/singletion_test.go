package Singleton

import (
	"testing"
)

func TestNewSingleton(t *testing.T) {
	getInt := NewSingleton(func() *int {
		a := 10
		return &a
	})
	s := getInt()
	b := getInt()
	if s != b {
		t.Fatalf("error")
	}
}
