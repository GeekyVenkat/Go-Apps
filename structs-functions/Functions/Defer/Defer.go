package main

import "fmt"

func main() {
	//fmt.Println()
	defer foo()
	defer test()
	bar()
	woo()
}
func test() {
	fmt.Println("Test")
}

func woo() {
	fmt.Println("this is woo")
}

func foo() {
	fmt.Println("This is foo")
}

func bar() {
	fmt.Println("This is bar")
}
