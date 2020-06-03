package main

import "fmt"

type person struct {
	fname string
	lname string
	age   int
}
type secretAgent struct {
	person
	hasLicenceGun bool
}

func (s secretAgent) operation1() {
	fmt.Println("Agent :", s.fname, s.lname)
	fmt.Println("Operation : Golden Eye")
}
func main() {

	sa1 := secretAgent{
		person: person{
			fname: "James",
			lname: "Bond",
			age:   32,
		},
		hasLicenceGun: true,
	}
	fmt.Println(sa1)
	sa1.operation1()
}
