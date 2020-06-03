package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	age       int
}
type SecretAgent struct {
	person
	hasLicencedGun bool
	//firstName      string
}

func main() {

	sa1 := SecretAgent{
		person: person{
			firstName: "James",
			lastName:  "Bond",
			age:       32,
		},
		hasLicencedGun: true,
		//firstName:      "James1",
	}

	p1 := person{
		firstName: "James",
		lastName:  "Bond",
		age:       32,
	}
	fmt.Println(sa1.firstName, sa1.lastName, sa1.age, sa1.hasLicencedGun)
	fmt.Println(sa1)
	fmt.Println(p1)
}
