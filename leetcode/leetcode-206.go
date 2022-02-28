<<<<<<< HEAD
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
=======
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
>>>>>>> 3101daa1d3c34d250c063e0fc3af2a0ec6cc4302
}