package main

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}
//
//func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
//	result := &l1
//	for {
//		if (*result).Next == nil && (*result).Val == 0 && l2 == nil {
//			if (*result).Val == 0 {
//				*result = nil
//			}
//			return l1
//		}
//		if (*result).Next == nil {
//			(*result).Next = &ListNode{Val: 0}
//		}
//		l2v := 0
//		if l2 != nil {
//			l2v = l2.Val
//		}
//		(*result).Next.Val = ((*result).Val+l2v)/10 + (*result).Next.Val
//		(*result).Val = ((*result).Val + l2v) % 10
//
//		if (*result) != nil {
//			result = &(*result).Next
//		}
//		if l2 != nil {
//			l2 = l2.Next
//		}
//	}
//}
//
//func main() {
//	l1 := &ListNode{
//		Val: 2,
//		Next: &ListNode{
//			Val: 4,
//			Next: &ListNode{
//				Val: 3,
//			},
//		},
//	}
//	l2 := &ListNode{
//		Val: 5,
//		Next: &ListNode{
//			Val: 6,
//			Next: &ListNode{
//				Val:  4,
//				Next: &ListNode{Val: 9},
//			},
//		},
//	}
//	rst := addTwoNumbers(l1, l2)
//	print(rst)
//}
