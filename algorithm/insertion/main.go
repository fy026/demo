package main

import "fmt"

// 插入排序
// 插入排序从左到右进行，每次都将当前元素插入到左侧已经排序的数组中，使得插入之后左部数组依然有序。
// 第 j 元素是通过不断向左比较并交换来实现插入过程：当第 j 元素小于第 j - 1 元素，就将它们的位置交换，
// 然后令 j 指针向左移动一个位置，不断进行以上操作
func sort(nums []int) {
	N := len(nums)
	for i:=1;i<N;i++{
		for j:= i ;j>0 && nums[j]< nums[j-1];j--{
			nums[j], nums[j-1] = nums[j-1], nums[j]
		}
	}
}


func debug(nums []int){
	fmt.Println("============================")
	for _, n := range nums {
		fmt.Println(n)
	}
}


func main(){
	t := []int{3, 7, 4,5,1}
	sort(t)
	debug(t)
}
