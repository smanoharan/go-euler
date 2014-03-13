package main

import "sort"

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

func problem44() string {
	
	const max = 100000000
	isPenta := NewBitSet(max+1)

	count := 0
	for d,p := 4,1; p < max; d,p = d+3,p+d {
		isPenta.Set(p)
		count++
	}
	
	penta := func(n int) int {
		return n*(3*n - 1) / 2 
	}

	minD := max
	for pi := 1; pi < count; pi++ {
		ni := penta(pi)
		for pj,nj := pi+1,penta(pi+1); (nj - ni < minD) && (nj + ni) < max; pj,nj = pj+1,penta(pj+1) {
			if isPenta.Get(nj - ni) && isPenta.Get(nj + ni) {
				minD = nj - ni
				break
			}
		}
	}
	
	return itoa(minD)
}

func problem45() string {
	// iterate thru hex numbers (which are already triangle numbers)
	i := int64(144) // last number was hexNum(143)
	for !IsPentaNum(HexNum(i)) { i++ }
	return i64toa(HexNum(i))
}


func problem46() string {

	sieveMax := 1000*1000 // 100*1000*1000 
	sieve := NewBitSet(sieveMax)

	for p := 2; p < sieveMax; p++ {
		if !sieve.Get(p) {
			for q := 2*p; q < sieveMax; q+=p {
				sieve.Set(q)
			}
		} else if (p & 1) == 1 {
			found := false
			for d := 1; p - 2*d*d > 1; d++ {
				if !sieve.Get(p - 2*d*d) { 
					found = true
					break
				}
			}

			if !found { return itoa(p) }
		}
	}

	return "Not found. Increase sieveMax"
}

func problem47() string {
	const max = 1000*1000
	numFactors := make([]int, max)
	
	for p := 2; p < max; p++ {
		if numFactors[p] == 0 {
			for q := 2*p; q < max; q+=p {
				numFactors[q]++
			}
		} else if p > 4 && 
				numFactors[p] == 4 && 
				numFactors[p-1] == 4 && 
				numFactors[p-2] == 4 && 
				numFactors[p-3] == 4 {
			return itoa(p-3)
		}
	}

	return "Not found. Increase max"
}

func problem48() string {
	const mod = int64(10*1000)*int64(1000*1000)

	expMod := func(base, exp int) int64 {
		base64 := int64(base)
		prod := base64
		for i := 1; i < exp; i++ { 
			prod = (prod * base64) % mod 
		}
		return prod
	}

	sum := int64(0)
	for i := 1; i <= 1000; i++ {
		sum = (sum + expMod(i,i)) % mod
	}

	return i64toa(sum)
}

func problem49() string {
	
	toKey := func(i int) int {
		d1 := i / 1000
		i -= d1 * 1000
		d2 := i / 100
		i -= d2 * 100
		d3 := i / 10
		d4 := i - d3*10
		
		ds := sort.IntSlice([]int{d1,d2,d3,d4})
		ds.Sort()
		return ds[0]*1000 + ds[1]*100 + ds[2]*10 + ds[3]
	}

	const max = 10000
	isComposite := BuildPrimeSieve(max+1)
	var perms [max][]int

	for i := 1000; i < max; i++ {
		if !isComposite.Get(i) {

			iKey := toKey(i)
			if perms[iKey] == nil {
				perms[iKey] = []int{i}
			} else {
				for _, j := range perms[iKey] {
					k := 2*i - j
					if k < max && !isComposite.Get(k) && iKey == toKey(k) {
						if toKey(i) != 1478 {
							return itoa(j) + itoa(i) + itoa(k) 
						}
					}
				}

				perms[iKey] = append(perms[iKey], i)
			}
		}
	}
		
	return "Error. Not Found."
}

func problem50() string {
	const max = 1000000

	var primes [max]int
	numPrimes := 0
	appendPrime := func(prime int) {
		primes[numPrimes] = prime
		numPrimes++
	}

	isComposite := IteratePrimes(max+1, appendPrime)

	maxJ, P := 0, -1
	for i := 0; i < numPrimes; i++ {
		maxj := Min2i(max, numPrimes - i)
		if maxj < maxJ { break }

		curSum := 0
		for j := 0; j < maxj; j++ {
			pj := primes[i+j]
			curSum += pj
			if curSum >= max { break }
			if !isComposite.Get(curSum) {
				if j+1 > maxJ {
					maxJ, P = j+1, curSum
				}
			}
		}
	}

	return itoa(P)
}


