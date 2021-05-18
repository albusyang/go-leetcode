package week01

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reversePrint(head *ListNode) []int {
	a := []int{}
	i := head
	for i != nil {
		a = append(a, i.Val)
		i = i.Next
	}
	r := make([]int, len(a))
	for j, v := range a {
		r[len(a)-j-1] = v
	}
	return r
}
func reversePrint1(head *ListNode) []int {
	if head == nil {
		return nil
	}
	a := []int{}
	i := head
	for i != nil {
		a = append(a, i.Val)
		i = i.Next
	}
	for m, n := 0, len(a)-1; m < n; {
		a[m], a[n] = a[n], a[m]
		m++
		n--
	}
	return a
}

func reversePrint2(head *ListNode) []int {
	a := make([]int, 10000)
	nodeSum := 0
	i := head
	for i != nil {
		a[10000-nodeSum-1] = i.Val
		i = i.Next
		nodeSum++
	}
	r := a[10000-nodeSum:]

	return r
}
