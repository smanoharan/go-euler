package main

func problem206() string {
	// let a be the number s.t. a*a = 1_2_3_4_5_6_7_8_9_0
	// since a*a ends with a 0, 10 | a*a, so 2 | a*a and 5 | a*a
	// so 2 | a and 5 | a (as 2,5 are prime). so 10 | a.
	// so a must be of the form ___ ___ __0
	// and a*a must be 1_2_3_4_5_6_7_8_900
	//
	// consider b = a / 10.
	// b must be an 8 digit number
	// b*b is 1_2_3_4_5_6_7_8_9
	// b*b ends with a 9, so the last digit of b must be a 3 or 7.
	// (3*3 == 9, 7*7 == 49, no other square of a single digit ends with a 9)
	// so b must be 1_2_3_4_5_6_7_8X9 where X is 3 or 7
	//
	// b*b is at most 19293949596979899 so b is at most ~139*1000*1000
	// let c be the first 7 digits of b. c is at most 14*1000*1000
	// Now only ~9mill numbers to search.

	const i10, i3, i7, i100 = int64(10), int64(3), int64(7), int64(100)
	const mill = int64(1000*1000)
	const start, max = mill, int64(14)*mill

	isB := func(b int64) bool {
		bb := b*b
		for i := 9; i > 0; i-- {
			if (bb % i10) != int64(i) { return false }
			bb /= i100
		}
		return true
	}


	for c := start; c < max; c++ {
		b1, b2 := c*i10 + i3, c*i10 + i7
		if isB(b1) { return i64toa(b1) + "0" }
		if isB(b2) { return i64toa(b2) + "0" }
	}

	panic("No Solutions found.")
}
