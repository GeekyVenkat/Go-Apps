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
type employee struct {
	Id   int
	name string
}

type human interface {
	speak()
}

func (s secretAgent) speak() {
	fmt.Println("I am ", s.fname, s.lname, " - the person speak")
}

func (p person) speak() {
	fmt.Println("I am", p.fname, p.lname, "- the person speak")
}

func bar(h human) {
	h.speak()
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

	p1 := person{
		fname: "Joker",
		lname: "",
		age:   35,
	}

	/*e := employee{
		Id:   1234,
		name: "Venkat",
	}*/
	fmt.Println(p1)
	bar(p1)
	bar(sa1)
	//bar(e)
}
