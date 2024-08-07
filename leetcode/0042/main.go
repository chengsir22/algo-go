package main

func trap(height []int) (ans int) {
	n := len(height)
	left := make([]int, n)
	right := make([]int, n)
	left[0], right[n-1] = height[0], height[n-1]
	for i := 1; i < n; i++ {
		left[i] = max(left[i-1], height[i])
	}
	for i := n - 2; i >= 0; i-- {
		right[i] = max(right[i+1], height[i])
	}
	for i, h := range height {
		ans += min(left[i], right[i]) - h
	}
	return
}
