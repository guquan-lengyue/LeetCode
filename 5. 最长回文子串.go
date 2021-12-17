package main

// import "fmt"

// func longestPalindrome(s string) string {

// 	palindrome := string(s[0])

// 	for i := 0; i < len(s)-len(palindrome); i++ {
// 		for j := i + 1; j <= len(s); j++ {
// 			if isPalindrome(s[i:j]) && j-i > len(palindrome) {
// 				palindrome = s[i:j]
// 			}
// 		}
// 	}
// 	return palindrome
// }

// func isPalindrome(s string) bool {
// 	for i := 0; i < len(s)/2; i++ {
// 		if s[i] != s[len(s)-i-1] {
// 			return false
// 		}
// 	}
// 	return true
// }

// func main() {
// 	fmt.Print(longestPalindrome("abb"))
// }
