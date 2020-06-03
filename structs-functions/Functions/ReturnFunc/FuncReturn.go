package main

import "fmt"

func bar() func(i int, j int) int {
	return func(x int, y int) int {
		return x * y
	}
}
func main() {
	fmt.Println()
	f := bar()
	i := f(4, 5)
	fmt.Println("i", i)
}
