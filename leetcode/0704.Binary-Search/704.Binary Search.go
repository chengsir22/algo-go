package main

import "fmt"

func search704(nums []int, target int) int {
	n := len(nums)
	l, r := 0, n-1
	for l < r {
		mid := (l + r) >> 1
		if nums[mid] >= target {
			r = mid
		} else {
			l = mid + 1
		}
	}
	if nums[l] == target {
		return l
	}
	return -1
}


func main(){
	ans := search704([]int{-1, 0, 3, 5, 9, 12}, 9)
	fmt.Println(ans)
}
