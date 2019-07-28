package main

import "fmt"

func singleNumber(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	m := make(map[int]int)
	for _, v := range nums {
		mv, ok := m[v]
		if !ok {
			m[v] = 1
		} else {
			m[v] = mv + 1
		}
	}

	for k, v := range m {

		if v == 1 {
			return k
		}
	}

	return 0
}

func main() {
	arr := []int{2, 2, 1, 3, 3, 4, 5, 5, 4}
	fmt.Println(singleNumber(arr))
}
