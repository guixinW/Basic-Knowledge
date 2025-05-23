// Package crp 责任链模式
package crp

import "fmt"

type Handler interface {
	SetNext(handler Handler)
	Handle(request string) string
}

type BaseHandler struct {
	Next Handler
}

func (b *BaseHandler) SetNext(handler Handler) {
	b.Next = handler
}

func (b *BaseHandler) Handle(request string) string {
	if b.Next != nil {
		return b.Next.Handle(request)
	}
	return "Request cannot be handled"
}

type ConcreteHandlerA struct {
	BaseHandler
}

func (c *ConcreteHandlerA) Handle(request string) string {
	if request != "A" {
		return c.BaseHandler.Handle(request)
	}
	fmt.Printf("ConcreteHandlerA:%v", request)
	return "Handle A"
}

type ConcreteHandlerB struct {
	BaseHandler
}

func (c *ConcreteHandlerB) Handle(request string) string {
	if request != "B" {
		return c.BaseHandler.Handle(request)
	}
	fmt.Printf("ConcreteHandlerB:%v", request)
	return "Handle B"
}
