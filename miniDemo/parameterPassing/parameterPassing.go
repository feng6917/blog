package main

import "fmt"

func main() {
	// 数组传参 值传递 函数内修改无效
	nums := [3]int{1, 2, 3}
	A1(nums)
	fmt.Println("array nums: ", nums)

	// slice传参 引用传递 函数内修改有效
	nums1 := []int{1, 2, 3}
	S1(nums1)
	fmt.Println("slice nums: ", nums1)

	// 数组传参 指针传递 函数内修改有效
	nums2 := [3]int{1, 2, 3}
	A2(&nums2)
	fmt.Println("array pointer nums: ", nums2)

}

func A1(nums [3]int) {
	nums[0] = 111
}

func S1(nums []int) {
	if len(nums) > 0 {
		nums[0] = 111
	}
}

func A2(nums *[3]int) {
	nums[0] = 111
}
