package main

import "fmt"

// 冒泡排序
// 通过从左到右不断交换相邻逆序的相邻元素，在一轮的交换之后，可以让未排序的元素上浮到右侧。
// 在一轮循环中，如果没有发生交换，就说明数组已经是有序的，此时可以直接退出
func sort(nums []int) {
	N := len(nums)
	hasSorted := false //检查第一轮是否数组有序
	for i := 0; i < N && !hasSorted; i++ {
		hasSorted = true

		for j := 0; j < N-i-1; j++ {
			if nums[j+1] < nums[j] {
				hasSorted = false
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
		debug(nums)
	}
}

func debug(nums []int){
	fmt.Println("============================")
	for _, n := range nums {
		fmt.Println(n)
	}
}

func main() {
	t := []int{3, 7, 4,5,1}
	sort(t)

}
