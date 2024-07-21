package main

import "fmt"

func moveZeroes(nums []int) {
	j := 0
	for i, x := range nums {
		if x != 0 {
			nums[i], nums[j] = nums[j], nums[i]
			j++
		}
	}
}

func main() {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
	for _, v := range nums {
		fmt.Print(v, " ")
	}
}
