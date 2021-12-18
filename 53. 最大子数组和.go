package main

const (
	MAX_INT = 0x3f3f3f3f
	MIN_INT = -MAX_INT
)

func max(arr ...int) int {
	maxint := MIN_INT
	for i := 0; i < len(arr); i++ {
		if maxint < arr[i] {
			maxint = arr[i]
		}
	}
	return maxint
}

func maxSubArray(nums []int) int {
	maxSum := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i]+nums[i-1] > nums[i] {
			nums[i] += nums[i-1]
		}
		maxSum = max(maxSum, nums[i])
	}
	return maxSum
}

// func main() {
// 	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
// 	rst := maxSubArray(nums)
// 	fmt.Print(rst)
// }
