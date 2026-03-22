package crp

import (
	"fmt"
	"testing"
)

func TestCrp(t *testing.T) {
	HandlerA := &ConcreteHandlerA{}
	HandlerB := &ConcreteHandlerB{}
	HandlerA.SetNext(HandlerB)
	fmt.Println(HandlerA.Handle("D"))
}
