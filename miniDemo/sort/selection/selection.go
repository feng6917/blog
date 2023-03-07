package main

import "fmt"

/*
选择排序法是在要排序的一组数中，选出最小（或最大）的一个数与第一个位置的数交换；
在剩下的数当中找最小的与第二个位置的数交换，即顺序放在已排好序的数列的最后，如此循环，直到全部数据元素排完为止。
*/
func main() {
	nums := []int{43, 1, 32, 45, 56, 78}
	Selection(nums)
	fmt.Println(nums)
}

func Selection(nums []int) {
	if len(nums) <= 1 {
		return
	}

	for i := 0; i < len(nums); i++ {
		// 当前下标
		min := i
		// 获取后续 比该值小的值的下标
		for j := i + 1; j < len(nums); j++ {
			if nums[min] > nums[j] {
				min = j
			}
		}
		// 判断 是否存在 后续比当前下标小的值，存在即交换位置
		if min != i {
			nums[min], nums[i] = nums[i], nums[min]
		}
	}
}
