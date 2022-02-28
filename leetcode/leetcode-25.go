/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func revert(head *ListNode) *ListNode {
	var pre *ListNode
	curr := head
	for curr != nil {
		pre, curr, curr.Next = curr, curr.Next, pre
	}
	return pre
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	dummy := &ListNode{Val: -1}
	dummy.Next = head
	pre, curr, end := dummy, dummy, dummy
	for end != nil {
		// 获取当前节点
		curr = pre.Next
		// 获取结束结点
		for i := 0; i < k; i++ {
			end = end.Next
			if end == nil {
				return dummy.Next
			}
		}
		// 暂存后节点
		next := end.Next
		end.Next = nil // 至空
		pre.Next = revert(curr)
		pre, end = curr, curr // 都是当前值
		curr.Next = next      //需要把链表串起来
	}
	return dummy.Next
}