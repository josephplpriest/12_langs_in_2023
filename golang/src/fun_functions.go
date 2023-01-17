package main

import "fmt"

func plus(a int, b int) int {
	return a + b
}

func plusPlus(a, b, c int) int {
	return a + b + c
}

func vals() (int, int) {
	return 4, 2
}

func bigNum(nums ...int) int {

	n := 1

	for _, num := range nums {
		n *= num
	}
	return n
}

func main() {
	res := plus(1, 2)
	fmt.Println("1 + 2 =", res)

	res = plusPlus(1, 2, 3)
	fmt.Println("1 + 2 + 3 =", res)

	a, b := vals()

	newRes := plus(a, b)
	fmt.Printf("%d + %d = %d\n\n", a, b, newRes)

	big := bigNum(1, 2, 3, 4, 5, 11)
	fmt.Println(big)
}
