<<<<<<< HEAD
func mergeTwoLists(l1 *ListNode, l2 *ListNode)  (ret *ListNode) {
	if l1 == nil && l2 != nil{
		return l2
	} else if{
		l1 != nil && l2 == nil{
			return l1
		}else if l1 == nil && l2 == nil
		return 
	}
	a := l1
	if l1.val <= l2.val {
		a = l1
		l1 = l1.Next
	}else{
		a = l2
		l2 = l2.Next
	}
	b := a
	for {
		if l1 != nil && l2 != nil{
			if l1 <= l2{
				a.Next = l1.Next
				l1 = l1.Next
				a = a.Next
			}else{
				a.Next = l2.Next
				l2 = l2.Next
				a = a.Next
			}
		}else if l1 == nil && l2 != nil{
			a.Next = l2
			break
		}else if l1 != nil && l2 = nil{
			a.Next = l1
			break
		}else{
			return b
		}
	}
	return b
=======
func mergeTwoLists(l1 *ListNode, l2 *ListNode)  (ret *ListNode) {
	if l1 == nil && l2 != nil{
		return l2
	} else if{
		l1 != nil && l2 == nil{
			return l1
		}else if l1 == nil && l2 == nil
		return 
	}
	a := l1
	if l1.val <= l2.val {
		a = l1
		l1 = l1.Next
	}else{
		a = l2
		l2 = l2.Next
	}
	b := a
	for {
		if l1 != nil && l2 != nil{
			if l1 <= l2{
				a.Next = l1.Next
				l1 = l1.Next
				a = a.Next
			}else{
				a.Next = l2.Next
				l2 = l2.Next
				a = a.Next
			}
		}else if l1 == nil && l2 != nil{
			a.Next = l2
			break
		}else if l1 != nil && l2 = nil{
			a.Next = l1
			break
		}else{
			return b
		}
	}
	return b
>>>>>>> 3101daa1d3c34d250c063e0fc3af2a0ec6cc4302
}