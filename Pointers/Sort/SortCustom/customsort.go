package main

import (
	"fmt"
	"sort"
)

type person struct {
	name string
	age  int
}

func (p person) String() string {
	return fmt.Sprintf("%s : %d", p.name, p.age)
}

type ByAge []person

func (a ByAge) Len() int {
	return len(a)
}
func (a ByAge) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
func (a ByAge) Less(i, j int) bool {
	return a[i].name < a[j].name
}

func main() {
	p1 := person{"James", 32}
	p2 := person{"SpiderMan", 28}
	p3 := person{"BatMan", 34}
	p4 := person{"IronMan", 36}

	people := []person{p4, p2, p3, p1}
	fmt.Println(people)
	fmt.Println("After Sorting")
	//sort.Sort(ByAge(people))
	//fmt.Println(people)

	sort.Slice(people, func(i, j int) bool {
		return people[i].age < people[j].age
	})
	fmt.Println(people)
}
