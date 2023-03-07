package main

import "fmt"

/*
冒泡排序（Bubble Sort），是一种计算机科学领域的较简单的排序算法。
它重复地走访过要排序的元素列，依次比较两个相邻的元素，如果顺序（如从大到小、首字母从Z到A）错误就把他们交换过来。
走访元素的工作是重复地进行直到没有相邻元素需要交换，也就是说该元素列已经排序完成。
这个算法的名字由来是因为越小的元素会经由交换慢慢"浮"到数列的顶端（升序或降序排列），
就如同碳酸饮料中二氧化碳的气泡最终会上浮到顶端一样，故名"冒泡排序"。
*/
func main() {
	nums := []int{43, 1, 32, 45, 56, 78}
	Bubble(nums)
	fmt.Println(nums)
}

func Bubble(nums []int) {
	for i := 0; i < len(nums)-1; i++ {
		ok := false
		for j := 0; j < len(nums)-1-i; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
				ok = true
			}
		}
		if !ok {
			break
		}
	}
}
