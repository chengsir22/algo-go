package main

import "github.com/chengsir22/algo-go/utils"

type TreeNode = utils.TreeNode

func findMode(root *TreeNode) []int {
	var pre *TreeNode
	var res []int
	count, maxCount := 0, 0

	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Left)
		if pre != nil && pre.Val == root.Val {
			count++
		} else {
			count = 1
		}
		if count > maxCount {
			maxCount = count
			res = []int{root.Val}
		} else if count == maxCount {
			res = append(res, root.Val)
		}
		pre = root
		dfs(root.Right)
	}
	dfs(root)

	return res
}

func main() {

	root := &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 2}
	root.Right.Left = &TreeNode{Val: 2}

	res := findMode(root)
	for _, v := range res {
		println(v)
	}
}
