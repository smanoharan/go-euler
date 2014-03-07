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

func GCD(a, b int) int {
	if b == 0 { return a }
	return GCD(b, a % b)
}

func problem33() string {
	numer, denom := 1, 1

	for c := 1; c < 10; c++ {
		for b := 1; b < 10; b++ {
			if b == c { continue }
			for a := 1; a < 10; a++ {
				if (10*a*b == 9*b*c + a*c) {
					println(10*b+a, "/", 10*a+c, " = ", b, "/", c)
					numer *= b
					denom *= c
				}
			}
		}
	}

	return itoa(denom / GCD(denom, numer))
}

func problem34() string {
	var digitFactorial [10]int
	digitFactorial[0] = 1
	for i := 1; i < 10; i++ {
		digitFactorial[i] = digitFactorial[i-1]*i		
	}

	// 9! = 362880
	// n digit number has a max digit factorial sum is n*362,880
	// 7 digit number (max 9,999,999) has a max digit factorial sum of (2,540,160)
	sum := 0
	for a := 0; a < 3; a++ {
		va, da := a, digitFactorial[a]
		if (a == 0) { da = 0 }

		for b := 0; b < 10; b++ {
			if (a==2) && b > 5 { break }
			vb, db := va*10 + b, da + digitFactorial[b]
			if (da == 0 && b == 0) { db = 0 }

			for c := 0; c < 10; c++ {
				vc, dc := vb*10 + c, db + digitFactorial[c]
				if (db == 0 && c == 0) { dc = 0 }

				for d := 0; d < 10; d++ {
					vd, dd := vc*10 + d, dc + digitFactorial[d]
					if (dc == 0 && d == 0) { dd = 0 }

					for e := 0; e < 10; e++ {
						ve, de := vd*10 + e, dd + digitFactorial[e]
						if (dd == 0 && e == 0) { de = 0 }

						for f := 0; f < 10; f++ {
							vf, df := ve*10 + f, de + digitFactorial[f]
							if (de == 0 && f == 0) { continue } // don't count 1 digit sums

							for g := 0; g < 10; g++ {
								vg, dg := vf*10 + g, df + digitFactorial[g]
								if vg == dg { sum += vg }
							}
						}
					}
				}
			}
		}
	}

	return itoa(sum)
}

func problem35() string {
	const max = 1000001
	count := 0
	isComposite := BuildPrimeSieve(max*5)
	
	for p := 2; p < max; p++ {
		if !isComposite.Get(p) {
			digits := toDigits(p)
			ld := len(digits)
			if ld == 1 {
				count++
			} else {
				mult := 1 // multiplier for right rotation (corresponds to value of MSD)
				for i := 1; i < ld; i++ { mult *= 10 }

				v, isCirc := p, true
				for rot := 0; rot < ld; rot++ {
					v = (v / 10) + (mult * (v % 10)) // right rotate: move last digit to front
					if isComposite.Get(v) { isCirc = false; break }
				}

				if isCirc { count++ }
			}
		}
	}

	return itoa(count)
}

func isBinaryPalindrome(v int, n uint) bool {
	for i, j := uint(0), n; i < j; i, j = i+1, j-1 {
		if ((v >> i) & 1) != ((v >> j) & 1) { return false }
	}
	return true
}

func problem36() string {
	sum := int64(0)
	for i, n := 1, uint(0); i < 1000000; i++ {
		if (i >> n) > 1 { n++ }
		if isPalindrome(i) && isBinaryPalindrome(i, n) {
			sum += int64(i)
		}
	}
	return i64toa(sum)
}

func problem37() string {
	const max = 1000000 // just a guess, refined until 11 such primes are found
	isComposite := BuildPrimeSieve(max)
	isLeftTruncatablePrime := NewBitSet(max)
	isRightTruncatablePrime := NewBitSet(max)

	// mark all 1-digit primes as left- & right- truncatable, (but don't include them in sum)
	for p := 2; p < 10; p++ {
		if !isComposite.Get(p) {
			isLeftTruncatablePrime.Set(p)
			isRightTruncatablePrime.Set(p)
		}
	}

	sum := 0
	for p, m := 10, 10; p < max; p++ {
		if p > 10*m { m *= 10 }

		if !isComposite.Get(p) {
			a := p / 10
			b := p - (p / m)*m

			left, right := isLeftTruncatablePrime.Get(a), isRightTruncatablePrime.Get(b)
			if left { isLeftTruncatablePrime.Set(p) }
			if right { isRightTruncatablePrime.Set(p) }
			
			if left && right {
				sum += p
			}
		}
	}

	return itoa(sum)
}

