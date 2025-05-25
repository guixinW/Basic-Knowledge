package main

import "fmt"

func A() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("A GET PANIC BEFORE EXIT:%v\n", r)
		}
		fmt.Println("A EXIT")
	}()
	func() {
		panic("NO NAME ERROR")
		fmt.Println("NO NAME FUNC, NEVER EXECUTE")
	}()
}

func main() {
	defer func() {
		//捕获函数仅仅能够捕获一次
		if r := recover(); r != nil {
			fmt.Printf("MAIN GET PANIC BEFORE EXIT:%v\n", r)
		}
		fmt.Println("MAIN EXIT")
	}()
	A()
}
