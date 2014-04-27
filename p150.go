package main

func problem145() string {
	// 9-digit numbers cannot be reversible,
	// as the middle (5th) digits would add together (to produce an even digit)
	// (and any carry from the (4,6)th digit would also carry to (7,3) digit and onwards)
	// so only need to check 8-digit numbers and below
	const max = 100 * 1000 * 1000

	reverse := func(i int) int {
		rev := 0
		for ; i > 0; i /= 10 {
			rev *= 10
			rev += i % 10
		}
		return rev
	}

	hasOddDigitsOnly := func(i int) bool {
		for ; i > 0; i /= 10 {
			if ((i % 10) & 1) == 0 {
				return false
			}
		}
		return true
	}

	count := 0
	for i := 10; i < max; i++ {
		if (i % 10) == 0 {
			continue
		}

		if hasOddDigitsOnly(i + reverse(i)) {
			count++
		}
	}
	return itoa(count)
}
