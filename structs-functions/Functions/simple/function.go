package main

import "fmt"

func main() {
	fmt.Println()
	//foo()
	x := 30
	bar("James", x)
	fmt.Println(x)
	//s, b := funMulReturn("James", "Bond")
	//fmt.Println(s, b)

}

// func (r Receiver) identifier(parameters) (return(s)) { ... }

func foo() {
	fmt.Println("In Foo function")
}

//Everything is passed by value
func bar(s string, x int) {
	x = x + 20
	fmt.Println("Hello,", s)
	fmt.Println(x)
}

func woo(s string) string {
	return fmt.Sprint("Form woo", s)
}

func funMulReturn(s1 string, s2 string) (string, bool) {
	s := fmt.Sprint(s1, s2, `, doing great!`)
	b := true
	return s, b
}
