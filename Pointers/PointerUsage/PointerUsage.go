package main

import "fmt"

func foo(y *int) {
	fmt.Println(y)
	*y = 43
}

func main() {
	x := 0
	foo(&x)
	fmt.Println(x)
}
