package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	var indexes = map[string]int{}
	var counter = map[string]int{}
	var result = make([]*TreeNode, 0)
	getId(root, indexes, counter, &result)
	return result
}

func getId(node *TreeNode, indexes, counter map[string]int, result *[]*TreeNode) int {
	if node == nil {
		return -1
	}
	s := fmt.Sprintf("%v#%v#%v", node.Val, getId(node.Left, indexes, counter, result), getId(node.Right, indexes, counter, result))
	if id, ok := indexes[s]; ok {
		counter[s]++
		if counter[s] == 2 {
			*result = append(*result, node)
		}
		return id
	}
	id := len(indexes)
	indexes[s] = id
	counter[s] = 1
	return id
}
