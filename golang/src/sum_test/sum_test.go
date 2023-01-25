package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 3, 5, 7, -9}

		got := Sum(numbers)
		want := 7

		if got != want {
			t.Errorf("got %d want %d given: %v", got, want, numbers)
		}
	})

	t.Run("collection of 3 numbers", func(t *testing.T) {
		numbers := []int{100, 300, -900}
		got := Sum(numbers)
		want := -500

		if got != want {
			t.Errorf("got %d want %d given: %v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {

	got := SumAll([]int{1, 2}, []int{0, 9}, []int{-5, -7})
	want := []int{3, 9, -12}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v\nwant: %v", got, want)
	}
}
