package main

func isPanDigital(i int) bool {
	mask := 0
	for i > 0 {
		lastDigit := i % 10
		if lastDigit == 0 { return false }

		dMask := 1 << uint(lastDigit - 1)
		if (mask & dMask) > 0 { return false }

		mask |= dMask 
		i /= 10
	}

	return ((mask & (mask + 1)) == 0)
}

func problem41() string {
	max := 1000*1000*10
	isComposite := BuildPrimeSieve(max+1)
	
	for i := max; i >= 2; i-- {
		if !isComposite.Get(i) && isPanDigital(i) {
			return itoa(i)
		}
	}

	return "<None Found>"
}

func problem42() string {
	count := 0
	max := 1000 // allows words of size 50 or so

	isTri := NewBitSet(max+1)
	for i, j := 2, 1; j < max ; i, j = i+1, j+i {
		isTri.Set(j)
	}

	line := ReadAllLines("data/p42.txt")[0]
	n := len(line)

	QUOTE := int('"')

	for i := 0; i < n; i++ {
		wordSum := 0
		for i++; int(line[i]) != QUOTE; i++ {
			wordSum += int(line[i]) - int('A') + 1
		}
		if isTri.Get(wordSum) { count++ }
		i++ // skip comma
	}

	return itoa(count)
}

func problem43() string {
	
	pDivides := func(i, j, k, mod int) bool {
		return ((i*100 + j*10 + k) % mod) == 0
	}

	isValid := func(p *Permutation, mid int) bool {
		pp := *p
		return (
			((pp[3] & 1) == 0) && // 2 div (d2 d3 d4) ==> d4 must be even
			pDivides(pp[2], pp[3], pp[4],  3) && // 5 div is implicit
			pDivides(pp[4],   mid, pp[5],  7) && 
			pDivides(  mid, pp[5], pp[6], 11) && 
			pDivides(pp[5], pp[6], pp[7], 13) && 
			pDivides(pp[6], pp[7], pp[8], 17))
	}

	toNum := func(p *Permutation, mid int) int64 {
		pp := *p
		return (
			int64(pp[0])*int64(1000000000) + 
			int64(pp[1]*100 + pp[2]*10 + pp[3])*int64(1000*1000) + 
			int64(pp[4]*100 +     5*10 + pp[5])*int64(1000) + 
			int64(pp[6]*100 + pp[7]*10 + pp[8]))
	}

	sum := int64(0)

	// since 5 div (d4 d5 d6) ; d6 must be 0 or 5.
	p := Permutation([]int {1, 0, 2, 3, 4, 6, 7, 8, 9})
	for p.NextPermutation() {
		if isValid(&p, 5) {
			sum += toNum(&p, 5)
		}
	}
	
	p = Permutation([]int {1, 2, 3, 4, 5, 6, 7, 8, 9})
	for p.NextPermutation() {
		if isValid(&p, 0) {
			sum += toNum(&p, 0)
		}
	}
	
	return i64toa(sum)
}


