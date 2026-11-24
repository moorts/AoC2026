package main

import (
	"testing"
	"slices"
)

func TestFmapAddOne(t *testing.T) {
	addOne := func(x int) int {
		return x + 1
	}

	list1 := []int{1,2,3,4,5}
	expected1 := []int{2,3,4,5,6}
	actual1 := fmap(list1, addOne)

	if !slices.Equal(actual1, expected1) {
		t.Errorf("Expected: %v, but got: %v", expected1, actual1)
	}
}
