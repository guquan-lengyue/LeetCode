package main

// import (
// 	"fmt"
// 	"sort"
// )

// func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
// 	nums2 = append(nums2, nums1...)
// 	sort.Ints(nums2)
// 	if len(nums2)%2 != 0 {
// 		return float64(nums2[len(nums2)/2])
// 	} else {
// 		rst := nums2[len(nums2)/2] + nums2[len(nums2)/2-1]
// 		return float64(rst) / 2
// 	}
// }

// func main() {
// 	input1 := []int{1, 2}
// 	input2 := []int{3, 4}
// 	arrays := findMedianSortedArrays(input2, input1)
// 	fmt.Println(arrays)
// }
