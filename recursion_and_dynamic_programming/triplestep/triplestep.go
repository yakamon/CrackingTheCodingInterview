package triplestep

// TripleStep counts way up stairs.
func TripleStep(n int) int {
	if n < 0 {
		return 0
	}
	if n == 0 {
		return 1
	}
	return TripleStep(n-3) + TripleStep(n-2) + TripleStep(n-1)
}

// TripleStep2 counts way up stairs.
func TripleStep2(n int) int {
	memo := make([]int, n+1, n+1)
	for i := range memo {
		memo[i] = -1
	}
	return tripleStep2(n, memo)

}

func tripleStep2(n int, memo []int) int {
	if n < 0 {
		return 0
	}
	if n == 0 {
		return 1
	}
	if memo[n] > -1 {
		return memo[n]
	}
	memo[n] = tripleStep2(n-3, memo) + tripleStep2(n-2, memo) + tripleStep2(n-1, memo)
	return memo[n]
}
