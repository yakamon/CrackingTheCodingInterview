package fib

// Fib1 computes fibonacci sequence.
func Fib1(n int) int {
	if n < 0 {
		return 0
	}
	if n < 2 {
		return n
	}
	return Fib1(n-1) + Fib1(n-2)
}

// Fib2 computes fibonacci sequence.
func Fib2(n int) int {
	return fib2(n, make([]int, n+1, n+1))
}

func fib2(n int, memo []int) int {
	if n < 0 {
		return 0
	}
	if n < 2 {
		return n
	}
	if memo[n] == 0 {
		memo[n] = fib2(n-1, memo) + fib2(n-2, memo)
	}
	return memo[n]
}

// Fib3 computes fibonacci sequence.
func Fib3(n int) int {
	if n < 0 {
		return 0
	}
	if n < 2 {
		return n
	}

	memo := make([]int, n, n)
	memo[0], memo[1] = 0, 1
	for i := 2; i < n; i++ {
		memo[i] = memo[i-2] + memo[i-1]
	}
	return memo[n-2] + memo[n-1]
}

// Fib4 computes fibonacci sequence.
func Fib4(n int) int {
	if n < 0 {
		return 0
	}
	a, b := 0, 1
	for i := 0; i < n; i++ {
		a, b = b, a+b
	}
	return a
}
