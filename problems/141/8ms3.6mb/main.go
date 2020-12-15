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
	if head == nil{
		return false
	}
	c := 0
	for{
		if head.Next == nil{
			break
		}
		head = head.Next
		c ++
		if c > 10000{
			return true
		}
	}
	return false
}

