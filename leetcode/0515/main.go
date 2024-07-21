package main

import "github.com/chengsir22/algo-go/utils"

type TreeNode = utils.TreeNode

func largestValues(root *TreeNode) []int {
	var ans []int
	if root == nil {
		return ans
	}
	q := []*TreeNode{root}
	for len(q) > 0 {
		t := q[0].Val
		n := len(q)
		for i := 0; i < n; i++ {
			node := q[0]
			q = q[1:]
			t = max(t, node.Val)
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		ans = append(ans, t)
	}
	return ans
}
