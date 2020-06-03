package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type ColorGroup struct {
	ID     int
	Name   string
	Colors []string
}

func main() {
	fmt.Println()
	group := ColorGroup{
		ID:     1,
		Name:   "Reds",
		Colors: []string{"crimsion", "Red", "Ruby", "Maroon"},
	}
	b, err := json.Marshal(group)

	if err != nil {
		fmt.Println("error", err)
	}
	os.Stdout.Write(b)
}