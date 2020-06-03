package main

import "fmt"

func main() {
	fmt.Println()

	func() {
		fmt.Println("Anonymous Func")
	}()

	func(x int) {
		fmt.Println("X", x)
	}(43)

	f := func() {
		fmt.Println("Func Expression")
	}
	f()
}
