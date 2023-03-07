package main

import "fmt"

/*
插入排序是指在待排序的元素中，
假设前面n-1(其中n>=2)个数已经是排好顺序的，现将第n个数插到前面已经排好的序列中，然后找到合适自己的位置，
使得插入第n个数的这个序列也是排好顺序的。 按照此法对所有元素进行插入，直到整个序列排为有序的过程，称为插入排序 。
*/

func main() {
	nums := []int{43, 1, 32, 45, 56, 78}
	Insertion(nums)
	fmt.Println(nums)
}

func Insertion(nums []int) {
	if len(nums) <= 1 {
		return
	}
	for i := 0; i < len(nums); i++ {
		// 未排序值
		value := nums[i]
		j := i - 1
		for ; j >= 0; j-- {
			// 大于 未排序值 后移
			if nums[j] > value {
				nums[j+1] = nums[j]
			} else {
				break
			}
		}
		// 插入未排序值
		nums[j+1] = value
	}
}
