package main

import "math"

func problem61() string {
	const prefixMax, typeFirst, typeLast = 100, 3, 8
	const numTypes = typeLast - typeFirst + 1
	const cycleSize = numTypes

	var nbpt [prefixMax][numTypes - 1][]int // numbers by prefix and type

	for p := 0; p < prefixMax; p++ {
		for t := 0; t < numTypes-1; t++ {
			nbpt[p][t] = make([]int, 0, 20)
		}
	}

	for t := 0; t < numTypes-1; t++ {
		cur, diff, diffdiff := 1, t-1+typeFirst, t+1

		// skip 3-digit numbers and below
		for ; cur < 1000; diff += diffdiff {
			cur += diff
		}

		for ; cur < 10000; diff += diffdiff {
			nbpt[cur/100][t] = append(nbpt[cur/100][t], cur)
			cur += diff
		}
	}

	tryCycleWith := func(num int) (bool, []int) {
		const maxDepth = numTypes
		cycle := make([]int, maxDepth)
		usedTypes := NewBitSet(typeLast)

		// function address place holder (to allow recursion)
		buildCycle := func(depth int) bool {
			return false
		}

		buildCycle = func(depth int) bool {
			if depth == maxDepth {
				// now, we have generated the entire cycle
				// check for consistency against the first element
				return cycle[0]/100 == (cycle[depth-1] % 100)
			} else {
				// generate the next element of the cycle
				nextPrefix := cycle[depth-1] % 100
				for t := 0; t < numTypes-1; t++ {
					ns := nbpt[nextPrefix][t]
					nns := len(ns)
					if nns > 0 && !usedTypes.Get(t) {
						usedTypes.Set(t)

						// try every number in the current list
						for i := 0; i < nns; i++ {
							cycle[depth] = ns[i]

							// check number is not already in use
							repeat := false
							for j := 0; j < depth; j++ {
								if cycle[depth] == cycle[j] {
									repeat = true
									break
								}
							}

							if !repeat && buildCycle(depth+1) {
								return true
							}
						}

						usedTypes.Unset(t)
					}
				}
				return false
			}
		}

		cycle[0] = num
		return buildCycle(1), cycle
	}

	// generate each number of the biggest type, as the starting point of the cycle
	cur, diff, diffdiff := 1, numTypes-1-1+typeFirst, numTypes-1+1
	for ; cur < 1000; diff += diffdiff {
		cur += diff
	}
	for ; cur < 10000; diff += diffdiff {
		if success, cycle := tryCycleWith(cur); success {
			sum := 0
			for i := 0; i < len(cycle); i++ {
				sum += cycle[i]
			}
			return itoa(sum)
		}
		cur += diff
	}

	panic("No solution found")
}

func problem62() string {
	toCanonOrder := func(c uint64) uint64 {

		// bucket-sort the digits
		digits := make([]int, 10)
		for c > 0 {
			digits[c%10]++
			c /= 10
		}

		res := uint64(0)
		for d := 9; d >= 0; d-- {
			dd, du := digits[d], uint64(d)
			for i := 0; i < dd; i++ {
				res = (res * 10) + du
			}
		}

		return res
	}

	i := 1
	cubes := make(map[uint64]int)
	mincube := make(map[uint64]int)

	for true {
		ii := uint64(i)
		ci := toCanonOrder(ii * ii * ii)
		cubes[ci]++

		if cubes[ci] == 5 {
			mi := int64(mincube[ci])
			return i64toa(mi * mi * mi)
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

func problem65() string {
	const max = 100

	rat := NewRat(0, 1)
	for i := max - 1; i > 0; i-- {
		a := 1
		if (i % 3) == 2 {
			a = 2 * ((i + 1) / 3)
		}

		rat.Add(rat, NewRat(a, 1))
		rat.Inv(rat)
	}

	rat.Add(rat, NewRat(2, 1))

	return itoa(DigitalSum(rat.Num()))
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
				j := i * PF_LEN
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
	k := i * PF_LEN
	mj := pf[k]
	for j := 1; j <= mj; j++ {
		phi *= (f1 - (f1 / float64(pf[k+j])))
	}

	return phi
}

func problem69() string {
	const max = 1000 * 1000

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

func problem70() string {
	const max = 10 * 1000 * 1000
	pf := ToDistinctPrimeFactors(max)

	isPermutation := func(i, j int64) bool {
		digits := make([]int, 10)
		for ; i > 0; i /= 10 {
			digits[i%10]++
		}

		for ; j > 0; j /= 10 {
			k := j % 10
			digits[k]--
			if digits[k] < 0 {
				return false
			}
		}

		for i := 0; i < 10; i++ {
			if digits[i] != 0 {
				return false
			}
		}

		return true
	}

	minPhi, minN := float64(1), float64(max)
	for i := 2; i < max; i++ {
		fi := float64(i)
		phi := Phi(i, max, pf)
		if (fi*minPhi < phi*minN) && isPermutation(int64(i), int64(math.Floor(phi+0.5))) {
			minPhi, minN = phi, fi
		}
	}

	return itoa(int(math.Floor(minN + 0.5)))
}

func problem243() string {
	const max = 150 * 1000 * 1000
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
		if (n & 0x4ffff) == 0 {
			println(n, phi[n], n-1, phi[n]*denom/(numer*float64(n-1)))
		}

		if !isComposite.Get(n) {
			for i := n; i < max; i += n {
				isComposite.Set(i)
				// n is a prime factor of i
				phi[i] *= (f1 - f1/float64(n))
			}
		}

		// By now, phi[n] has been fully computed:
		// resilience R(n) = Phi(n) / n-1
		if denom*phi[n] < numer*float64(n-1) {
			println(n, phi[n], n-1)
			return itoa(n)
		}
	}

	return "Not found. Increase Max"
}
