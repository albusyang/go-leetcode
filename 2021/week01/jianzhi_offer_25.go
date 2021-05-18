package week01

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil { // 一个链表为空直接返回另一个
		return l2
	}
	head := l1

	for l1 != nil && l2 != nil {
		if l2.Val >= l1.Val { // 此时将l2接在l1后面
			if l1.Next == nil || l2.Val < l1.Next.Val {
				t1 := l1.Next
				t2 := l2.Next
				l2.Next = l1.Next
				l1.Next = l2
				l2 = t2
				if t1 != nil && l2 != nil && t1.Val < l2.Val { // [4,5,6][1,2,3]
					l1 = t1
				}
				continue
			}
			l1 = l1.Next
		} else { // 将l2拼在l1前面
			t1 := l2
			t2 := l2.Next
			l2.Next = l1
			l2 = t2
			l1 = t1
			if head.Val > t1.Val {
				head = t1 // head 留在最前面
			}
		}
	}
	return head
}
