// https://leetcode.com/problems/add-two-numbers/description/
// 002 Add Two Numbers
package main

import (
	"fmt"
)

// ListNode は単一リンクリストのノードを表します。
type ListNode struct {
	Val  int
	Next *ListNode
}


func main() {
	// 例として、2つのリストを作成します。
	l1 := &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}} // 342
	l2 := &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}} // 465

	result := addTwoNumbers(l1, l2) // 結果は807を表すリスト

	// 結果の表示
	for result != nil {
		fmt.Print(result.Val)
		if result.Next != nil {
			fmt.Print(" -> ")
		}
		result = result.Next
	}
	fmt.Println()
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummyHead := &ListNode{}
	current := dummyHead
	carry := 0

	for l1 != nil || l2 != nil || carry != 0 {
		var val1, val2 int
		if l1 != nil {
			val1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			val2 = l2.Val
			l2 = l2.Next
		}

		sum := val1 + val2 + carry
		carry = sum / 10

		current.Next = &ListNode{Val: sum % 10}
		current = current.Next
	}

	return dummyHead.Next
}
