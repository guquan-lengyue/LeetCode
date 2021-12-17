package main

//import "fmt"
//
//func twoSum(nums []int, target int) []int {
//	//return hashFunc(nums, target)
//	return foreach(nums, target)
//}
//
//func hashFunc(nums []int, target int) []int {
//	num2index := map[int]int{}
//	for i, num := range nums {
//		num2index[num] = i
//	}
//	for i, num := range nums {
//		if num2index[target-num] > i {
//			return []int{i, num2index[target-num]}
//		}
//	}
//	return []int{}
//}
//
//func foreach(nums []int, target int) []int {
//	for i := 0; i < len(nums)-1; i++ {
//		for j := i + 1; j < len(nums); j++ {
//			if nums[i]+nums[j] == target {
//				return []int{i, j}
//			}
//
//		}
//	}
//	return []int{}
//}
//
//func main() {
//	nums := []int{1, 3, 4, 2}
//	fmt.Println(twoSum(nums, 6))
//}
