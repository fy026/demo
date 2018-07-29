// 选择排序
package main

import (
	"fmt"
)

// 选择排序
// 选择出数组中的最小元素，将它与数组的第一个元素交换位置。再从剩下的元素中选择出最小的元素，
// 将它与数组的第二个元素交换位置。不断进行这样的操作，直到将整个数组排序
func sort(nums []int) {
	N := len(nums)
	for i := 0; i < N; i++ {
		min := i
		for j := i + 1; j < N; j++ {
			if nums[j] < nums[min] {
				min = j
			}

		}
		nums[i], nums[min] = nums[min], nums[i]
	}

	for _, n := range nums {
		fmt.Println(n)
	}
}

func main() {
	t := []int{3, 4, 7, 29, 3, 5, 1, 0, 7}
	sort(t)
}
