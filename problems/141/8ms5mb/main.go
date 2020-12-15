/**
2 * @Author: Nico
3 * @Date: 2020/12/15 20:44
4 */
package main

type ListNode struct {
	    Val int
	    Next *ListNode
}
func hasCycle(head *ListNode) bool {
	mark := map[*ListNode]bool{}
	if head == nil{
		return false
	}

	for{
		if head.Next == nil{
			break
		}
		if _, ok := mark[head.Next]; ok{
			return true
		}
		mark[head] = true
		head = head.Next
	}
	return false
}
