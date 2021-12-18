package main

import "fmt"

func canJump(nums []int) bool {
	if len(nums) == 1 {
		return true
	}
	if nums[0] == 0 {
		return false
	}
	// 不可到达当前的最后元素的元素个数
	n := 0
	// 当前最后元素的下标
	lastIdx := len(nums) - 1
	for i := lastIdx - 1; i > -1; i-- {
		if lastIdx-i <= nums[i] {

			// 如果可以出现可以到达当前最后元素,
			// 则把n清零
			n = 0
			// 然后移动最后元素下标为当前元素下标
			lastIdx = i
		} else {
			// 如果无法到达当前最后元素, 则累计个数
			n++
		}
	}
	return n == 0
}

func main() {
	nums := []int{3, 2, 1, 0, 4}
	fmt.Print(canJump(nums))
}
