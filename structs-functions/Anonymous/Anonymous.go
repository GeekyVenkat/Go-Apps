package main

import "fmt"

func main() {

	p1 := struct {
		firstName string
		lastName  string
		age       int
	}{
		firstName: "James",
		lastName:  "Bond",
		age:       32,
	}
	fmt.Println(p1.lastName, p1.firstName, p1.age)
}
