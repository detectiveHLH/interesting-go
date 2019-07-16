package main

import "fmt"

func main() {
	arr := []int{6, 5, 2, 3, 1, 7}
	insertionSort(arr)
	fmt.Println(arr)
}

func insertionSort(arr []int) {
	// 控制循环的次数
	for i := 1; i < len(arr); i++ {
		// 每次循环确定一个有序元素
		// 并且 从已经有序的数组中 从后向前比较，
		for j := i - 1; j >= 0; j-- {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			} else {
				// 没有必要比较了，当前如何有序直接结束当前的循环
				// 因为前面的元素肯定都是有序的
				break
			}
		}
	}
}
