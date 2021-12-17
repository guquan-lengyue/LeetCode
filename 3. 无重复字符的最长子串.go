package main

//import (
//	"fmt"
//)
//
//type window struct {
//	str    []uint8
//	Length int
//}
//
//func (w *window) clear() {
//	for i := range w.str {
//		w.str[i] = 0
//	}
//	w.Length = 0
//}
//
//func (w *window) expend(charStr uint8) bool {
//	if w.str[charStr] < 1 {
//		w.str[charStr]++
//		w.Length++
//		return true
//	} else {
//		return false
//	}
//}
//
//func lengthOfLongestSubstring(s string) int {
//	w := window{str: make([]uint8, 128)}
//
//	maxLength := 0
//	// window begin in s index
//	for wbi := 0; wbi+maxLength < len(s); wbi++ {
//		for i := wbi; i < len(s); i++ {
//			if w.expend(s[i]) {
//				if maxLength < w.Length {
//					maxLength = w.Length
//				}
//			} else {
//				w.clear()
//				break
//			}
//		}
//	}
//	return maxLength
//}
//
//func main() {
//	input := "abcabcbb"
//	rst := lengthOfLongestSubstring(input)
//	fmt.Println(rst)
//}
