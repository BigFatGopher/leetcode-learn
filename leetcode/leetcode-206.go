func reverseList(head *ListNode) *ListNode {
	cur := head
	var prev *ListNode
	for cur != nil {
		Next := cur.Next
		cur.Next = prev
		prev = cur //重新初始化
		cur = Next

	}
	return prev
}