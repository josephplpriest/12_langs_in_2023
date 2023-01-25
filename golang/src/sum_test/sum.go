package main

func Sum(numbers []int) int {
	sum := 0
	for _, n := range numbers {
		sum += n
	}

	return sum
}

func SumAll(numsToSum ...[]int) []int {
	result := []int{}
	for _, nums := range numsToSum {
		result = append(result, Sum(nums))
	}
	return result
}
