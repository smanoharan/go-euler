package main

func problem62() string {
	toCanonOrder := func(c uint64) uint64 {
		
		// bucket-sort the digits
		digits := make([]int, 10)
		for c > 0 {
			digits[c % 10]++
			c /= 10
		}

		res := uint64(0)
		for d := 9; d >= 0; d-- {
			dd, du := digits[d], uint64(d)
			for i := 0; i < dd; i++ { res = (res * 10) + du }
		}
		
		return res
	}

	i := 1
	cubes := make(map[uint64]int)
	mincube := make(map[uint64]int)

	for true {
		ii := uint64(i)
		ci := toCanonOrder(ii*ii*ii)
		cubes[ci]++
		
		if cubes[ci] == 5 {
			mi := int64(mincube[ci])
			return i64toa(mi*mi*mi) 
		} else if cubes[ci] == 1 {
			mincube[ci] = i
		}

		i++
	}
	panic("Impossible.")
}

func problem63() string {
	count := 0
	for base := 1; base < 10; base++ {
		val, baseBig := NewBig(base), NewBig(base)
		for exp := 1; exp == len(val.String()); exp++ {
			println("Incl", base, exp, val.String())
			MultBy(val, baseBig)
			count++
		}
	}
	return itoa(count)
}

func problem67() string {
	return i64toa(sumMax(ReadGrid("data/p67.txt")))
}

func problem69() string {
	const max = 1000*1000
	const f1 = float64(1)
	
	maxInvPhi, maxN := float64(-1), -1

	pf := make([][]int, max) // stores prime factors
	for i := 0; i < max; i++ {
	 	// 2^20 == 1024*1024 >= 1000000 (so num prime factors < 20 for all n <= 1000000)
		pf[i] = make([]int, 22)
		pf[i][0] = 0
	}

	isComposite := NewBitSet(max)
	for n := 2; n < max; n++ {
		if !isComposite.Get(n) {
			for i := n; i < max; i += n {
				isComposite.Set(i)
				pf[i][0]++
				pf[i][pf[i][0]] = n
			}
		}
	}
	
	for i := 2; i < max; i++ {
		phi := float64(i)
		mj := pf[i][0]
		for j := 1; j <= mj; j++ {
			phi *= (f1 - (f1 / float64(pf[i][j])))
		}
		
		invPhi := float64(i) / phi
		if invPhi > maxInvPhi {
			maxInvPhi = invPhi
			maxN = i
		}
	}
	
	return itoa(maxN)
}
