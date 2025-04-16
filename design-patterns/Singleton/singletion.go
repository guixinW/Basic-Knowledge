package Singleton

import "sync"

type Singleton[T any] struct {
	once     sync.Once
	instance *T
}

func NewSingleton[T any](init func() *T) func() *T {
	var s Singleton[T]
	return func() *T {
		s.once.Do(func() {
			s.instance = init()
		})
		return s.instance
	}
}
