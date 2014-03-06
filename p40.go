package main

func problem31() string {
	const max = 200
	var numCoins [2][max+1]int
	coinVal := []int{ 2, 5, 10, 20, 50, 100, 200 }
	
	for i := 0; i <= max; i++ {
		numCoins[0][i] = 1 // using 1p coins, there is exactly one way to make any amount
	}

	cur := 1
	last := 0
	for _, v := range coinVal {
	
		// for any amount less than v, no coins of value v can be used,
		// so the number of ways of making the amount with coin v or smaller
		// is the same as the number of ways of making the amount without coin v or bigger.
		for c := 0; c < v; c++ {
			numCoins[cur][c] = numCoins[last][c]
		}

		// for any other amount, the number of ways of making the amount with coin v or smaller
		// is the number of ways of making the amount without using coin v or bigger, plus
		// the number of ways of making amount - v using coin v or smaller (i.e. the last coin is v).
		for c := v; c <= max; c++ {
			numCoins[cur][c] = numCoins[last][c] + numCoins[cur][c-v]
		}
		cur, last = last, cur
	}
	return itoa(numCoins[last][max])
}

func problem32() string {
	// product can be at most 4 digits long (5 digit product means max of 99 x 99 < 10000)
	
	const (
		max = 10000
		z = uint(0)
		m = uint(10)
	)
	
	digitMasks := make([]int, max)
	pandigitalMask := (1 << 10) - 2 // last bit (corresponding to the 0-digit) is zero.
	
	// find the digitMask for all numbers with non-repeating digits 
	// (and no-significant-zero digits) under 10000
	for a := z; a < m; a++ {

		for b := z; b < m; b++ {
			if a == b && b > 0 { continue }
			if a > 0 && b == 0 { continue }

			for c := z; c < m; c++ {
				if (a == c || b == c) && c > 0  { continue }
				if (a > 0 && c == 0) { continue }

				for d := z; d < m; d++ {
					if (a == d || b == d || c == d) && d > 0 { continue }
					if a > 0 && d == 0 { continue }
					
					prod := a*1000 + b*100 + c*10 + d
					mask := ((1 << a) | (1 << b) | (1 << c) | (1 << d)) & pandigitalMask
					digitMasks[prod] = mask
				}
			}
		}
	}
	
	visProd := NewBitSet(max+1)
	sum := 0
	
	// try all pairs of numbers which might result in a pandigital product
	for a := 1; a < max; a++ {
		for b := a; b < max; b++ {
			prod := a * b
			if ((prod < max) && (!visProd.Get(prod)) && 
				((digitMasks[a] & digitMasks[b]) == 0) &&
				((digitMasks[prod] & digitMasks[a]) == 0) && 
				((digitMasks[prod] & digitMasks[b]) == 0) &&
				((digitMasks[a] | digitMasks[b] | digitMasks[prod]) == pandigitalMask)) {
				visProd.Set(prod)
				sum += prod
			}
		}
	}

	return itoa(sum)
}
