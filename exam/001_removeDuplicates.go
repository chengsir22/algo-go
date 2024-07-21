package main

import (
	"fmt"
	"sort"
)

// removeDuplicates 用于删除无序数组中的重复项并返回一个有序数组
func removeDuplicates(nums []int) []int {
	// 使用map来标记已经出现过的元素
	mp := map[int]bool{}
	var ans []int

	// 遍历数组，如果元素未在map中出现，则加入到结果slice中
	for _, num := range nums {
		if _, ok := mp[num]; !ok {
			mp[num] = true
			ans = append(ans, num)
		}
	}

	// 对结果slice进行排序
	sort.Ints(ans)
	return ans
}

func main() {
	// 示例数组
	nums := []int{4, 5, 9, 4, 2, 5, 1, 3, 9, 8}
	fmt.Println("Original array:", nums)

	// 删除重复项并排序
	result := removeDuplicates(nums)
	fmt.Println("Sorted without duplicates:", result)
}
