package main

import "fmt"

func longestConsecutive(nums []int) (res int) {
	st := map[int]bool{}
	for _, v := range nums {
		st[v] = true
	}
	for _, x := range nums {
		if !st[x-1] {
			y := x + 1
			for st[y] {
				y++
			}
			res = max(res, y-x)
		}
	}
	return
}

func main() {
	fmt.Print(longestConsecutive([]int{100, 4, 200, 1, 3, 2}))
}
