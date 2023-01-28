package main

import (
	"fmt"
	"reflect"
	"testing"
)

// Define a test to check if a string is a palindrome

func TestRecursiveStrings(t *testing.T) {
	word1 := "radar"
	t.Run(
		fmt.Sprintf("Testing a valid palindrome: %s", word1),
		func(t *testing.T) {

			got := recursiveStrings(word1)
			want := true

			if got != want {
				t.Errorf("input: %s, got: %t, want: %t", word1, got, want)
			}
		})

	word2 := "notapalindromeordinlapaton"
	t.Run(
		fmt.Sprintf("Testing a valid palindrome: %s", word2),
		func(t *testing.T) {

			got := recursiveStrings(word2)
			want := false

			if got != want {
				t.Errorf("input: %s, got: %t, want: %t", word2, got, want)
			}
		})

}

// Define a test for creating a new slice/array
// with all products from a given range

func TestCreateProducts(t *testing.T) {
	n1 := 1
	n2 := 4
	t.Run(
		fmt.Sprintf("Testing array creation with %d, %d", n1, n2),
		func(t *testing.T) {
			got1, got2 := CreateProducts(n1, n2)
			want1 := [][]int{
				{1, 1},
				{1, 2},
				{1, 3},
				{1, 4},
				{2, 2},
				{2, 3},
				{2, 4},
				{3, 3},
				{3, 4},
				{4, 4},
			}

			want2 := []int{1, 2, 3, 4, 4, 6, 8, 9, 12, 16}
			if !reflect.DeepEqual(got1, want1) {
				t.Errorf("got: %v\n want: %v\n", got1, want1)
			}
			if !reflect.DeepEqual(got2, want2) {
				t.Errorf("got: %v\n want: %v\n", got2, want2)
			}
		})
}
