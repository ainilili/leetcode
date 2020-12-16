/**
2 * @Author: Nico
3 * @Date: 2020/12/16 10:16
4 */
package _ms2_5mb

 type ListNode struct {
	 Val int
	 Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var p, q *ListNode = nil, head
	if head == nil || head.Next == nil{
		return head
	}
	for{
		if q == nil{
			break
		}
		n := q.Next
		q.Next = p
		p = q
		q = n
	}
	return p
}