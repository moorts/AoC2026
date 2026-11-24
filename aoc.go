package main

import (
	"fmt"
)

func addOne(x int) int {
	return x + 1;
}

func main() {
	test := []int{1,2,3,4}
	test2 := []string{"hello", "there"}
	test3 := []string{}
	test4 := []int{}

	out := fmap(test[:], addOne)

	out2 := sum(test[:])
	out3 := sum(test2[:])
	out4 := sum(test3[:])
	out5 := sum(test4[:])

	fmt.Println(out)
	fmt.Println(out2)
	fmt.Println(out3)
	fmt.Println(out4)
	fmt.Println(out5)
}
