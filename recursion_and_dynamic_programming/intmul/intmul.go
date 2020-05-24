package intmul

// Mul multiples a and b.
func Mul(a, b int) int {
	ans := 0
	for i := 0; i < b; i++ {
		ans += a
	}
	return ans
}

// Mul2 multiples a and b.
func Mul2(a, b int) int {
	multiplicand, multiplier := a, b
	if multiplicand < multiplier {
		multiplicand, multiplier = multiplier, multiplicand
	}
	ans := 0
	for i := 0; i < multiplier; i++ {
		ans += multiplicand
	}
	return ans
}

// Mul3 multiplies a and b.
func Mul3(a, b int) int {
	smaller, bigger := a, b
	if smaller > bigger {
		smaller, bigger = bigger, smaller
	}
	return mul3(smaller, bigger)
}

func mul3(smaller, bigger int) int {
	if smaller == 0 {
		return 0
	} else if smaller == 1 {
		return bigger
	}
	next := smaller >> 1
	side := mul3(next, bigger)
	if smaller%2 == 1 {
		return side + side + bigger
	}
	return side + side
}
