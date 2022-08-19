func climbStairs(n int) int {
    last := [2]int{1, 1}

	for i := 0; i < n-1; i++ {
		last[0], last[1] = last[0]+last[1], last[0]
	}

	return last[0]
}
