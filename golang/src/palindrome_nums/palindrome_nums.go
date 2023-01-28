package main

import (
	"fmt"
	"strconv"
)

func recursiveStrings(word string) bool {

	if len(word) == 1 {
		return true
	} else if word[0] == word[len(word)-1] {
		s := word[1 : len(word)-1]
		if len(word) == 2 {
			return true
		} else {
			return recursiveStrings(s)
		}
	} else {
		return false
	}
}

func CreateProducts(num1, num2 int) ([][]int, []int) {
	arr1 := [][]int{}

	arr2 := []int{}

	for i := num1; i < num2+1; i++ {
		for j := i; j < num2+1; j++ {
			arr2 = append(arr2, i*j)
			arr1 = append(arr1, []int{i, j})
		}
	}
	return arr1, arr2
}

func main() {

	n1 := 1
	n2 := 100

	nums1, nums2 := CreateProducts(n1, n2)

	largest := 0
	smallest := 1000000

	largest_factors := [][]int{}
	smallest_factors := [][]int{}

	for _, num := range nums2 {
		word := strconv.Itoa(num)
		if recursiveStrings(word) {
			if num <= smallest {
				smallest = num
				if num >= largest {
					largest = num
				}
			}
		}
	}
	for i, num := range nums2 {
		if num == smallest {
			smallest_factors = append(smallest_factors, nums1[i])
			if num == largest {
				largest_factors = append(largest_factors, nums1[i])
			}
		}
	}
	fmt.Println(largest, largest_factors, smallest, smallest_factors)

}
