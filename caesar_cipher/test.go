package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	name string
	age  int
}

func main() {
	m := map[string]int{"a": 1, "b": 2}
	s := []user{{"bob", 20}, {"alice", 30}}

	// %v default
	fmt.Printf("%%v map: %v\n", m)
	fmt.Printf("%%v slice: %v\n\n", s)

	// %+v fields
	fmt.Printf("%%+v slice: %+v\n\n", s)

	// %#v syntax
	fmt.Printf("%%#v slice: %#v\n\n", s)

	// json pretty print
	b, _ := json.MarshalIndent(s, "", "  ")
	fmt.Printf("json:\n%s\n", string(b))
}
