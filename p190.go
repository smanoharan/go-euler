package main

func problem187() string {
	const max = 100 * 1000 * 1000
	const halfMax = max / 2
	isComposite := BuildPrimeSieve(halfMax + 1)

	sqrtMax := SqrtI(max)
	primes := make([]int, 0, max/1000)

	for p := 2; p <= sqrtMax; p++ {
		if !isComposite.Get(p) {
			primes = append(primes, p)
		}
	}
	primeCount := len(primes)

	for p := primes[primeCount-1] + 1; p <= halfMax; p++ {
		if !isComposite.Get(p) {
			primes = append(primes, p)
		}
	}

	// the product of any two primes under sqrtMax is under max,
	// and is a semi-prime
	semiPrimeCount := ((primeCount + 1) * primeCount) / 2

	// any extra semi-primes are found by product of a prime under sqrtMax
	// and another over sqrtMax
	for pi, pj := 0, len(primes)-1; pi < primeCount; pi++ {
		ppi := primes[pi]
		for primes[pj]*ppi >= max {
			pj--
		}
		semiPrimeCount += pj - primeCount + 1
	}

	return itoa(semiPrimeCount)
}
