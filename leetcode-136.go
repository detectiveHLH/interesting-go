package main

import (
	"fmt"
	"sort"
)

func singleNumber(nums []int) int {
	if len(nums) == 0 {
		return nums[0]
	}
	sort.Ints(nums)
	fmt.Println(nums)
	for i := 0; i < len(nums)-1; i = i + 2 {
		if nums[i] != nums[i+1] {
			return nums[i]
		}

	}

	return nums[len(nums)-1]
}

func main() {
	arr := []int{2, 2, 1, 3, 3, 4, 5, 5, 4}
	fmt.Println(singleNumber(arr))
}
