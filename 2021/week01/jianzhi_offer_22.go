package week01

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getKthFromEnd(head *ListNode, k int) *ListNode {
	if head == nil || k == 0 {
		return head
	}
	lenListNode := 0
	myHead := head
	for myHead != nil {
		lenListNode++
		myHead = myHead.Next
	}
	myHead = head
	for i := 0; i < lenListNode-k; i++ {
		myHead = myHead.Next
	}
	return myHead
}

// 内存消耗更多
func getKthFromEnd1(head *ListNode, k int) *ListNode {
	if head == nil || k == 0 {
		return head
	}
	lenListNode := 0
	myHead := head
	nodeArr := []*ListNode{}
	for myHead != nil {
		lenListNode++
		nodeArr = append(nodeArr, myHead)
		myHead = myHead.Next
	}
	if k >= 0 && k <= lenListNode {
		myHead = nodeArr[lenListNode-k]
	}
	return myHead
}
