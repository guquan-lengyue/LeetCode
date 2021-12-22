package main

import "fmt"

func isPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	revertedHalfNumber := 0
	for revertedHalfNumber < x {
		revertedHalfNumber = revertedHalfNumber*10 + x%10
		x /= 10
	}
	return x == revertedHalfNumber || x == revertedHalfNumber/10
}

// v1.0
// func isPalindrome(x int) bool {
// 	if x < 0 || (x%10==0&&x!=0) {
// 		return false
// 	}

// 	nums := make([]int, 11)
// 	countLength := 0
// 	for s := x; ; {
// 		y := s % 10
// 		s /= 10
// 		nums[countLength] = y

// 		if s == 0 && y == 0 {
// 			break
// 		}
// 		countLength++

// 	}
// 	for i := 0; i < countLength/2; i++ {
// 		if nums[i]^nums[countLength-1-i] != 0 {
// 			return false
// 		}
// 	}
// 	return true
// }

func main() {
	rst := isPalindrome(121)
	fmt.Print(rst)
}
