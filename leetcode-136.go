package main

import (
	"fmt"
)

func singleNumber(nums []int) int {
	var t int
	for k, num := range nums {
		if k == 0 {
			t = num
		} else {
			t = t ^ num
		}
	}

	return t
}

func main() {
	arr := []int{2, 2, 1, 3, 3, 4, 5, 5, 4}
	fmt.Println(singleNumber(arr))
}
