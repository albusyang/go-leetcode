package week01

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 先拿到链表长度；
//
// 倒数第k位相当于 正数的链表长度减k位
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

// 直接两个指针，第一个先走k位，然后两个一起走，第一个到nil时返回第二个
func getKthFromEnd2(head *ListNode, k int) *ListNode {
	kHead := head
	myHead := head
	for i := 0; i < k; i++ {
		if kHead == nil {
			break
		}
		kHead = kHead.Next
	}
	for kHead != nil {
		myHead = myHead.Next
		kHead = kHead.Next
	}
	return myHead
}
