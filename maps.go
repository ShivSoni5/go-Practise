package main

import "fmt"

// map is similar to dictionary in python
func main() {
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:",m)

	v1 := m["k1"]
	fmt.Println("v1: ", v1)

	fmt.Println("len:", len(m))

	delete(m, "k2")
	fmt.Println("map:", m)

	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	n := map[string]int{"foo":1, "bar":2}
	fmt.Println("map:", n)

	k := map[int]string{1:"foo", 2:"bar"}
	fmt.Println("map_inverse:", k)
}
