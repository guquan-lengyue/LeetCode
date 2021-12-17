package main

// import "fmt"

// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

// func reverseList(head *ListNode) *ListNode {
// 	return forEach(head)
// }

// func recursive(current *ListNode) *ListNode {
// 	if current == nil || current.Next == nil {
// 		return current
// 	}
// 	res := recursive(current.Next)
// 	current.Next.Next = current
// 	current.Next = nil
// 	return res
// }

// func forEach(head *ListNode) *ListNode {
// 	var pre *ListNode
// 	cur := head
// 	for cur != nil {
// 		next := cur.Next
// 		cur.Next = pre
// 		pre = cur
// 		cur = next
// 	}
// 	return pre
// }

// func main() {
// 	rst := reverseList(
// 		&ListNode{
// 			Val: 1,
// 			Next: &ListNode{
// 				Val: 2,
// 				Next: &ListNode{
// 					Val: 3,
// 				},
// 			},
// 		},
// 	)
// 	fmt.Print(rst)
// }
