package main

func lastRemaining(n int) int {
	head := 1
	step := 1
	left := true
	for n > 1 {
		if left || n%2 == 1 {
			head += step
		}

		step<<=1
		n>>=1
		left=!left
	}
	return head
}


