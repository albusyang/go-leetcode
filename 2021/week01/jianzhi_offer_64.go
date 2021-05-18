package week01

// 求 1+2+...+n ，
//要求不能使用乘除法、for、while、if、else、switch、case等关键字及条件判断语句（A?B:C）。
// 1 <= n <= 10000
func sumNums(n int) int {
	m := 0

	var sumCalculator func(n int) bool
	sumCalculator = func(n int) bool {
		m += n
		// return n > 0 && sumCalculator(n-1)
		return n <= 0 || sumCalculator(n-1)
	}
	sumCalculator(n)
	return m
}

func sumNums1(n int) int {
	ans, A, B := 0, n, n+1
	addGreatZero := func() bool {
		ans += A
		return ans > 0
	}

	_ = ((B & 1) > 0) && addGreatZero()
	A <<= 1
	B >>= 1

	_ = ((B & 1) > 0) && addGreatZero()
	A <<= 1
	B >>= 1

	_ = ((B & 1) > 0) && addGreatZero()
	A <<= 1
	B >>= 1

	_ = ((B & 1) > 0) && addGreatZero()
	A <<= 1
	B >>= 1

	_ = ((B & 1) > 0) && addGreatZero()
	A <<= 1
	B >>= 1

	_ = ((B & 1) > 0) && addGreatZero()
	A <<= 1
	B >>= 1

	_ = ((B & 1) > 0) && addGreatZero()
	A <<= 1
	B >>= 1

	_ = ((B & 1) > 0) && addGreatZero()
	A <<= 1
	B >>= 1

	_ = ((B & 1) > 0) && addGreatZero()
	A <<= 1
	B >>= 1

	_ = ((B & 1) > 0) && addGreatZero()
	A <<= 1
	B >>= 1

	_ = ((B & 1) > 0) && addGreatZero()
	A <<= 1
	B >>= 1

	_ = ((B & 1) > 0) && addGreatZero()
	A <<= 1
	B >>= 1

	_ = ((B & 1) > 0) && addGreatZero()
	A <<= 1
	B >>= 1

	_ = ((B & 1) > 0) && addGreatZero()
	A <<= 1
	B >>= 1
	return ans >> 1
}
