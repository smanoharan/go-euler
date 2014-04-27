package main

import "math"

func problem71() string {
	const max = 1000 * 1000
	a, b := 2, 7 // 2/7 < 3/7
	for d := max; d > 0; d-- {
		if d == 7 {
			continue
		}
		i, t := 3*d/7, int(int64(d)*int64(a)/int64(b))
		for i > t && gcd(i, d) != 1 {
			i--
		}
		if i > t {
			a, b = i, d
		}
	}
	return itoa(a)
}

func problem72() string {
	const max = 1000*1000 + 1
	phi := make([]int, max)
	for i := range phi {
		phi[i] = i
	}

	isComposite := NewBitSet(max)

	phi[1] = 0
	for p := 2; p < max; p++ {
		if !isComposite.Get(p) {
			phi[p] = p - 1
			for i := 2 * p; i < max; i += p {
				isComposite.Set(i)
				phi[i] /= p
				phi[i] *= (p - 1)
			}
		}
	}

	phiSum := int64(0)
	for _, p := range phi {
		phiSum += int64(p)
	}
	return i64toa(phiSum)
}

func problem73() string {
	const max = 12000
	count := 0
	for d := 4; d <= max; d++ {
		s, f := int(math.Ceil(float64(d)/3.0)), d/2
		for n := s; n <= f; n++ {
			if gcd(n, d) == 1 {
				count++
			}
		}
	}
	return itoa(count)
}

func problem74() string {

	fac := make([]int, 10)
	fac[0] = 1
	for i := 1; i < 10; i++ {
		fac[i] = fac[i-1] * i
	}

	next := func(i int) int {
		res := 0
		for j := i; j > 0; j /= 10 {
			res += fac[j%10]
		}
		return res
	}

	endpt := make(map[int]int)
	endpt[169] = 3
	endpt[363601] = 3
	endpt[1454] = 3
	endpt[871] = 2
	endpt[45361] = 2
	endpt[872] = 2
	endpt[45362] = 2

	const max = 1000 * 1000
	count := 0
	for i := 1; i < max; i++ {
		j := 0
		for cur := i; true; j++ {
			_, exists := endpt[cur]
			if exists {
				j += endpt[cur]
			}

			nx := next(cur)
			if exists || cur == nx {
				if j == 60 {
					count++
				}
				break
			}
			cur = nx
		}
	}

	return itoa(count)
}

func problem75() string {
	// using Euclid's formula for generating pythagorean triples

	const max = 1500*1000 + 1
	const lim = 868 // sqrt(max/2)+1
	foundOne, foundTwo := NewBitSet(max), NewBitSet(max)

	for m := 1; m < lim; m++ {
		mm := m * m
		for n := (m & 1) + 1; n < m; n += 2 {
			if gcd(m, n) == 1 {
				nn := n * n
				a, b, c := mm-nn, 2*m*n, mm+nn
				if a*a+b*b != c*c {
					println(">>", m, n, a, b, c)
					panic("Failed")
				}
				L := a + b + c
				for kL := L; kL < max; kL += L {
					if foundOne.Get(kL) {
						foundTwo.Set(kL)
					} else {
						foundOne.Set(kL)
					}
				}
			}
		}
	}

	count := 0
	for i := 1; i < max; i++ {
		if foundOne.Get(i) && !foundTwo.Get(i) {
			count++
		}
	}

	return itoa(count)
}

func problem76() string {
	// using Euler's Pentagonal Number Formula for Partitions
	const max = 100
	const f1, f6 = float64(1), float64(6)

	pv := make([]int, max+1)
	pv[0] = 1
	for n := 1; n <= max; n++ {
		pv[n] = 0
		sqrtDet := math.Sqrt(float64(24*n + 1))
		ks := int(math.Ceil((f1 - sqrtDet) / f6))
		kf := int(math.Floor((f1 + sqrtDet) / f6))

		for k := ks; k <= kf; k++ {
			if k == 0 {
				continue
			}

			i := n - ((3*k - 1) * k / 2)
			m := 2*(k&1) - 1
			pv[n] += m * pv[i]
		}
	}

	return itoa(pv[max] - 1) // exclude the single number sum
}

func problem77() string {
	const pmax, cmax, targetCount = 100 * 1000, 100 * 1000, 5000
	ccmin := cmax - 1
	arr := make([]int, cmax)

	// using prime 2: there is 1 way iff the number is even (else 0)
	for i := 0; i < cmax; i++ {
		arr[i] = 1 - (i & 1)
	}

	// find the next primes
	isComposite := BuildPrimeSieve(pmax)

	for p := 3; p < pmax && p <= ccmin; p++ {
		if isComposite.Get(p) {
			continue
		}

		for j := p; j <= ccmin; j++ {
			arr[j] += arr[j-p]
			if arr[j] >= targetCount {
				ccmin = Min2i(j, ccmin)
			}
		}
	}

	if ccmin == cmax-1 {
		panic("No solutions found. Increase pMax and/or cMax. Last Elem: " + itoa(arr[cmax-1]))
	} else {
		return itoa(ccmin)
	}
}

func problem78() string {
	// using Euler's Pentagonal Number Formula for Partitions
	const mod = int64(1000 * 1000)
	const f1, f6 = float64(1), float64(6)

	pv := make([]int64, 1, 1000)
	pv[0] = 1

	for n := 1; true; n++ {
		pvn := int64(0)
		sqrtDet := math.Sqrt(float64(24*n + 1))
		ks := int(math.Ceil((f1 - sqrtDet) / f6))
		kf := int(math.Floor((f1 + sqrtDet) / f6))

		for k := ks; k <= kf; k++ {
			if k == 0 {
				continue
			}

			i := n - ((3*k - 1) * k / 2)
			m := int64(2*(k&1) - 1)
			pvn = (pvn + m*pv[i]) % mod
		}

		pv = append(pv, pvn)
		if pvn == 0 {
			return itoa(n)
		}
	}

	return "Impossible" // necessary for compilation.
}

func problem79() string {
	// done by hand, near trivial (since no digit is repeated)
	// effectively a topological sort of the poset defined by the attempts
	return "73162890"
}
