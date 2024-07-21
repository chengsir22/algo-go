package utils

type ListNode struct {
	Val  int
	Next *ListNode
}
func CreateList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	head := &ListNode{Val: nums[0]}
	cur := head

	for _, num := range nums[1:] {
		cur.Next = &ListNode{Val: num}
		cur = cur.Next
	}

	return head
}
