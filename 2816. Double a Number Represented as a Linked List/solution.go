package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	doubleIt(&ListNode{1, &ListNode{8, &ListNode{9, nil}}})
}

func doubleIt(head *ListNode) *ListNode {
	cur := head
	head = &ListNode{0, head}
	prev := head
	for cur != nil {
		temp := cur.Val * 2
		cur.Val = temp % 10
		prev.Val += temp / 10
		prev = cur
		cur = cur.Next
	}
	if head.Val == 0 {
		return head.Next
	}
	return head
}
