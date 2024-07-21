package leetcode

import "fmt"

func twoSum(nums []int, target int) []int {
	m := map[int]int{}
	for i,v:=range nums{
		if j,ok:=m[target-v];ok{
			return []int{j,i}
		}
		m[v]=i
	}
	return nil
}

func main() {
	twoSum := twoSum([]int{3, 2, 4}, 6)
	fmt.Println(twoSum[0], twoSum[1])
}
