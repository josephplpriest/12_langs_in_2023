package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["key1"] = 5
	m["key2"] = 7

	fmt.Println("map:", m)

	v1 := m["key1"]

	fmt.Println("len:", len(m))

	fmt.Println(v1)

	names := []string{"Paul", "Dave", "Dan"}

	heights := []int{71, 74, 73}

	brosByHeight := make(map[string]int)

	for i := range names {
		brosByHeight[names[i]] = heights[i]
	}
	fmt.Println(brosByHeight)
}
