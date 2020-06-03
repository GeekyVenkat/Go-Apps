package main

import (
	"encoding/json"
	"fmt"
)

type Animal struct {
	Name  string
	Order string
}

func (a Animal) String() string {
	return fmt.Sprintf("%s : %s", a.Name, a.Order)
}
func main() {

	var jsonBlob = []byte(`[
	{"Name": "Platypus", "Order": "Monotremata"},
	{"Name": "Quoll",    "Order": "Dasyuromorphia"}
]`)

	var animals []Animal
	err := json.Unmarshal(jsonBlob, &animals)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("%+v", animals)
	fmt.Println()
	fmt.Println(animals[0])

}
