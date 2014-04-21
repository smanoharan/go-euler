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

const PF_LEN = 13 // enough for 12 slots. 2*3*...*13 = 13! ~= 6B > 4.3B; so enough for all 32-bit numbers

func ToDistinctPrimeFactors(max int) []int {

	pf := make([]int, max*PF_LEN) // stores prime factors
	for i := 0; i < max; i++ {
		pf[i*PF_LEN] = 0
	}

	isComposite := NewBitSet(max)
	for n := 2; n < max; n++ {
		if !isComposite.Get(n) {
			for i := n; i < max; i += n {
				isComposite.Set(i)
				j := i*PF_LEN
				pf[j]++
				pf[j+pf[j]] = n
			}
		}
	}

	return pf
}

func Phi(i, max int, pf []int) float64 {
	const f1 = float64(1)
	
	phi := float64(i)
	k := i*PF_LEN
	mj := pf[k]
	for j := 1; j <= mj; j++ {
		phi *= (f1 - (f1 / float64(pf[k+j])))
	}

	return phi
}

func problem69() string {
	const max = 1000*1000
	
	maxInvPhi, maxN := float64(-1), -1
	pf := ToDistinctPrimeFactors(max)
	
	for i := 2; i < max; i++ {
		invPhi := float64(i) / Phi(i, max, pf)
		if invPhi > maxInvPhi {
			maxInvPhi = invPhi
			maxN = i
		}
	}
	
	return itoa(maxN)
}

func problem243() string {
	const max = 150*1000*1000
	const numer, denom = float64(15499), float64(94744)
	//const numer, denom = float64(4), float64(10)
	
	// memory-optimised version for computing Phi(n)
	// result is identical to calling ToDistinctPrimeFactors then Phi on each element
	const f1 = float64(1)

	phi := make([]float64, max) 
	for i := 0; i < max; i++ {
		phi[i] = float64(i)
	}

	isComposite := NewBitSet(max)
	for n := 2; n < max; n++ {
		if ((n & 0x4ffff) == 0) {
			println(n, phi[n], n-1, phi[n]*denom / (numer*float64(n-1)))
		}

		if !isComposite.Get(n) {
			for i := n; i < max; i += n {
				isComposite.Set(i)
				// n is a prime factor of i
				phi[i] *= (f1 - f1 / float64(n))
			}
		}

		// By now, phi[n] has been fully computed:
		// resilience R(n) = Phi(n) / n-1
		if (denom*phi[n] < numer*float64(n-1)) {
			println(n, phi[n], n-1)
			return itoa(n)
		}
	}

	return "Not found. Increase Max"
}
